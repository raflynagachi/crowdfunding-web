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
	err := repository.DB.Debug().Preload(
		"CampaignImages", "campaign_images.is_primary = 1",
	).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (repository *CampaignRepositoryImpl) FindByUserID(userID int) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	err := repository.DB.Debug().Where("user_id = ?", userID).Preload(
		"CampaignImages", "campaign_images.is_primary = 1",
	).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (repository *CampaignRepositoryImpl) FindByID(campaignID int) (models.Campaign, error) {
	var campaign models.Campaign
	err := repository.DB.Debug().Where("id = ?", campaignID).Preload(
		"CampaignImages").Preload("User").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
