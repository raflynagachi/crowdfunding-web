package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type UserControllerImpl struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (controller *UserControllerImpl) Create(c *gin.Context) {
	webResponse := web.WebResponse{}

	var input web.UserCreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponse.Code = http.StatusUnprocessableEntity
		webResponse.Status = "error"
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
		c.JSON(http.StatusUnprocessableEntity, webResponse)
		return
	}

	userResponse := controller.service.Create(input)

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = userResponse

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindByEmail(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindAll(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
