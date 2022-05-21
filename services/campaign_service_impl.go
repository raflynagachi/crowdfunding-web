package services

import (
	"github.com/raflynagachi/crowdfunding-web/models"
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

func (service *CampaignServiceImpl) FindCampaigns(userID int) ([]models.Campaign, error) {
	if userID != 0 {
		campaigns, err := service.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := service.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
