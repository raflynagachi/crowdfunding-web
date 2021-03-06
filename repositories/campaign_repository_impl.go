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

func (repository *CampaignRepositoryImpl) Create(campaign models.Campaign) (models.Campaign, error) {
	err := repository.DB.Debug().Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (repository *CampaignRepositoryImpl) Update(campaign models.Campaign) (models.Campaign, error) {
	err := repository.DB.Debug().Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (repository *CampaignRepositoryImpl) CreateImage(campaignImage models.CampaignImage) (models.CampaignImage, error) {
	err := repository.DB.Debug().Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (repository *CampaignRepositoryImpl) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := repository.DB.Debug().Model(&models.CampaignImage{}).Where("campaign_id = ?", campaignID).Update(
		"is_primary", false,
	).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
