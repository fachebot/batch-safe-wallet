package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Key struct {
	ID         uint   `gorm:"column:id;primarykey"`
	Address    string `gorm:"column:address;index;not null"`
	Contract   string `gorm:"column:contract;not null"`
	PrivateKey string `gorm:"column:privateKey;not null"`
}

func (Key) TableName() string {
	return "keys"
}

func NewKey() (Key, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Key{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Key{}, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	contractAddr := crypto.CreateAddress(address, 0)
	privateKeyHex := "0x" + hexutil.Encode(privateKeyBytes)[2:]

	return Key{Address: address.Hex(), PrivateKey: privateKeyHex, Contract: contractAddr.Hex()}, nil
}
