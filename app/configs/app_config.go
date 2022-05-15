package config

import "github.com/raflynagachi/crowdfunding-web/helpers"

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (appConfig *AppConfig) SetupEnv() {
	appConfig.AppName = helpers.GetEnv("APP_NAME", "GoFund")
	appConfig.AppEnv = helpers.GetEnv("APP_ENV", "GoFund")
	appConfig.AppPort = helpers.GetEnv("APP_PORT", "GoFund")
}
