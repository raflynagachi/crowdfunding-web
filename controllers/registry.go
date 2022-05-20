package controllers

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	AuthController
	UserController
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware, auth AuthController, user UserController) Controller {
	return Controller{
		AuthMiddleware: *middleware,
		AuthController: auth,
		UserController: user,
	}
}
