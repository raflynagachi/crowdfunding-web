package models

import "time"

type CampaignImages struct {
	ID         int    `gorm:"primaryKey"`
	CampaignID int    `gorm:"not null;index"`
	Filename   string `gorm:"not null"`
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
