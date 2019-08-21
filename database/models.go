package database

import (
	"github.com/jinzhu/gorm"
)

type Data struct {
	gorm.Model
	CityName  string
	Community string
	Name      string
	Cover     string
	Address   string
	Price     string
	PriceTxt  string
	Url       string
	Date      string
	Status    int `gorm:"type:tinyint(1);DEFAULT:0;NOT NULL"`
}

type FangPrice struct {
	gorm.Model
	HashKey        string `gorm:"type:char(32);unique_index;NOT NULL"`
	CityName       string `gorm:"type:varchar(32);NOT NULL"`
	Community      string `gorm:"type:varchar(32);NOT NULL"`
	Name           string `gorm:"type:varchar(32);index:fang_name;NOT NULL"`
	Cover          string `gorm:"type:varchar(128);NOT NULL"`
	Address        string `gorm:"type:varchar(128);NOT NULL"`
	Price          string `gorm:"type:varchar(32);NOT NULL"`
	PriceTxt       string `gorm:"type:varchar(32);NOT NULL"`
	PriceTxtStatus int    `gorm:"type:tinyint(1);DEFAULT:0;NOT NULL"`
	Url            string `gorm:"type:varchar(128);NOT NULL"`
	Date           string `gorm:"type:varchar(32);NOT NULL"`
}
