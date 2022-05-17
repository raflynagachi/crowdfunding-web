package controllers

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
