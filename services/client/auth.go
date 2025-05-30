package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GomdimApps/gmail-2fa/database"
	"github.com/GomdimApps/gmail-2fa/model"
	"github.com/GomdimApps/gmail-2fa/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(input LoginInput) (*model.Client, error) {
	var client model.Client
	if err := database.DB.Where("email = ?", input.Email).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Clear password
	client.Password = ""
	return &client, nil
}

func LoginClientHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedInClient, err := Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed", "details": err.Error()})
		return
	}

	// Generate JWT token
	token, err := util.GenerateUserToken(fmt.Sprintf("%d", loggedInClient.ID), loggedInClient.Email, loggedInClient.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"client": gin.H{
			"id":    loggedInClient.ID,
			"name":  loggedInClient.Name,
			"email": loggedInClient.Email,
			"role":  loggedInClient.Role,
		},
	})
}
