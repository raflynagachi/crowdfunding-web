package handler

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	User     userHandler
	Campaign campaignHandler
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware,
	user userHandler,
	campaign campaignHandler,
) Controller {
	return Controller{
		AuthMiddleware: *middleware,
		User:           user,
		Campaign:       campaign,
	}
}
