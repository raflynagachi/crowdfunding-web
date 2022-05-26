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
	midtransConf := config.GetMidtransKey()
	db := app.OpenDB(dbConfig)

	flag.Parse()
	narg := flag.NArg()
	if narg != 0 {
		cmd.RootCmd.Execute()
		return
	}

	userRepository := repositories.NewUserRepository(db)
	campaignRepository := repositories.NewCampaignRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)

	authService := services.NewAuthService(userRepository)
	jwtService := jwt.NewJwtService()
	authController := controllers.NewAuthController(authService, jwtService)

	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	campaignService := services.NewCampaignService(campaignRepository)
	campaignController := controllers.NewCampaignController(campaignService)

	paymentService := services.NewPaymentService(transactionRepository, midtransConf)
	transactionService := services.NewTransactionService(transactionRepository, campaignRepository, paymentService)
	transactionController := controllers.NewTransactionController(transactionService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, userService)
	controller := controllers.RegisterController(authMiddleware,
		authController,
		userController,
		campaignController,
		transactionController,
	)

	router := app.NewRouter(controller)
	router.Run(":" + appConfig.AppPort)
}
