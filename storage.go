package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const BatchSize = 2000

func init() {
	d := sqlite.Open("keys.db")
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
	if db.Migrator().HasTable(Key{}) {
		return
	}

	if err = db.Migrator().CreateTable(Key{}); err != nil {
		panic(err)
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
