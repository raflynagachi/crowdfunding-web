package handler

import "github.com/raflynagachi/crowdfunding-web/middleware"

type Controller struct {
	User        userHandler
	Campaign    campaignHandler
	Transaction transactionHandler
	Auth        authHandler
	middleware.AuthMiddleware
}

func RegisterController(middleware *middleware.AuthMiddleware,
	user userHandler,
	campaign campaignHandler,
	transaction transactionHandler,
	auth authHandler,
) Controller {
	return Controller{
		AuthMiddleware: *middleware,
		User:           user,
		Campaign:       campaign,
		Transaction:    transaction,
		Auth:           auth,
	}
}
