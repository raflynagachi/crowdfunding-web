package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/auth/jwt"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type AuthMiddleware struct {
	jwtService  jwt.JwtService
	userService services.UserService
}

func NewAuthMiddleware(jwtService jwt.JwtService, userService services.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (m *AuthMiddleware) Serve(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
		return
	}

	var tokenString string
	tokens := strings.Split(authHeader, " ")
	if len(tokens) == 2 {
		tokenString = tokens[1]
	}

	token, err := m.jwtService.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
		return
	}

	claims, err := m.jwtService.JwtTokenToMap(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
		return
	}

	userID := int(claims["user_id"].(float64))
	user, err := m.userService.FindById(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
		return
	}

	c.Set("user", user)
}
