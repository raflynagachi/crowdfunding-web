package database

import "github.com/raflynagachi/crowdfunding-web/models"

type Model struct {
	TableName string
	Model     interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}, TableName: "users"},
		{Model: models.Campaign{}, TableName: "campaigns"},
		{Model: models.CampaignImages{}, TableName: "campaign_images"},
		{Model: models.Transaction{}, TableName: "transactions"},
	}
}
