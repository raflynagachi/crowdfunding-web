package services

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type UserService interface {
	UpdateAvatar(userID int, fileLocation string) (models.User, error)
	FindById(userID int) (models.User, error)
	FindAll() ([]models.User, error)
	Update(user web.UserUpdateRequest) (models.User, error)
}
