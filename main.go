package main

import (
	"flag"

	"github.com/raflynagachi/crowdfunding-web/app/cmd"
	"github.com/raflynagachi/crowdfunding-web/app/config"
)

func main() {
	dbConfig, _ := config.ConfigEnv()
	config.OpenDB(dbConfig)

	flag.Parse()
	narg := flag.NArg()
	if narg != 0 {
		cmd.RootCmd.Execute()
	}

	// userRepository := repositories.NewUserRepository(db)
	// userRepository.Create(models.User{
	// 	Name: "Nagachi",
	// })

}
