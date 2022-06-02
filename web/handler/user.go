package handler

import (
	"net/http"

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
