package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type CampaignService interface {
	FindCampaigns(userID int) ([]web.CampaignResponse, error)
	FindCampaign(campaignID int) (web.CampaignDetailResponse, error)
	Create(campaign web.CampaignCreateRequest) (web.CampaignResponse, error)
	Update(campaignID int, campaign web.CampaignUpdateRequest) (web.CampaignResponse, error)
	CreateCampaignImage(campaignImage web.CampaignImageCreateRequest, fileLocation string) (web.CampaignImageResponse, error)
}
