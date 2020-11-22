package main

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/opyn"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/uniswapv1"
)

var (
	clientOpynCommand                           = clientToken2Command.Command("opyn", "Opyn operations")
	clientOpynOTokenAddress                     = clientToken2Command.Flag("otoken", "OToken Address").String()
	clientOpynListOptionsContracts              = clientOpynCommand.Command("list-options-contracts", "List options contracts")
	clientOpynListOptionsContractsActive        = clientOpynListOptionsContracts.Flag("active", "Only show active contracts").Bool()
	clientOpynOpenVaultCommand                  = clientOpynCommand.Command("open-vault", "Open a vault")
	clientOpynRedeemVaultBalanceCommand         = clientOpynCommand.Command("redeem-vault-balance", "Redeem vault balance")
	clientOpynEstimateCommand                   = clientOpynCommand.Command("estimate", "Estimate contract operation")
	clientOpynEstimateAmount                    = clientOpynEstimateCommand.Arg("amount", "Amount of collateral").Required().String()
	clientOpynAddCollateralOption               = clientOpynCommand.Command("add-collateral-option", "Add collateral Option")
	clientOpynAddCollateralAmtToCreate          = clientOpynAddCollateralOption.Arg("amt-to-create", "Amount of otoken to create").Required().String()
	clientOpynAddCollateralCollateralAmt        = clientOpynAddCollateralOption.Arg("collateral-amt", "Collateral amount").Required().String()
	clientOpynAddCollateralReceiver             = clientOpynAddCollateralOption.Arg("receiver", "Address of receiver").Required().String()
	clientOpynAddAndSellCollateralOption        = clientOpynCommand.Command("add-and-sell-collateral-option", "Add and sell collateral Option")
	clientOpynAddAndSellCollateralAmtToCreate   = clientOpynAddAndSellCollateralOption.Arg("amt-to-create", "Amount of otoken to create").Required().String()
	clientOpynAddAndSellCollateralCollateralAmt = clientOpynAddAndSellCollateralOption.Arg("collateral-amt", "Collateral amount").Required().String()
	clientOpynAddAndSellCollateralReceiver      = clientOpynAddAndSellCollateralOption.Arg("receiver", "Address of receiver").Required().String()
	clientOpynGetVault                          = clientOpynCommand.Command("get-vault", "Get info about the given vault")
)

func opynCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	opynClient, err := opyn.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "list-options-contracts":
		return true, opynListOptionsContracts(ctx, opynClient, reg)
	case "open-vault":
		return true, opynOpenVault(ctx, opynClient, reg)
	case "redeem-vault-balance":
		return true, opynRedeemVaultBalance(ctx, opynClient, reg)
	case "estimate":
		return true, opynEstimateContract(ctx, opynClient, reg)
	case "add-collateral-option":
		return true, opynAddCollateralOption(ctx, opynClient, reg)
	case "add-and-sell-collateral-option":
		return true, opynAddAndSellCollateralOption(ctx, opynClient, reg)
	case "get-vault":
		return true, opynGetVault(ctx, opynClient, reg)
	}
	return false, nil
}

func getOToken(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) (*opyn.OTokenClient, error) {
	if *clientOpynOTokenAddress == "" {
		return nil, errors.New("--otoken flag is required")
	}
	var optionsContractAddress common.Address
	if strings.HasPrefix(*clientOpynOTokenAddress, "0x") {
		// Contract address.
		optionsContractAddress = common.HexToAddress(*clientOpynOTokenAddress)
	} else {
		val, err := strconv.ParseInt(*clientOpynOTokenAddress, 10, 64)
		if err == nil {
			// Contract index.
			optionsContractAddress, err = opynClient.OptionsContract(ctx, big.NewInt(val))
			if err != nil {
				return nil, err
			}
		} else {
			tok, err := reg.ByName(*clientOpynOTokenAddress)
			if err != nil {
				return nil, err
			}
			optionsContractAddress = tok.Address
		}
	}
	return opynClient.GetOToken(ctx, optionsContractAddress)
}

func opynListOptionsContracts(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	numberOfOptionsContracts, err := opynClient.GetNumberOfOptionsContracts(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Number of options contracts: %v\n", numberOfOptionsContracts)
	for i := int64(0); i < numberOfOptionsContracts.Int64(); i++ {
		optionsContractAddress, err := opynClient.OptionsContract(ctx, big.NewInt(i))
		if err != nil {
			return err
		}
		otoken, err := opynClient.GetOToken(ctx, optionsContractAddress)
		if err != nil {
			return err
		}
		hasExpired, err := otoken.HasExpired(ctx)
		if err != nil {
			return err
		}
		name, err := otoken.Name(ctx)
		if err != nil {
			return err
		}
		if *clientOpynListOptionsContractsActive && (hasExpired || name == "") {
			continue
		}
		info, err := otoken.ContractInfo(ctx)
		if err != nil {
			return err
		}
		currencyName := info.Currency.Hex()
		if info.Currency == opyn.EthContract {
			currencyName = "eth"
		} else {
			name, _, err := reg.ByAddress(info.Currency)
			if err == nil {
				currencyName = name
			}
		}
		status := ""
		if hasExpired {
			status = " EXPIRED"
		}
		fmt.Printf("%d %s %s %s: %s %s (%s)%s\n", i, currencyName, info.Type, info.StrikePrice, info.Name, info.Expiry, info.Address.Hex(), status)
	}
	return nil
}

func opynOpenVault(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet is locked")
	}
	tx, err := otoken.OpenVault(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())
	return nil
}

