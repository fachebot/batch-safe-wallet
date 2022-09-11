.PHONY: build

all: build

build:
	go build -o bin/batch-safe-wallet .
