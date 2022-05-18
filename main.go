package main

import (
	"flag"

	"github.com/raflynagachi/crowdfunding-web/app"
	"github.com/raflynagachi/crowdfunding-web/app/cmd"
	"github.com/raflynagachi/crowdfunding-web/app/config"
	"github.com/raflynagachi/crowdfunding-web/controllers"
	"github.com/raflynagachi/crowdfunding-web/repositories"
	"github.com/raflynagachi/crowdfunding-web/services"
)

func main() {
	dbConfig, appConfig := config.ConfigEnv()
	db := app.OpenDB(dbConfig)

	flag.Parse()
	narg := flag.NArg()
	if narg != 0 {
		cmd.RootCmd.Execute()
		return
	}

	userRepository := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	controller := controllers.RegisterController(authController, userController)

	router := app.NewRouter(controller)
	router.Run(":" + appConfig.AppPort)
}
