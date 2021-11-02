package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

func InitDB() *gorm.DB {
	if os.Getenv("DB_DRIVER") == "mysql" {
		DSN := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

		if err != nil {
			panic("connectionString error")
		}
		return DB
	}
	fmt.Println("DB_DRIVER not supported")
	return nil
}

func GetDBInstance() *gorm.DB {
	return DB
}
