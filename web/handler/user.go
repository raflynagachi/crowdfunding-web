package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type userHandler struct {
	userService services.UserService
	authService services.AuthService
}

func NewUserHandler(userService services.UserService, auth services.AuthService) *userHandler {
	return &userHandler{userService, auth}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_create.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input web.UserRegisterRequest
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_create.html", input)
		return
	}

	var authInput web.AuthRegisterRequest
	authInput.Name = input.Name
	authInput.Email = input.Email
	authInput.Occupation = input.Occupation
	authInput.Password = input.Password

	_, err = h.authService.Register(authInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Edit(c *gin.Context) {
	idParam := c.Param("userID")
	id, _ := strconv.Atoi(idParam)
	registeredUser, err := h.userService.FindById(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	input := web.UserUpdateRequest{
		ID:         id,
		Name:       registeredUser.Name,
		Email:      registeredUser.Email,
		Occupation: registeredUser.Occupation,
	}
	c.HTML(http.StatusOK, "user_edit.html", input)
}

func (h *userHandler) Update(c *gin.Context) {
	idParam := c.Param("userID")
	id, _ := strconv.Atoi(idParam)

	var input web.UserUpdateRequest
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	input.ID = id
	_, err = h.userService.Update(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) NewAvatar(c *gin.Context) {
	userID := c.Param("userID")
	id, _ := strconv.Atoi(userID)

	c.HTML(http.StatusOK, "user_avatar.html", gin.H{"ID": id})
}

func (h *userHandler) CreateAvatar(c *gin.Context) {
	userID := c.Param("userID")
	id, _ := strconv.Atoi(userID)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	path := fmt.Sprintf("avatar-images/%d-%s", id, file.Filename)
	fullpath := "assets/" + path
	err = c.SaveUploadedFile(file, fullpath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	_, err = h.userService.UpdateAvatar(id, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}
