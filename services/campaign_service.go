package services

import "github.com/raflynagachi/crowdfunding-web/models"

type CampaignService interface {
	FindCampaigns(userID int) ([]models.Campaign, error)
}
