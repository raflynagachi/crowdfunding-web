package database

import "github.com/raflynagachi/crowdfunding-web/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Campaign{}},
		{Model: models.CampaignImages{}},
		{Model: models.Transaction{}},
	}
}
