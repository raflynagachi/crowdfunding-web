package controllers

import "github.com/gin-gonic/gin"

type CampaignController interface {
	FindCampaigns(c *gin.Context)
	FindCampaign(c *gin.Context)
	Create(c *gin.Context)
}
