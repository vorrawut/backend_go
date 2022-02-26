// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venusbep20delegator

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

// Venusbep20delegatorABI is the input ABI used to generate the binding from.
const Venusbep20delegatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"MintBehalf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"RedeemFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractVTokenInterface\",\"name\":\"vTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Venusbep20delegatorBin is the compiled bytecode used for deploying new contracts.
var Venusbep20delegatorBin = "0x60806040523480156200001157600080fd5b50604051620048253803806200482583398181016040526101408110156200003857600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190805160405193929190846401000000008211156200008157600080fd5b838201915060208201858111156200009857600080fd5b8251866001820283011164010000000082111715620000b657600080fd5b8083526020830192505050908051906020019080838360005b83811015620000ec578082015181840152602081019050620000cf565b50505050905090810190601f1680156200011a5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200013e57600080fd5b838201915060208201858111156200015557600080fd5b82518660018202830111640100000000821117156200017357600080fd5b8083526020830192505050908051906020019080838360005b83811015620001a95780820151818401526020810190506200018c565b50505050905090810190601f168015620001d75780820380516001836020036101000a031916815260200191505b50604052602001805190602001909291908051906020019092919080519060200190929190805160405193929190846401000000008211156200021957600080fd5b838201915060208201858111156200023057600080fd5b82518660018202830111640100000000821117156200024e57600080fd5b8083526020830192505050908051906020019080838360005b838110156200028457808201518184015260208101905062000267565b50505050905090810190601f168015620002b25780820380516001836020036101000a031916815260200191505b5060405250505033600360016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000523828b8b8b8b8b8b8b604051602401808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200180602001806020018460ff1660ff168152602001838103835286818151815260200191508051906020019080838360005b83811015620003f7578082015181840152602081019050620003da565b50505050905090810190601f168015620004255780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156200046057808201518184015260208101905062000443565b50505050905090810190601f1680156200048e5780820380516001836020036101000a031916815260200191505b5099505050505050505050506040516020818303038152906040527f1a31d465000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200058960201b60201c565b5062000538826000836200066460201b60201c565b82600360016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050505062000a13565b6060600060608473ffffffffffffffffffffffffffffffffffffffff16846040518082805190602001908083835b60208310620005dc5780518252602082019150602081019050602083039250620005b7565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146200063e576040519150601f19603f3d011682016040523d82523d6000602084013e62000643565b606091505b5091509150600082141562000659573d60208201fd5b809250505092915050565b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200070c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526039815260200180620047ec6039913960400191505060405180910390fd5b8115620007a857620007a66040516024016040516020818303038152906040527f153ab505000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050620009d660201b60201c565b505b6000601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083601260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000916826040516024018080602001828103825283818151815260200191508051906020019080838360005b838110156200085a5780820151818401526020810190506200083d565b50505050905090810190601f168015620008885780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040527f56e67728000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050620009d660201b60201c565b507fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a81601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150505050565b606062000a0c601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836200058960201b60201c565b9050919050565b613dc98062000a236000396000f3fe6080604052600436106102e45760003560e01c80636f307dc311610190578063b71d1a0c116100dc578063e9c714f211610095578063f5e3c4621161006f578063f5e3c46214611596578063f851a44014611625578063f8f9da281461167c578063fca7820b146116a7576102e4565b8063e9c714f2146114af578063f2b3abbd146114da578063f3fdb15a1461153f576102e4565b8063b71d1a0c14611282578063bd6d894d146112e7578063c37f68e214611312578063c5ebeaec1461138c578063db006a75146113db578063dd62ed3e1461142a576102e4565b806395dd919311610149578063a9059cbb11610123578063a9059cbb1461112a578063aa5af0fd1461119d578063ae9d70b0146111c8578063b2a02ff1146111f3576102e4565b806395dd91931461104b578063a0712d68146110b0578063a6afed95146110ff576102e4565b80636f307dc314610e5a57806370a0823114610eb157806373acee9814610f16578063852a12e314610f415780638f840ddd14610f9057806395d89b4114610fbb576102e4565b8063313ce5671161024f5780634576b5db116102085780635c60da1b116101e25780635c60da1b14610d325780635fe3b56714610d89578063601a0bf114610de05780636c540baf14610e2f576102e4565b80634576b5db14610bae57806347bd371814610c13578063555bcc4014610c3e576102e4565b8063313ce5671461092e5780633af9e6691461095f5780633b1d21a2146109c45780633d9ea3a1146109ef5780633e94101014610a1e5780634487152f14610a6d576102e4565b806318160ddd116102a157806318160ddd14610710578063182df0f51461073b57806323323e031461076657806323b872dd146107d55780632608f8181461086857806326782247146108d7576102e4565b806306fdde03146103ed5780630933c1ed1461047d578063095ea7b3146105be5780630e75270214610631578063173b99041461068057806317bfdfbc146106ab575b6000341461033d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526037815260200180613d5e6037913960400191505060405180910390fd5b6000601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600036604051808383808284378083019250505092505050600060405180830381855af49150503d80600081146103cb576040519150601f19603f3d011682016040523d82523d6000602084013e6103d0565b606091505b505090506040513d6000823e81600081146103e9573d82f35b3d82fd5b3480156103f957600080fd5b506104026116f6565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610442578082015181840152602081019050610427565b50505050905090810190601f16801561046f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561048957600080fd5b50610543600480360360208110156104a057600080fd5b81019080803590602001906401000000008111156104bd57600080fd5b8201836020820111156104cf57600080fd5b803590602001918460018302840111640100000000831117156104f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611794565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610583578082015181840152602081019050610568565b50505050905090810190601f1680156105b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105ca57600080fd5b50610617600480360360408110156105e157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506117c9565b604051808215151515815260200191505060405180910390f35b34801561063d57600080fd5b5061066a6004803603602081101561065457600080fd5b81019080803590602001909291905050506118c8565b6040518082815260200191505060405180910390f35b34801561068c57600080fd5b50610695611992565b6040518082815260200191505060405180910390f35b3480156106b757600080fd5b506106fa600480360360208110156106ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611998565b6040518082815260200191505060405180910390f35b34801561071c57600080fd5b50610725611a8e565b6040518082815260200191505060405180910390f35b34801561074757600080fd5b50610750611a94565b6040518082815260200191505060405180910390f35b34801561077257600080fd5b506107bf6004803603604081101561078957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611b51565b6040518082815260200191505060405180910390f35b3480156107e157600080fd5b5061084e600480360360608110156107f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611c50565b604051808215151515815260200191505060405180910390f35b34801561087457600080fd5b506108c16004803603604081101561088b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611d84565b6040518082815260200191505060405180910390f35b3480156108e357600080fd5b506108ec611e83565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561093a57600080fd5b50610943611ea9565b604051808260ff1660ff16815260200191505060405180910390f35b34801561096b57600080fd5b506109ae6004803603602081101561098257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ebc565b6040518082815260200191505060405180910390f35b3480156109d057600080fd5b506109d9611fb2565b6040518082815260200191505060405180910390f35b3480156109fb57600080fd5b50610a0461206f565b604051808215151515815260200191505060405180910390f35b348015610a2a57600080fd5b50610a5760048036036020811015610a4157600080fd5b8101908080359060200190929190505050612074565b6040518082815260200191505060405180910390f35b348015610a7957600080fd5b50610b3360048036036020811015610a9057600080fd5b8101908080359060200190640100000000811115610aad57600080fd5b820183602082011115610abf57600080fd5b80359060200191846001830284011164010000000083111715610ae157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061213e565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b73578082015181840152602081019050610b58565b50505050905090810190601f168015610ba05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610bba57600080fd5b50610bfd60048036036020811015610bd157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506123d4565b6040518082815260200191505060405180910390f35b348015610c1f57600080fd5b50610c286124ca565b6040518082815260200191505060405180910390f35b348015610c4a57600080fd5b50610d3060048036036060811015610c6157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080351515906020019092919080359060200190640100000000811115610caa57600080fd5b820183602082011115610cbc57600080fd5b80359060200191846001830284011164010000000083111715610cde57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506124d0565b005b348015610d3e57600080fd5b50610d4761282c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610d9557600080fd5b50610d9e612852565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610dec57600080fd5b50610e1960048036036020811015610e0357600080fd5b8101908080359060200190929190505050612878565b6040518082815260200191505060405180910390f35b348015610e3b57600080fd5b50610e44612942565b6040518082815260200191505060405180910390f35b348015610e6657600080fd5b50610e6f612948565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610ebd57600080fd5b50610f0060048036036020811015610ed457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061296e565b6040518082815260200191505060405180910390f35b348015610f2257600080fd5b50610f2b612a64565b6040518082815260200191505060405180910390f35b348015610f4d57600080fd5b50610f7a60048036036020811015610f6457600080fd5b8101908080359060200190929190505050612b21565b6040518082815260200191505060405180910390f35b348015610f9c57600080fd5b50610fa5612beb565b6040518082815260200191505060405180910390f35b348015610fc757600080fd5b50610fd0612bf1565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015611010578082015181840152602081019050610ff5565b50505050905090810190601f16801561103d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561105757600080fd5b5061109a6004803603602081101561106e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612c8f565b6040518082815260200191505060405180910390f35b3480156110bc57600080fd5b506110e9600480360360208110156110d357600080fd5b8101908080359060200190929190505050612d85565b6040518082815260200191505060405180910390f35b34801561110b57600080fd5b50611114612e4f565b6040518082815260200191505060405180910390f35b34801561113657600080fd5b506111836004803603604081101561114d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612f0c565b604051808215151515815260200191505060405180910390f35b3480156111a957600080fd5b506111b261300b565b6040518082815260200191505060405180910390f35b3480156111d457600080fd5b506111dd613011565b6040518082815260200191505060405180910390f35b3480156111ff57600080fd5b5061126c6004803603606081101561121657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506130ce565b6040518082815260200191505060405180910390f35b34801561128e57600080fd5b506112d1600480360360208110156112a557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613202565b6040518082815260200191505060405180910390f35b3480156112f357600080fd5b506112fc6132f8565b6040518082815260200191505060405180910390f35b34801561131e57600080fd5b506113616004803603602081101561133557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506133b5565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b34801561139857600080fd5b506113c5600480360360208110156113af57600080fd5b81019080803590602001909291905050506134d5565b6040518082815260200191505060405180910390f35b3480156113e757600080fd5b50611414600480360360208110156113fe57600080fd5b810190808035906020019092919050505061359f565b6040518082815260200191505060405180910390f35b34801561143657600080fd5b506114996004803603604081101561144d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613669565b6040518082815260200191505060405180910390f35b3480156114bb57600080fd5b506114c4613794565b6040518082815260200191505060405180910390f35b3480156114e657600080fd5b50611529600480360360208110156114fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613851565b6040518082815260200191505060405180910390f35b34801561154b57600080fd5b50611554613947565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156115a257600080fd5b5061160f600480360360608110156115b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061396d565b6040518082815260200191505060405180910390f35b34801561163157600080fd5b5061163a613aa1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561168857600080fd5b50611691613ac7565b6040518082815260200191505060405180910390f35b3480156116b357600080fd5b506116e0600480360360208110156116ca57600080fd5b8101908080359060200190929190505050613b84565b6040518082815260200191505060405180910390f35b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561178c5780601f106117615761010080835404028352916020019161178c565b820191906000526020600020905b81548152906001019060200180831161176f57829003601f168201915b505050505081565b60606117c2601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683613c4e565b9050919050565b600060606118978484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040527f095ea7b3000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b90508080602001905160208110156118ae57600080fd5b810190808051906020019092919050505091505092915050565b6000606061196283604051602401808281526020019150506040516020818303038152906040527f0e752702000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561197957600080fd5b8101908080519060200190929190505050915050919050565b60085481565b60006060611a5e83604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527f17bfdfbc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015611a7557600080fd5b8101908080519060200190929190505050915050919050565b600d5481565b60006060611b236040516024016040516020818303038152906040527f182df0f5000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b9050808060200190516020811015611b3a57600080fd5b810190808051906020019092919050505091505090565b60006060611c1f8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040527f23323e03000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015611c3657600080fd5b810190808051906020019092919050505091505092915050565b60006060611d52858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200193505050506040516020818303038152906040527f23b872dd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015611d6957600080fd5b81019080805190602001909291905050509150509392505050565b60006060611e528484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040527f2608f818000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015611e6957600080fd5b810190808051906020019092919050505091505092915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900460ff1681565b60006060611f8283604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527f3af9e669000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015611f9957600080fd5b8101908080519060200190929190505050915050919050565b600060606120416040516024016040516020818303038152906040527f3b1d21a2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b905080806020019051602081101561205857600080fd5b810190808051906020019092919050505091505090565b600181565b6000606061210e83604051602401808281526020019150506040516020818303038152906040527f3e941010000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561212557600080fd5b8101908080519060200190929190505050915050919050565b6060600060603073ffffffffffffffffffffffffffffffffffffffff16846040516024018080602001828103825283818151815260200191508051906020019080838360005b8381101561219f578082015181840152602081019050612184565b50505050905090810190601f1680156121cc5780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040527f0933c1ed000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b60208310612281578051825260208201915060208101905060208303925061225e565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146122e1576040519150601f19603f3d011682016040523d82523d6000602084013e6122e6565b606091505b509150915060008214156122fb573d60208201fd5b80806020019051602081101561231057600080fd5b810190808051604051939291908464010000000082111561233057600080fd5b8382019150602082018581111561234657600080fd5b825186600182028301116401000000008211171561236357600080fd5b8083526020830192505050908051906020019080838360005b8381101561239757808201518184015260208101905061237c565b50505050905090810190601f1680156123c45780820380516001836020036101000a031916815260200191505b5060405250505092505050919050565b6000606061249a83604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527f4576b5db000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b90508080602001905160208110156124b157600080fd5b8101908080519060200190929190505050915050919050565b600b5481565b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612576576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526039815260200180613d256039913960400191505060405180910390fd5b8115612609576126076040516024016040516020818303038152906040527f153ab505000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b505b6000601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083601260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061276c826040516024018080602001828103825283818151815260200191508051906020019080838360005b838110156126b857808201518184015260208101905061269d565b50505050905090810190601f1680156126e55780820380516001836020036101000a031916815260200191505b50925050506040516020818303038152906040527f56e67728000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b507fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a81601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150505050565b601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000606061291283604051602401808281526020019150506040516020818303038152906040527f601a0bf1000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561292957600080fd5b8101908080519060200190929190505050915050919050565b60095481565b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006060612a3483604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527f70a08231000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b9050808060200190516020811015612a4b57600080fd5b8101908080519060200190929190505050915050919050565b60006060612af36040516024016040516020818303038152906040527f73acee98000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015612b0a57600080fd5b810190808051906020019092919050505091505090565b60006060612bbb83604051602401808281526020019150506040516020818303038152906040527f852a12e3000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015612bd257600080fd5b8101908080519060200190929190505050915050919050565b600c5481565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612c875780601f10612c5c57610100808354040283529160200191612c87565b820191906000526020600020905b815481529060010190602001808311612c6a57829003601f168201915b505050505081565b60006060612d5583604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527f95dd9193000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b9050808060200190516020811015612d6c57600080fd5b8101908080519060200190929190505050915050919050565b60006060612e1f83604051602401808281526020019150506040516020818303038152906040527fa0712d68000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015612e3657600080fd5b8101908080519060200190929190505050915050919050565b60006060612ede6040516024016040516020818303038152906040527fa6afed95000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015612ef557600080fd5b810190808051906020019092919050505091505090565b60006060612fda8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040527fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015612ff157600080fd5b810190808051906020019092919050505091505092915050565b600a5481565b600060606130a06040516024016040516020818303038152906040527fae9d70b0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b90508080602001905160208110156130b757600080fd5b810190808051906020019092919050505091505090565b600060606131d0858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200193505050506040516020818303038152906040527fb2a02ff1000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b90508080602001905160208110156131e757600080fd5b81019080805190602001909291905050509150509392505050565b600060606132c883604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527fb71d1a0c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b90508080602001905160208110156132df57600080fd5b8101908080519060200190929190505050915050919050565b600060606133876040516024016040516020818303038152906040527fbd6d894d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561339e57600080fd5b810190808051906020019092919050505091505090565b600080600080606061347f86604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527fc37f68e2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b905080806020019051608081101561349657600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291905050509450945094509450509193509193565b6000606061356f83604051602401808281526020019150506040516020818303038152906040527fc5ebeaec000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561358657600080fd5b8101908080519060200190929190505050915050919050565b6000606061363983604051602401808281526020019150506040516020818303038152906040527fdb006a75000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561365057600080fd5b8101908080519060200190929190505050915050919050565b600060606137638484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001925050506040516020818303038152906040527fdd62ed3e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b905080806020019051602081101561377a57600080fd5b810190808051906020019092919050505091505092915050565b600060606138236040516024016040516020818303038152906040527fe9c714f2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561383a57600080fd5b810190808051906020019092919050505091505090565b6000606061391783604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019150506040516020818303038152906040527ff2b3abbd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b905080806020019051602081101561392e57600080fd5b8101908080519060200190929190505050915050919050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006060613a6f858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200193505050506040516020818303038152906040527ff5e3c462000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015613a8657600080fd5b81019080805190602001909291905050509150509392505050565b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006060613b566040516024016040516020818303038152906040527ff8f9da28000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061213e565b9050808060200190516020811015613b6d57600080fd5b810190808051906020019092919050505091505090565b60006060613c1e83604051602401808281526020019150506040516020818303038152906040527ffca7820b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611794565b9050808060200190516020811015613c3557600080fd5b8101908080519060200190929190505050915050919050565b6060600060608473ffffffffffffffffffffffffffffffffffffffff16846040518082805190602001908083835b60208310613c9f5780518252602082019150602081019050602083039250613c7c565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114613cff576040519150601f19603f3d011682016040523d82523d6000602084013e613d04565b606091505b50915091506000821415613d19573d60208201fd5b80925050509291505056fe56426570323044656c656761746f723a3a5f736574496d706c656d656e746174696f6e3a2043616c6c6572206d7573742062652061646d696e56426570323044656c656761746f723a66616c6c6261636b3a2063616e6e6f742073656e642076616c756520746f2066616c6c6261636ba265627a7a723158206ea6df130da33001517db93ff70eade07339f0e0cb53b7506f4dccf105ccd54c64736f6c6343000510003256426570323044656c656761746f723a3a5f736574496d706c656d656e746174696f6e3a2043616c6c6572206d7573742062652061646d696e"

