package my_farm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"safebsc/features/shared"
	"safebsc/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type connectionClient struct {
	db mongo.DB
}

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

func myFarmHandler(c *gin.Context) {
	log.Println(">>>>>>>>>>>>   myFarmHandler")
	db, err := mongo.Connect()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	conn := connectionClient{
		db: db,
	}

	defer db.Client.Disconnect(db.Ctx)

	supportedFarms, err := conn.getSupporttedFarms()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	walletAddress := c.Param("walletAddress")
	myFarms, err := conn.findMyFarms(walletAddress, supportedFarms)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	contents, err := json.Marshal(myFarms)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func (dbClient connectionClient) findMyFarms(walletAddress string, supportedFarms []shared.SupportedFarm) ([]myFarm, error) {
	defer timeTrack(time.Now(), "findMyFarms")
	snaps, err := dbClient.findUserFarmsByWallet(walletAddress)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	tokensRate, err := getTokenPrice(dbClient)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	farmsApr, err := getFarmApr(dbClient)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	myFarms, err := transformToFarms(walletAddress, snaps, supportedFarms, tokensRate, farmsApr)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	return myFarms, nil
}

func (dbClient connectionClient) findUserFarmsByWallet(walletAddress string) ([]UserPortFolio, error) {
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("user_portfolio_latests")

	snaps := make([]UserPortFolio, 0)
	cursor, err := collection.Find(dbClient.db.Ctx, bson.M{"wallet_address": walletAddress})
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return snaps, err
	}
	if err = cursor.All(dbClient.db.Ctx, &snaps); err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return snaps, err
	}

	return snaps, nil
}

func (dbClient connectionClient) getSupporttedFarms() ([]shared.SupportedFarm, error) {
	var sf []shared.SupportedFarm
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("supported_farms")

	cur, err := collection.Find(dbClient.db.Ctx, bson.M{})
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(dbClient.db.Ctx) {
		var elem shared.SupportedFarm
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		sf = append(sf, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(dbClient.db.Ctx)
	return sf, err
}

func (dbClient connectionClient) findSnapShortByWallet(walletAddress string) ([]UserPortFolio, error) {
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("user_portfolios")

	var lastSnapshot UserPortFolio
	snaps := make([]UserPortFolio, 0)
	findOneOptions := options.FindOneOptions{}
	findOneOptions.SetSort(bson.D{{"batch_no", -1}})
	err := collection.FindOne(dbClient.db.Ctx, bson.M{"wallet_address": walletAddress}).Decode(&lastSnapshot)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return snaps, err
	}

	condition := bson.M{
		"$and": []bson.M{
			{"wallet_address": walletAddress},
			{"batch_no": lastSnapshot.BatchNo},
		},
	}
	cursor, err := collection.Find(dbClient.db.Ctx, condition)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return snaps, err
	}
	if err = cursor.All(dbClient.db.Ctx, &snaps); err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return snaps, err
	}

	return snaps, nil
}

