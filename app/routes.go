package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/controllers"
)

func NewRouter(controller controllers.Controller) *gin.Engine {
	router := gin.Default()
	apiRoot := router.Group("/api/v1")

	apiRoot.POST("/register", controller.Register)
	apiRoot.POST("/login", controller.Login)
	apiRoot.POST("/email-checker", controller.IsEmailAvailable)
	apiRoot.PUT("/avatars", controller.AuthMiddleware.Serve, controller.UpdateAvatar)

	return router
}
