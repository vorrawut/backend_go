// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venuscomptroller

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VenuscomptrollerABI is the input ABI used to generate the binding from.
const VenuscomptrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"ActionProtocolPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusVAIMintIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIMinterVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIVaultVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"oldVAIController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"newVAIController\",\"type\":\"address\"}],\"name\":\"NewVAIController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVAIMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseInterval_\",\"type\":\"uint256\"}],\"name\":\"NewVAIVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIVaultRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIVaultRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIVaultRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"VenusSpeedUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"vaiController_\",\"type\":\"address\"}],\"name\":\"_setVAIController\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"_setVAIMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReleaseAmount_\",\"type\":\"uint256\"}],\"name\":\"_setVAIVaultInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"venusSpeed\",\"type\":\"uint256\"}],\"name\":\"_setVenusSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIVaultRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIVaultRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"}],\"name\":\"distributeVAIMinterVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VenuscomptrollerBin is the compiled bytecode used for deploying new contracts.
var VenuscomptrollerBin = "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061a7f480620000626000396000f3fe608060405234801561001057600080fd5b50600436106104f85760003560e01c80637d172bd511610299578063bf32442d11610167578063ddfd287e116100d9578063eabe7d9111610092578063eabe7d911461249f578063ede4edd014612521578063f4b8d5fe14612579578063f851a440146125bd578063fa6331d814612607578063fd51a3ad14612625576104f8565b8063ddfd287e1461231e578063e37d4b7914612328578063e4028eee146123cf578063e6653f3d14612431578063e6de652814612453578063e875544614612481576104f8565b8063d02f73511161012b578063d02f735114611f1c578063d09c54ba14611fde578063d24febad14612142578063da3d454c146121c4578063dce1544914612246578063dcfbc0c7146122d4576104f8565b8063bf32442d14611cd7578063c299823814611d21578063c488847b14611def578063c5b4db5514611e78578063c5f956af14611ed2576104f8565b8063a06a87f11161020b578063b0772d0b116101c4578063b0772d0b14611a7d578063b2eafc3914611adc578063b8324c7c14611b26578063bb82aa5e14611bcd578063bdcdc25814611c17578063bec04f7214611cb9576104f8565b8063a06a87f11461186f578063a76b3fda146118bd578063a78dc77514611915578063abfceffc1461197e578063ac0b0bb714611a17578063adcd5fb914611a39576104f8565b80638a7dc1651161025d5780638a7dc1651461166d5780638e8f294b146116c55780639254f5e514611733578063929fe9a11461177d57806394b2294b146117f95780639cfdd9e614611817576104f8565b80637d172bd5146114c15780637dc0d1d01461150b57806386df31ee14611555578063879c2e1d1461162d57806387f763031461164b576104f8565b806342cbb15c116103d65780635c778605116103485780636a56947e116103015780636a56947e1461128d5780636d154ea51461131b5780636d35bf9114611377578063719f701b14611425578063731f0c2b14611443578063765513831461149f576104f8565b80635c77860514610fa35780635ec88c79146110115780635f5af1aa146110775780635fc7e71e146110cf578063607ef6c1146111915780636662c7c91461125f576104f8565b80634e79238f1161039a5780634e79238f14610d075780634ef4c3e114610da15780634fd42e1714610e2357806351dff98914610e6557806352d84d1e14610edd57806355ee1fe114610f4b576104f8565b806342cbb15c14610b6357806347ef3b3b14610b815780634a58443214610c395780634ada90af14610c915780634e0853db14610caf576104f8565b80632a6a60651161046f578063399cc80c11610433578063399cc80c146109ef5780633c94786f14610a0d5780634088c73e14610a2f57806341a18d2c14610a5157806341c728b914610ac9578063425fad5814610b41576104f8565b80632a6a6065146108875780632bc7e29e146108cf5780632ec0412414610927578063317b0b7714610969578063391957d7146109ab576104f8565b80631d504dc6116104c15780631d504dc61461062b5780631ededc911461066f57806321af45691461070757806324008a621461075157806324a3d622146107f3578063267822471461083d576104f8565b80627e3dd2146104fd57806304ef9d581461051f57806308e0225c1461053d5780630db4b4e5146105b55780631abcaa77146105d3575b600080fd5b610505612687565b604051808215151515815260200191505060405180910390f35b61052761268c565b6040518082815260200191505060405180910390f35b61059f6004803603604081101561055357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612692565b6040518082815260200191505060405180910390f35b6105bd6126b7565b6040518082815260200191505060405180910390f35b610615600480360360208110156105e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506126bd565b6040518082815260200191505060405180910390f35b61066d6004803603602081101561064157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506126d5565b005b610705600480360360a081101561068557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506128f1565b005b61070f612909565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6107dd6004803603608081101561076757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061292f565b6040518082815260200191505060405180910390f35b6107fb612ae3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610845612b09565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6108b56004803603602081101561089d57600080fd5b81019080803515159060200190929190505050612b2f565b604051808215151515815260200191505060405180910390f35b610911600480360360208110156108e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612d5e565b6040518082815260200191505060405180910390f35b6109536004803603602081101561093d57600080fd5b8101908080359060200190929190505050612d76565b6040518082815260200191505060405180910390f35b6109956004803603602081101561097f57600080fd5b8101908080359060200190929190505050612e42565b6040518082815260200191505060405180910390f35b6109ed600480360360208110156109c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612f68565b005b6109f761312d565b6040518082815260200191505060405180910390f35b610a15613133565b604051808215151515815260200191505060405180910390f35b610a37613146565b604051808215151515815260200191505060405180910390f35b610ab360048036036040811015610a6757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613159565b6040518082815260200191505060405180910390f35b610b3f60048036036080811015610adf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061317e565b005b610b49613184565b604051808215151515815260200191505060405180910390f35b610b6b613197565b6040518082815260200191505060405180910390f35b610c37600480360360c0811015610b9757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061319f565b005b610c7b60048036036020811015610c4f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506131b8565b6040518082815260200191505060405180910390f35b610c996131d0565b6040518082815260200191505060405180910390f35b610d0560048036036060811015610cc557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506131d6565b005b610d7d60048036036080811015610d1d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061335f565b60405180848152602001838152602001828152602001935050505060405180910390f35b610e0d60048036036060811015610db757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061339b565b6040518082815260200191505060405180910390f35b610e4f60048036036020811015610e3957600080fd5b8101908080359060200190929190505050613572565b6040518082815260200191505060405180910390f35b610edb60048036036080811015610e7b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061363e565b005b610f0960048036036020811015610ef357600080fd5b81019080803590602001909291905050506136c6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610f8d60048036036020811015610f6157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613702565b6040518082815260200191505060405180910390f35b61100f60048036036060811015610fb957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613880565b005b6110536004803603602081101561102757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613896565b60405180848152602001838152602001828152602001935050505060405180910390f35b6110b96004803603602081101561108d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506138d1565b6040518082815260200191505060405180910390f35b61117b600480360360a08110156110e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613a4f565b6040518082815260200191505060405180910390f35b61125d600480360360408110156111a757600080fd5b81019080803590602001906401000000008111156111c457600080fd5b8201836020820111156111d657600080fd5b803590602001918460208302840111640100000000831117156111f857600080fd5b90919293919293908035906020019064010000000081111561121957600080fd5b82018360208201111561122b57600080fd5b8035906020019184602083028401116401000000008311171561124d57600080fd5b9091929391929390505050613e0a565b005b61128b6004803603602081101561127557600080fd5b81019080803590602001909291905050506140c3565b005b611319600480360360808110156112a357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506141d6565b005b61135d6004803603602081101561133157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506141ed565b604051808215151515815260200191505060405180910390f35b611423600480360360a081101561138d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061420d565b005b61142d614225565b6040518082815260200191505060405180910390f35b6114856004803603602081101561145957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061422b565b604051808215151515815260200191505060405180910390f35b6114a761424b565b604051808215151515815260200191505060405180910390f35b6114c961425e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611513614284565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61162b6004803603604081101561156b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156115a857600080fd5b8201836020820111156115ba57600080fd5b803590602001918460208302840111640100000000831117156115dc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506142aa565b005b611635614339565b6040518082815260200191505060405180910390f35b61165361433f565b604051808215151515815260200191505060405180910390f35b6116af6004803603602081101561168357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614352565b6040518082815260200191505060405180910390f35b611707600480360360208110156116db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061436a565b604051808415151515815260200183815260200182151515158152602001935050505060405180910390f35b61173b6143ae565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6117df6004803603604081101561179357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506143d4565b604051808215151515815260200191505060405180910390f35b61180161446b565b6040518082815260200191505060405180910390f35b6118596004803603602081101561182d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614471565b6040518082815260200191505060405180910390f35b6118bb6004803603604081101561188557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506145e0565b005b6118ff600480360360208110156118d357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614668565b6040518082815260200191505060405180910390f35b6119616004803603604081101561192b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506148ea565b604051808381526020018281526020019250505060405180910390f35b6119c06004803603602081101561199457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614b28565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611a035780820151818401526020810190506119e8565b505050509050019250505060405180910390f35b611a1f614bf5565b604051808215151515815260200191505060405180910390f35b611a7b60048036036020811015611a4f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614c08565b005b611a85614c9b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611ac8578082015181840152602081019050611aad565b505050509050019250505060405180910390f35b611ae4614d29565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611b6860048036036020811015611b3c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050614d4f565b60405180837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681526020018263ffffffff1663ffffffff1681526020019250505060405180910390f35b611bd5614dab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611ca360048036036080811015611c2d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050614dd1565b6040518082815260200191505060405180910390f35b611cc1614f3a565b6040518082815260200191505060405180910390f35b611cdf614f40565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611d9860048036036020811015611d3757600080fd5b8101908080359060200190640100000000811115611d5457600080fd5b820183602082011115611d6657600080fd5b80359060200191846020830284011164010000000083111715611d8857600080fd5b9091929391929390505050614f5c565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611ddb578082015181840152602081019050611dc0565b505050509050019250505060405180910390f35b611e5b60048036036060811015611e0557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050615016565b604051808381526020018281526020019250505060405180910390f35b611e80615331565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611eda615344565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b611fc8600480360360a0811015611f3257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061536a565b6040518082815260200191505060405180910390f35b61214060048036036080811015611ff457600080fd5b810190808035906020019064010000000081111561201157600080fd5b82018360208201111561202357600080fd5b8035906020019184602083028401116401000000008311171561204557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156120a557600080fd5b8201836020820111156120b757600080fd5b803590602001918460208302840111640100000000831117156120d957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803515159060200190929190803515159060200190929190505050615709565b005b6121ae6004803603606081101561215857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050615cb8565b6040518082815260200191505060405180910390f35b612230600480360360608110156121da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061605b565b6040518082815260200191505060405180910390f35b6122926004803603604081101561225c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506166fa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6122dc616745565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61232661676b565b005b61236a6004803603602081101561233e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050616a69565b60405180837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681526020018263ffffffff1663ffffffff1681526020019250505060405180910390f35b61241b600480360360408110156123e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050616ac5565b6040518082815260200191505060405180910390f35b612439616d9e565b604051808215151515815260200191505060405180910390f35b61247f6004803603602081101561246957600080fd5b8101908080359060200190929190505050616db1565b005b612489616ec4565b6040518082815260200191505060405180910390f35b61250b600480360360608110156124b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050616eca565b6040518082815260200191505060405180910390f35b6125636004803603602081101561253757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050616fa5565b6040518082815260200191505060405180910390f35b6125bb6004803603602081101561258f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050617487565b005b6125c5617703565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61260f617728565b6040518082815260200191505060405180910390f35b6126716004803603604081101561263b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061772e565b6040518082815260200191505060405180910390f35b600181565b60225481565b6013602052816000526040600020602052806000526040600020600091509150505481565b601d5481565b600f6020528060005260406000206000915090505481565b8073ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b815260040160206040518083038186803b15801561271b57600080fd5b505afa15801561272f573d6000803e3d6000fd5b505050506040513d602081101561274557600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146127f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6f6e6c7920756e6974726f6c6c65722061646d696e2063616e0000000000000081525060200191505060405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff1663c1e803346040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561284057600080fd5b505af1158015612854573d6000803e3d6000fd5b505050506040513d602081101561286a57600080fd5b8101908080519060200190929190505050146128ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6e6f7420617574686f72697a656400000000000000000000000000000000000081525060200191505060405180910390fd5b50565b600015612902576007546007819055505b5050505050565b601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000601860029054906101000a900460ff16156129b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16612a1c5760096012811115612a1557fe5b9050612adb565b612a2461a674565b60405180602001604052808773ffffffffffffffffffffffffffffffffffffffff1663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b158015612a7557600080fd5b505afa158015612a89573d6000803e3d6000fd5b505050506040513d6020811015612a9f57600080fd5b81019080805190602001909291905050508152509050612abf8682617911565b612aca868583617d11565b60006012811115612ad757fe5b9150505b949350505050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612bda57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b612c2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061a76a6021913960400191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612c8e575060011515811515145b612d00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6f6e6c792061646d696e2063616e20756e70617573650000000000000000000081525060200191505060405180910390fd5b82601860026101000a81548160ff0219169083151502179055507fd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b3883604051808215151515815260200191505060405180910390a182915050919050565b60166020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612ddf57612dd860016014618133565b9050612e3d565b60006017549050826017819055507f73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a48184604051808381526020018281526020019250505060405180910390a160006012811115612e3957fe5b9150505b919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612f06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6f6e6c792061646d696e2063616e2073657420636c6f736520666163746f720081525060200191505060405180910390fd5b60006005549050826005819055507f3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd98184604051808381526020018281526020019250505060405180910390a160006012811115612f6057fe5b915050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461302a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6f6e6c792061646d696e2063616e00000000000000000000000000000000000081525060200191505060405180910390fd5b6000601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081601e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507feda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e298183604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a15050565b60195481565b600a60149054906101000a900460ff1681565b601860009054906101000a900460ff1681565b6012602052816000526040600020602052806000526040600020600091509150505481565b50505050565b601860029054906101000a900460ff1681565b600043905090565b6000156131b0576007546007819055505b505050505050565b601f6020528060005260406000206000915090505481565b60065481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613298576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6f6e6c792061646d696e2063616e00000000000000000000000000000000000081525060200191505060405180910390fd5b82601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081601c8190555080601d819055507f7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a1505050565b6000806000806000806133748a8a8a8a6181a7565b92509250925082601281111561338657fe5b82829550955095505050509450945094915050565b6000601860029054906101000a900460ff1615613420576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156134e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6d696e742069732070617573656400000000000000000000000000000000000081525060200191505060405180910390fd5b600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16613548576009601281111561354157fe5b905061356b565b613551846186ec565b61355b8484618ae2565b6000601281111561356857fe5b90505b9392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146135db576135d46001600b618133565b9050613639565b60006006549050826006819055507faeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec13168184604051808381526020018281526020019250505060405180910390a16000601281111561363557fe5b9150505b919050565b60008114158061364e5750600082145b6136c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f72656465656d546f6b656e73207a65726f00000000000000000000000000000081525060200191505060405180910390fd5b50505050565b600d81815481106136d357fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461376b5761376460016010618133565b905061387b565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e228184604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a16000601281111561387757fe5b9150505b919050565b600015613891576007546007819055505b505050565b6000806000806000806138ad8760008060006181a7565b9250925092508260128111156138bf57fe5b82829550955095505050509193909250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461393a5761393360016013618133565b9050613a4a565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e8184604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a160006012811115613a4657fe5b9150505b919050565b6000601860029054906101000a900460ff1615613ad4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1680613b7c5750601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16145b1580613bd55750600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16155b15613bee5760096012811115613be757fe5b9050613e01565b600080613bff8560008060006181a7565b925050915060006012811115613c1157fe5b826012811115613c1d57fe5b14613c3757816012811115613c2e57fe5b92505050613e01565b6000811415613c565760036012811115613c4d57fe5b92505050613e01565b6000601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614613d6c578873ffffffffffffffffffffffffffffffffffffffff166395dd9193876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015613d2a57600080fd5b505afa158015613d3e573d6000803e3d6000fd5b505050506040513d6020811015613d5457600080fd5b81019080805190602001909291905050509050613daf565b601660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b6000613dcb604051806020016040528060055481525083618f41565b905080861115613ded5760116012811115613de257fe5b945050505050613e01565b60006012811115613dfa57fe5b9450505050505b95945050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480613eb25750601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b613f07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603581526020018061a78b6035913960400191505060405180910390fd5b6000848490509050600083839050905060008214158015613f2757508082145b613f99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f696e76616c696420696e7075740000000000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b828110156140ba57848482818110613fb357fe5b90506020020135601f6000898985818110613fca57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555086868281811061403157fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f686868481811061409157fe5b905060200201356040518082815260200191505060405180910390a28080600101915050613f9f565b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6f6e6c792061646d696e2063616e00000000000000000000000000000000000081525060200191505060405180910390fd5b6000601a54905081601a819055507fe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b8183604051808381526020018281526020019250505060405180910390a15050565b6000156141e7576007546007819055505b50505050565b600c6020528060005260406000206000915054906101000a900460ff1681565b60001561421e576007546007819055505b5050505050565b601c5481565b600b6020528060005260406000206000915054906101000a900460ff1681565b601860019054906101000a900460ff1681565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060016040519080825280602002602001820160405280156142dc5781602001602082028038833980820191505090505b50905082816000815181106142ed57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506143348183600180615709565b505050565b600e5481565b600a60169054906101000a900460ff1681565b60146020528060005260406000206000915090505481565b60096020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154908060030160009054906101000a900460ff16905083565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60075481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146144da576144d360016015618133565b90506145db565b6000601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b28184604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505b919050565b6145e8618f69565b61465a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c792061646d696e2063616e207365742076656e7573207370656564000081525060200191505060405180910390fd5b6146648282619018565b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146146d1576146ca60016012618133565b90506148e5565b600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff161561473957614732600a6011618133565b90506148e5565b8173ffffffffffffffffffffffffffffffffffffffff16633d9ea3a16040518163ffffffff1660e01b815260040160206040518083038186803b15801561477f57600080fd5b505afa158015614793573d6000803e3d6000fd5b505050506040513d60208110156147a957600080fd5b81019080805190602001909291905050505060405180606001604052806001151581526020016000815260200160001515815250600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160030160006101000a81548160ff02191690831515021790555090505061487282619714565b7fcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600060128111156148e257fe5b90505b919050565b6000806000670de0b6b3a764000090506000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561499b57600080fd5b505afa1580156149af573d6000803e3d6000fd5b505050506040513d60208110156149c557600080fd5b8101908080519060200190929190505050905060008114156149fe57600d60128111156149ee57fe5b6000809050935093505050614b21565b60008673ffffffffffffffffffffffffffffffffffffffff1663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b158015614a4657600080fd5b505afa158015614a5a573d6000803e3d6000fd5b505050506040513d6020811015614a7057600080fd5b810190808051906020019092919050505090506000614a8d61a674565b614a9561a674565b614a9d61a674565b614ac5604051806020016040528060065481525060405180602001604052808a815250619877565b9250614aed604051806020016040528088815250604051806020016040528088815250619877565b9150614af983836198b8565b9050614b05818b618f41565b935060006012811115614b1457fe5b8498509850505050505050505b9250929050565b6060600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015614be957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311614b9f575b50505050509050919050565b600a60179054906101000a900460ff1681565b614c9881600d805480602002602001604051908101604052809291908181526020018280548015614c8e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311614c44575b50505050506142aa565b50565b6060600d805480602002602001604051908101604052809291908181526020018280548015614d1f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311614cd5575b5050505050905090565b602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60106020528060005260406000206000915090508060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169080600001601c9054906101000a900463ffffffff16905082565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000601860029054906101000a900460ff1615614e56576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600a60169054906101000a900460ff1615614ed9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f7472616e7366657220697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b6000614ee68686856198f9565b905060006012811115614ef557fe5b8114614f045780915050614f32565b614f0d866186ec565b614f178686618ae2565b614f218685618ae2565b60006012811115614f2e57fe5b9150505b949350505050565b60175481565b600073cf6bb5389c92bdda8a3747ddb454cb7a64626c63905090565b60606000838390509050606081604051908082528060200260200182016040528015614f975781602001602082028038833980820191505090505b50905060008090505b8281101561500a57614fda868683818110614fb757fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1633619a87565b6012811115614fe557fe5b828281518110614ff157fe5b6020026020010181815250508080600101915050614fa0565b50809250505092915050565b6000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156150ba57600080fd5b505afa1580156150ce573d6000803e3d6000fd5b505050506040513d60208110156150e457600080fd5b810190808051906020019092919050505090506000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561519857600080fd5b505afa1580156151ac573d6000803e3d6000fd5b505050506040513d60208110156151c257600080fd5b8101908080519060200190929190505050905060008214806151e45750600081145b1561520657600d60128111156151f657fe5b6000809050935093505050615329565b60008673ffffffffffffffffffffffffffffffffffffffff1663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b15801561524e57600080fd5b505afa158015615262573d6000803e3d6000fd5b505050506040513d602081101561527857600080fd5b81019080805190602001909291905050509050600061529561a674565b61529d61a674565b6152a561a674565b6152cd604051806020016040528060065481525060405180602001604052808a815250619877565b92506152f5604051806020016040528088815250604051806020016040528088815250619877565b915061530183836198b8565b905061530d818b618f41565b93506000601281111561531c57fe5b8498509850505050505050505b935093915050565b6ec097ce7bc90715b34b9f100000000081565b602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000601860029054906101000a900460ff16156153ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600a60179054906101000a900460ff1615615472576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f7365697a6520697320706175736564000000000000000000000000000000000081525060200191505060405180910390fd5b600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1615806155735750600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16806155715750601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b155b1561558c576009601281111561558557fe5b9050615700565b8473ffffffffffffffffffffffffffffffffffffffff16635fe3b5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156155d257600080fd5b505afa1580156155e6573d6000803e3d6000fd5b505050506040513d60208110156155fc57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16635fe3b5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561566957600080fd5b505afa15801561567d573d6000803e3d6000fd5b505050506040513d602081101561569357600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16146156d357600260128111156156cc57fe5b9050615700565b6156dc866186ec565b6156e68684618ae2565b6156f08685618ae2565b600060128111156156fd57fe5b90505b95945050505050565b60008073ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461580757601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d244b0816040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156157ca57600080fd5b505af11580156157de573d6000803e3d6000fd5b505050506040513d60208110156157f457600080fd5b8101908080519060200190929190505050505b600090505b84518110156159035761583185828151811061582457fe5b6020026020010151617487565b6158a085828151811061584057fe5b60200260200101516014600088858151811061585857fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054619ce9565b601460008784815181106158b057fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808060010191505061580c565b60008090505b8451811015615cb057600085828151811061592057fe5b60200260200101519050600960008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff166159ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f6e6f74206c6973746564206d61726b657400000000000000000000000000000081525060200191505060405180910390fd5b8415615b95576159fa61a674565b60405180602001604052808373ffffffffffffffffffffffffffffffffffffffff1663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b158015615a4b57600080fd5b505afa158015615a5f573d6000803e3d6000fd5b505050506040513d6020811015615a7557600080fd5b81019080805190602001909291905050508152509050615a958282617911565b600093505b8751841015615b9357615ac182898681518110615ab357fe5b602002602001015183617d11565b615b30888581518110615ad057fe5b6020026020010151601460008b8881518110615ae857fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054619ce9565b601460008a8781518110615b4057fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508380600101945050615a9a565b505b8315615ca257615ba4816186ec565b600092505b8651831015615ca157615bcf81888581518110615bc257fe5b6020026020010151618ae2565b615c3e878481518110615bde57fe5b6020026020010151601460008a8781518110615bf657fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054619ce9565b60146000898681518110615c4e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508280600101935050615ba9565b5b508080600101915050615909565b505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480615d625750602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b615d7957615d7260016017618133565b9050616054565b670de0b6b3a76400008210615df6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f74726561737572792070657263656e7420636170206f766572666c6f7700000081525060200191505060405180910390fd5b6000602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000602254905086602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846022819055507f29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e88388604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a17f8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac8287604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a17f0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed78186604051808381526020018281526020019250505060405180910390a16000601281111561604e57fe5b93505050505b9392505050565b6000601860029054906101000a900460ff16156160e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156161a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f626f72726f77206973207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16616208576009601281111561620157fe5b90506166f3565b600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661637c578373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461633a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73656e646572206d7573742062652076546f6b656e000000000000000000000081525060200191505060405180910390fd5b60006163468585619a87565b90506000601281111561635557fe5b81601281111561636157fe5b1461637a5780601281111561637257fe5b9150506166f3565b505b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561641d57600080fd5b505afa158015616431573d6000803e3d6000fd5b505050506040513d602081101561644757600080fd5b8101908080519060200190929190505050141561647257600d601281111561646b57fe5b90506166f3565b6000601f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081146165c95760008573ffffffffffffffffffffffffffffffffffffffff166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b15801561650657600080fd5b505afa15801561651a573d6000803e3d6000fd5b505050506040513d602081101561653057600080fd5b81019080805190602001909291905050509050600061654f8286619ea1565b90508281106165c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6d61726b657420626f72726f772063617020726561636865640000000000000081525060200191505060405180910390fd5b50505b6000806165d986886000886181a7565b9250509150600060128111156165eb57fe5b8260128111156165f757fe5b146166125781601281111561660857fe5b93505050506166f3565b60008114616631576004601281111561662757fe5b93505050506166f3565b61663961a674565b60405180602001604052808973ffffffffffffffffffffffffffffffffffffffff1663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561668a57600080fd5b505afa15801561669e573d6000803e3d6000fd5b505050506040513d60208110156166b457600080fd5b810190808051906020019092919050505081525090506166d48882617911565b6166df888883617d11565b600060128111156166ec57fe5b9450505050505b9392505050565b6008602052816000526040600020818154811061671357fe5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000601c5414806167845750601c54616782613197565b105b1561678e57616a67565b6000616798614f40565b905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561681957600080fd5b505afa15801561682d573d6000803e3d6000fd5b505050506040513d602081101561684357600080fd5b810190808051906020019092919050505090506000811415616866575050616a67565b60008061687c616874613197565b601c54619eeb565b9050600061688c601a5483619f35565b9050601d548110156168a2575050505050616a67565b8084106168b1578092506168b5565b8392505b6168bd613197565b601c819055508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561696c57600080fd5b505af1158015616980573d6000803e3d6000fd5b505050506040513d602081101561699657600080fd5b8101908080519060200190929190505050507ff6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d836040518082815260200191505060405180910390a1601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663faa1809e6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015616a4957600080fd5b505af1158015616a5d573d6000803e3d6000fd5b5050505050505050505b565b60116020528060005260406000206000915090508060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169080600001601c9054906101000a900463ffffffff16905082565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614616b2e57616b2760016006618133565b9050616d98565b6000600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900460ff16616b9b57616b9360096007618133565b915050616d98565b616ba361a674565b6040518060200160405280858152509050616bbc61a674565b6040518060200160405280670c7d713b49da00008152509050616bdf8183619f7f565b15616bfa57616bf060066008618133565b9350505050616d98565b60008514158015616ce457506000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df886040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015616ca757600080fd5b505afa158015616cbb573d6000803e3d6000fd5b505050506040513d6020811015616cd157600080fd5b8101908080519060200190929190505050145b15616cff57616cf5600d6009618133565b9350505050616d98565b6000836001015490508584600101819055507f70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5878288604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a160006012811115616d9157fe5b9450505050505b92915050565b600a60159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614616e73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f6f6e6c792061646d696e2063616e00000000000000000000000000000000000081525060200191505060405180910390fd5b60006019549050816019819055507f75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c198183604051808381526020018281526020019250505060405180910390a15050565b60055481565b6000601860029054906101000a900460ff1615616f4f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b6000616f5c8585856198f9565b905060006012811115616f6b57fe5b8114616f7a5780915050616f9e565b616f83856186ec565b616f8d8585618ae2565b60006012811115616f9a57fe5b9150505b9392505050565b60008082905060008060008373ffffffffffffffffffffffffffffffffffffffff1663c37f68e2336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060806040518083038186803b15801561702d57600080fd5b505afa158015617041573d6000803e3d6000fd5b505050506040513d608081101561705757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291905050505092509250925060008314617103576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6765744163636f756e74536e617073686f74206661696c65640000000000000081525060200191505060405180910390fd5b6000811461712257617117600c6002618133565b945050505050617482565b600061712f8733856198f9565b90506000811461715257617146600e600383619f94565b95505050505050617482565b6000600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661720257600060128111156171f557fe5b9650505050505050617482565b8060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008180549050905060005b818110156173c9578873ffffffffffffffffffffffffffffffffffffffff168382815481106172cd57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156173bc5782600183038154811061732357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683828154811061735a57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828054809190600190036173b6919061a687565b506173c9565b80806001019150506172a2565b8181106173d257fe5b7fe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d8933604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a16000601281111561747657fe5b99505050505050505050505b919050565b600073ffffffffffffffffffffffffffffffffffffffff16601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146174e6576174e561676b565b5b600073ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461770057600080600080601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d178273b866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050608060405180830381600087803b1580156175e357600080fd5b505af11580156175f7573d6000803e3d6000fd5b505050506040513d608081101561760d57600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809550819650829750839450505050506000601281111561765957fe5b8114156176fb5783601460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff167f2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f8484604051808381526020018281526020019250505060405180910390a25b505050505b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601a5481565b6000601860029054906101000a900460ff16156177b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f746f636f6c20697320706175736564000000000000000000000000000081525060200191505060405180910390fd5b601860009054906101000a900460ff161580156177dd5750601860019054906101000a900460ff16155b61784f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f564149206973207061757365640000000000000000000000000000000000000081525060200191505060405180910390fd5b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146178b7576178b0600e6016618133565b905061790b565b81601660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000601281111561790857fe5b90505b92915050565b6000601160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006179a2613197565b905060006179ca8285600001601c9054906101000a900463ffffffff1663ffffffff16619eeb565b90506000811180156179dc5750600083115b15617c9d576000617a6d8773ffffffffffffffffffffffffffffffffffffffff166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b158015617a2c57600080fd5b505afa158015617a40573d6000803e3d6000fd5b505050506040513d6020811015617a5657600080fd5b81019080805190602001909291905050508761a008565b90506000617a7b8386619f35565b9050617a8561a6b3565b60008311617aa25760405180602001604052806000815250617aad565b617aac828461a031565b5b9050617ab761a6b3565b617b1960405180602001604052808a60000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152508361a071565b90506040518060400160405280617b6983600001516040518060400160405280601381526020017f6e657720696e646578206f766572666c6f77730000000000000000000000000081525061a0a1565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152602001617bcb886040518060400160405280601681526020017f626c6f636b206e756d626572206f766572666c6f77730000000000000000000081525061a174565b63ffffffff16815250601160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550602082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505050505050617d09565b6000811115617d0857617ce5826040518060400160405280601681526020017f626c6f636b206e756d626572206f766572666c6f77730000000000000000000081525061a174565b84600001601c6101000a81548163ffffffff021916908363ffffffff1602179055505b5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff16601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614617d7057617d6f61676b565b5b6000601160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050617dbb61a6b3565b60405180602001604052808360000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152509050617e1e61a6b3565b6040518060200160405280601360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481525090508160000151601360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008160000151111561812b57617f4561a6b3565b617f4f838361a22f565b905060006180148873ffffffffffffffffffffffffffffffffffffffff166395dd9193896040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015617fd357600080fd5b505afa158015617fe7573d6000803e3d6000fd5b505050506040513d6020811015617ffd57600080fd5b81019080805190602001909291905050508761a008565b90506000618022828461a25f565b9050600061806f601460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483619ea1565b905080601460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167f837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb848960000151604051808381526020018281526020019250505060405180910390a3505050505b505050505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601281111561816257fe5b83601781111561816e57fe5b600060405180848152602001838152602001828152602001935050505060405180910390a182601281111561819f57fe5b905092915050565b60008060006181b461a6c6565b60006060600860008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561827757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161822d575b5050505050905060008090505b815181101561864057600082828151811061829b57fe5b602002602001015190508073ffffffffffffffffffffffffffffffffffffffff1663c37f68e28d6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060806040518083038186803b15801561832257600080fd5b505afa158015618336573d6000803e3d6000fd5b505050506040513d608081101561834c57600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919050505088604001896060018a60800183815250838152508381525083975050505050600084146183bd57600f60008081915080905097509750975050505050506186e2565b6040518060200160405280600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101548152508560c00181905250604051806020016040528086608001518152508560e00181905250600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc57d4df826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156184d057600080fd5b505afa1580156184e4573d6000803e3d6000fd5b505050506040513d60208110156184fa57600080fd5b81019080805190602001909291905050508560a001818152505060008560a00151141561853c57600d60008081915080905097509750975050505050506186e2565b60405180602001604052808660a0015181525085610100018190525061857861856d8660c001518760e00151619877565b866101000151619877565b8561012001819052506185998561012001518660400151876000015161a28f565b8560000181815250506185ba8561010001518660600151876020015161a28f565b8560200181815250508a73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156186325761860b8561012001518b876020015161a28f565b8560200181815250506186288561010001518a876020015161a28f565b8560200181815250505b508080600101915050618284565b5061868e8360200151601660008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054619ea1565b8360200181815250508260200151836000015111156186c7576000836020015184600001510360008090509550955095505050506186e2565b60008084600001518560200151038191509550955095505050505b9450945094915050565b6000601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600061877d613197565b905060006187a58285600001601c9054906101000a900463ffffffff1663ffffffff16619eeb565b90506000811180156187b75750600083115b15618a6f5760008573ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561880457600080fd5b505afa158015618818573d6000803e3d6000fd5b505050506040513d602081101561882e57600080fd5b81019080805190602001909291905050509050600061884d8386619f35565b905061885761a6b3565b60008311618874576040518060200160405280600081525061887f565b61887e828461a031565b5b905061888961a6b3565b6188eb60405180602001604052808a60000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152508361a071565b9050604051806040016040528061893b83600001516040518060400160405280601381526020017f6e657720696e646578206f766572666c6f77730000000000000000000000000081525061a0a1565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200161899d886040518060400160405280601681526020017f626c6f636b206e756d626572206f766572666c6f77730000000000000000000081525061a174565b63ffffffff16815250601060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550602082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505050505050618adb565b6000811115618ada57618ab7826040518060400160405280601681526020017f626c6f636b206e756d626572206f766572666c6f77730000000000000000000081525061a174565b84600001601c6101000a81548163ffffffff021916908363ffffffff1602179055505b5b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff16601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614618b4157618b4061676b565b5b6000601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050618b8c61a6b3565b60405180602001604052808360000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168152509050618bef61a6b3565b6040518060200160405280601260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481525090508160000151601260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008160000151148015618d19575060008260000151115b15618d56576ec097ce7bc90715b34b9f10000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168160000181815250505b618d5e61a6b3565b618d68838361a22f565b905060008673ffffffffffffffffffffffffffffffffffffffff166370a08231876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015618de957600080fd5b505afa158015618dfd573d6000803e3d6000fd5b505050506040513d6020811015618e1357600080fd5b810190808051906020019092919050505090506000618e32828461a25f565b90506000618e7f601460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483619ea1565b905080601460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167ffa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9848960000151604051808381526020018281526020019250505060405180910390a3505050505050505050565b6000618f4b61a674565b618f55848461a2c1565b9050618f608161a2ed565b91505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806190135750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b905090565b6000600f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081146191165761906c61a674565b60405180602001604052808573ffffffffffffffffffffffffffffffffffffffff1663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156190bd57600080fd5b505afa1580156190d1573d6000803e3d6000fd5b505050506040513d60208110156190e757600080fd5b81019080805190602001909291905050508152509050619106846186ec565b6191108482617911565b50619675565b60008214619674576000600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160000160009054906101000a900460ff161515146191ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f76656e7573206d61726b6574206973206e6f74206c697374656400000000000081525060200191505060405180910390fd5b6000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161480156192dc57506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c9054906101000a900463ffffffff1663ffffffff16145b1561942f5760405180604001604052806ec097ce7bc90715b34b9f10000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200161936561932a613197565b6040518060400160405280601c81526020017f626c6f636b206e756d626572206578636565647320333220626974730000000081525061a174565b63ffffffff16815250601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550602082015181600001601c6101000a81548163ffffffff021916908363ffffffff1602179055509050505b6000601160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1614801561951f57506000601160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c9054906101000a900463ffffffff1663ffffffff16145b156196725760405180604001604052806ec097ce7bc90715b34b9f10000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681526020016195a861956d613197565b6040518060400160405280601c81526020017f626c6f636b206e756d626572206578636565647320333220626974730000000081525061a174565b63ffffffff16815250601160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550602082015181600001601c6101000a81548163ffffffff021916908363ffffffff1602179055509050505b505b5b81811461970f5781600f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548836040518082815260200191505060405180910390a25b505050565b60008090505b600d8054905081101561980d578173ffffffffffffffffffffffffffffffffffffffff16600d828154811061974b57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415619800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6d61726b657420616c726561647920616464656400000000000000000000000081525060200191505060405180910390fd5b808060010191505061971a565b50600d8190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61987f61a674565b6040518060200160405280670de0b6b3a76400006198a586600001518660000151619f35565b816198ac57fe5b04815250905092915050565b6198c061a674565b60405180602001604052806198ee6198e48660000151670de0b6b3a7640000619f35565b856000015161a30c565b815250905092915050565b6000600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16619963576009601281111561995c57fe5b9050619a80565b600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16619a085760006012811115619a0157fe5b9050619a80565b600080619a1885878660006181a7565b925050915060006012811115619a2a57fe5b826012811115619a3657fe5b14619a5057816012811115619a4757fe5b92505050619a80565b60008114619a6e5760046012811115619a6557fe5b92505050619a80565b60006012811115619a7b57fe5b925050505b9392505050565b600080600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900460ff16619aeb576009915050619ce3565b8060020160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615619b49576000915050619ce3565b60018160020160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a58484604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a160009150505b92915050565b600080619cf4614f40565b905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015619d7557600080fd5b505afa158015619d89573d6000803e3d6000fd5b505050506040513d6020811015619d9f57600080fd5b81019080805190602001909291905050509050600084118015619dc25750808411155b15619e95578173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86866040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015619e4e57600080fd5b505af1158015619e62573d6000803e3d6000fd5b505050506040513d6020811015619e7857600080fd5b810190808051906020019092919050505050600092505050619e9b565b83925050505b92915050565b6000619ee383836040518060400160405280601181526020017f6164646974696f6e206f766572666c6f7700000000000000000000000000000081525061a356565b905092915050565b6000619f2d83836040518060400160405280601581526020017f7375627472616374696f6e20756e646572666c6f77000000000000000000000081525061a415565b905092915050565b6000619f7783836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f7700000000000000000081525061a4cf565b905092915050565b60008160000151836000015110905092915050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0846012811115619fc357fe5b846017811115619fcf57fe5b8460405180848152602001838152602001828152602001935050505060405180910390a1836012811115619fff57fe5b90509392505050565b600061a02961a01f84670de0b6b3a7640000619f35565b836000015161a30c565b905092915050565b61a03961a6b3565b604051806020016040528061a06661a060866ec097ce7bc90715b34b9f1000000000619f35565b8561a30c565b815250905092915050565b61a07961a6b3565b604051806020016040528061a09685600001518560000151619ea1565b815250905092915050565b60007c01000000000000000000000000000000000000000000000000000000008310829061a16a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a12f57808201518184015260208101905061a114565b50505050905090810190601f16801561a15c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082905092915050565b60006401000000008310829061a225576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a1ea57808201518184015260208101905061a1cf565b50505050905090810190601f16801561a2175780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082905092915050565b61a23761a6b3565b604051806020016040528061a25485600001518560000151619eeb565b815250905092915050565b60006ec097ce7bc90715b34b9f100000000061a27f848460000151619f35565b8161a28657fe5b04905092915050565b600061a29961a674565b61a2a3858561a2c1565b905061a2b761a2b18261a2ed565b84619ea1565b9150509392505050565b61a2c961a674565b604051806020016040528061a2e2856000015185619f35565b815250905092915050565b6000670de0b6b3a764000082600001518161a30457fe5b049050919050565b600061a34e83836040518060400160405280600e81526020017f646976696465206279207a65726f00000000000000000000000000000000000081525061a5b4565b905092915050565b600080838501905084811015839061a409576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a3ce57808201518184015260208101905061a3b3565b50505050905090810190601f16801561a3fb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50809150509392505050565b600083831115829061a4c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a48757808201518184015260208101905061a46c565b50505050905090810190601f16801561a4b45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b60008084148061a4df5750600083145b1561a4ed576000905061a5ad565b600083850290508385828161a4fe57fe5b0414839061a5a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a56c57808201518184015260208101905061a551565b50505050905090810190601f16801561a5995780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50809150505b9392505050565b6000808311829061a660576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561a62557808201518184015260208101905061a60a565b50505050905090810190601f16801561a6525780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082848161a66a57fe5b0490509392505050565b6040518060200160405280600081525090565b81548183558181111561a6ae5781836000526020600020918201910161a6ad919061a731565b5b505050565b6040518060200160405280600081525090565b60405180610140016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200161a70461a756565b815260200161a71161a756565b815260200161a71e61a756565b815260200161a72b61a756565b81525090565b61a75391905b8082111561a74f57600081600090555060010161a737565b5090565b90565b604051806020016040528060008152509056fe6f6e6c7920706175736520677561726469616e20616e642061646d696e2063616e6f6e6c792061646d696e206f7220626f72726f772063617020677561726469616e2063616e2073657420626f72726f772063617073a265627a7a723158206efe0c84c7a590b0880116eceb926fbb539557f2188894eb34b61d4046bfd84d64736f6c63430005100032"

