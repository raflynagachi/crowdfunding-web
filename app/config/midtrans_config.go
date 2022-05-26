package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/raflynagachi/crowdfunding-web/helpers"
)

type MidtransConfig struct {
	ServerKey string
	ClientKey string
}

func (conf *MidtransConfig) SetupEnv() {
	conf.ServerKey = helpers.GetEnv("MIDTRANS_SERVER_KEY", "SB-XXX-XXX")
	conf.ClientKey = helpers.GetEnv("MIDTRANS_CLIENT_KEY", "SB-XXX-XXX")
}

func GetMidtransKey() MidtransConfig {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	midtransConf := MidtransConfig{}
	midtransConf.SetupEnv()

	return midtransConf
}
