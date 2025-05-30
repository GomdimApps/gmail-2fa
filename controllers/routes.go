package controllers

import (
	"net/http"

	"github.com/GomdimApps/gmail-2fa/services/client"
	"github.com/gin-gonic/gin"
)

func CreateClientHandler(c *gin.Context) {
	var input client.CreateClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdClient, err := client.CreateClient(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdClient)
}

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		clientRoutes := v1.Group("/clients")
		{
			clientRoutes.POST("", CreateClientHandler)
		}
	}
}
