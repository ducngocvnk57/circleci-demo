package app

import (
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	DNS  string
	Host string
}

var once sync.Once
var gormOnce sync.Once
var db *gorm.DB

func Db() *gorm.DB {
	gormOnce.Do(func() {
		dbDriver := os.Getenv("DB_DRIVER")
		dbDns := os.Getenv("DB_DNS")
		var err error
		db, err = gorm.Open(dbDriver, dbDns)
		if err != nil {
			log.Fatalln("connect database failed")
		}
	})
	return db
}
