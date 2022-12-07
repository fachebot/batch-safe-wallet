package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	tronAddress "github.com/fbsobreira/gotron-sdk/pkg/address"
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

func NewKey(typ string) (Key, error) {
	switch typ {
	case "evm":
		return NewEvmKey()
	case "tron":
		return NewTronKey()
	default:
		return Key{}, errors.New("invalid key type")
	}
}

func NewEvmKey() (Key, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Key{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Key{}, errors.New("invalid public key")
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	contractAddr := crypto.CreateAddress(address, 0)
	privateKeyHex := "0x" + hexutil.Encode(privateKeyBytes)[2:]

	return Key{Address: address.Hex(), PrivateKey: privateKeyHex, Contract: contractAddr.Hex()}, nil
}

func NewTronKey() (Key, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Key{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Key{}, errors.New("invalid public key")
	}

	address := tronAddress.PubkeyToAddress(*publicKeyECDSA).String()
	privateKeyHex := hex.EncodeToString(privateKey.D.Bytes())
	return Key{Address: address, PrivateKey: privateKeyHex}, nil
}

type Create2Key struct {
	ID        uint   `gorm:"column:id;primarykey"`
	Address   string `gorm:"column:address;index;not null"`
	Contract  string `gorm:"column:contract;not null"`
	SaltNonce uint64 `gorm:"column:salt_nonce;not null"`
	InitHash  string `gorm:"column:init_hash;not null"`
}

func NewCreate2Key(deployer common.Address, initHash common.Hash, saltNonce uint64) (Create2Key, error) {
	var salt [32]byte
	copy(salt[:], common.LeftPadBytes(big.NewInt(0).SetUint64(saltNonce).Bytes(), 32))
	address := crypto.CreateAddress2(deployer, salt, initHash.Bytes())

	return Create2Key{Address: deployer.Hex(), Contract: address.Hex(), SaltNonce: saltNonce, InitHash: initHash.Hex()}, nil
}
