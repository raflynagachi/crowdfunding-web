package services

import "github.com/raflynagachi/crowdfunding-web/models"

type PaymentService interface {
	GetPaymentURL(transaction models.Transaction, user models.User) (string, error)
}