// DeployVenuscomptroller deploys a new Ethereum contract, binding an instance of Venuscomptroller to it.
func DeployVenuscomptroller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Venuscomptroller, error) {
	parsed, err := abi.JSON(strings.NewReader(VenuscomptrollerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VenuscomptrollerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Venuscomptroller{VenuscomptrollerCaller: VenuscomptrollerCaller{contract: contract}, VenuscomptrollerTransactor: VenuscomptrollerTransactor{contract: contract}, VenuscomptrollerFilterer: VenuscomptrollerFilterer{contract: contract}}, nil
}

// Venuscomptroller is an auto generated Go binding around an Ethereum contract.
type Venuscomptroller struct {
	VenuscomptrollerCaller     // Read-only binding to the contract
	VenuscomptrollerTransactor // Write-only binding to the contract
	VenuscomptrollerFilterer   // Log filterer for contract events
}

// VenuscomptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenuscomptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenuscomptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenuscomptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenuscomptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenuscomptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenuscomptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenuscomptrollerSession struct {
	Contract     *Venuscomptroller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenuscomptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenuscomptrollerCallerSession struct {
	Contract *VenuscomptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VenuscomptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenuscomptrollerTransactorSession struct {
	Contract     *VenuscomptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VenuscomptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenuscomptrollerRaw struct {
	Contract *Venuscomptroller // Generic contract binding to access the raw methods on
}

// VenuscomptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenuscomptrollerCallerRaw struct {
	Contract *VenuscomptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// VenuscomptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenuscomptrollerTransactorRaw struct {
	Contract *VenuscomptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenuscomptroller creates a new instance of Venuscomptroller, bound to a specific deployed contract.
func NewVenuscomptroller(address common.Address, backend bind.ContractBackend) (*Venuscomptroller, error) {
	contract, err := bindVenuscomptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Venuscomptroller{VenuscomptrollerCaller: VenuscomptrollerCaller{contract: contract}, VenuscomptrollerTransactor: VenuscomptrollerTransactor{contract: contract}, VenuscomptrollerFilterer: VenuscomptrollerFilterer{contract: contract}}, nil
}

// NewVenuscomptrollerCaller creates a new read-only instance of Venuscomptroller, bound to a specific deployed contract.
func NewVenuscomptrollerCaller(address common.Address, caller bind.ContractCaller) (*VenuscomptrollerCaller, error) {
	contract, err := bindVenuscomptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerCaller{contract: contract}, nil
}

// NewVenuscomptrollerTransactor creates a new write-only instance of Venuscomptroller, bound to a specific deployed contract.
func NewVenuscomptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*VenuscomptrollerTransactor, error) {
	contract, err := bindVenuscomptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerTransactor{contract: contract}, nil
}

// NewVenuscomptrollerFilterer creates a new log filterer instance of Venuscomptroller, bound to a specific deployed contract.
func NewVenuscomptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*VenuscomptrollerFilterer, error) {
	contract, err := bindVenuscomptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerFilterer{contract: contract}, nil
}

// bindVenuscomptroller binds a generic wrapper to an already deployed contract.
func bindVenuscomptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VenuscomptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venuscomptroller *VenuscomptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venuscomptroller.Contract.VenuscomptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venuscomptroller *VenuscomptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.VenuscomptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venuscomptroller *VenuscomptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.VenuscomptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venuscomptroller *VenuscomptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venuscomptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venuscomptroller *VenuscomptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venuscomptroller *VenuscomptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) BorrowGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.BorrowGuardianPaused(&_Venuscomptroller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) BorrowGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.BorrowGuardianPaused(&_Venuscomptroller.CallOpts)
}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) MintGuardianPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) MintGuardianPaused1() (bool, error) {
	return _Venuscomptroller.Contract.MintGuardianPaused1(&_Venuscomptroller.CallOpts)
}

