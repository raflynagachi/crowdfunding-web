package helpers

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

func UserToUserResponse(user models.User) web.UserResponse {
	return web.UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Occupation:    user.Occupation,
		TokenRemember: user.RememberToken,
	}
}

func CampaignImageToCampaignImageResponse(campaignImage models.CampaignImage) web.CampaignImageResponse {
	return web.CampaignImageResponse{
		ImageUrl:  campaignImage.Filename,
		IsPrimary: campaignImage.IsPrimary,
	}
}

func CampaignImagesToCampaignImageResponses(campaignImages []models.CampaignImage) []web.CampaignImageResponse {
	campaignImagesFormatter := []web.CampaignImageResponse{}
	for _, campaign := range campaignImages {
		campaignFormatted := CampaignImageToCampaignImageResponse(campaign)
		campaignImagesFormatter = append(campaignImagesFormatter, campaignFormatted)
	}
	return campaignImagesFormatter
}

func CampaignToCampaignResponse(campaign models.Campaign) web.CampaignResponse {
	var imageUrl string
	if len(campaign.CampaignImages) == 0 {
		imageUrl = "default.jpg"
	} else {
		imageUrl = campaign.CampaignImages[0].Filename
	}

	return web.CampaignResponse{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount:      campaign.BackerCount,
		Slug:             campaign.Slug,
		ImageUrl:         imageUrl,
	}
}

func CampaignToCampaignDetailResponse(campaign models.Campaign) web.CampaignDetailResponse {
	var imageUrl string
	if len(campaign.CampaignImages) == 0 {
		imageUrl = "default.jpg"
	} else {
		imageUrl = campaign.CampaignImages[0].Filename
	}

	perks := PerksToSlice(campaign.Perks)

	campaignUserResponse := web.CampaignUserResponse{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.AvatarFilename,
	}
	campaignImageResponse := CampaignImagesToCampaignImageResponses(campaign.CampaignImages)

	return web.CampaignDetailResponse{
		ID:                     campaign.ID,
		UserID:                 campaign.UserID,
		Name:                   campaign.Name,
		ShortDescription:       campaign.ShortDescription,
		GoalAmount:             campaign.GoalAmount,
		CurrentAmount:          campaign.CurrentAmount,
		BackerCount:            campaign.BackerCount,
		Slug:                   campaign.Slug,
		ImageUrl:               imageUrl,
		Perks:                  perks,
		CampaignUserResponse:   campaignUserResponse,
		CampaignImageResponses: campaignImageResponse,
	}
}

func CampaignsToCampaignResponses(campaigns []models.Campaign) []web.CampaignResponse {
	campaignsFormatter := []web.CampaignResponse{}
	for _, campaign := range campaigns {
		campaignFormatted := CampaignToCampaignResponse(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatted)
	}
	return campaignsFormatter
}

func CampaignToCampaignCreateResponse(campaign models.Campaign) web.CampaignResponse {
	return web.CampaignResponse{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		Slug:             campaign.Slug,
	}
}
