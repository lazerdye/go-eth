
PACKAGE=github.com/lazerdye/go-eth

all: token/erc20/erc20.go kyber/kyber.go zeroex/ether_token/ether_token.go zeroex/exchange/exchange.go 

token/TokenERC20.abi:
	solc --abi token/erc20.sol -o sol

token/erc20/erc20.go: token/erc20/TokenERC20.abi
	abigen --abi=token/erc20/TokenERC20.abi --pkg=erc20 --out=token/erc20/erc20.go

kyber/kyber.go: kyber/kyber_network_proxy.abi
	abigen --abi=kyber/kyber_network_proxy.abi --pkg=kyber --out=kyber/kyber.go

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

uniswapv2/router01.go: uniswapv2/router01.abi
	abigen --abi=uniswapv2/router01.abi --pkg=uniswapv2 --type=Router01 --out=uniswapv2/router01.go

compound/compound.go: compound/compound.abi
	abigen --abi=compound/compound.abi --pkg=compound --type=Compound --out=compound/compound.go

gofmt:
	go fmt $(PACKAGE)/...


