package main

import (
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const BatchSize = 2000

func init() {
	openDatabase("keys.db")
}

var table = []interface{}{
	new(Key),
	new(Create2Key),
}

func openDatabase(path string) {
	d := sqlite.Open(path)
	db, err := gorm.Open(d, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				IgnoreRecordNotFoundError: true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	state.Db = db

	for _, table := range table {
		if db.Migrator().HasTable(table) {
			err := db.Migrator().AutoMigrate(table)
			if err != nil {
				panic(err)
			}
			continue
		}

		if err = db.Migrator().CreateTable(table); err != nil {
			panic(err)
		}
	}
}

type Keys struct{}

func (Keys) Save(keys []Key) error {
	state.DbMutex.Lock()
	defer state.DbMutex.Unlock()

	return state.Db.CreateInBatches(keys, BatchSize).Error
}

func (Keys) Scan(offset, limit int) ([]Key, error) {
	state.DbMutex.Lock()
	defer state.DbMutex.Unlock()

	var result []Key
	err := state.Db.Model(Key{}).Order("id").Offset(offset).Limit(limit).Find(&result).Error
	return result, err
}

type Create2Keys struct{}

func (Create2Keys) Save(keys []Create2Key) error {
	state.DbMutex.Lock()
	defer state.DbMutex.Unlock()

	return state.Db.CreateInBatches(keys, BatchSize).Error
}

func (Create2Keys) Scan(offset, limit int) ([]Create2Key, error) {
	state.DbMutex.Lock()
	defer state.DbMutex.Unlock()

	var result []Create2Key
	err := state.Db.Model(Create2Key{}).Order("id").Offset(offset).Limit(limit).Find(&result).Error
	return result, err
}

func (Create2Keys) LastNonce(address common.Address) (uint64, error) {
	state.DbMutex.Lock()
	defer state.DbMutex.Unlock()

	var record Create2Key
	err := state.Db.Model(Create2Key{}).Where(`"address" = ?`, address.Hex()).Last(&record).Error
	return record.SaltNonce, err
}
