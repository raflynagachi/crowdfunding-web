package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type AuthControllerImpl struct {
	service services.UserService
}

func NewAuthController(service services.UserService) AuthController {
	return &AuthControllerImpl{
		service: service,
	}
}

func (controller *AuthControllerImpl) Register(c *gin.Context) {
	webResponse := web.WebResponse{
		Status: "error",
	}

	var input web.AuthRegisterRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponse.Code = http.StatusUnprocessableEntity
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
		c.JSON(http.StatusUnprocessableEntity, webResponse)
		return
	}

	userResponse, err := controller.service.Register(input)
	if err != nil {
		webResponse.Code = http.StatusUnprocessableEntity
		webResponse.Data = gin.H{"errors": err}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = userResponse

	c.JSON(http.StatusOK, webResponse)
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {
	webResponse := web.WebResponse{
		Status: "error",
	}

	var input web.AuthLoginRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponse.Code = http.StatusUnprocessableEntity
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
		c.JSON(http.StatusUnprocessableEntity, webResponse)
		return
	}

	userResponse, err := controller.service.Login(input)
	if err != nil {
		webResponse.Code = http.StatusUnauthorized
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusUnauthorized, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = userResponse

	c.JSON(http.StatusOK, webResponse)
}
