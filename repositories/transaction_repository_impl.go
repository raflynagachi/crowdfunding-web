package repositories

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

func (r *TransactionRepositoryImpl) FindByCampaignID(campaignID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where("campaign_id = ?", campaignID).Order("id DESC").Preload("User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
