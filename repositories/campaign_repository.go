package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type CampaignRepository interface {
	FindAll() ([]models.Campaign, error)
	FindByUserID(userID int) ([]models.Campaign, error)
}