// DeployVenusbep20delegator deploys a new Ethereum contract, binding an instance of Venusbep20delegator to it.
func DeployVenusbep20delegator(auth *bind.TransactOpts, backend bind.ContractBackend, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8, admin_ common.Address, implementation_ common.Address, becomeImplementationData []byte) (common.Address, *types.Transaction, *Venusbep20delegator, error) {
	parsed, err := abi.JSON(strings.NewReader(Venusbep20delegatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Venusbep20delegatorBin), backend, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_, admin_, implementation_, becomeImplementationData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Venusbep20delegator{Venusbep20delegatorCaller: Venusbep20delegatorCaller{contract: contract}, Venusbep20delegatorTransactor: Venusbep20delegatorTransactor{contract: contract}, Venusbep20delegatorFilterer: Venusbep20delegatorFilterer{contract: contract}}, nil
}

// Venusbep20delegator is an auto generated Go binding around an Ethereum contract.
type Venusbep20delegator struct {
	Venusbep20delegatorCaller     // Read-only binding to the contract
	Venusbep20delegatorTransactor // Write-only binding to the contract
	Venusbep20delegatorFilterer   // Log filterer for contract events
}

// Venusbep20delegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type Venusbep20delegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Venusbep20delegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Venusbep20delegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Venusbep20delegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Venusbep20delegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Venusbep20delegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Venusbep20delegatorSession struct {
	Contract     *Venusbep20delegator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Venusbep20delegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Venusbep20delegatorCallerSession struct {
	Contract *Venusbep20delegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Venusbep20delegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Venusbep20delegatorTransactorSession struct {
	Contract     *Venusbep20delegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Venusbep20delegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type Venusbep20delegatorRaw struct {
	Contract *Venusbep20delegator // Generic contract binding to access the raw methods on
}

// Venusbep20delegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Venusbep20delegatorCallerRaw struct {
	Contract *Venusbep20delegatorCaller // Generic read-only contract binding to access the raw methods on
}

// Venusbep20delegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Venusbep20delegatorTransactorRaw struct {
	Contract *Venusbep20delegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusbep20delegator creates a new instance of Venusbep20delegator, bound to a specific deployed contract.
func NewVenusbep20delegator(address common.Address, backend bind.ContractBackend) (*Venusbep20delegator, error) {
	contract, err := bindVenusbep20delegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegator{Venusbep20delegatorCaller: Venusbep20delegatorCaller{contract: contract}, Venusbep20delegatorTransactor: Venusbep20delegatorTransactor{contract: contract}, Venusbep20delegatorFilterer: Venusbep20delegatorFilterer{contract: contract}}, nil
}

// NewVenusbep20delegatorCaller creates a new read-only instance of Venusbep20delegator, bound to a specific deployed contract.
func NewVenusbep20delegatorCaller(address common.Address, caller bind.ContractCaller) (*Venusbep20delegatorCaller, error) {
	contract, err := bindVenusbep20delegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorCaller{contract: contract}, nil
}

// NewVenusbep20delegatorTransactor creates a new write-only instance of Venusbep20delegator, bound to a specific deployed contract.
func NewVenusbep20delegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*Venusbep20delegatorTransactor, error) {
	contract, err := bindVenusbep20delegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorTransactor{contract: contract}, nil
}

// NewVenusbep20delegatorFilterer creates a new log filterer instance of Venusbep20delegator, bound to a specific deployed contract.
func NewVenusbep20delegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*Venusbep20delegatorFilterer, error) {
	contract, err := bindVenusbep20delegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorFilterer{contract: contract}, nil
}

