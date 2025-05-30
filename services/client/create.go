package client

import (
	"github.com/GomdimApps/gmail-2fa/database"
	"github.com/GomdimApps/gmail-2fa/model"
)

type CreateClientInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func CreateClient(input CreateClientInput) (*model.Client, error) {
	client := model.Client{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := database.DB.Create(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}