func opynRedeemVaultBalance(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet is locked")
	}
	tx, err := otoken.RedeemVaultBalance(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())
	return nil
}

func opynEstimateContract(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	amount, err := decimal.NewFromString(*clientOpynEstimateAmount)
	if err != nil {
		return err
	}
	info, err := otoken.ContractInfo(ctx)
	if err != nil {
		return err
	}
	tokensIssuable, err := otoken.MaxOTokensIssuable(ctx, amount)
	if err != nil {
		return err
	}
	currencyName := info.Currency.Hex()
	if info.Currency == opyn.EthContract {
		currencyName = "eth"
	} else {
		name, _, err := reg.ByAddress(info.Currency)
		if err == nil {
			currencyName = name
		}
	}
	switch info.Type {
	case opyn.PutType:
		// XX USDC buys XX contracts.
		fmt.Printf("%s usdc mints %s contracts\n", amount, tokensIssuable)
	case opyn.CallType:
		// XXX currency buys XX contracts
		fmt.Printf("%s %s mints %s contracts\n", amount, currencyName, tokensIssuable)
	default:
		return errors.Errorf("Contract id %s with type %s not supported", otoken.Contract.Hex(), info.Type)
	}

	uniswapv1Client, err := uniswapv1.NewClient(opynClient.Client)
	if err != nil {
		return err
	}

	otoken2, err := token2.NewClient(otoken.Client, otoken.Contract, otoken.Decimals)
	if err != nil {
		return err
	}

	uniswapv1Exchange, err := uniswapv1Client.GetExchange(ctx, otoken2)
	if err != nil {
		return err
	}

	fmt.Printf("Uniswapv1 exchange: %s\n", uniswapv1Exchange.ContractAddress().Hex())

	sellEth, err := uniswapv1Exchange.GetTokenToEthInputPrice(ctx, tokensIssuable)
	if err != nil {
		return err
	}
	fmt.Printf("%s contracts sell for %s eth\n", tokensIssuable, sellEth)

	buyEth, err := uniswapv1Exchange.GetEthToTokenOutputPrice(ctx, tokensIssuable)
	if err != nil {
		return err
	}
	fmt.Printf("%s contracts cost %s eth to buy\n", tokensIssuable, buyEth)

	return nil
}

func opynAddCollateralOption(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	amountToCreate, err := decimal.NewFromString(*clientOpynAddCollateralAmtToCreate)
	if err != nil {
		return err
	}
	collateralAmount, err := decimal.NewFromString(*clientOpynAddCollateralCollateralAmt)
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet is locked")
	}

	receiver := common.HexToAddress(*clientOpynAddCollateralReceiver)
	collateral, err := otoken.Collateral(ctx)
	if err != nil {
		return err
	}
	if collateral == opyn.EthContract {
		return errors.New("ETH as collateral not yet supported")
	}

	tx, err := otoken.AddERC20CollateralOption(ctx, account, amountToCreate, collateralAmount, receiver)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", tx.Hash().Hex())

	return nil
}

func opynAddAndSellCollateralOption(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	amountToCreate, err := decimal.NewFromString(*clientOpynAddAndSellCollateralAmtToCreate)
	if err != nil {
		return err
	}
	collateralAmount, err := decimal.NewFromString(*clientOpynAddAndSellCollateralCollateralAmt)
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet is locked")
	}

	receiver := common.HexToAddress(*clientOpynAddAndSellCollateralReceiver)
	collateral, err := otoken.Collateral(ctx)
	if err != nil {
		return err
	}
	if collateral == opyn.EthContract {
		return errors.New("ETH as collateral not yet supported")
	}

	tx, err := otoken.AddAndSellERC20CollateralOption(ctx, account, amountToCreate, collateralAmount, receiver)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", tx.Hash().Hex())

	return nil
}

func opynGetVault(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	account, _, err := getAccount()
	if err != nil {
		return err
	}

	a, b, c, d, err := otoken.GetVault(ctx, account.Address())
	if err != nil {
		return err
	}

	fmt.Printf("%s %s %s %t\n", a, b, c, d)

	return nil
}
