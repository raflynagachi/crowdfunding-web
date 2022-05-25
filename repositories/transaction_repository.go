package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type TransactionRepository interface {
	FindCampaignByID(campaignID int) ([]models.Transaction, error)
}
