.PHONY: build

all: build

tools:
	add-apt-repository -y ppa:ethereum/ethereum
	apt-get update -y
	apt-get install ethereum -y

build:
	solc --abi contracts/GnosisSafeProxy.sol -o build --gas --optimize --optimize-runs 200 --overwrite
	abigen --abi=./build/GnosisSafeProxy.abi --pkg=proxies --out=proxies/GnosisSafeProxy.go
	go build -o bin/batch-safe-wallet .
