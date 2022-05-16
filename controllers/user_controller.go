package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}
