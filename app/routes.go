package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/controllers"
)

func NewRouter(userController controllers.UserController) *gin.Engine {
	router := gin.Default()
	apiRoot := router.Group("/api/v1")

	apiRoot.POST("/users", userController.Create)
	// apiRoot.PUT("/users", userController.Update)
	// apiRoot.DELETE("/users", userController.Delete)
	// apiRoot.GET("/users", userController.FindById)
	// apiRoot.GET("/users", userController.FindAll)

	return router
}
