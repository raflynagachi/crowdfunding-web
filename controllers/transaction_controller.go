package controllers

import "github.com/gin-gonic/gin"

type TransactionController interface {
	FindByCampaignID(c *gin.Context)
	FindByUserID(c *gin.Context)
}
