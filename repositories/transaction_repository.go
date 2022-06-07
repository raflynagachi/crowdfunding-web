package repositories

import "github.com/raflynagachi/crowdfunding-web/models"

type TransactionRepository interface {
	FindByCampaignID(campaignID int) ([]models.Transaction, error)
	FindByUserID(userID int) ([]models.Transaction, error)
	FindByID(ID int) (models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	Create(transaction models.Transaction) (models.Transaction, error)
	Update(transaction models.Transaction) (models.Transaction, error)
}
