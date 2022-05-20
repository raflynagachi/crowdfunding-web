package repositories

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{
		DB: db,
	}
}

func (repository *CampaignRepositoryImpl) FindAll() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	err := repository.DB.Debug().Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (repository *CampaignRepositoryImpl) FindByUserID(userID int) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	err := repository.DB.Debug().Where("user_id = ?", userID).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
