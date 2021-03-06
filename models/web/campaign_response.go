package web

import "github.com/leekchan/accounting"

type CampaignResponse struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	GoalAmount       float64 `json:"goal_amount"`
	CurrentAmount    float64 `json:"current_amount"`
	BackerCount      int     `json:"backer_count"`
	Slug             string  `json:"slug"`
	ImageUrl         string  `json:"image_url"`
}

func (camp *CampaignResponse) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(camp.GoalAmount)
}

type CampaignDetailResponse struct {
	ID                     int                     `json:"id"`
	UserID                 int                     `json:"user_id"`
	Name                   string                  `json:"name"`
	ShortDescription       string                  `json:"short_description"`
	Description            string                  `json:"description"`
	GoalAmount             float64                 `json:"goal_amount"`
	CurrentAmount          float64                 `json:"current_amount"`
	BackerCount            int                     `json:"backer_count"`
	Slug                   string                  `json:"slug"`
	Perks                  []string                `json:"perks"`
	ImageUrl               string                  `json:"image_url"`
	CampaignUserResponse   CampaignUserResponse    `json:"user"`
	CampaignImageResponses []CampaignImageResponse `json:"images"`
}

func (camp CampaignDetailResponse) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(camp.GoalAmount)
}
func (camp CampaignDetailResponse) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(camp.CurrentAmount)
}

type CampaignUserResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type CampaignImageResponse struct {
	IsPrimary bool   `json:"is_primary"`
	ImageUrl  string `json:"image_url"`
}
