
PACKAGE=gitlab.com/lazerdye/go-eth

all: token/erc20.go

token/TokenERC20.abi:
	solc --abi token/erc20.sol -o sol

token/erc20.go: token/TokenERC20.abi
	abigen --abi=token/TokenERC20.abi --pkg=token --out=token/erc20.go

gofmt:
	go fmt $(PACKAGE)/...


