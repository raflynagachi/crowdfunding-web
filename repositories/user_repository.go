package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type UserRepository interface {
	Register(user models.User) (models.User, error)
	Login(userEmail string) (models.User, error)
}
