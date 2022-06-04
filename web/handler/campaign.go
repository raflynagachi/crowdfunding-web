package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type campaignHandler struct {
	campaignService services.CampaignService
}

func NewCampaignHandler(campaignService services.CampaignService) *campaignHandler {
	return &campaignHandler{
		campaignService: campaignService,
	}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.FindCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}
