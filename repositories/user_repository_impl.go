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
	err := repository.DB.Debug().Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(userEmail string) (models.User, error) {
	var user models.User
	err := repository.DB.Debug().Where("email = ?", userEmail).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindById(userId int) (models.User, error) {
	var user models.User
	err := repository.DB.Debug().Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
