package services

import (
	"errors"

	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/repositories"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (service *UserServiceImpl) UpdateAvatar(userID int, fileLocation string) (models.User, error) {
	user, err := service.repository.FindById(userID)
	if err != nil {
		return user, err
	}

	user.AvatarFilename = fileLocation
	user, err = service.repository.Update(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *UserServiceImpl) FindById(userID int) (models.User, error) {
	user, err := service.repository.FindById(userID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found")
	}
	return user, nil
}

func (service *UserServiceImpl) FindAll() ([]models.User, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
