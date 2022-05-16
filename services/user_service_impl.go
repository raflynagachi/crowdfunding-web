package services

import (
	"log"

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

func (service *UserServiceImpl) Create(r web.UserCreateRequest) web.UserResponse {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Name:         r.Name,
		Email:        r.Email,
		Occupation:   r.Occupation,
		PasswordHash: string(passwordHash),
		Role:         models.UserRole,
	}

	user, err = service.repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	return helpers.UserToUserResponse(user)
}

func (service *UserServiceImpl) Update(r web.UserUpdateRequest) web.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) Delete(r web.UserCreateRequest) {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) FindById(r web.UserCreateRequest) web.UserResponse {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) FindAll() []web.UserResponse {
	panic("not implemented") // TODO: Implement
}
