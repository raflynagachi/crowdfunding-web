package services

import (
	"strconv"

	"github.com/raflynagachi/crowdfunding-web/app/config"
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/repositories"
	"github.com/veritrans/go-midtrans"
)

type PaymentServiceImpl struct {
	repository   repositories.TransactionRepository
	campaignRepo repositories.CampaignRepository
	midtransConf config.MidtransConfig
}

func NewPaymentService(repository repositories.TransactionRepository,
	campaignRepo repositories.CampaignRepository,
	midtransConf config.MidtransConfig) PaymentService {
	return &PaymentServiceImpl{
		repository:   repository,
		campaignRepo: campaignRepo,
		midtransConf: midtransConf,
	}
}

func (s *PaymentServiceImpl) GetPaymentURL(transaction models.Transaction, user models.User) (string, error) {
	midClient := midtrans.NewClient()
	midClient.ServerKey = s.midtransConf.ServerKey
	midClient.ClientKey = s.midtransConf.ClientKey
	midClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *PaymentServiceImpl) ProcessPayment(transactionReq web.NotificationTransactionRequest) error {
	transaction_id, _ := strconv.Atoi(transactionReq.OrderID)
	transaction, err := s.repository.FindByID(transaction_id)
	if err != nil {
		return err
	}

	if transactionReq.PaymentType == "credit_card" &&
		transactionReq.TransactionStatus == "capture" &&
		transactionReq.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if transactionReq.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if transactionReq.TransactionStatus == "deny" {
		transaction.Status = "cancelled"
	} else if transactionReq.TransactionStatus == "expire" {
		transaction.Status = "cancelled"
	} else if transactionReq.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepo.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount += 1
		campaign.CurrentAmount += updatedTransaction.Amount

		_, err := s.campaignRepo.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
