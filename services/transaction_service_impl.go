package services

import (
	"errors"

	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/repositories"
)

type TransactionServiceImpl struct {
	repository         repositories.TransactionRepository
	campaignRepository repositories.CampaignRepository
}

func NewTransactionService(repository repositories.TransactionRepository, campaignRepo repositories.CampaignRepository) TransactionService {
	return &TransactionServiceImpl{
		repository:         repository,
		campaignRepository: campaignRepo,
	}
}

func (s *TransactionServiceImpl) FindByCampaignID(campaignID int, userID int) ([]web.TransactionResponse, error) {
	transactionResponses := []web.TransactionResponse{}

	campaign, err := s.campaignRepository.FindByID(campaignID)
	if err != nil {
		return transactionResponses, err
	}
	if campaign.UserID != userID {
		return transactionResponses, errors.New("unauthorized user")
	}

	transaction, err := s.repository.FindByCampaignID(campaignID)
	if err != nil {
		return transactionResponses, err
	}
	return helpers.TransactionsToTransactionResponses(transaction), nil
}

func (s *TransactionServiceImpl) FindByUserID(UserID int) ([]web.TransactionUserResponse, error) {
	var transactionResponse []web.TransactionUserResponse

	transactions, err := s.repository.FindByUserID(UserID)
	if err != nil {
		return transactionResponse, err
	}
	return helpers.TransactionsToTransactionUserResponses(transactions), nil
}
