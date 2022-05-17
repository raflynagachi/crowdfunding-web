package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type AuthService interface {
	Register(r web.AuthRegisterRequest) (web.UserResponse, error)
	Login(r web.AuthLoginRequest) (web.UserResponse, error)
	IsEmailAvailable(r web.AuthIsEmailAvailableRequest) (bool, error)
}
