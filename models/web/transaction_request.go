package web

import "github.com/raflynagachi/crowdfunding-web/models"

type CreateTransactionRequest struct {
	Amount     float64     `json:"amount" binding:"required"`
	CampaignID int         `json:"campaign_id" binding:"required"`
	User       models.User `json:"user"`
}
