package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/controllers"
)

func NewRouter(controller controllers.Controller) *gin.Engine {
	router := gin.Default()
	router.Static("/avatar-images", "./assets/avatar_images")
	router.Static("/campaign-images", "./assets/campaign_images")

	apiRoot := router.Group("/api/v1")

	apiRoot.POST("/register", controller.Register)
	apiRoot.POST("/login", controller.Login)
	apiRoot.POST("/email-checker", controller.IsEmailAvailable)
	apiRoot.PUT("/avatars", controller.AuthMiddleware.Serve, controller.UpdateAvatar)

	apiRoot.POST("/campaigns", controller.AuthMiddleware.Serve, controller.Create)
	apiRoot.PUT("/campaigns/:campaignID", controller.AuthMiddleware.Serve, controller.Update)
	apiRoot.GET("/campaigns", controller.FindCampaigns)
	apiRoot.GET("/campaigns/:campaignID", controller.FindCampaign)

	apiRoot.POST("/campaign-image", controller.AuthMiddleware.Serve, controller.CreateImage)

	apiRoot.GET("/campaigns/:campaignID/transactions",
		controller.AuthMiddleware.Serve,
		controller.TransactionController.FindByCampaignID)

	return router
}
