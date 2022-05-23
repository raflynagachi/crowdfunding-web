package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type CampaignRepository interface {
	FindAll() ([]models.Campaign, error)
	FindByID(campaignID int) (models.Campaign, error)
	FindByUserID(userID int) ([]models.Campaign, error)
	Create(campaign models.Campaign) (models.Campaign, error)
}
