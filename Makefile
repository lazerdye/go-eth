.PHONY: bin/cli

PACKAGE=github.com/lazerdye/go-eth

all: token/erc20/erc20.go kyber/kyber.go zeroex/ether_token/ether_token.go zeroex/exchange/exchange.go

bin/cli:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o bin/cli cli/*.go

token/TokenERC20.abi:
	solc --abi token/erc20.sol -o sol

token/erc20/erc20.go: token/erc20/TokenERC20.abi
	abigen --abi=token/erc20/TokenERC20.abi --pkg=erc20 --out=token/erc20/erc20.go

token2/erc20.go: token2/erc20.abi
	abigen --abi=token2/erc20.abi --pkg=token2 --type=Erc20 --out=token2/erc20.go

kyber/kyber.go: kyber/kyber.abi
	abigen --abi=kyber/kyber.abi --pkg=kyber --out=kyber/kyber.go

zeroex/ether_token/ether_token.go: zeroex/ether_token/ether_token.abi
	abigen --abi=zeroex/ether_token/ether_token.abi --pkg=ether_token --out=zeroex/ether_token/ether_token.go

zeroex/exchange/exchange.go: zeroex/exchange/exchange.abi
	abigen --abi=zeroex/exchange/exchange.abi --pkg=exchange --out=zeroex/exchange/exchange.go

uniswapv1/factory.go: uniswapv1/factory.abi
	abigen --abi=uniswapv1/factory.abi --pkg=uniswapv1 --type=Factory --out=uniswapv1/factory.go

uniswapv1/exchange.go: uniswapv1/exchange.abi
	abigen --abi=uniswapv1/exchange.abi --pkg=uniswapv1 --type=Exchange --out=uniswapv1/exchange.go

uniswapv2/factory.go: uniswapv2/factory.abi
	abigen --abi=uniswapv2/factory.abi --pkg=uniswapv2 --type=Factory --out=uniswapv2/factory.go

uniswapv2/pair.go: uniswapv2/pair.abi
	abigen --abi=uniswapv2/pair.abi --pkg=uniswapv2 --type=Pair --out=uniswapv2/pair.go

uniswapv2/router02.go: uniswapv2/router02.abi
	abigen --abi=uniswapv2/router02.abi --pkg=uniswapv2 --type=Router02 --out=uniswapv2/router02.go

uniswapv2/token_distributor.go: uniswapv2/token_distributor.abi
	abigen --abi=uniswapv2/token_distributor.abi --pkg=uniswapv2 --type=TokenDistributor --out=uniswapv2/token_distributor.go

compound/ceth.go: compound/cETH.abi
	abigen --abi=compound/cETH.abi --pkg=compound --type=Ceth --out=compound/ceth.go

compound/cerc20.go: compound/cERC20.abi
	abigen --abi=compound/cERC20.abi --pkg=compound --type=CErc20 --out=compound/cerc20.go

augur/repv2.go: augur/repv2.abi
	abigen --abi=augur/repv2.abi --pkg=augur --type=Repv2 --out=augur/repv2.go

bancor/contract_registry.go: bancor/contract_registry.abi
	abigen --abi=bancor/contract_registry.abi --pkg=bancor --type=ContractRegistry --out=bancor/contract_registry.go

bancor/converter_registry.go: bancor/converter_registry.abi
	abigen --abi=bancor/converter_registry.abi --pkg=bancor --type=ConverterRegistry --out=bancor/converter_registry.go

bancor/network.go: bancor/network.abi
	abigen --abi=bancor/network.abi --pkg=bancor --type=Network --out=bancor/network.go

opyn/options_factory.go: opyn/options_factory.abi
	abigen --abi=opyn/options_factory.abi --pkg=opyn --type=OptionsFactory --out=opyn/options_factory.go

opyn/otoken.go: opyn/otoken.abi
	abigen --abi=opyn/otoken.abi --pkg=opyn --type=OToken --out=opyn/otoken.go

eth2/deposit.go: eth2/deposit.abi
	abigen --abi=eth2/deposit.abi --pkg=eth2 --type=Deposit --out=eth2/deposit.go

digixdao/dgdinterface.go: digixdao/dgdinterface.abi
	abigen --abi=digixdao/dgdinterface.abi --pkg=digixdao --type=DGDInterface --out=digixdao/dgdinterface.go

sushi/sushiv2factory.go: sushi/sushiv2factory.abi
	abigen --abi=sushi/sushiv2factory.abi --pkg=sushi --type=SushiV2Factory --out=sushi/sushiv2factory.go

sushi/pair.go: sushi/pair.abi
	abigen --abi=sushi/pair.abi --pkg=sushi --type=Pair --out=sushi/pair.go

gofmt:
	go fmt $(PACKAGE)/...


