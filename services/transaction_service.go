package services

import (
	"github.com/raflynagachi/crowdfunding-web/models/web"
)

type TransactionService interface {
	FindByCampaignID(campaignID int, userID int) ([]web.TransactionResponse, error)
	FindByUserID(UserID int) ([]web.TransactionUserResponse, error)
	Create(transactionReq web.CreateTransactionRequest) (web.TransactionCreateResponse, error)
}
