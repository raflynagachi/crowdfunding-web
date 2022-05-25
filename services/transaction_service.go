package services

import "github.com/raflynagachi/crowdfunding-web/models/web"

type TransactionService interface {
	FindByCampaignID(campaignID int, userID int) ([]web.TransactionResponse, error)
}