func transformToFarms(walletAddress string, snaps []UserPortFolio,
	supportedFarms []shared.SupportedFarm,
	tokensRate []tokenPrice,
	farmsApr []farmApr) ([]myFarm, error) {
	myFarms := make([]myFarm, 0)
	for _, supportedFarm := range supportedFarms {
		farmSnaps := filterByFarmID(snaps, supportedFarm.OrderID)
		if len(farmSnaps) > 0 {
			farmDetails := make([]farmDetail, 0)
			var updatedAt time.Time
			tokenPrice := getRateStrBySymbol(supportedFarm.TokenSymbol, tokensRate)
			totalFarmDailyRewardAmount := 0.0
			for _, fs := range farmSnaps {
				updatedAt = fs.UpdatedAt
				var tokens []token
				var names []string
				for _, lp := range fs.LpTokens {
					amount, err := strconv.ParseFloat(lp.Amount, 64)
					if err != nil {
						return nil, err
					}
					names = append(names, lp.Symbol)
					imgNameUrl := strings.ToLower(lp.Symbol)
					if imgNameUrl == "4belt" {
						imgNameUrl = "belt"
					}
					imgUrl := fmt.Sprintf("assets/icons/token/%s.png", imgNameUrl)
					balanceUsd := amount * getRateBySymbol(lp.Symbol, tokensRate)
					tokens = append(tokens, token{
						Symbol:        lp.Symbol,
						Name:          lp.Symbol,
						ImageURL:      imgUrl,
						Balance:       amount,
						BalanceUsd:    balanceUsd,
						BalanceStr:    formatNumber(lp.Amount, 10),
						BalanceUsdStr: float64ToStrFormat(balanceUsd, 10),
						RawBalance:    lp.Amount,
					})
				}

				balanceUsd, balance, rewards, err := getRewardAndBalance(fs, supportedFarm, tokensRate, tokens)
				if err != nil {
					return nil, err
				}

				apr := getAprByNameAndPoolID(supportedFarm.FarmName, fs.PoolID, farmsApr)
				aprNum, err := strconv.ParseFloat("0"+apr.Pool.Apr, 64)
				if err != nil {
					return nil, err
				}
				aprDaily := aprNum / 365
				dailyRewardAmount := aprDaily * balanceUsd
				totalFarmDailyRewardAmount += dailyRewardAmount

				farmDetails = append(farmDetails, farmDetail{
					Name:              strings.Join(names, "-"),
					PoolID:            fs.PoolID,
					Reward:            rewards,
					Tokens:            tokens,
					BalanceUsd:        balanceUsd,
					BalanceUsdStr:     float64ToStrFormat(balanceUsd, 10),
					BalanceStr:        float64ToStrFormat(balance, 10),
					Apr:               apr.Pool.Apr,
					AprDaily:          float64ToStr(aprDaily),
					DailyRewardAmount: float64ToStr(dailyRewardAmount),
					LpAddress:         fs.LpAddress,
					PortAmount:        fs.PortAmount,
				})

			}
			totalPrice := sumTotalPriceOfFarm(farmDetails)

			myFarms = append(myFarms, myFarm{
				FarmID:            supportedFarm.OrderID,
				Name:              supportedFarm.FarmName,
				TotalPrice:        float64ToStrFormat(totalPrice, 10),
				WalletAddress:     walletAddress,
				FarmDetails:       farmDetails,
				ImageURL:          supportedFarm.FarmIcon,
				FarmInfo:          supportedFarm,
				UpdatedAt:         updatedAt,
				TokenPrice:        formatNumber(tokenPrice, 10),
				DailyRewardAmount: float64ToStr(totalFarmDailyRewardAmount),
			})
		}
	}

	return myFarms, nil
}

func getRewardAndBalance(fs UserPortFolio, supportedFarm shared.SupportedFarm, tokensRate []tokenPrice, tokens []token) (float64, float64, []rewardResponse, error) {
	rewards := make([]rewardResponse, 0)
	var reward float64
	var rewardUSD float64

	for _, rw := range fs.Reward {
		reward, err := strconv.ParseFloat(rw.Amount, 64)
		if err != nil {
			return 0, 0, nil, err
		}

		rewardUSD := reward * getRateBySymbol(supportedFarm.TokenSymbol, tokensRate)

		rewards = append(rewards, rewardResponse{
			Symbol:       rw.Symbol,
			Reward:       reward,
			RewardStr:    formatNumber(rw.Amount, 3),
			RewardUsd:    rewardUSD,
			RewardUsdStr: float64ToStrFormat(rewardUSD, 2),
		})
	}

	balanceUsd := sumTotalPriceOfDetail(rewardUSD, tokens)
	balance := sumTotalBalanceOfDetail(reward, tokens)

	return balanceUsd, balance, rewards, nil
}

func getTokenPrice(dbClient connectionClient) ([]tokenPrice, error) {
	conf := make([]tokenPrice, 0)
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("token_prices")
	cursor, err := collection.Find(dbClient.db.Ctx, bson.M{})

	if err != nil {
		log.Println(err.Error())
		return conf, err
	}

	if err = cursor.All(dbClient.db.Ctx, &conf); err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	return conf, err
}

