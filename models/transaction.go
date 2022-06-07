package models

import (
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID         int     `gorm:"primaryKey"`
	UserID     int     `gorm:"not null;index"`
	CampaignID int     `gorm:"not null;index"`
	Amount     float64 `gorm:"not null"`
	Status     string  `gorm:"not null"`
	Code       string  `gorm:"not null;index"`
	PaymentUrl string  `gorm:"not null"`
	User       User
	Campaign   Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (t *Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
