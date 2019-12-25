
PACKAGE=github.com/lazerdye/go-eth

all: token/erc20.go

token/TokenERC20.abi:
	solc --abi token/erc20.sol -o sol

token/erc20.go: token/TokenERC20.abi
	abigen --abi=token/TokenERC20.abi --pkg=token --out=token/erc20.go

kyber/kyber.go: kyber/kyber_network_proxy.abi
	abigen --abi=kyber/kyber_network_proxy.abi --pkg=kyber --out=kyber/kyber.go

gofmt:
	go fmt $(PACKAGE)/...


