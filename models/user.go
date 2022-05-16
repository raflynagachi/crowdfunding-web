package models

import "time"

const UserRole = "user"
const AdminRole = "admin"

type User struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Email          string `gorm:"not null;index"`
	PasswordHash   string `gorm:"not null"`
	Occupation     string
	AvatarFilename string
	Role           string `gorm:"not null"`
	RememberToken  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
