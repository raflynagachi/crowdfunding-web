package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (h *campaignHandler) NewImage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	c.HTML(http.StatusOK, "campaign_image.html", gin.H{"ID": id})
}

func (h *campaignHandler) CreateImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingCampaign, err := h.campaignService.FindCampaign(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userID := existingCampaign.UserID
	path := fmt.Sprintf("campaign-images/%d-%s", userID, file.Filename)
	fullpath := "assets/" + path

	err = c.SaveUploadedFile(file, fullpath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userCampaign, err := h.userService.FindById(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createCampaignReq := web.CampaignImageCreateRequest{
		CampaignID: id,
		IsPrimary:  true,
		User:       userCampaign,
	}

	_, err = h.campaignService.CreateCampaignImage(createCampaignReq, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingCampaign, err := h.campaignService.FindCampaign(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	formUpdateCampaign := web.CampaignFormUpdateRequest{
		ID:               existingCampaign.ID,
		Name:             existingCampaign.Name,
		ShortDescription: existingCampaign.ShortDescription,
		Description:      existingCampaign.Description,
		GoalAmount:       existingCampaign.GoalAmount,
		Perks:            strings.Join(existingCampaign.Perks, ","),
	}

	c.HTML(http.StatusOK, "campaign_edit.html", formUpdateCampaign)
}

func (h *campaignHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input web.CampaignFormUpdateRequest
	err := c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = err
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	existingCampaign, err := h.campaignService.FindCampaign(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userID := existingCampaign.UserID
	userCampaign, err := h.userService.FindById(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateInput := web.CampaignUpdateRequest{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		User:             userCampaign,
	}

	_, err = h.campaignService.Update(id, updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}
