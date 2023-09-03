package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	dbname := "flights"
	username := "root"
	password := ""

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