// MintGuardianPaused1 is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MintGuardianPaused1() (bool, error) {
	return _Venuscomptroller.Contract.MintGuardianPaused1(&_Venuscomptroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Venuscomptroller.Contract.AccountAssets(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Venuscomptroller.Contract.AccountAssets(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) Admin() (common.Address, error) {
	return _Venuscomptroller.Contract.Admin(&_Venuscomptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) Admin() (common.Address, error) {
	return _Venuscomptroller.Contract.Admin(&_Venuscomptroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Venuscomptroller.Contract.AllMarkets(&_Venuscomptroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Venuscomptroller.Contract.AllMarkets(&_Venuscomptroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.BorrowCapGuardian(&_Venuscomptroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.BorrowCapGuardian(&_Venuscomptroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.BorrowCaps(&_Venuscomptroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.BorrowCaps(&_Venuscomptroller.CallOpts, arg0)
}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) BorrowGuardianPaused1(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) BorrowGuardianPaused1(arg0 common.Address) (bool, error) {
	return _Venuscomptroller.Contract.BorrowGuardianPaused1(&_Venuscomptroller.CallOpts, arg0)
}

// BorrowGuardianPaused1 is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) BorrowGuardianPaused1(arg0 common.Address) (bool, error) {
	return _Venuscomptroller.Contract.BorrowGuardianPaused1(&_Venuscomptroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, vToken common.Address) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "checkMembership", account, vToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Venuscomptroller.Contract.CheckMembership(&_Venuscomptroller.CallOpts, account, vToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Venuscomptroller.Contract.CheckMembership(&_Venuscomptroller.CallOpts, account, vToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Venuscomptroller.Contract.CloseFactorMantissa(&_Venuscomptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Venuscomptroller.Contract.CloseFactorMantissa(&_Venuscomptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Venuscomptroller.Contract.ComptrollerImplementation(&_Venuscomptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Venuscomptroller.Contract.ComptrollerImplementation(&_Venuscomptroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.GetAccountLiquidity(&_Venuscomptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.GetAccountLiquidity(&_Venuscomptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Venuscomptroller *VenuscomptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Venuscomptroller *VenuscomptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Venuscomptroller.Contract.GetAllMarkets(&_Venuscomptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Venuscomptroller.Contract.GetAllMarkets(&_Venuscomptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Venuscomptroller *VenuscomptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Venuscomptroller *VenuscomptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Venuscomptroller.Contract.GetAssetsIn(&_Venuscomptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Venuscomptroller.Contract.GetAssetsIn(&_Venuscomptroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Venuscomptroller.Contract.GetBlockNumber(&_Venuscomptroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Venuscomptroller.Contract.GetBlockNumber(&_Venuscomptroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, vTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.GetHypotheticalAccountLiquidity(&_Venuscomptroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.GetHypotheticalAccountLiquidity(&_Venuscomptroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) GetXVSAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "getXVSAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) GetXVSAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.GetXVSAddress(&_Venuscomptroller.CallOpts)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) GetXVSAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.GetXVSAddress(&_Venuscomptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) IsComptroller() (bool, error) {
	return _Venuscomptroller.Contract.IsComptroller(&_Venuscomptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) IsComptroller() (bool, error) {
	return _Venuscomptroller.Contract.IsComptroller(&_Venuscomptroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", vTokenBorrowed, vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.LiquidateCalculateSeizeTokens(&_Venuscomptroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.LiquidateCalculateSeizeTokens(&_Venuscomptroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) LiquidateVAICalculateSeizeTokens(opts *bind.CallOpts, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "liquidateVAICalculateSeizeTokens", vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.LiquidateVAICalculateSeizeTokens(&_Venuscomptroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Venuscomptroller.Contract.LiquidateVAICalculateSeizeTokens(&_Venuscomptroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Venuscomptroller.Contract.LiquidationIncentiveMantissa(&_Venuscomptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Venuscomptroller.Contract.LiquidationIncentiveMantissa(&_Venuscomptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Venuscomptroller *VenuscomptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsVenus                  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsVenus = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Venuscomptroller *VenuscomptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Venuscomptroller.Contract.Markets(&_Venuscomptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Venuscomptroller *VenuscomptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Venuscomptroller.Contract.Markets(&_Venuscomptroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) MaxAssets() (*big.Int, error) {
	return _Venuscomptroller.Contract.MaxAssets(&_Venuscomptroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Venuscomptroller.Contract.MaxAssets(&_Venuscomptroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) MinReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "minReleaseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) MinReleaseAmount() (*big.Int, error) {
	return _Venuscomptroller.Contract.MinReleaseAmount(&_Venuscomptroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MinReleaseAmount() (*big.Int, error) {
	return _Venuscomptroller.Contract.MinReleaseAmount(&_Venuscomptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Venuscomptroller.Contract.MintGuardianPaused(&_Venuscomptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Venuscomptroller.Contract.MintGuardianPaused(&_Venuscomptroller.CallOpts, arg0)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) MintVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "mintVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) MintVAIGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.MintVAIGuardianPaused(&_Venuscomptroller.CallOpts)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MintVAIGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.MintVAIGuardianPaused(&_Venuscomptroller.CallOpts)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) MintedVAIs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "mintedVAIs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.MintedVAIs(&_Venuscomptroller.CallOpts, arg0)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.MintedVAIs(&_Venuscomptroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) Oracle() (common.Address, error) {
	return _Venuscomptroller.Contract.Oracle(&_Venuscomptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) Oracle() (common.Address, error) {
	return _Venuscomptroller.Contract.Oracle(&_Venuscomptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) PauseGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.PauseGuardian(&_Venuscomptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.PauseGuardian(&_Venuscomptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) PendingAdmin() (common.Address, error) {
	return _Venuscomptroller.Contract.PendingAdmin(&_Venuscomptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Venuscomptroller.Contract.PendingAdmin(&_Venuscomptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Venuscomptroller.Contract.PendingComptrollerImplementation(&_Venuscomptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Venuscomptroller.Contract.PendingComptrollerImplementation(&_Venuscomptroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) ProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "protocolPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) ProtocolPaused() (bool, error) {
	return _Venuscomptroller.Contract.ProtocolPaused(&_Venuscomptroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) ProtocolPaused() (bool, error) {
	return _Venuscomptroller.Contract.ProtocolPaused(&_Venuscomptroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) ReleaseStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "releaseStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Venuscomptroller.Contract.ReleaseStartBlock(&_Venuscomptroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Venuscomptroller.Contract.ReleaseStartBlock(&_Venuscomptroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) RepayVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "repayVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.RepayVAIGuardianPaused(&_Venuscomptroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.RepayVAIGuardianPaused(&_Venuscomptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.SeizeGuardianPaused(&_Venuscomptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.SeizeGuardianPaused(&_Venuscomptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) TransferGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.TransferGuardianPaused(&_Venuscomptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Venuscomptroller *VenuscomptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Venuscomptroller.Contract.TransferGuardianPaused(&_Venuscomptroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) TreasuryAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.TreasuryAddress(&_Venuscomptroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.TreasuryAddress(&_Venuscomptroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) TreasuryGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.TreasuryGuardian(&_Venuscomptroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _Venuscomptroller.Contract.TreasuryGuardian(&_Venuscomptroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) TreasuryPercent() (*big.Int, error) {
	return _Venuscomptroller.Contract.TreasuryPercent(&_Venuscomptroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _Venuscomptroller.Contract.TreasuryPercent(&_Venuscomptroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) VaiController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "vaiController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) VaiController() (common.Address, error) {
	return _Venuscomptroller.Contract.VaiController(&_Venuscomptroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VaiController() (common.Address, error) {
	return _Venuscomptroller.Contract.VaiController(&_Venuscomptroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VaiMintRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "vaiMintRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VaiMintRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VaiMintRate(&_Venuscomptroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VaiMintRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VaiMintRate(&_Venuscomptroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCaller) VaiVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "vaiVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerSession) VaiVaultAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.VaiVaultAddress(&_Venuscomptroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VaiVaultAddress() (common.Address, error) {
	return _Venuscomptroller.Contract.VaiVaultAddress(&_Venuscomptroller.CallOpts)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusAccrued(&_Venuscomptroller.CallOpts, arg0)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusAccrued(&_Venuscomptroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Venuscomptroller.Contract.VenusBorrowState(&_Venuscomptroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Venuscomptroller.Contract.VenusBorrowState(&_Venuscomptroller.CallOpts, arg0)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusBorrowerIndex(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusBorrowerIndex(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Venuscomptroller *VenuscomptrollerSession) VenusInitialIndex() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusInitialIndex(&_Venuscomptroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusInitialIndex(&_Venuscomptroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusRate(&_Venuscomptroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusRate(&_Venuscomptroller.CallOpts)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusSpeeds(&_Venuscomptroller.CallOpts, arg0)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusSpeeds(&_Venuscomptroller.CallOpts, arg0)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusSupplierIndex(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusSupplierIndex(&_Venuscomptroller.CallOpts, arg0, arg1)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Venuscomptroller.Contract.VenusSupplyState(&_Venuscomptroller.CallOpts, arg0)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Venuscomptroller.Contract.VenusSupplyState(&_Venuscomptroller.CallOpts, arg0)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusVAIRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusVAIRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusVAIRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusVAIRate(&_Venuscomptroller.CallOpts)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusVAIRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusVAIRate(&_Venuscomptroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCaller) VenusVAIVaultRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venuscomptroller.contract.Call(opts, &out, "venusVAIVaultRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusVAIVaultRate(&_Venuscomptroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Venuscomptroller *VenuscomptrollerCallerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Venuscomptroller.Contract.VenusVAIVaultRate(&_Venuscomptroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Venuscomptroller *VenuscomptrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.Become(&_Venuscomptroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.Become(&_Venuscomptroller.TransactOpts, unitroller)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetBorrowCapGuardian(&_Venuscomptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetBorrowCapGuardian(&_Venuscomptroller.TransactOpts, newBorrowCapGuardian)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetCloseFactor(&_Venuscomptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetCloseFactor(&_Venuscomptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setCollateralFactor", vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetCollateralFactor(&_Venuscomptroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetCollateralFactor(&_Venuscomptroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetLiquidationIncentive(&_Venuscomptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetLiquidationIncentive(&_Venuscomptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setMarketBorrowCaps", vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetMarketBorrowCaps(&_Venuscomptroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetMarketBorrowCaps(&_Venuscomptroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetPauseGuardian(&_Venuscomptroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetPauseGuardian(&_Venuscomptroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetPriceOracle(&_Venuscomptroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetPriceOracle(&_Venuscomptroller.TransactOpts, newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetProtocolPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setProtocolPaused", state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Venuscomptroller *VenuscomptrollerSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetProtocolPaused(&_Venuscomptroller.TransactOpts, state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetProtocolPaused(&_Venuscomptroller.TransactOpts, state)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetTreasuryData(&_Venuscomptroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetTreasuryData(&_Venuscomptroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVAIController(opts *bind.TransactOpts, vaiController_ common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVAIController", vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIController(&_Venuscomptroller.TransactOpts, vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIController(&_Venuscomptroller.TransactOpts, vaiController_)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVAIMintRate(opts *bind.TransactOpts, newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVAIMintRate", newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIMintRate(&_Venuscomptroller.TransactOpts, newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIMintRate(&_Venuscomptroller.TransactOpts, newVAIMintRate)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVAIVaultInfo(opts *bind.TransactOpts, vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVAIVaultInfo", vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIVaultInfo(&_Venuscomptroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVAIVaultInfo(&_Venuscomptroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVenusSpeed(opts *bind.TransactOpts, vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVenusSpeed", vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusSpeed(&_Venuscomptroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusSpeed(&_Venuscomptroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVenusVAIRate(opts *bind.TransactOpts, venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVenusVAIRate", venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusVAIRate(&_Venuscomptroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusVAIRate(&_Venuscomptroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SetVenusVAIVaultRate(opts *bind.TransactOpts, venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_setVenusVAIVaultRate", venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusVAIVaultRate(&_Venuscomptroller.TransactOpts, venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetVenusVAIVaultRate(&_Venuscomptroller.TransactOpts, venusVAIVaultRate_)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SupportMarket(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "_supportMarket", vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SupportMarket(&_Venuscomptroller.TransactOpts, vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SupportMarket(&_Venuscomptroller.TransactOpts, vToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "borrowAllowed", vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.BorrowAllowed(&_Venuscomptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.BorrowAllowed(&_Venuscomptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "borrowVerify", vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Venuscomptroller *VenuscomptrollerSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.BorrowVerify(&_Venuscomptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.BorrowVerify(&_Venuscomptroller.TransactOpts, vToken, borrower, borrowAmount)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) ClaimVenus(opts *bind.TransactOpts, holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "claimVenus", holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus(&_Venuscomptroller.TransactOpts, holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus(&_Venuscomptroller.TransactOpts, holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) ClaimVenus0(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "claimVenus0", holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Venuscomptroller *VenuscomptrollerSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus0(&_Venuscomptroller.TransactOpts, holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus0(&_Venuscomptroller.TransactOpts, holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) ClaimVenus1(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "claimVenus1", holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Venuscomptroller *VenuscomptrollerSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus1(&_Venuscomptroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ClaimVenus1(&_Venuscomptroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) DistributeVAIMinterVenus(opts *bind.TransactOpts, vaiMinter common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "distributeVAIMinterVenus", vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Venuscomptroller *VenuscomptrollerSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.DistributeVAIMinterVenus(&_Venuscomptroller.TransactOpts, vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.DistributeVAIMinterVenus(&_Venuscomptroller.TransactOpts, vaiMinter)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Venuscomptroller *VenuscomptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "enterMarkets", vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Venuscomptroller *VenuscomptrollerSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.EnterMarkets(&_Venuscomptroller.TransactOpts, vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Venuscomptroller *VenuscomptrollerTransactorSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.EnterMarkets(&_Venuscomptroller.TransactOpts, vTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) ExitMarket(opts *bind.TransactOpts, vTokenAddress common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "exitMarket", vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ExitMarket(&_Venuscomptroller.TransactOpts, vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ExitMarket(&_Venuscomptroller.TransactOpts, vTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "liquidateBorrowAllowed", vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.LiquidateBorrowAllowed(&_Venuscomptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.LiquidateBorrowAllowed(&_Venuscomptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "liquidateBorrowVerify", vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.LiquidateBorrowVerify(&_Venuscomptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.LiquidateBorrowVerify(&_Venuscomptroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) MintAllowed(opts *bind.TransactOpts, vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "mintAllowed", vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.MintAllowed(&_Venuscomptroller.TransactOpts, vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.MintAllowed(&_Venuscomptroller.TransactOpts, vToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) MintVerify(opts *bind.TransactOpts, vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "mintVerify", vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.MintVerify(&_Venuscomptroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.MintVerify(&_Venuscomptroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "redeemAllowed", vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RedeemAllowed(&_Venuscomptroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RedeemAllowed(&_Venuscomptroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "redeemVerify", vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RedeemVerify(&_Venuscomptroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RedeemVerify(&_Venuscomptroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) ReleaseToVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "releaseToVault")
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Venuscomptroller *VenuscomptrollerSession) ReleaseToVault() (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ReleaseToVault(&_Venuscomptroller.TransactOpts)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) ReleaseToVault() (*types.Transaction, error) {
	return _Venuscomptroller.Contract.ReleaseToVault(&_Venuscomptroller.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "repayBorrowAllowed", vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RepayBorrowAllowed(&_Venuscomptroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RepayBorrowAllowed(&_Venuscomptroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "repayBorrowVerify", vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Venuscomptroller *VenuscomptrollerSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RepayBorrowVerify(&_Venuscomptroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.RepayBorrowVerify(&_Venuscomptroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "seizeAllowed", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SeizeAllowed(&_Venuscomptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SeizeAllowed(&_Venuscomptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "seizeVerify", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SeizeVerify(&_Venuscomptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SeizeVerify(&_Venuscomptroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) SetMintedVAIOf(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "setMintedVAIOf", owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetMintedVAIOf(&_Venuscomptroller.TransactOpts, owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.SetMintedVAIOf(&_Venuscomptroller.TransactOpts, owner, amount)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "transferAllowed", vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.TransferAllowed(&_Venuscomptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Venuscomptroller *VenuscomptrollerTransactorSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.TransferAllowed(&_Venuscomptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactor) TransferVerify(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.contract.Transact(opts, "transferVerify", vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Venuscomptroller *VenuscomptrollerSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.TransferVerify(&_Venuscomptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Venuscomptroller *VenuscomptrollerTransactorSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Venuscomptroller.Contract.TransferVerify(&_Venuscomptroller.TransactOpts, vToken, src, dst, transferTokens)
}

// VenuscomptrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Venuscomptroller contract.
type VenuscomptrollerActionPausedIterator struct {
	Event *VenuscomptrollerActionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerActionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerActionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerActionPaused represents a ActionPaused event raised by the Venuscomptroller contract.
type VenuscomptrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*VenuscomptrollerActionPausedIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerActionPausedIterator{contract: _Venuscomptroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerActionPaused)
				if err := _Venuscomptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseActionPaused(log types.Log) (*VenuscomptrollerActionPaused, error) {
	event := new(VenuscomptrollerActionPaused)
	if err := _Venuscomptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Venuscomptroller contract.
type VenuscomptrollerActionPaused0Iterator struct {
	Event *VenuscomptrollerActionPaused0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerActionPaused0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerActionPaused0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerActionPaused0 represents a ActionPaused0 event raised by the Venuscomptroller contract.
type VenuscomptrollerActionPaused0 struct {
	VToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*VenuscomptrollerActionPaused0Iterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerActionPaused0Iterator{contract: _Venuscomptroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerActionPaused0)
				if err := _Venuscomptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseActionPaused0(log types.Log) (*VenuscomptrollerActionPaused0, error) {
	event := new(VenuscomptrollerActionPaused0)
	if err := _Venuscomptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerActionProtocolPausedIterator is returned from FilterActionProtocolPaused and is used to iterate over the raw logs and unpacked data for ActionProtocolPaused events raised by the Venuscomptroller contract.
type VenuscomptrollerActionProtocolPausedIterator struct {
	Event *VenuscomptrollerActionProtocolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerActionProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerActionProtocolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerActionProtocolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerActionProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerActionProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerActionProtocolPaused represents a ActionProtocolPaused event raised by the Venuscomptroller contract.
type VenuscomptrollerActionProtocolPaused struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActionProtocolPaused is a free log retrieval operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterActionProtocolPaused(opts *bind.FilterOpts) (*VenuscomptrollerActionProtocolPausedIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerActionProtocolPausedIterator{contract: _Venuscomptroller.contract, event: "ActionProtocolPaused", logs: logs, sub: sub}, nil
}

// WatchActionProtocolPaused is a free log subscription operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchActionProtocolPaused(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerActionProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerActionProtocolPaused)
				if err := _Venuscomptroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionProtocolPaused is a log parse operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseActionProtocolPaused(log types.Log) (*VenuscomptrollerActionProtocolPaused, error) {
	event := new(VenuscomptrollerActionProtocolPaused)
	if err := _Venuscomptroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerDistributedBorrowerVenusIterator is returned from FilterDistributedBorrowerVenus and is used to iterate over the raw logs and unpacked data for DistributedBorrowerVenus events raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedBorrowerVenusIterator struct {
	Event *VenuscomptrollerDistributedBorrowerVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerDistributedBorrowerVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerDistributedBorrowerVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerDistributedBorrowerVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerDistributedBorrowerVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerDistributedBorrowerVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerDistributedBorrowerVenus represents a DistributedBorrowerVenus event raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedBorrowerVenus struct {
	VToken           common.Address
	Borrower         common.Address
	VenusDelta       *big.Int
	VenusBorrowIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerVenus is a free log retrieval operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterDistributedBorrowerVenus(opts *bind.FilterOpts, vToken []common.Address, borrower []common.Address) (*VenuscomptrollerDistributedBorrowerVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerDistributedBorrowerVenusIterator{contract: _Venuscomptroller.contract, event: "DistributedBorrowerVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerVenus is a free log subscription operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchDistributedBorrowerVenus(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerDistributedBorrowerVenus, vToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerDistributedBorrowerVenus)
				if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedBorrowerVenus is a log parse operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseDistributedBorrowerVenus(log types.Log) (*VenuscomptrollerDistributedBorrowerVenus, error) {
	event := new(VenuscomptrollerDistributedBorrowerVenus)
	if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerDistributedSupplierVenusIterator is returned from FilterDistributedSupplierVenus and is used to iterate over the raw logs and unpacked data for DistributedSupplierVenus events raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedSupplierVenusIterator struct {
	Event *VenuscomptrollerDistributedSupplierVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerDistributedSupplierVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerDistributedSupplierVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerDistributedSupplierVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerDistributedSupplierVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerDistributedSupplierVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerDistributedSupplierVenus represents a DistributedSupplierVenus event raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedSupplierVenus struct {
	VToken           common.Address
	Supplier         common.Address
	VenusDelta       *big.Int
	VenusSupplyIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierVenus is a free log retrieval operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterDistributedSupplierVenus(opts *bind.FilterOpts, vToken []common.Address, supplier []common.Address) (*VenuscomptrollerDistributedSupplierVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerDistributedSupplierVenusIterator{contract: _Venuscomptroller.contract, event: "DistributedSupplierVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierVenus is a free log subscription operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchDistributedSupplierVenus(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerDistributedSupplierVenus, vToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerDistributedSupplierVenus)
				if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedSupplierVenus is a log parse operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseDistributedSupplierVenus(log types.Log) (*VenuscomptrollerDistributedSupplierVenus, error) {
	event := new(VenuscomptrollerDistributedSupplierVenus)
	if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerDistributedVAIMinterVenusIterator is returned from FilterDistributedVAIMinterVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIMinterVenus events raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedVAIMinterVenusIterator struct {
	Event *VenuscomptrollerDistributedVAIMinterVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerDistributedVAIMinterVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerDistributedVAIMinterVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerDistributedVAIMinterVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerDistributedVAIMinterVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerDistributedVAIMinterVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerDistributedVAIMinterVenus represents a DistributedVAIMinterVenus event raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedVAIMinterVenus struct {
	VaiMinter         common.Address
	VenusDelta        *big.Int
	VenusVAIMintIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIMinterVenus is a free log retrieval operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterDistributedVAIMinterVenus(opts *bind.FilterOpts, vaiMinter []common.Address) (*VenuscomptrollerDistributedVAIMinterVenusIterator, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerDistributedVAIMinterVenusIterator{contract: _Venuscomptroller.contract, event: "DistributedVAIMinterVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIMinterVenus is a free log subscription operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchDistributedVAIMinterVenus(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerDistributedVAIMinterVenus, vaiMinter []common.Address) (event.Subscription, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerDistributedVAIMinterVenus)
				if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIMinterVenus is a log parse operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseDistributedVAIMinterVenus(log types.Log) (*VenuscomptrollerDistributedVAIMinterVenus, error) {
	event := new(VenuscomptrollerDistributedVAIMinterVenus)
	if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerDistributedVAIVaultVenusIterator is returned from FilterDistributedVAIVaultVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIVaultVenus events raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedVAIVaultVenusIterator struct {
	Event *VenuscomptrollerDistributedVAIVaultVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerDistributedVAIVaultVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerDistributedVAIVaultVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerDistributedVAIVaultVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerDistributedVAIVaultVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerDistributedVAIVaultVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerDistributedVAIVaultVenus represents a DistributedVAIVaultVenus event raised by the Venuscomptroller contract.
type VenuscomptrollerDistributedVAIVaultVenus struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIVaultVenus is a free log retrieval operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterDistributedVAIVaultVenus(opts *bind.FilterOpts) (*VenuscomptrollerDistributedVAIVaultVenusIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerDistributedVAIVaultVenusIterator{contract: _Venuscomptroller.contract, event: "DistributedVAIVaultVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIVaultVenus is a free log subscription operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchDistributedVAIVaultVenus(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerDistributedVAIVaultVenus) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerDistributedVAIVaultVenus)
				if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIVaultVenus is a log parse operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseDistributedVAIVaultVenus(log types.Log) (*VenuscomptrollerDistributedVAIVaultVenus, error) {
	event := new(VenuscomptrollerDistributedVAIVaultVenus)
	if err := _Venuscomptroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Venuscomptroller contract.
type VenuscomptrollerFailureIterator struct {
	Event *VenuscomptrollerFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerFailure represents a Failure event raised by the Venuscomptroller contract.
type VenuscomptrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*VenuscomptrollerFailureIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerFailureIterator{contract: _Venuscomptroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerFailure)
				if err := _Venuscomptroller.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseFailure(log types.Log) (*VenuscomptrollerFailure, error) {
	event := new(VenuscomptrollerFailure)
	if err := _Venuscomptroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Venuscomptroller contract.
type VenuscomptrollerMarketEnteredIterator struct {
	Event *VenuscomptrollerMarketEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerMarketEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerMarketEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerMarketEntered represents a MarketEntered event raised by the Venuscomptroller contract.
type VenuscomptrollerMarketEntered struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*VenuscomptrollerMarketEnteredIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerMarketEnteredIterator{contract: _Venuscomptroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerMarketEntered)
				if err := _Venuscomptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseMarketEntered(log types.Log) (*VenuscomptrollerMarketEntered, error) {
	event := new(VenuscomptrollerMarketEntered)
	if err := _Venuscomptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Venuscomptroller contract.
type VenuscomptrollerMarketExitedIterator struct {
	Event *VenuscomptrollerMarketExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerMarketExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerMarketExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerMarketExited represents a MarketExited event raised by the Venuscomptroller contract.
type VenuscomptrollerMarketExited struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*VenuscomptrollerMarketExitedIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerMarketExitedIterator{contract: _Venuscomptroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerMarketExited)
				if err := _Venuscomptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseMarketExited(log types.Log) (*VenuscomptrollerMarketExited, error) {
	event := new(VenuscomptrollerMarketExited)
	if err := _Venuscomptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Venuscomptroller contract.
type VenuscomptrollerMarketListedIterator struct {
	Event *VenuscomptrollerMarketListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerMarketListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerMarketListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerMarketListed represents a MarketListed event raised by the Venuscomptroller contract.
type VenuscomptrollerMarketListed struct {
	VToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*VenuscomptrollerMarketListedIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerMarketListedIterator{contract: _Venuscomptroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerMarketListed)
				if err := _Venuscomptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseMarketListed(log types.Log) (*VenuscomptrollerMarketListed, error) {
	event := new(VenuscomptrollerMarketListed)
	if err := _Venuscomptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Venuscomptroller contract.
type VenuscomptrollerNewBorrowCapIterator struct {
	Event *VenuscomptrollerNewBorrowCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewBorrowCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewBorrowCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewBorrowCap represents a NewBorrowCap event raised by the Venuscomptroller contract.
type VenuscomptrollerNewBorrowCap struct {
	VToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, vToken []common.Address) (*VenuscomptrollerNewBorrowCapIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewBorrowCapIterator{contract: _Venuscomptroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewBorrowCap, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewBorrowCap)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewBorrowCap(log types.Log) (*VenuscomptrollerNewBorrowCap, error) {
	event := new(VenuscomptrollerNewBorrowCap)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Venuscomptroller contract.
type VenuscomptrollerNewBorrowCapGuardianIterator struct {
	Event *VenuscomptrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewBorrowCapGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewBorrowCapGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Venuscomptroller contract.
type VenuscomptrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*VenuscomptrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewBorrowCapGuardianIterator{contract: _Venuscomptroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewBorrowCapGuardian)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*VenuscomptrollerNewBorrowCapGuardian, error) {
	event := new(VenuscomptrollerNewBorrowCapGuardian)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Venuscomptroller contract.
type VenuscomptrollerNewCloseFactorIterator struct {
	Event *VenuscomptrollerNewCloseFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewCloseFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewCloseFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewCloseFactor represents a NewCloseFactor event raised by the Venuscomptroller contract.
type VenuscomptrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*VenuscomptrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewCloseFactorIterator{contract: _Venuscomptroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewCloseFactor)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewCloseFactor(log types.Log) (*VenuscomptrollerNewCloseFactor, error) {
	event := new(VenuscomptrollerNewCloseFactor)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Venuscomptroller contract.
type VenuscomptrollerNewCollateralFactorIterator struct {
	Event *VenuscomptrollerNewCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Venuscomptroller contract.
type VenuscomptrollerNewCollateralFactor struct {
	VToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*VenuscomptrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewCollateralFactorIterator{contract: _Venuscomptroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewCollateralFactor)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewCollateralFactor(log types.Log) (*VenuscomptrollerNewCollateralFactor, error) {
	event := new(VenuscomptrollerNewCollateralFactor)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Venuscomptroller contract.
type VenuscomptrollerNewLiquidationIncentiveIterator struct {
	Event *VenuscomptrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewLiquidationIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewLiquidationIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Venuscomptroller contract.
type VenuscomptrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*VenuscomptrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewLiquidationIncentiveIterator{contract: _Venuscomptroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewLiquidationIncentive)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*VenuscomptrollerNewLiquidationIncentive, error) {
	event := new(VenuscomptrollerNewLiquidationIncentive)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Venuscomptroller contract.
type VenuscomptrollerNewPauseGuardianIterator struct {
	Event *VenuscomptrollerNewPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Venuscomptroller contract.
type VenuscomptrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*VenuscomptrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewPauseGuardianIterator{contract: _Venuscomptroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewPauseGuardian)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewPauseGuardian(log types.Log) (*VenuscomptrollerNewPauseGuardian, error) {
	event := new(VenuscomptrollerNewPauseGuardian)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Venuscomptroller contract.
type VenuscomptrollerNewPriceOracleIterator struct {
	Event *VenuscomptrollerNewPriceOracle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewPriceOracle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewPriceOracle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewPriceOracle represents a NewPriceOracle event raised by the Venuscomptroller contract.
type VenuscomptrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*VenuscomptrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewPriceOracleIterator{contract: _Venuscomptroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewPriceOracle)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewPriceOracle(log types.Log) (*VenuscomptrollerNewPriceOracle, error) {
	event := new(VenuscomptrollerNewPriceOracle)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryAddressIterator struct {
	Event *VenuscomptrollerNewTreasuryAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewTreasuryAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewTreasuryAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*VenuscomptrollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewTreasuryAddressIterator{contract: _Venuscomptroller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewTreasuryAddress)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewTreasuryAddress(log types.Log) (*VenuscomptrollerNewTreasuryAddress, error) {
	event := new(VenuscomptrollerNewTreasuryAddress)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryGuardianIterator struct {
	Event *VenuscomptrollerNewTreasuryGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewTreasuryGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewTreasuryGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*VenuscomptrollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewTreasuryGuardianIterator{contract: _Venuscomptroller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewTreasuryGuardian)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*VenuscomptrollerNewTreasuryGuardian, error) {
	event := new(VenuscomptrollerNewTreasuryGuardian)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryPercentIterator struct {
	Event *VenuscomptrollerNewTreasuryPercent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewTreasuryPercent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewTreasuryPercent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the Venuscomptroller contract.
type VenuscomptrollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*VenuscomptrollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewTreasuryPercentIterator{contract: _Venuscomptroller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewTreasuryPercent)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewTreasuryPercent(log types.Log) (*VenuscomptrollerNewTreasuryPercent, error) {
	event := new(VenuscomptrollerNewTreasuryPercent)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewVAIControllerIterator is returned from FilterNewVAIController and is used to iterate over the raw logs and unpacked data for NewVAIController events raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIControllerIterator struct {
	Event *VenuscomptrollerNewVAIController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewVAIControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewVAIController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewVAIController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewVAIControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewVAIControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewVAIController represents a NewVAIController event raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIController struct {
	OldVAIController common.Address
	NewVAIController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewVAIController is a free log retrieval operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewVAIController(opts *bind.FilterOpts) (*VenuscomptrollerNewVAIControllerIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewVAIControllerIterator{contract: _Venuscomptroller.contract, event: "NewVAIController", logs: logs, sub: sub}, nil
}

// WatchNewVAIController is a free log subscription operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewVAIController(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewVAIController) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewVAIController)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIController is a log parse operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewVAIController(log types.Log) (*VenuscomptrollerNewVAIController, error) {
	event := new(VenuscomptrollerNewVAIController)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewVAIMintRateIterator is returned from FilterNewVAIMintRate and is used to iterate over the raw logs and unpacked data for NewVAIMintRate events raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIMintRateIterator struct {
	Event *VenuscomptrollerNewVAIMintRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewVAIMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewVAIMintRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewVAIMintRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewVAIMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewVAIMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewVAIMintRate represents a NewVAIMintRate event raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIMintRate struct {
	OldVAIMintRate *big.Int
	NewVAIMintRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintRate is a free log retrieval operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewVAIMintRate(opts *bind.FilterOpts) (*VenuscomptrollerNewVAIMintRateIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewVAIMintRateIterator{contract: _Venuscomptroller.contract, event: "NewVAIMintRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintRate is a free log subscription operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewVAIMintRate(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewVAIMintRate) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewVAIMintRate)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIMintRate is a log parse operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewVAIMintRate(log types.Log) (*VenuscomptrollerNewVAIMintRate, error) {
	event := new(VenuscomptrollerNewVAIMintRate)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewVAIVaultInfoIterator is returned from FilterNewVAIVaultInfo and is used to iterate over the raw logs and unpacked data for NewVAIVaultInfo events raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIVaultInfoIterator struct {
	Event *VenuscomptrollerNewVAIVaultInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewVAIVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewVAIVaultInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewVAIVaultInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewVAIVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewVAIVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewVAIVaultInfo represents a NewVAIVaultInfo event raised by the Venuscomptroller contract.
type VenuscomptrollerNewVAIVaultInfo struct {
	Vault             common.Address
	ReleaseStartBlock *big.Int
	ReleaseInterval   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVAIVaultInfo is a free log retrieval operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewVAIVaultInfo(opts *bind.FilterOpts) (*VenuscomptrollerNewVAIVaultInfoIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewVAIVaultInfoIterator{contract: _Venuscomptroller.contract, event: "NewVAIVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVAIVaultInfo is a free log subscription operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewVAIVaultInfo(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewVAIVaultInfo) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewVAIVaultInfo)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIVaultInfo is a log parse operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewVAIVaultInfo(log types.Log) (*VenuscomptrollerNewVAIVaultInfo, error) {
	event := new(VenuscomptrollerNewVAIVaultInfo)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewVenusVAIRateIterator is returned from FilterNewVenusVAIRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIRate events raised by the Venuscomptroller contract.
type VenuscomptrollerNewVenusVAIRateIterator struct {
	Event *VenuscomptrollerNewVenusVAIRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewVenusVAIRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewVenusVAIRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewVenusVAIRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewVenusVAIRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewVenusVAIRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewVenusVAIRate represents a NewVenusVAIRate event raised by the Venuscomptroller contract.
type VenuscomptrollerNewVenusVAIRate struct {
	OldVenusVAIRate *big.Int
	NewVenusVAIRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIRate is a free log retrieval operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewVenusVAIRate(opts *bind.FilterOpts) (*VenuscomptrollerNewVenusVAIRateIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewVenusVAIRateIterator{contract: _Venuscomptroller.contract, event: "NewVenusVAIRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIRate is a free log subscription operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewVenusVAIRate(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewVenusVAIRate) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewVenusVAIRate)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIRate is a log parse operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewVenusVAIRate(log types.Log) (*VenuscomptrollerNewVenusVAIRate, error) {
	event := new(VenuscomptrollerNewVenusVAIRate)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerNewVenusVAIVaultRateIterator is returned from FilterNewVenusVAIVaultRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIVaultRate events raised by the Venuscomptroller contract.
type VenuscomptrollerNewVenusVAIVaultRateIterator struct {
	Event *VenuscomptrollerNewVenusVAIVaultRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerNewVenusVAIVaultRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerNewVenusVAIVaultRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerNewVenusVAIVaultRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerNewVenusVAIVaultRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerNewVenusVAIVaultRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerNewVenusVAIVaultRate represents a NewVenusVAIVaultRate event raised by the Venuscomptroller contract.
type VenuscomptrollerNewVenusVAIVaultRate struct {
	OldVenusVAIVaultRate *big.Int
	NewVenusVAIVaultRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIVaultRate is a free log retrieval operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterNewVenusVAIVaultRate(opts *bind.FilterOpts) (*VenuscomptrollerNewVenusVAIVaultRateIterator, error) {

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerNewVenusVAIVaultRateIterator{contract: _Venuscomptroller.contract, event: "NewVenusVAIVaultRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIVaultRate is a free log subscription operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchNewVenusVAIVaultRate(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerNewVenusVAIVaultRate) (event.Subscription, error) {

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerNewVenusVAIVaultRate)
				if err := _Venuscomptroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIVaultRate is a log parse operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseNewVenusVAIVaultRate(log types.Log) (*VenuscomptrollerNewVenusVAIVaultRate, error) {
	event := new(VenuscomptrollerNewVenusVAIVaultRate)
	if err := _Venuscomptroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenuscomptrollerVenusSpeedUpdatedIterator is returned from FilterVenusSpeedUpdated and is used to iterate over the raw logs and unpacked data for VenusSpeedUpdated events raised by the Venuscomptroller contract.
type VenuscomptrollerVenusSpeedUpdatedIterator struct {
	Event *VenuscomptrollerVenusSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VenuscomptrollerVenusSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenuscomptrollerVenusSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VenuscomptrollerVenusSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VenuscomptrollerVenusSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenuscomptrollerVenusSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenuscomptrollerVenusSpeedUpdated represents a VenusSpeedUpdated event raised by the Venuscomptroller contract.
type VenuscomptrollerVenusSpeedUpdated struct {
	VToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVenusSpeedUpdated is a free log retrieval operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Venuscomptroller *VenuscomptrollerFilterer) FilterVenusSpeedUpdated(opts *bind.FilterOpts, vToken []common.Address) (*VenuscomptrollerVenusSpeedUpdatedIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Venuscomptroller.contract.FilterLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenuscomptrollerVenusSpeedUpdatedIterator{contract: _Venuscomptroller.contract, event: "VenusSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchVenusSpeedUpdated is a free log subscription operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Venuscomptroller *VenuscomptrollerFilterer) WatchVenusSpeedUpdated(opts *bind.WatchOpts, sink chan<- *VenuscomptrollerVenusSpeedUpdated, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Venuscomptroller.contract.WatchLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenuscomptrollerVenusSpeedUpdated)
				if err := _Venuscomptroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVenusSpeedUpdated is a log parse operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Venuscomptroller *VenuscomptrollerFilterer) ParseVenusSpeedUpdated(log types.Log) (*VenuscomptrollerVenusSpeedUpdated, error) {
	event := new(VenuscomptrollerVenusSpeedUpdated)
	if err := _Venuscomptroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
