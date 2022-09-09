package main

import (
	"gorm.io/gorm"
	"sync"
)

var state struct {
	Db      *gorm.DB
	DbMutex sync.Mutex
}
