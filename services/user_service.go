package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type UserService interface {
	Create(r web.UserCreateRequest) web.UserResponse
	Update(r web.UserUpdateRequest) web.UserResponse
	Delete(r web.UserCreateRequest)
	FindById(r web.UserCreateRequest) web.UserResponse
	FindAll() []web.UserResponse
}
