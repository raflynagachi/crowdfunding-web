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
