package handler

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	userHandler
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware,
	user userHandler,
) Controller {
	return Controller{
		AuthMiddleware: *middleware,
		userHandler:    user,
	}
}
