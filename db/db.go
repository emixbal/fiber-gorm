package db

import (
	"fiber-gorm/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	db  *gorm.DB
)

func Init() {
	conf := config.GetConfig()
	DSN := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("connectionString error")
	}
}

func CreateCon() *gorm.DB {
	return db
}
