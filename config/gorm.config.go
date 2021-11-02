package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

func InitDB() *gorm.DB {
	conf := GetConfig()
	DSN := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("connectionString error")
	}
	return DB
}
