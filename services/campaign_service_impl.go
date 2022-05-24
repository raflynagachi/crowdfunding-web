package services

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models"
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

func (service *CampaignServiceImpl) Create(campaignCreateReq web.CampaignCreateRequest) (web.CampaignResponse, error) {
	campaignResponse := web.CampaignResponse{}

	campaign := models.Campaign{
		UserID:           campaignCreateReq.UserID,
		Name:             campaignCreateReq.Name,
		ShortDescription: campaignCreateReq.ShortDescription,
		Description:      campaignCreateReq.Description,
		GoalAmount:       campaignCreateReq.GoalAmount,
		Perks:            campaignCreateReq.Perks,
	}
	slugRaw := fmt.Sprintf("%s %d", campaign.Name, campaign.UserID)
	campaign.Slug = slug.Make(slugRaw)

	campaign, err := service.repository.Create(campaign)
	if err != nil {
		return campaignResponse, err
	}

	return helpers.CampaignToCampaignResponse(campaign), nil
}

func (service *CampaignServiceImpl) Update(campaignID int, campaignUpdateReq web.CampaignUpdateRequest) (web.CampaignResponse, error) {
	var campaignResponse web.CampaignResponse

	campaign, err := service.repository.FindByID(campaignID)
	if err != nil {
		return campaignResponse, err
	}

	if campaign.UserID != campaignUpdateReq.User.ID {
		return campaignResponse, errors.New("unauthorized user")
	}

	campaign.Name = campaignUpdateReq.Name
	campaign.ShortDescription = campaignUpdateReq.ShortDescription
	campaign.Description = campaignUpdateReq.Description
	campaign.GoalAmount = campaignUpdateReq.GoalAmount
	campaign.Perks = campaignUpdateReq.Perks
	campaign.User = campaignUpdateReq.User

	campaignUpdated, err := service.repository.Update(campaign)
	if err != nil {
		return campaignResponse, err
	}

	return helpers.CampaignToCampaignResponse(campaignUpdated), nil
}

func (service *CampaignServiceImpl) CreateCampaignImage(
	campaignImage web.CampaignImageCreateRequest,
	fileLocation string) (web.CampaignImageResponse, error) {
	imageResponse := web.CampaignImageResponse{}

	campaign, err := service.repository.FindByID(campaignImage.CampaignID)
	if err != nil {
		return imageResponse, err
	}
	if campaign.UserID != campaignImage.User.ID {
		return imageResponse, errors.New("unauthorized user")
	}

	if campaignImage.IsPrimary {
		_, err := service.repository.MarkAllImagesAsNonPrimary(campaignImage.CampaignID)
		if err != nil {
			return imageResponse, err
		}
	}

	campaignImageNew := models.CampaignImage{}
	campaignImageNew.CampaignID = campaignImage.CampaignID
	campaignImageNew.IsPrimary = campaignImage.IsPrimary
	campaignImageNew.Filename = fileLocation

	campaignImageNew, err = service.repository.CreateImage(campaignImageNew)
	if err != nil {
		return imageResponse, err
	}
	return helpers.CampaignImageToCampaignImageResponse(campaignImageNew), nil
}
