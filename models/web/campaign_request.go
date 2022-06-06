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
type CampaignUpdateRequest struct {
	Name             string  `json:"name" binding:"required"`
	ShortDescription string  `json:"short_description" binding:"required"`
	Description      string  `json:"description" binding:"required"`
	GoalAmount       float64 `json:"goal_amount" binding:"required"`
	Perks            string  `json:"perks" binding:"required"`
	User             models.User
}

type CampaignImageCreateRequest struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       models.User
}

type CampaignFormCreateRequest struct {
	UserID           int     `form:"user_id" binding:"required"`
	Name             string  `form:"name" binding:"required"`
	ShortDescription string  `form:"short_description" binding:"required"`
	Description      string  `form:"description" binding:"required"`
	GoalAmount       float64 `form:"goal_amount" binding:"required"`
	Perks            string  `form:"perks" binding:"required"`
	Users            []models.User
	Error            error
}
