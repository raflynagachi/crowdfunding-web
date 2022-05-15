package config

import (
	"fmt"
	"log"

	"github.com/raflynagachi/crowdfunding-web/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBDriver string
	DBHost   string
	DBPort   string
	DBName   string
	DBUser   string
	DBPass   string
}

func (dbConfig *DBConfig) SetupEnv() {
	dbConfig.DBDriver = helpers.GetEnv("DB_DRIVER", "mysql")
	dbConfig.DBHost = helpers.GetEnv("DB_HOST", "localhost")
	dbConfig.DBPort = helpers.GetEnv("DB_PORT", "3306")
	dbConfig.DBName = helpers.GetEnv("DB_NAME", "go-fund")
	dbConfig.DBUser = helpers.GetEnv("DB_USERNAME", "admin")
	dbConfig.DBPass = helpers.GetEnv("DB_PASSWORD", "password")
}

func OpenDB(dbConfig DBConfig) *gorm.DB {
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
