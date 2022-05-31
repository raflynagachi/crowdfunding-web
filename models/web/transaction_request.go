package web

import "github.com/raflynagachi/crowdfunding-web/models"

type CreateTransactionRequest struct {
	Amount     float64     `json:"amount" binding:"required"`
	CampaignID int         `json:"campaign_id" binding:"required"`
	User       models.User `json:"user"`
}

type NotificationTransactionRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
