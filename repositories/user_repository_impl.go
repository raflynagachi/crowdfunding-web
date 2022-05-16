package repositories

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Create(user models.User) (models.User, error) {
	err := repository.DB.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Update(user models.User) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) Delete(user models.User) error {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) FindById(userID int) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) FindAll() ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}
