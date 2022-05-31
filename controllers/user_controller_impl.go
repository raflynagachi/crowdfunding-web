package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models"
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

func (controller *UserControllerImpl) UpdateAvatar(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "error",
		Data:   gin.H{"is_uploaded": false},
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	user := c.MustGet("user").(models.User)
	userId := user.ID

	path := fmt.Sprintf("assets/avatar_images/%d-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	_, err = controller.service.UpdateAvatar(userId, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = gin.H{"is_uploaded": true}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) FetchUser(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	currentUser := c.MustGet("user").(models.User)

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = helpers.UserToUserResponse(currentUser)
	c.JSON(http.StatusOK, webResponse)
}
