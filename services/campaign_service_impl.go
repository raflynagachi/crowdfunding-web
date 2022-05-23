package services

import (
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/repositories"
)

type CampaignServiceImpl struct {
	repository repositories.CampaignRepository
}

func NewCampaignService(repository repositories.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{
		repository: repository,
	}
}

func (service *CampaignServiceImpl) FindCampaigns(userID int) ([]web.CampaignResponse, error) {
	campaignResponses := []web.CampaignResponse{}

	if userID != 0 {
		campaigns, err := service.repository.FindByUserID(userID)
		if err != nil {
			return campaignResponses, err
		}
		return helpers.CampaignsToCampaignResponses(campaigns), nil
	}

	campaigns, err := service.repository.FindAll()
	if err != nil {
		return campaignResponses, err
	}
	return helpers.CampaignsToCampaignResponses(campaigns), nil
}

func (service *CampaignServiceImpl) FindCampaign(campaignID int) (web.CampaignDetailResponse, error) {
	campaignResponse := web.CampaignDetailResponse{}

	campaign, err := service.repository.FindByID(campaignID)
	if err != nil {
		return campaignResponse, err
	}
	return helpers.CampaignToCampaignDetailResponse(campaign), nil
}
