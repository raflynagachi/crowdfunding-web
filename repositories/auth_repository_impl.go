package repositories

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repository *AuthRepositoryImpl) Register(user models.User) (models.User, error) {
	err := repository.DB.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *AuthRepositoryImpl) Login(userEmail string) (models.User, error) {
	var user models.User
	err := repository.DB.Debug().Where("email = ?", userEmail).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
