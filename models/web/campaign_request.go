package web

import "github.com/raflynagachi/crowdfunding-web/models"

type CampaignCreateRequest struct {
	UserID           int     `json:"user_id" binding:"required"`
	Name             string  `json:"name" binding:"required"`
	ShortDescription string  `json:"short_description" binding:"required"`
	Description      string  `json:"description" binding:"required"`
	GoalAmount       float64 `json:"goal_amount" binding:"required"`
	Perks            string  `json:"perks" binding:"required"`
	User             models.User
}
