package database

import (
	"github.com/cnbattle/anjuke/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	Local *gorm.DB
)

func init() {
	var err error
	args := "database-" + utils.GetYmd() + ".db"
	log.Println(args)
	Local, err = gorm.Open("sqlite3", args)
	if err != nil {
		log.Panic(err)
	}
	Local.LogMode(false)
	Local.DB().SetMaxOpenConns(10)
	Local.DB().SetMaxIdleConns(20)
	Local.AutoMigrate(&Data{})
}
