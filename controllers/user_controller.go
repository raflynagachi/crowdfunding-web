package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	UpdateAvatar(c *gin.Context)
}
