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
	"github.com/lazerdye/go-eth/uniswapv2"
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
	clientOpynListVaults                        = clientOpynCommand.Command("list-vaults", "List vaults")
	clientOpynExerciseCommand                   = clientOpynCommand.Command("exercise", "Exercise token")
	clientOpynExerciseAmount                    = clientOpynExerciseCommand.Arg("amount", "Amount of token to exercise").Required().String()
	clientOpynExerciseVaults                    = clientOpynExerciseCommand.Arg("vualts", "List of vaults to exercise").Required().String()
)

func opynCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) error {
	opynClient, err := opyn.NewClient(client)
	if err != nil {
		return err
	}
	switch commands[0] {
	case "list-options-contracts":
		return opynListOptionsContracts(ctx, opynClient, reg)
	case "open-vault":
		return opynOpenVault(ctx, opynClient, reg)
	case "redeem-vault-balance":
		return opynRedeemVaultBalance(ctx, opynClient, reg)
	case "estimate":
		return opynEstimateContract(ctx, opynClient, reg)
	case "add-collateral-option":
		return opynAddCollateralOption(ctx, opynClient, reg)
	case "add-and-sell-collateral-option":
		return opynAddAndSellCollateralOption(ctx, opynClient, reg)
	case "get-vault":
		return opynGetVault(ctx, opynClient, reg)
	case "list-vaults":
		return opynListVaults(ctx, opynClient, reg)
	case "exercise":
		return opynExercise(ctx, opynClient, reg)
	default:
		return errors.Errorf("Unknown opyn subcommand: %s", commands[0])
	}
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
	account, err := getAccount(true)
	if err != nil {
		return err
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
	account, err := getAccount(true)
	if err != nil {
		return err
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

	wethToken, err := reg.ByName("weth")
	if err != nil {
		return err
	}
	usdcToken, err := reg.ByName("usdc")
	if err != nil {
		return err
	}

	uniswapv2Client, err := uniswapv2.NewClient(opynClient.Client)
	if err != nil {
		return err
	}

	if info.Type == opyn.PutType {
		amountOut, err := uniswapv2Client.GetAmountOut(ctx, buyEth, wethToken, usdcToken)
		if err != nil {
			return err
		}

		fmt.Printf("Uniswapv2: %s eth converts to %s usdc\n", buyEth, amountOut)
	} else if info.Type == opyn.CallType {

		amountOut, err := uniswapv2Client.GetAmountOut(ctx, sellEth, wethToken, usdcToken)
		if err != nil {
			return err
		}

		totalSell := tokensIssuable.Add(amountOut)
		fmt.Printf("Uniswapv2: %s eth converts to %s usdc (%s total sell)\n", sellEth, amountOut, totalSell)

		currencyToken := wethToken
		if info.Currency != opyn.EthContract {
			_, currencyToken, err = reg.ByAddress(info.Currency)
			if err != nil {
				return err
			}
		}

		usdcIn, err := uniswapv2Client.GetAmountIn(ctx, amount, usdcToken, currencyToken)
		if err != nil {
			return err
		}

		fmt.Printf("Uniswapv2: %s usdc to buy %s %s\n", usdcIn, amount, currencyName)
	}

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
	account, err := getAccount(true)
	if err != nil {
		return err
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
	account, err := getAccount(true)
	if err != nil {
		return err
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
	account, err := getAccount(false)
	if err != nil {
		return err
	}

	a, b, c, d, err := otoken.GetVault(ctx, account.Address())
	if err != nil {
		return err
	}

	collateralExp, err := otoken.CollateralExp(ctx)
	if err != nil {
		return err
	}
	//underlyingExp, err := otoken.UnderlyingExp(ctx)
	//if err != nil {
	//	return err
	//}

	fmt.Printf("%s %s %s %t\n", decimal.NewFromBigInt(a, collateralExp), decimal.NewFromBigInt(b, -int32(otoken.Decimals)), c, d)

	return nil
}

func opynListVaults(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	collateralExp, err := otoken.CollateralExp(ctx)
	if err != nil {
		return err
	}
	underlyingExp, err := otoken.UnderlyingExp(ctx)
	if err != nil {
		return err
	}
	graph := opyn.NewGraph()
	opened, err := graph.OpenedVaults(ctx, otoken.Contract.Hex())
	if err != nil {
		return err
	}
	for _, vaultInfo := range opened {
		fmt.Printf("Vault Owner: %s\n", vaultInfo.Owner)
		collateralAmtWei, tokenAmtWei, underlyingAmtWei, d, err := otoken.GetVault(ctx, common.HexToAddress(vaultInfo.Owner))
		if err != nil {
			return err
		}
		collateralAmt := decimal.NewFromBigInt(collateralAmtWei, collateralExp)
		tokenAmt := decimal.NewFromBigInt(tokenAmtWei, -int32(otoken.Decimals))
		underlyingAmt := decimal.NewFromBigInt(underlyingAmtWei, underlyingExp)

		fmt.Printf("%s - %s / %s / %s / %t\n", vaultInfo.Owner, collateralAmt, tokenAmt, underlyingAmt, d)
	}

	return errors.New("Not Implemented")
}

func opynExercise(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	otoken, err := getOToken(ctx, opynClient, reg)
	if err != nil {
		return err
	}
	account, err := getAccount(true)
	if err != nil {
		return err
	}
	amount, err := decimal.NewFromString(*clientOpynExerciseAmount)
	if err != nil {
		return err
	}
	vaults := strings.Split(*clientOpynExerciseVaults, ",")
	vaultAddresses := make([]common.Address, len(vaults))
	for i, vault := range vaults {
		vaultAddresses[i] = common.HexToAddress(vault)
	}
	tx, err := otoken.Exercise(ctx, account, amount, vaultAddresses)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())
	return nil
}
