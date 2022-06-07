package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/crowdfunding-web/services"
)

type transactionHandler struct {
	transaction services.TransactionService
}

func NewTransactionHandler(transaction services.TransactionService) *transactionHandler {
	return &transactionHandler{
		transaction: transaction,
	}
}

func (h *transactionHandler) Index(c *gin.Context) {
	transactions, err := h.transaction.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}
