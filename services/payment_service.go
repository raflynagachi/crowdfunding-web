package services

import (
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type PaymentService interface {
	GetPaymentURL(transaction models.Transaction, user models.User) (string, error)
	ProcessPayment(transaction web.NotificationTransactionRequest) error
}
