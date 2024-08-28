package fake

import (
	"task-app/helpers"
	"task-app/models"
	"time"

	"github.com/google/uuid"
)

func User(email, firstName, lastName string) *models.User {
	return &models.User{
		Id:           uuid.New().String(),
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: helpers.Hash("password"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

}
