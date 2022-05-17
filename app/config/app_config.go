package config

import "github.com/raflynagachi/crowdfunding-web/helpers"

type AppConfig struct {
	AppName string
	AppEnv  string
	AppHost string
	AppPort string
}

func (appConfig *AppConfig) SetupEnv() {
	appConfig.AppName = helpers.GetEnv("APP_NAME", "GoFund")
	appConfig.AppEnv = helpers.GetEnv("APP_ENV", "development")
	appConfig.AppHost = helpers.GetEnv("APP_HOST", "localhost")
	appConfig.AppPort = helpers.GetEnv("APP_PORT", "8080")
}
