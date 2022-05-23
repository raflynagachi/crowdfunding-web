package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type CampaignService interface {
	FindCampaigns(userID int) ([]web.CampaignResponse, error)
	FindCampaign(campaignID int) (web.CampaignDetailResponse, error)
}
