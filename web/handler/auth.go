package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type authHandler struct {
	auth services.AuthService
	user services.UserService
}

func NewAuthHandler(auth services.AuthService, user services.UserService) *authHandler {
	return &authHandler{
		auth: auth,
		user: user,
	}
}

func (h *authHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *authHandler) Create(c *gin.Context) {
	var input web.AuthLoginRequest

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userLoggedIn, err := h.auth.Login(input)
	if err != nil {
		// user not found
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := h.user.FindById(userLoggedIn.ID)
	if err != nil || user.Role != "admin" {
		// unauthorized user
		c.Redirect(http.StatusFound, "/login")
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *authHandler) Destroy(c *gin.Context) {
	sessions := sessions.Default(c)
	sessions.Clear()
	sessions.Save()
	c.Redirect(http.StatusFound, "/login")
}
