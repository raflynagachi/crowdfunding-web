package config

import (
	"github.com/raflynagachi/crowdfunding-web/helpers"
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
