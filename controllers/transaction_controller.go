package controllers

import "github.com/gin-gonic/gin"

type TransactionController interface {
	FindCampaignByID(c *gin.Context)
}
