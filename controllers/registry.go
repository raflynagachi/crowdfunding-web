package controllers

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	AuthController
	UserController
	CampaignController
	TransactionController
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware,
	auth AuthController,
	user UserController,
	campaign CampaignController,
	transaction TransactionController,
) Controller {
	return Controller{
		AuthMiddleware:        *middleware,
		AuthController:        auth,
		UserController:        user,
		CampaignController:    campaign,
		TransactionController: transaction,
	}
}
