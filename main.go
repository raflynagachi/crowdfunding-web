package main

import (
	"github.com/raflynagachi/crowdfunding-web/app/config"
)

func main() {
	dbConfig, _ := config.ConfigEnv()

	config.OpenDB(dbConfig)
}
