package main

import (
	"flag"
	"fmt"

	"github.com/raflynagachi/crowdfunding-web/app"
	"github.com/raflynagachi/crowdfunding-web/app/cmd"
	"github.com/raflynagachi/crowdfunding-web/app/config"
	"github.com/raflynagachi/crowdfunding-web/auth/jwt"
	"github.com/raflynagachi/crowdfunding-web/controllers"
	"github.com/raflynagachi/crowdfunding-web/middleware"
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
	campaignRepository := repositories.NewCampaignRepository(db)
	camp, _ := campaignRepository.FindAll()
	fmt.Println("=======================")
	fmt.Println(camp)

	authService := services.NewAuthService(userRepository)
	jwtService := jwt.NewJwtService()
	authController := controllers.NewAuthController(authService, jwtService)

	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, userService)
	controller := controllers.RegisterController(authMiddleware, authController, userController)

	router := app.NewRouter(controller)
	router.Run(":" + appConfig.AppPort)
}
