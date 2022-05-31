package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/helpers"
	"github.com/raflynagachi/crowdfunding-web/models"
	"github.com/raflynagachi/crowdfunding-web/models/web"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type TransactionControllerImpl struct {
	service        services.TransactionService
	paymentService services.PaymentService
}

func NewTransactionController(service services.TransactionService, payment services.PaymentService) TransactionController {
	return &TransactionControllerImpl{
		service:        service,
		paymentService: payment,
	}
}

func (controller *TransactionControllerImpl) FindByCampaignID(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	campaignID, _ := strconv.Atoi(c.Param("campaignID"))
	if campaignID == 0 {
		webResponse.Data = gin.H{"errors": "campaignID doesn't exist"}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	userID := c.MustGet("user").(models.User).ID

	transaction, err := controller.service.FindByCampaignID(campaignID, userID)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = transaction
	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) FindByUserID(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	currentUser := c.MustGet("user").(models.User)
	transactions, err := controller.service.FindByUserID(currentUser.ID)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = transactions
	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) Create(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	var createReq web.CreateTransactionRequest
	err := c.ShouldBindJSON(&createReq)
	if err != nil {
		webResponse.Data = gin.H{"errors": helpers.ValidationErrorsToSlice(err)}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	currentUser := c.MustGet("user").(models.User)
	createReq.User = currentUser

	transaction, err := controller.service.Create(createReq)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = transaction
	c.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionControllerImpl) GetNotification(c *gin.Context) {
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
	}

	var transactionReq web.NotificationTransactionRequest
	err := c.ShouldBindJSON(&transactionReq)

	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	err = controller.paymentService.ProcessPayment(transactionReq)
	if err != nil {
		webResponse.Data = gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse.Code = http.StatusOK
	webResponse.Status = "OK"
	webResponse.Data = ""
	c.JSON(http.StatusOK, webResponse)
}
