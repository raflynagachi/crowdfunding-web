package web

import (
	"time"
)

type TransactionResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
type TransactionUserResponse struct {
	ID            int                   `json:"id"`
	Amount        float64               `json:"amount"`
	CreatedAt     time.Time             `json:"created_at"`
	CampaignName  string                `json:"campaign_name"`
	CampaignImage CampaignImageResponse `json:"campaign_image"`
}

type TransactionCreateResponse struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	UserID     int       `json:"user_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentUrl string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
}
