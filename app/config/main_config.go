package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ConfigEnv() (DBConfig, AppConfig) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := DBConfig{}
	dbConfig.SetupEnv()
	appConfig := AppConfig{}
	appConfig.SetupEnv()

	return dbConfig, appConfig
}
