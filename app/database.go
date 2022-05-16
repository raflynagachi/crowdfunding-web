package app

import (
	"fmt"
	"log"

	"github.com/raflynagachi/crowdfunding-web/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(dbConfig config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected to localhost:3306")
	return db
}
