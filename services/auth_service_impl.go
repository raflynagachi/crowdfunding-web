package services

import (
	"errors"

	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (service *UserServiceImpl) Register(r web.AuthRegisterRequest) (web.UserResponse, error) {
	userResponse := web.UserResponse{}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	if err != nil {
		return userResponse, err
	}

	user := models.User{
		Name:         r.Name,
		Email:        r.Email,
		Occupation:   r.Occupation,
		PasswordHash: string(passwordHash),
		Role:         models.UserRole,
	}

	user, err = service.repository.Register(user)
	if err != nil {
		return userResponse, err
	}
	return helpers.UserToUserResponse(user), nil
}

func (service *UserServiceImpl) Login(r web.AuthLoginRequest) (web.UserResponse, error) {
	userResponse := web.UserResponse{}
	email := r.Email
	password := r.Password

	user, err := service.repository.Login(email)
	if err != nil {
		return userResponse, err
	}

	if user.ID == 0 {
		return userResponse, errors.New("email doesn't match any user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return userResponse, errors.New("password or email doesn't match")
	}

	return helpers.UserToUserResponse(user), nil
}
