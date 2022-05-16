package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type UserRepository interface {
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) error
	FindById(userID int) (models.User, error)
	FindAll() ([]models.User, error)
}
