package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "crowdfunding-web"

func ConfigEnv() (DBConfig, AppConfig) {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := DBConfig{}
	dbConfig.SetupEnv()
	appConfig := AppConfig{}
	appConfig.SetupEnv()

	return dbConfig, appConfig
}