// bindVenusbep20delegator binds a generic wrapper to an already deployed contract.
func bindVenusbep20delegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Venusbep20delegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venusbep20delegator *Venusbep20delegatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venusbep20delegator.Contract.Venusbep20delegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venusbep20delegator *Venusbep20delegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Venusbep20delegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venusbep20delegator *Venusbep20delegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Venusbep20delegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venusbep20delegator *Venusbep20delegatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venusbep20delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venusbep20delegator *Venusbep20delegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venusbep20delegator *Venusbep20delegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) AccrualBlockNumber() (*big.Int, error) {
	return _Venusbep20delegator.Contract.AccrualBlockNumber(&_Venusbep20delegator.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Venusbep20delegator.Contract.AccrualBlockNumber(&_Venusbep20delegator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) Admin() (common.Address, error) {
	return _Venusbep20delegator.Contract.Admin(&_Venusbep20delegator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Admin() (common.Address, error) {
	return _Venusbep20delegator.Contract.Admin(&_Venusbep20delegator.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.Allowance(&_Venusbep20delegator.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.Allowance(&_Venusbep20delegator.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.BalanceOf(&_Venusbep20delegator.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.BalanceOf(&_Venusbep20delegator.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowBalanceStored(&_Venusbep20delegator.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowBalanceStored(&_Venusbep20delegator.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BorrowIndex() (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowIndex(&_Venusbep20delegator.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) BorrowIndex() (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowIndex(&_Venusbep20delegator.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowRatePerBlock(&_Venusbep20delegator.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Venusbep20delegator.Contract.BorrowRatePerBlock(&_Venusbep20delegator.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) Comptroller() (common.Address, error) {
	return _Venusbep20delegator.Contract.Comptroller(&_Venusbep20delegator.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Comptroller() (common.Address, error) {
	return _Venusbep20delegator.Contract.Comptroller(&_Venusbep20delegator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Venusbep20delegator *Venusbep20delegatorSession) Decimals() (uint8, error) {
	return _Venusbep20delegator.Contract.Decimals(&_Venusbep20delegator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Decimals() (uint8, error) {
	return _Venusbep20delegator.Contract.Decimals(&_Venusbep20delegator.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorCaller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Venusbep20delegator.Contract.DelegateToViewImplementation(&_Venusbep20delegator.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Venusbep20delegator.Contract.DelegateToViewImplementation(&_Venusbep20delegator.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) ExchangeRateStored() (*big.Int, error) {
	return _Venusbep20delegator.Contract.ExchangeRateStored(&_Venusbep20delegator.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Venusbep20delegator.Contract.ExchangeRateStored(&_Venusbep20delegator.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Venusbep20delegator.Contract.GetAccountSnapshot(&_Venusbep20delegator.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Venusbep20delegator.Contract.GetAccountSnapshot(&_Venusbep20delegator.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) GetCash() (*big.Int, error) {
	return _Venusbep20delegator.Contract.GetCash(&_Venusbep20delegator.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) GetCash() (*big.Int, error) {
	return _Venusbep20delegator.Contract.GetCash(&_Venusbep20delegator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) Implementation() (common.Address, error) {
	return _Venusbep20delegator.Contract.Implementation(&_Venusbep20delegator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Implementation() (common.Address, error) {
	return _Venusbep20delegator.Contract.Implementation(&_Venusbep20delegator.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) InterestRateModel() (common.Address, error) {
	return _Venusbep20delegator.Contract.InterestRateModel(&_Venusbep20delegator.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) InterestRateModel() (common.Address, error) {
	return _Venusbep20delegator.Contract.InterestRateModel(&_Venusbep20delegator.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorCaller) IsVToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "isVToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorSession) IsVToken() (bool, error) {
	return _Venusbep20delegator.Contract.IsVToken(&_Venusbep20delegator.CallOpts)
}

// IsVToken is a free data retrieval call binding the contract method 0x3d9ea3a1.
//
// Solidity: function isVToken() view returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) IsVToken() (bool, error) {
	return _Venusbep20delegator.Contract.IsVToken(&_Venusbep20delegator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorSession) Name() (string, error) {
	return _Venusbep20delegator.Contract.Name(&_Venusbep20delegator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Name() (string, error) {
	return _Venusbep20delegator.Contract.Name(&_Venusbep20delegator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) PendingAdmin() (common.Address, error) {
	return _Venusbep20delegator.Contract.PendingAdmin(&_Venusbep20delegator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) PendingAdmin() (common.Address, error) {
	return _Venusbep20delegator.Contract.PendingAdmin(&_Venusbep20delegator.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Venusbep20delegator.Contract.ReserveFactorMantissa(&_Venusbep20delegator.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Venusbep20delegator.Contract.ReserveFactorMantissa(&_Venusbep20delegator.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Venusbep20delegator.Contract.SupplyRatePerBlock(&_Venusbep20delegator.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Venusbep20delegator.Contract.SupplyRatePerBlock(&_Venusbep20delegator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorSession) Symbol() (string, error) {
	return _Venusbep20delegator.Contract.Symbol(&_Venusbep20delegator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Symbol() (string, error) {
	return _Venusbep20delegator.Contract.Symbol(&_Venusbep20delegator.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) TotalBorrows() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalBorrows(&_Venusbep20delegator.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) TotalBorrows() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalBorrows(&_Venusbep20delegator.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) TotalReserves() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalReserves(&_Venusbep20delegator.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) TotalReserves() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalReserves(&_Venusbep20delegator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) TotalSupply() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalSupply(&_Venusbep20delegator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) TotalSupply() (*big.Int, error) {
	return _Venusbep20delegator.Contract.TotalSupply(&_Venusbep20delegator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venusbep20delegator.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorSession) Underlying() (common.Address, error) {
	return _Venusbep20delegator.Contract.Underlying(&_Venusbep20delegator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Venusbep20delegator *Venusbep20delegatorCallerSession) Underlying() (common.Address, error) {
	return _Venusbep20delegator.Contract.Underlying(&_Venusbep20delegator.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AcceptAdmin(&_Venusbep20delegator.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AcceptAdmin(&_Venusbep20delegator.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AddReserves(&_Venusbep20delegator.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AddReserves(&_Venusbep20delegator.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.ReduceReserves(&_Venusbep20delegator.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.ReduceReserves(&_Venusbep20delegator.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetComptroller(&_Venusbep20delegator.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetComptroller(&_Venusbep20delegator.TransactOpts, newComptroller)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Venusbep20delegator *Venusbep20delegatorTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Venusbep20delegator *Venusbep20delegatorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetImplementation(&_Venusbep20delegator.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetImplementation(&_Venusbep20delegator.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetInterestRateModel(&_Venusbep20delegator.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetInterestRateModel(&_Venusbep20delegator.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetPendingAdmin(&_Venusbep20delegator.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetPendingAdmin(&_Venusbep20delegator.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetReserveFactor(&_Venusbep20delegator.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.SetReserveFactor(&_Venusbep20delegator.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) AccrueInterest() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AccrueInterest(&_Venusbep20delegator.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.AccrueInterest(&_Venusbep20delegator.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Approve(&_Venusbep20delegator.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Approve(&_Venusbep20delegator.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.BalanceOfUnderlying(&_Venusbep20delegator.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.BalanceOfUnderlying(&_Venusbep20delegator.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Borrow(&_Venusbep20delegator.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Borrow(&_Venusbep20delegator.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.BorrowBalanceCurrent(&_Venusbep20delegator.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.BorrowBalanceCurrent(&_Venusbep20delegator.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.DelegateToImplementation(&_Venusbep20delegator.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.DelegateToImplementation(&_Venusbep20delegator.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.ExchangeRateCurrent(&_Venusbep20delegator.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.ExchangeRateCurrent(&_Venusbep20delegator.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.LiquidateBorrow(&_Venusbep20delegator.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address vTokenCollateral) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, vTokenCollateral common.Address) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.LiquidateBorrow(&_Venusbep20delegator.TransactOpts, borrower, repayAmount, vTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Mint(&_Venusbep20delegator.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Mint(&_Venusbep20delegator.TransactOpts, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) MintBehalf(opts *bind.TransactOpts, receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "mintBehalf", receiver, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) MintBehalf(receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.MintBehalf(&_Venusbep20delegator.TransactOpts, receiver, mintAmount)
}

// MintBehalf is a paid mutator transaction binding the contract method 0x23323e03.
//
// Solidity: function mintBehalf(address receiver, uint256 mintAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) MintBehalf(receiver common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.MintBehalf(&_Venusbep20delegator.TransactOpts, receiver, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Redeem(&_Venusbep20delegator.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Redeem(&_Venusbep20delegator.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RedeemUnderlying(&_Venusbep20delegator.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RedeemUnderlying(&_Venusbep20delegator.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RepayBorrow(&_Venusbep20delegator.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RepayBorrow(&_Venusbep20delegator.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RepayBorrowBehalf(&_Venusbep20delegator.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.RepayBorrowBehalf(&_Venusbep20delegator.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Seize(&_Venusbep20delegator.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Seize(&_Venusbep20delegator.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.TotalBorrowsCurrent(&_Venusbep20delegator.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.TotalBorrowsCurrent(&_Venusbep20delegator.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Transfer(&_Venusbep20delegator.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Transfer(&_Venusbep20delegator.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.TransferFrom(&_Venusbep20delegator.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.TransferFrom(&_Venusbep20delegator.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Venusbep20delegator *Venusbep20delegatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Venusbep20delegator *Venusbep20delegatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Fallback(&_Venusbep20delegator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Venusbep20delegator *Venusbep20delegatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Venusbep20delegator.Contract.Fallback(&_Venusbep20delegator.TransactOpts, calldata)
}

// Venusbep20delegatorAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Venusbep20delegator contract.
type Venusbep20delegatorAccrueInterestIterator struct {
	Event *Venusbep20delegatorAccrueInterest // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorAccrueInterest)
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
		it.Event = new(Venusbep20delegatorAccrueInterest)
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
func (it *Venusbep20delegatorAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorAccrueInterest represents a AccrueInterest event raised by the Venusbep20delegator contract.
type Venusbep20delegatorAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*Venusbep20delegatorAccrueInterestIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorAccrueInterestIterator{contract: _Venusbep20delegator.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorAccrueInterest)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseAccrueInterest(log types.Log) (*Venusbep20delegatorAccrueInterest, error) {
	event := new(Venusbep20delegatorAccrueInterest)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Venusbep20delegator contract.
type Venusbep20delegatorApprovalIterator struct {
	Event *Venusbep20delegatorApproval // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorApproval)
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
		it.Event = new(Venusbep20delegatorApproval)
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
func (it *Venusbep20delegatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorApproval represents a Approval event raised by the Venusbep20delegator contract.
type Venusbep20delegatorApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Venusbep20delegatorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorApprovalIterator{contract: _Venusbep20delegator.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorApproval)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseApproval(log types.Log) (*Venusbep20delegatorApproval, error) {
	event := new(Venusbep20delegatorApproval)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Venusbep20delegator contract.
type Venusbep20delegatorBorrowIterator struct {
	Event *Venusbep20delegatorBorrow // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorBorrow)
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
		it.Event = new(Venusbep20delegatorBorrow)
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
func (it *Venusbep20delegatorBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorBorrow represents a Borrow event raised by the Venusbep20delegator contract.
type Venusbep20delegatorBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterBorrow(opts *bind.FilterOpts) (*Venusbep20delegatorBorrowIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorBorrowIterator{contract: _Venusbep20delegator.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorBorrow) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorBorrow)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseBorrow(log types.Log) (*Venusbep20delegatorBorrow, error) {
	event := new(Venusbep20delegatorBorrow)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Venusbep20delegator contract.
type Venusbep20delegatorFailureIterator struct {
	Event *Venusbep20delegatorFailure // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorFailure)
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
		it.Event = new(Venusbep20delegatorFailure)
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
func (it *Venusbep20delegatorFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorFailure represents a Failure event raised by the Venusbep20delegator contract.
type Venusbep20delegatorFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterFailure(opts *bind.FilterOpts) (*Venusbep20delegatorFailureIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorFailureIterator{contract: _Venusbep20delegator.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorFailure) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorFailure)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseFailure(log types.Log) (*Venusbep20delegatorFailure, error) {
	event := new(Venusbep20delegatorFailure)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Venusbep20delegator contract.
type Venusbep20delegatorLiquidateBorrowIterator struct {
	Event *Venusbep20delegatorLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorLiquidateBorrow)
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
		it.Event = new(Venusbep20delegatorLiquidateBorrow)
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
func (it *Venusbep20delegatorLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorLiquidateBorrow represents a LiquidateBorrow event raised by the Venusbep20delegator contract.
type Venusbep20delegatorLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	VTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*Venusbep20delegatorLiquidateBorrowIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorLiquidateBorrowIterator{contract: _Venusbep20delegator.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorLiquidateBorrow)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address vTokenCollateral, uint256 seizeTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseLiquidateBorrow(log types.Log) (*Venusbep20delegatorLiquidateBorrow, error) {
	event := new(Venusbep20delegatorLiquidateBorrow)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Venusbep20delegator contract.
type Venusbep20delegatorMintIterator struct {
	Event *Venusbep20delegatorMint // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorMint)
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
		it.Event = new(Venusbep20delegatorMint)
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
func (it *Venusbep20delegatorMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorMint represents a Mint event raised by the Venusbep20delegator contract.
type Venusbep20delegatorMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterMint(opts *bind.FilterOpts) (*Venusbep20delegatorMintIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorMintIterator{contract: _Venusbep20delegator.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorMint) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorMint)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseMint(log types.Log) (*Venusbep20delegatorMint, error) {
	event := new(Venusbep20delegatorMint)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorMintBehalfIterator is returned from FilterMintBehalf and is used to iterate over the raw logs and unpacked data for MintBehalf events raised by the Venusbep20delegator contract.
type Venusbep20delegatorMintBehalfIterator struct {
	Event *Venusbep20delegatorMintBehalf // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorMintBehalfIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorMintBehalf)
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
		it.Event = new(Venusbep20delegatorMintBehalf)
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
func (it *Venusbep20delegatorMintBehalfIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorMintBehalfIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorMintBehalf represents a MintBehalf event raised by the Venusbep20delegator contract.
type Venusbep20delegatorMintBehalf struct {
	Payer      common.Address
	Receiver   common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintBehalf is a free log retrieval operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterMintBehalf(opts *bind.FilterOpts) (*Venusbep20delegatorMintBehalfIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorMintBehalfIterator{contract: _Venusbep20delegator.contract, event: "MintBehalf", logs: logs, sub: sub}, nil
}

// WatchMintBehalf is a free log subscription operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchMintBehalf(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorMintBehalf) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "MintBehalf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorMintBehalf)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "MintBehalf", log); err != nil {
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

// ParseMintBehalf is a log parse operation binding the contract event 0x297989b84a5f5b82d2ee0c266504c19bd9b10b410f187dc72ca4b0f0faecb345.
//
// Solidity: event MintBehalf(address payer, address receiver, uint256 mintAmount, uint256 mintTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseMintBehalf(log types.Log) (*Venusbep20delegatorMintBehalf, error) {
	event := new(Venusbep20delegatorMintBehalf)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "MintBehalf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewAdminIterator struct {
	Event *Venusbep20delegatorNewAdmin // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewAdmin)
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
		it.Event = new(Venusbep20delegatorNewAdmin)
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
func (it *Venusbep20delegatorNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewAdmin represents a NewAdmin event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*Venusbep20delegatorNewAdminIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewAdminIterator{contract: _Venusbep20delegator.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewAdmin)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewAdmin(log types.Log) (*Venusbep20delegatorNewAdmin, error) {
	event := new(Venusbep20delegatorNewAdmin)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewComptrollerIterator struct {
	Event *Venusbep20delegatorNewComptroller // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewComptroller)
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
		it.Event = new(Venusbep20delegatorNewComptroller)
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
func (it *Venusbep20delegatorNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewComptroller represents a NewComptroller event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*Venusbep20delegatorNewComptrollerIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewComptrollerIterator{contract: _Venusbep20delegator.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewComptroller)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewComptroller(log types.Log) (*Venusbep20delegatorNewComptroller, error) {
	event := new(Venusbep20delegatorNewComptroller)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewImplementationIterator struct {
	Event *Venusbep20delegatorNewImplementation // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewImplementation)
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
		it.Event = new(Venusbep20delegatorNewImplementation)
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
func (it *Venusbep20delegatorNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewImplementation represents a NewImplementation event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*Venusbep20delegatorNewImplementationIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewImplementationIterator{contract: _Venusbep20delegator.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewImplementation) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewImplementation)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewImplementation(log types.Log) (*Venusbep20delegatorNewImplementation, error) {
	event := new(Venusbep20delegatorNewImplementation)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewMarketInterestRateModelIterator struct {
	Event *Venusbep20delegatorNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewMarketInterestRateModel)
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
		it.Event = new(Venusbep20delegatorNewMarketInterestRateModel)
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
func (it *Venusbep20delegatorNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*Venusbep20delegatorNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewMarketInterestRateModelIterator{contract: _Venusbep20delegator.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewMarketInterestRateModel)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewMarketInterestRateModel(log types.Log) (*Venusbep20delegatorNewMarketInterestRateModel, error) {
	event := new(Venusbep20delegatorNewMarketInterestRateModel)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewPendingAdminIterator struct {
	Event *Venusbep20delegatorNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewPendingAdmin)
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
		it.Event = new(Venusbep20delegatorNewPendingAdmin)
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
func (it *Venusbep20delegatorNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewPendingAdmin represents a NewPendingAdmin event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*Venusbep20delegatorNewPendingAdminIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewPendingAdminIterator{contract: _Venusbep20delegator.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewPendingAdmin)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewPendingAdmin(log types.Log) (*Venusbep20delegatorNewPendingAdmin, error) {
	event := new(Venusbep20delegatorNewPendingAdmin)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewReserveFactorIterator struct {
	Event *Venusbep20delegatorNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorNewReserveFactor)
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
		it.Event = new(Venusbep20delegatorNewReserveFactor)
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
func (it *Venusbep20delegatorNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorNewReserveFactor represents a NewReserveFactor event raised by the Venusbep20delegator contract.
type Venusbep20delegatorNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*Venusbep20delegatorNewReserveFactorIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorNewReserveFactorIterator{contract: _Venusbep20delegator.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorNewReserveFactor)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseNewReserveFactor(log types.Log) (*Venusbep20delegatorNewReserveFactor, error) {
	event := new(Venusbep20delegatorNewReserveFactor)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Venusbep20delegator contract.
type Venusbep20delegatorRedeemIterator struct {
	Event *Venusbep20delegatorRedeem // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorRedeem)
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
		it.Event = new(Venusbep20delegatorRedeem)
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
func (it *Venusbep20delegatorRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorRedeem represents a Redeem event raised by the Venusbep20delegator contract.
type Venusbep20delegatorRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterRedeem(opts *bind.FilterOpts) (*Venusbep20delegatorRedeemIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorRedeemIterator{contract: _Venusbep20delegator.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorRedeem) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorRedeem)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseRedeem(log types.Log) (*Venusbep20delegatorRedeem, error) {
	event := new(Venusbep20delegatorRedeem)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorRedeemFeeIterator is returned from FilterRedeemFee and is used to iterate over the raw logs and unpacked data for RedeemFee events raised by the Venusbep20delegator contract.
type Venusbep20delegatorRedeemFeeIterator struct {
	Event *Venusbep20delegatorRedeemFee // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorRedeemFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorRedeemFee)
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
		it.Event = new(Venusbep20delegatorRedeemFee)
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
func (it *Venusbep20delegatorRedeemFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorRedeemFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorRedeemFee represents a RedeemFee event raised by the Venusbep20delegator contract.
type Venusbep20delegatorRedeemFee struct {
	Redeemer     common.Address
	FeeAmount    *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemFee is a free log retrieval operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterRedeemFee(opts *bind.FilterOpts) (*Venusbep20delegatorRedeemFeeIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorRedeemFeeIterator{contract: _Venusbep20delegator.contract, event: "RedeemFee", logs: logs, sub: sub}, nil
}

// WatchRedeemFee is a free log subscription operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchRedeemFee(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorRedeemFee) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "RedeemFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorRedeemFee)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "RedeemFee", log); err != nil {
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

// ParseRedeemFee is a log parse operation binding the contract event 0xccf8e53b86a99b7e9ecf796342c165764d66154780f638c08e6241d711fba6d4.
//
// Solidity: event RedeemFee(address redeemer, uint256 feeAmount, uint256 redeemTokens)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseRedeemFee(log types.Log) (*Venusbep20delegatorRedeemFee, error) {
	event := new(Venusbep20delegatorRedeemFee)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "RedeemFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Venusbep20delegator contract.
type Venusbep20delegatorRepayBorrowIterator struct {
	Event *Venusbep20delegatorRepayBorrow // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorRepayBorrow)
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
		it.Event = new(Venusbep20delegatorRepayBorrow)
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
func (it *Venusbep20delegatorRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorRepayBorrow represents a RepayBorrow event raised by the Venusbep20delegator contract.
type Venusbep20delegatorRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*Venusbep20delegatorRepayBorrowIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorRepayBorrowIterator{contract: _Venusbep20delegator.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorRepayBorrow)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseRepayBorrow(log types.Log) (*Venusbep20delegatorRepayBorrow, error) {
	event := new(Venusbep20delegatorRepayBorrow)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Venusbep20delegator contract.
type Venusbep20delegatorReservesAddedIterator struct {
	Event *Venusbep20delegatorReservesAdded // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorReservesAdded)
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
		it.Event = new(Venusbep20delegatorReservesAdded)
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
func (it *Venusbep20delegatorReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorReservesAdded represents a ReservesAdded event raised by the Venusbep20delegator contract.
type Venusbep20delegatorReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*Venusbep20delegatorReservesAddedIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorReservesAddedIterator{contract: _Venusbep20delegator.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorReservesAdded)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseReservesAdded(log types.Log) (*Venusbep20delegatorReservesAdded, error) {
	event := new(Venusbep20delegatorReservesAdded)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Venusbep20delegator contract.
type Venusbep20delegatorReservesReducedIterator struct {
	Event *Venusbep20delegatorReservesReduced // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorReservesReduced)
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
		it.Event = new(Venusbep20delegatorReservesReduced)
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
func (it *Venusbep20delegatorReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorReservesReduced represents a ReservesReduced event raised by the Venusbep20delegator contract.
type Venusbep20delegatorReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*Venusbep20delegatorReservesReducedIterator, error) {

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorReservesReducedIterator{contract: _Venusbep20delegator.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorReservesReduced)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseReservesReduced(log types.Log) (*Venusbep20delegatorReservesReduced, error) {
	event := new(Venusbep20delegatorReservesReduced)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Venusbep20delegatorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Venusbep20delegator contract.
type Venusbep20delegatorTransferIterator struct {
	Event *Venusbep20delegatorTransfer // Event containing the contract specifics and raw log

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
func (it *Venusbep20delegatorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Venusbep20delegatorTransfer)
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
		it.Event = new(Venusbep20delegatorTransfer)
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
func (it *Venusbep20delegatorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Venusbep20delegatorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Venusbep20delegatorTransfer represents a Transfer event raised by the Venusbep20delegator contract.
type Venusbep20delegatorTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Venusbep20delegatorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Venusbep20delegator.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Venusbep20delegatorTransferIterator{contract: _Venusbep20delegator.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Venusbep20delegatorTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Venusbep20delegator.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Venusbep20delegatorTransfer)
				if err := _Venusbep20delegator.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Venusbep20delegator *Venusbep20delegatorFilterer) ParseTransfer(log types.Log) (*Venusbep20delegatorTransfer, error) {
	event := new(Venusbep20delegatorTransfer)
	if err := _Venusbep20delegator.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
