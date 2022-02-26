package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"safebsc/features/pro_account/bep20"
	"safebsc/features/pro_account/masterchef"
	pancakefactory "safebsc/features/pro_account/pancake-factory"
	pancakepair "safebsc/features/pro_account/pancake-pair"
	stdreference "safebsc/features/pro_account/std-reference"
	"safebsc/mongo"
	"safebsc/postgres"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var globalClient *ethclient.Client
var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

func init() {
	var err error
	globalClient, err = ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		panic(err)
	}
}

const proValueMinimum = 100
const SCZ = "0x39f1014a88c8ec087cedf1bfc7064d24f507b894"
const mSCZ = "0x7f0ecd7d237efc78499da30cf74ba55a87310735"
const SCZ_BUSD = "0xb3d2C0cb104CBfA2167Af0D82A12475946B22386"
const SCZ_DBM = "0xF31455AE22DfE3637b87D62371F6fc945F9ea301"
const SCZ_BNB = "0xad468bcbcb33037e061ed5ec905ddea06cafc67f"
const SCZ_EVE = "0x91b8be48e6f20226ad4e70d5ec7db476151707b5"

const chefMoon = "0xbe739a112ef6278ceb374bcad977252bc3918ca9"
const chefFoodCourt = "0x07b48c393cababf6458e72d57b7ccd8f73e94c03"
const chefNeverSwap = "0x8fb60dd3557c491e04d00a06fdc0618423a3c618"

const pancakeRouter = "0x10ED43C718714eb63d5aA57B78B54704E256024E"

func bandPrice(vSymbol string) float64 {

	var rate = float64(0)
	stdInstance, err := stdreference.NewStdreference(common.HexToAddress("0xDA7a001b254CD22e46d3eAB04d937489c93174C3"), globalClient)
	if err != nil {
		return float64(0)
	}

	rates, err := stdInstance.GetReferenceData(&bind.CallOpts{}, vSymbol, "USD")
	if err != nil {
		return rate
	}

	if float64(rates.Rate.Uint64()) != 3963877391197344453575983046348115674221700746820753546331534351508065746944 {
		rate = float64(rates.Rate.Uint64()) / 1e18
	}

	return rate
}

func checkVirtualToken(tokenAddress string) (address string, rate float64) {

	if tokenAddress == "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c" {
	}

	var virtualToken = virtualTokens
	var tmp []virtualStruct

	for _, v := range virtualToken {
		if v.sourceAddr == tokenAddress {
			tmp = append(tmp, v)
		}
	}

	var bandPriceRate = float64(0)
	if len(tmp) > 0 {
		bandPriceRate = bandPrice(tmp[0].destName)
		if tmp[0].destAddr != "" {
			tokenAddress = tmp[0].destAddr
		}
	}

	return tokenAddress, bandPriceRate
}

