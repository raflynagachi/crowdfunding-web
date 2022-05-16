package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	var input web.UserCreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	userResponse := controller.service.Create(input)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindById(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindAll(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
