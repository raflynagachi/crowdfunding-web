package services

import (
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

func (service *UserServiceImpl) UpdateAvatar(userId int, fileLocation string) (models.User, error) {
	user, err := service.repository.FindById(userId)
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