func getRateUSD(tokenAddress string) float64 {

	virtualAddress, rates := checkVirtualToken(tokenAddress)

	if rates > 0 {
		return rates
	} else {
		tokenAddress = virtualAddress
	}

	const DOPLP = "0x9116f04092828390799514bac9986529d70c3791"
	const DOPUSTLP = "0x7edcdc8cd062948ce9a9bc38c477e6aa244dd545"
	const DOP2PLP = "0x124166103814e5a033869c88e0f40c61700fca17"
	const DOP3PLP = "0xaa5509ce0ecea324bff504a46fc61eb75cb68b0c"

	address := common.HexToAddress(tokenAddress)
	tokenContract, err := bep20.NewBep20(address, globalClient)
	if err != nil {
		log.Fatal(err)
	}

	const factoryCakeV2 = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"

	const SIX = "0x070a9867ea49ce7afc4505817204860e823489fe"
	const FINIX = "0x0f02b1f5af54e04fb6dd6550f009ac2429c4e30d"
	const JDI = "0x0491648c910ad2c1afaab733faf71d30313df7fc"

	const SCZ = "0x39f1014a88c8ec087cedf1bfc7064d24f507b894"
	const WEX = "0xa9c41a46a6b3531d28d5c32f6633dd2ff05dfb90"
	const PRIVACY = "0x7762a14082ab475c06d3868b385e46ae27017231"
	const ALESWAP = "0x99ca242f20424a6565cc17a409e557fa95e25bd7"

	const WAD = "0x0feadcc3824e7f3c12f40e324a60c23ca51627fc"
	const BIDR = "0x9a2f5556e9a637e8fbce886d8e3cf8b316a1d8a2"

	const BANANA = "0x603c7f932ed1fc6575303d8fb018fdcbb0f39a95"

	const DOLLY = "0xff54da7caf3bc3d34664891fc8f3c9b6dea6c7a5"
	const TWIN = "0x3806aae953a3a873d02595f76c7698a57d4c7a57"

	const TSLA = "0x17ace02e5c8814bf2ee9eaaff7902d52c15fb0f4"
	const AAPL = "0xc10b2ce6a2bcfdfdc8100ba1602c1689997299d3"
	const AMZN = "0x1085b90544ff5c421d528aaf79cc65afc920ac79"
	const GOOGL = "0x9c169647471c1c6a72773cffc50f6ba285684803"

	const COUPON = "0x084bb94e93891d74579b54ab63ed24c4ef9cd5ef"
	const kMATIC = "0x032574b64bf6fa42951f836cc8c5099d1c5747d3"
	const MATIC = "0xcc42724c6683b7e57334c4e856f4c9965ed682bd"
	const wMMP = "0x422d0a431d8fb752e3697e90ba04b3324ea0cb4a"
	const wSAFEMARS = "0x40733abc9acb7d48caa632ee83e4e7b3d0008d9d"
	const wSAFEMOON = "0xa3863434a1fc699185b3e6809a933056d1178366"

	const FAT = "0x73280e2951785f17acc6cb2a1d0c4d65031d54b3"

	const CSS = "0x3bc5798416c1122BcFd7cb0e055d50061F23850d"

	const BNB = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
	const BUSD = "0xe9e7cea3dedca5984780bafc599bd69add087d56"

	targetToken := BNB

	contractAddress := common.HexToAddress(factoryCakeV2)
	pancakefactory, err := pancakefactory.NewPancakefactory(contractAddress, globalClient)
	pairAddress, err := pancakefactory.GetPair(&bind.CallOpts{}, common.HexToAddress(BNB), common.HexToAddress(BUSD))
	if err != nil {
		log.Fatal(err)
	}

	pairInstance, err := pancakepair.NewPancakepair(pairAddress, globalClient)
	reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	BNBRate := float64(reserve.Reserve1.Uint64()) / float64(reserve.Reserve0.Uint64())

	var isHaveAddress = false
	support := []string{SCZ, WEX, PRIVACY, ALESWAP}

	for _, v := range support {
		if v == tokenAddress {
			isHaveAddress = true
		}
	}

	if isHaveAddress {
		targetToken = BUSD
	}

	pairAddress, err = pancakefactory.GetPair(&bind.CallOpts{}, common.HexToAddress(tokenAddress), common.HexToAddress(targetToken))
	if err != nil {
		log.Fatal(err)
	}
	rate := float64(0)

	if pairAddress.String() != "0x0000000000000000000000000000000000000000" {

		pairInstance, err := pancakepair.NewPancakepair(pairAddress, globalClient)
		reserves, err := pairInstance.GetReserves(&bind.CallOpts{})

		if err != nil {
			log.Fatal(err)
		}

		tokenA, err := pairInstance.Token0(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		reserve1 := new(big.Float)
		reserve1.SetInt(reserves.Reserve1)

		reserve0 := new(big.Float)
		reserve0.SetInt(reserves.Reserve0)

		if strings.ToLower(tokenA.String()) == strings.ToLower(tokenAddress) {
			balance := new(big.Float).Quo(reserve1, reserve0)
			bigval, _ := balance.Float64()
			rate = bigval
		} else {
			balance := new(big.Float).Quo(reserve0, reserve1)
			bigval, _ := balance.Float64()
			rate = bigval
		}

		decimals, err := tokenContract.Decimals(&bind.CallOpts{})
		rate = rate / float64(new(big.Int).Exp(big.NewInt(10), big.NewInt(18-int64(decimals)), big.NewInt(0)).Uint64())

		if targetToken == BNB {
			rate = rate * BNBRate
		} else if targetToken == BUSD {
			return rate
		} else {
			_rate := getRateUSD(targetToken)
			rate = rate * _rate
		}
	}
	return rate
}

func updateIsPremium(account string, level string, c *gin.Context) {

	db := postgres.Connect()
	defer db.Close()

	var user user
	location, _ := time.LoadLocation("Asia/Bangkok")
	user.Level = level
	user.UpdatedDate = time.Now().In(location)

	if _, err := db.Model(&user).Where("wallet_address = ?", account).Column("level", "updated_date").Returning("*").Update(); err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}
	Response("", http.StatusOK, c)
}

