package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func GetSecret() []byte {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	key, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok {
		log.Println("key doesn't match anything")
	}

	return []byte(key)
}
