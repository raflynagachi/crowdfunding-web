package main

import (
	"flag"

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

	authService := services.NewAuthService(userRepository)
	jwtService := jwt.NewJwtService()
	authController := controllers.NewAuthController(authService, jwtService)

	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	campaignService := services.NewCampaignService(campaignRepository)
	campaignController := controllers.NewCampaignController(campaignService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, userService)
	controller := controllers.RegisterController(authMiddleware, authController, userController, campaignController)

	router := app.NewRouter(controller)
	router.Run(":" + appConfig.AppPort)
}
