package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type UserService interface {
	Register(r web.AuthRegisterRequest) (web.UserResponse, error)
	Login(r web.AuthLoginRequest) (web.UserResponse, error)
}
