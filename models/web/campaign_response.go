package web

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
type CampaignDetailResponse struct {
	ID                     int                     `json:"id"`
	UserID                 int                     `json:"user_id"`
	Name                   string                  `json:"name"`
	ShortDescription       string                  `json:"short_description"`
	GoalAmount             float64                 `json:"goal_amount"`
	CurrentAmount          float64                 `json:"current_amount"`
	BackerCount            int                     `json:"backer_count"`
	Slug                   string                  `json:"slug"`
	Perks                  []string                `json:"perks"`
	ImageUrl               string                  `json:"image_url"`
	CampaignUserResponse   CampaignUserResponse    `json:"user"`
	CampaignImageResponses []CampaignImageResponse `json:"images"`
}

type CampaignUserResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type CampaignImageResponse struct {
	IsPrimary bool   `json:"is_primary"`
	ImageUrl  string `json:"image_url"`
}
