package client

import (
	"net/http"

	"github.com/GomdimApps/gmail-2fa/database"
	"github.com/GomdimApps/gmail-2fa/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateClientInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func CreateClient(input CreateClientInput) (*model.Client, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	client := model.Client{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "client",
	}

	if err := database.DB.Create(&client).Error; err != nil {
		return nil, err
	}

	// Clear password
	client.Password = ""
	return &client, nil
}

func CreateClientHandler(c *gin.Context) {
	var input CreateClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdClient, err := CreateClient(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdClient)
}
