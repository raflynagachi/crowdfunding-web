package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
