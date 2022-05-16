package helpers

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

func UserToUserResponse(user models.User) web.UserResponse {
	return web.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Occupation:    user.Occupation,
		TokenRemember: user.RememberToken,
	}
}
