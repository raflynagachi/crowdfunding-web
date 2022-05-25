package models

import "time"

type Transaction struct {
	ID         int     `gorm:"primaryKey"`
	UserID     int     `gorm:"not null;index"`
	CampaignID int     `gorm:"not null;index"`
	Amount     float64 `gorm:"not null"`
	Status     string  `gorm:"not null"`
	Code       string  `gorm:"not null;index"`
	User       User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
