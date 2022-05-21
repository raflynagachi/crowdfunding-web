package controllers

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	AuthController
	UserController
	CampaignController
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware,
	auth AuthController,
	user UserController,
	campaign CampaignController,
) Controller {
	return Controller{
		AuthMiddleware:     *middleware,
		AuthController:     auth,
		UserController:     user,
		CampaignController: campaign,
	}
}
