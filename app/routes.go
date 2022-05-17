package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/controllers"
)

func NewRouter(authController controllers.AuthController) *gin.Engine {
	router := gin.Default()
	apiRoot := router.Group("/api/v1")

	apiRoot.POST("/register", authController.Register)
	apiRoot.POST("/login", authController.Login)
	apiRoot.POST("/email-checker", authController.IsEmailAvailable)

	return router
}
