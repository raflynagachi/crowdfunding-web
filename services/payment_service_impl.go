package services

import (
	"strconv"

	"github.com/raflynagachi/crowdfunding-web/app/config"
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/repositories"
	"github.com/veritrans/go-midtrans"
)

type PaymentServiceImpl struct {
	repository   repositories.TransactionRepository
	midtransConf config.MidtransConfig
}

func NewPaymentService(repository repositories.TransactionRepository, midtransConf config.MidtransConfig) PaymentService {
	return &PaymentServiceImpl{
		repository:   repository,
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