func getFarmApr(dbClient connectionClient) ([]farmApr, error) {
	conf := make([]farmApr, 0)
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("pool_aprs")
	log.Println(">>>>>>>>> >>>   getFarmApr1")
	cursor, err := collection.Aggregate(dbClient.db.Ctx,
		[]bson.M{
			bson.M{
				"$group": bson.M{
					"_id":  []string{"$farm_name", "$pool.poolId"},
					"f":    bson.M{"$last": "$farm_name"},
					"date": bson.M{"$last": "$updatedAt"},
					"data": bson.M{"$last": "$pool"},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id":       "$_id",
					"farm_name": "$f",
					"pool":      "$data",
					"updatedAt": "$date",
				},
			},
		})
	if err != nil {
		log.Println(">>>>>>>>> >>>   getFarmApr Error")
		log.Println(err.Error())
	}
	log.Println(">>>>>>>>> >>>   getFarmApr2")
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(">>>>>>>>> >>>   getFarmApr3")
	if err = cursor.All(dbClient.db.Ctx, &conf); err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(">>>>>>>>> >>>   getFarmApr4")
	return conf, err
}

func sumTotalPriceOfFarm(dts []farmDetail) float64 {
	var sum float64
	for _, dt := range dts {
		for _, rw := range dt.Reward {
			sum += rw.RewardUsd
		}
		for _, tk := range dt.Tokens {
			sum += tk.BalanceUsd
		}
	}
	return sum
}

func sumTotalPriceOfDetail(rewardUsd float64, tks []token) float64 {

	var sum float64
	sum += rewardUsd
	for _, tk := range tks {
		sum += tk.BalanceUsd
	}
	return sum
}

func sumTotalBalanceOfDetail(reward float64, tks []token) float64 {

	var sum float64
	sum += reward
	for _, tk := range tks {
		sum += tk.Balance
	}
	return sum
}

func filterByFarmID(ss []UserPortFolio, farmID string) (ret []UserPortFolio) {
	for _, s := range ss {
		if s.FarmID == farmID {
			ret = append(ret, s)
		}
	}
	return
}

func getRateBySymbol(symbol string, tkp []tokenPrice) float64 {
	for _, tk := range tkp {
		if tk.Symbol == symbol {

			if s, err := strconv.ParseFloat(tk.Price, 64); err == nil {
				return s
			}
		}
	}
	return 0
}

func getRateStrBySymbol(symbol string, tkp []tokenPrice) string {
	for _, tk := range tkp {
		if tk.Symbol == symbol {

			return tk.Price
		}
	}
	return "0"
}

func getRateByTokenAddress(tokenAddress string, tkp []tokenPrice) float64 {
	for _, tk := range tkp {
		if tk.Token == tokenAddress {

			if s, err := strconv.ParseFloat(tk.Price, 64); err == nil {
				return s
			}
		}
	}
	return 0
}

func getAprByLpToken(lpToken string, farmsApr []farmApr) farmApr {
	for _, tk := range farmsApr {
		if tk.Pool.PoolAddress == lpToken {

			return tk
		}
	}

	return farmApr{}
}

func getAprByNameAndPoolID(farmName string, poolId int, farmsApr []farmApr) farmApr {
	for _, tk := range farmsApr {
		if tk.FarmName == farmName && tk.Pool.PoolId == poolId {

			return tk
		}
	}

	return farmApr{}
}

func formatNumber(number string, floatingDigit int) string {
	fd := strconv.Itoa(floatingDigit)

	_, err := strconv.ParseFloat(number, 64)
	if err != nil {
		log.Println(err)
	}

	if s, err := strconv.ParseFloat(number, 64); err == nil {
		fNum := roundDown(s, floatingDigit)
		return fmt.Sprintf("%."+fd+"f", fNum)
	}
	return fmt.Sprintf("%."+fd+"f", 0.00)
}

func float64ToStrFormat(number float64, floatingDigit int) string {
	return formatNumber(float64ToStr(number), floatingDigit)
}

func float64ToStr(number float64) string {
	return fmt.Sprintf("%v", number)
}

func roundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %dms", name, elapsed.Nanoseconds()/1000000)
}
