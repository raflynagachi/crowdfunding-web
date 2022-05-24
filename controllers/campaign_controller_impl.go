package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type CampaignControllerImpl struct {
	service services.CampaignService
}

func NewCampaignController(service services.CampaignService) CampaignController {
	return &CampaignControllerImpl{
		service: service,
	}
}

func (controller *CampaignControllerImpl) FindCampaigns(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusUnprocessableEntity,
		Status: "error",
	}

	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := controller.service.FindCampaigns(userID)
	if err != nil {
		webResponse.Code = http.StatusBadRequest
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = campaigns
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CampaignControllerImpl) FindCampaign(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusUnprocessableEntity,
		Status: "error",
	}

	campaignID, _ := strconv.Atoi(c.Param("campaignID"))
	if campaignID == 0 {
		webResponse.Data = gin.H{"errors": "empty campaignID"}
		c.JSON(http.StatusUnprocessableEntity, webResponse)
		return
	}

	campaign, err := controller.service.FindCampaign(campaignID)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusUnprocessableEntity, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = campaign
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CampaignControllerImpl) Create(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	var input web.CampaignCreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	currentUser := c.MustGet("user").(models.User)
	input.User = currentUser

	campaignResponse, err := controller.service.Create(input)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = campaignResponse
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CampaignControllerImpl) Update(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	var input web.CampaignUpdateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	currentUser := c.MustGet("user").(models.User)
	input.User = currentUser

	campaignID, _ := strconv.Atoi(c.Param("campaignID"))
	campaign, err := controller.service.Update(campaignID, input)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = campaign
	c.JSON(http.StatusOK, webResponse)
}
