package controllers

import (
	"github.com/GomdimApps/gmail-2fa/services/client"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		clientRoutes := v1.Group("/clients")
		{
			clientRoutes.POST("/create", client.CreateClientHandler)
			clientRoutes.POST("/login", client.LoginClientHandler)
		}
	}
}
