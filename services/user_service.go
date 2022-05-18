package services

import "github.com/raflynagachi/crowdfunding-web/models"

type UserService interface {
	UpdateAvatar(userId int, fileLocation string) (models.User, error)
}