func getProAccount(c *gin.Context) {

	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("function take %v to complete \n", elapsed)
	}()

	account := c.Query("account")
	syncMode := c.Query("sync")
	if account == "" {
		content, err := json.Marshal(false)
		if err != nil {
			log.Fatal(err)
		}
		Response(string(content), 200, c)
	}

	db := postgres.Connect()
	defer db.Close()
	var user Users
	if err := db.Model(&user).Where("wallet_address = ?", account).Select(); err != nil {
		log.Println("Query Users By Wallet Address error ", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var userSub []UsersSubscriptions
	if err := db.Model(&userSub).
		Where("users_id = ?", user.ID).
		Where("is_active = ?", true).
		Select(); err != nil {
		log.Println("usersSubscription")
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var now = time.Now().UTC()
	log.Printf("now : %s", now)
	if len(userSub) > 0 {
		for _, u := range userSub {
			if now.Before(u.EndDate) {
				log.Printf("u.EndDate : %s", u.EndDate)
				log.Printf("!!!!! Valid subscription")
				results := response{
					Level:       "SUBSCRIPTION",
					MsczBalance: "0",
					SczBalance:  "0",
					Rate:        "0",
					Value:       "0",
				}

				content, err := json.Marshal(results)
				if err != nil {
					log.Println(err.Error())
					Response(err.Error(), http.StatusInternalServerError, c)
				}
				c.String(http.StatusOK, string(content))
			} else {
				updateIsPremium(account, "FREE", c)
				// if err != nil {
				// 	log.Println("updateIsPremium error : ", err.Error())
				// 	c.String(http.StatusInternalServerError, err.Error())
				// }

				u.IsActive = false
				_, err := db.Model(&u).WherePK().Update()
				if err != nil {
					log.Println("User subscription Update Is active error : ", err.Error())
					c.String(http.StatusInternalServerError, err.Error())
				}
			}
		}
	}

	userAddress := common.HexToAddress(account)
	sczAddress := common.HexToAddress(SCZ)
	contract, err := bep20.NewBep20(sczAddress, globalClient)
	if err != nil {
		log.Println(err.Error())
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	mSCZAddress := common.HexToAddress(mSCZ)
	msczContract, err := bep20.NewBep20(mSCZAddress, globalClient)
	if err != nil {
		log.Println(err.Error())
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	var msczBalance, sczBalance *big.Int
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		msczBalance, _ = msczContract.BalanceOf(&bind.CallOpts{}, userAddress)
		if err != nil {
			log.Fatal(err.Error())

		}
		elapsed := time.Since(start)
		fmt.Printf("msczContract function take %v to complete \n", elapsed)
		waitgroup.Done()

	}()

	go func() {
		sczBalance, err = contract.BalanceOf(&bind.CallOpts{}, userAddress)
		if err != nil {
			log.Fatal(err.Error())
		}
		waitgroup.Done()
		elapsed := time.Since(start)
		fmt.Printf("sczBalance function take %v to complete \n", elapsed)
	}()

	waitgroup.Wait()

	balance := new(big.Int).Div(sczBalance, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))
	msczBalance = new(big.Int).Div(msczBalance, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0)))

	rateUSD := GetRateSCZ()
	value := float64(balance.Uint64()) * rateUSD

	expireDate, _ := time.Parse("2006/01/02 15:04:05", "2022/10/31 23:59:59")
	today, _ := time.Parse("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:04:05"))

	results := response{}
	results.Level = "PRO"
	results.MsczBalance = strconv.FormatFloat(float64(msczBalance.Uint64()), 'f', -1, 64)
	results.SczBalance = strconv.FormatFloat(float64(balance.Uint64()), 'f', -1, 64)
	results.Rate = strconv.FormatFloat(rateUSD, 'f', -1, 64)
	results.Value = strconv.FormatFloat(value, 'f', -1, 64)

	if (msczBalance.Cmp(big.NewInt(10)) >= 0) && (today.Before(expireDate) || today.Equal(expireDate)) {
		updateIsPremium(account, "PRO", c)
		content, err := json.Marshal(results)
		if err != nil {
			log.Println(err.Error())
			Response(err.Error(), http.StatusInternalServerError, c)
		}
		Response(string(content), 200, c)
	}

	if syncMode == "true" {
		rateUSD = getRateUSD(SCZ)
	}
	value = float64(balance.Uint64()) * rateUSD

	var proLevel = "PRO"
	if value < proValueMinimum {
		waitgroup.Add(1)
		go func() {
			sczAmount := checkSCZLPv2(userAddress, SCZ_BUSD)
			value = value + (sczAmount * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczAmountDBM := checkSCZLPv2(userAddress, SCZ_DBM)
			value = value + (sczAmountDBM * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczAmountStakeDBM := checkSCZStakeV2(userAddress, chefMoon, big.NewInt(5))
			value = value + (sczAmountStakeDBM * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczAmountBnb := checkSCZLPv2(userAddress, SCZ_BNB)
			value = value + (sczAmountBnb * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczStakeFoodCourt := checkSCZStakeV2(userAddress, chefFoodCourt, big.NewInt(45))
			value = value + (sczStakeFoodCourt * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczAmountEve := checkSCZLPv2(userAddress, SCZ_EVE)
			value = value + (sczAmountEve * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Add(1)
		go func() {
			sczStakeNeverSwap := checkSCZStakeV2(userAddress, chefNeverSwap, big.NewInt(14))
			value = value + (sczStakeNeverSwap * rateUSD)
			waitgroup.Done()
		}()

		waitgroup.Wait()

		if value < proValueMinimum {
			if value < proValueMinimum {
				proLevel = "FREE"
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("pro check function take %v to complete \n", elapsed)

	updateIsPremium(account, proLevel, c)
	results.Level = proLevel
	results.Value = strconv.FormatFloat(value, 'f', -1, 64)
	results.Rate = strconv.FormatFloat(rateUSD, 'f', -1, 64)
	content, err := json.Marshal(results)
	if err != nil {
		Response(err.Error(), 500, c)
	}
	Response(string(content), 200, c)
}

func GetRateSCZ() float64 {
	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println(err.Error())
	}

	var tokenPrice tokenPrice
	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("token_prices")
	if err != nil {
		log.Println("token_prices Not Found error ", err.Error())
	}

	err = collection.FindOne(mongoDB.Ctx, bson.M{"token": SCZ}).Decode(&tokenPrice)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println("SCZ Not Found error ", err.Error())
	}

	rateUSD, err := strconv.ParseFloat(tokenPrice.Price, 64)
	if err != nil {
		log.Println("rateUSD error ", err.Error())
	}

	return rateUSD
}

func checkSCZLP(userAddress common.Address, lpTokenAddress string) float64 {
	// client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	contractAddress := common.HexToAddress(lpTokenAddress)
	rate := getRateUSD(SCZ)

	pairInstance, err := pancakepair.NewPancakepair(contractAddress, globalClient)
	if err != nil {
		log.Fatal(err)
	}

	_balance, err := pairInstance.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		log.Fatal(err)
	}

	reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	_totalSupply, err := pairInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	token0, err := pairInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	result := big.NewInt(0)
	decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0))
	if token0 == common.HexToAddress(SCZ) {
		return float64(result.Mul(_balance, reserve.Reserve0).Div(result, _totalSupply).Div(result, decimal).Uint64()) * rate
	}
	return float64(result.Mul(_balance, reserve.Reserve1).Div(result, _totalSupply).Div(result, decimal).Uint64()) * rate
}

func checkSCZLPv2(userAddress common.Address, lpTokenAddress string) float64 {

	contractAddress := common.HexToAddress(lpTokenAddress)

	pairInstance, err := pancakepair.NewPancakepair(contractAddress, globalClient)
	if err != nil {
		log.Fatal(err)
	}

	var _balance, _reserve0, _reserve1, _totalSupply *big.Int
	var waitgroup sync.WaitGroup
	waitgroup.Add(4)

	go func() {
		_balance, err = pairInstance.BalanceOf(&bind.CallOpts{}, userAddress)
		if err != nil {
			log.Fatal(err)
		}
		waitgroup.Done()
	}()

	go func() {
		reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		_reserve0 = reserve.Reserve0
		_reserve1 = reserve.Reserve1
		waitgroup.Done()
	}()

	go func() {
		_totalSupply, err = pairInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		waitgroup.Done()
	}()

	var token0 common.Address
	go func() {
		token0, err = pairInstance.Token0(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		waitgroup.Done()
	}()

	waitgroup.Wait()

	result := big.NewInt(0)
	decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0))
	if token0 == common.HexToAddress(SCZ) {
		return float64(result.Mul(_balance, _reserve0).Div(result, _totalSupply).Div(result, decimal).Uint64())
	}
	return float64(result.Mul(_balance, _reserve1).Div(result, _totalSupply).Div(result, decimal).Uint64())
}

func checkSCZStake(userAddress common.Address, chefAddress string, poolId *big.Int) float64 {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	rate := getRateUSD(SCZ)
	chefContract := common.HexToAddress(chefAddress)
	masterchefInstance, err := masterchef.NewMasterchef(chefContract, client)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := masterchefInstance.UserInfo(&bind.CallOpts{}, poolId, userAddress)
	if err != nil {
		log.Fatal(err)
	}

	_balance := userInfo.Amount
	poolInfo, err := masterchefInstance.PoolInfo(&bind.CallOpts{}, poolId)
	pairInstance, err := pancakepair.NewPancakepair(poolInfo.LpToken, client)
	if err != nil {
		log.Fatal(err)
	}

	reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	_totalSupply, err := pairInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	token0, err := pairInstance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	result := big.NewInt(0)
	decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0))
	if token0 == common.HexToAddress(SCZ) {
		return float64(result.Mul(_balance, reserve.Reserve0).Div(result, _totalSupply).Div(result, decimal).Uint64()) * rate
	}
	return float64(result.Mul(_balance, reserve.Reserve1).Div(result, _totalSupply).Div(result, decimal).Uint64()) * rate
}

func checkSCZStakeV2(userAddress common.Address, chefAddress string, poolId *big.Int) float64 {

	chefContract := common.HexToAddress(chefAddress)
	masterchefInstance, err := masterchef.NewMasterchef(chefContract, globalClient)
	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := masterchefInstance.UserInfo(&bind.CallOpts{}, poolId, userAddress)
	if err != nil {
		log.Fatal(err)
	}

	var _balance, _reserve0, _reserve1, _totalSupply *big.Int

	_balance = userInfo.Amount
	poolInfo, _ := masterchefInstance.PoolInfo(&bind.CallOpts{}, poolId)
	pairInstance, err := pancakepair.NewPancakepair(poolInfo.LpToken, globalClient)
	if err != nil {
		log.Fatal(err)
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(3)

	go func() {
		reserve, err := pairInstance.GetReserves(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		_reserve0 = reserve.Reserve0
		_reserve1 = reserve.Reserve1
		waitgroup.Done()
	}()

	go func() {
		_totalSupply, err = pairInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		waitgroup.Done()
	}()

	var token0 common.Address
	go func() {
		token0, err = pairInstance.Token0(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		waitgroup.Done()
	}()

	waitgroup.Wait()

	result := big.NewInt(0)
	decimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), big.NewInt(0))
	if token0 == common.HexToAddress(SCZ) {
		return float64(result.Mul(_balance, _reserve0).Div(result, _totalSupply).Div(result, decimal).Uint64())
	}
	return float64(result.Mul(_balance, _reserve1).Div(result, _totalSupply).Div(result, decimal).Uint64())
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		getProAccount(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
