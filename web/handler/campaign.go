package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type campaignHandler struct {
	campaignService services.CampaignService
	userService     services.UserService
}

func NewCampaignHandler(campaign services.CampaignService, user services.UserService) *campaignHandler {
	return &campaignHandler{
		campaignService: campaign,
		userService:     user,
	}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.FindCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}

func (h *campaignHandler) New(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := web.CampaignFormCreateRequest{
		Users: users,
	}

	c.HTML(http.StatusOK, "campaign_new.html", input)
}

func (h *campaignHandler) Create(c *gin.Context) {
	var input web.CampaignFormCreateRequest

	err := c.ShouldBind(&input)
	if err != nil {
		users, e := h.userService.FindAll()
		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		input.Users = users
		input.Error = err
		c.HTML(http.StatusOK, "campaign_new.html", input)
		return
	}

	user, err := h.userService.FindById(input.UserID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createCampaign := web.CampaignCreateRequest{
		UserID:           input.UserID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		User:             user,
	}

	_, err = h.campaignService.Create(createCampaign)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}
