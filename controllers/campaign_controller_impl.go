package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/helpers"
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
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
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
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
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

func (controller *CampaignControllerImpl) CreateImage(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	var input web.CampaignImageCreateRequest

	err := c.ShouldBind(&input)
	if err != nil {
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err), "is_uploaded": false}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	file, err := c.FormFile("campaign_image")
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error(), "is_uploaded": false}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	currentUser := c.MustGet("user").(models.User)
	input.User = currentUser

	path := fmt.Sprintf("assets/campaign-images/%d-%s", input.User.ID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error(), "is_uploaded": false}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	campaignImage, err := controller.service.CreateCampaignImage(input, path)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error(), "is_uploaded": false}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = gin.H{"is_primary": campaignImage.IsPrimary, "is_uploaded": true}
	c.JSON(http.StatusOK, webResponse)
}
