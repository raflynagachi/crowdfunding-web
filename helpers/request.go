package helpers

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

func CampaignToCampaignCreateRequest(campaign models.Campaign) web.CampaignCreateRequest {
	return web.CampaignCreateRequest{
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		Perks:            campaign.Perks,
	}
}
