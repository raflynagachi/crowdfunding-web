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

func (r *TransactionRepositoryImpl) FindByUserID(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Where("user_id = ?", userID).Order("id DESC").Preload(
		"Campaign.CampaignImages", "campaign_images.is_primary = 1",
	).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) Create(transaction models.Transaction) (models.Transaction, error) {
	err := r.DB.Create(&transaction).Preload("User").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepositoryImpl) Update(transaction models.Transaction) (models.Transaction, error) {
	err := r.DB.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
