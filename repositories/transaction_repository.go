package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type TransactionRepository interface {
	FindByCampaignID(campaignID int) ([]models.Transaction, error)
	FindByUserID(userID int) ([]models.Transaction, error)
}
