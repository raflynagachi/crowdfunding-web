package models

import "time"

type Campaign struct {
	ID               int    `gorm:"primaryKey"`
	UserID           int    `gorm:"not null;index"`
	Name             string `gorm:"not null"`
	Slug             string `gorm:"not null"`
	ShortDescription string
	Description      string
	GoalAmount       float64 `gorm:"not null"`
	CurrentAmount    float64
	BackerCount      int
	Perks            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}
