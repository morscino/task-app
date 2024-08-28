package controllers

import (
	"net/http"
	"strings"

	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"

	"github.com/google/uuid"
)

// RegisterUser signs up users
func (c *Controller) RegisterUser(data *models.SignUpDto) *models.ResponseObject {
	// check if user with email exists
	existingUser, err := c.userRepo.GetOneUserByField("email", strings.ToLower(data.Email))
	if err != nil && err != messages.ErrUserNotFound {
		return &models.ResponseObject{
			Code:    http.StatusInternalServerError,
			Status:  "server-error",
			Error:   messages.ErrServerError,
			Message: messages.ErrServerError.Error(),
		}
	}
	if existingUser != nil {
		return &models.ResponseObject{
			Code:    http.StatusBadRequest,
			Status:  "bad-request",
			Error:   messages.ErrUserWithEmailAlreadyExists,
			Message: messages.ErrUserWithEmailAlreadyExists.Error(),
		}
	}
	newUser := &models.User{
		Id:           uuid.New().String(),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        strings.ToLower(data.Email),
		PasswordHash: helpers.Hash(data.Password),
	}

	user := c.userRepo.CreateUser(newUser)
	return &models.ResponseObject{Code: http.StatusCreated, Data: user, Status: "succes", Message: "user signed up successfully"}
}

// Login logs user in
func (c *Controller) Login(data *models.SignInDto) *models.ResponseObject {
	// get user
	user, err := c.userRepo.GetOneUserByField("email", strings.ToLower(data.Email))
	if err != nil {
		return &models.ResponseObject{
			Code:    http.StatusOK,
			Status:  "no-data-found",
			Message: err.Error(),
		}
	}

	// verify password
	if isValid := helpers.CompareHash(user.PasswordHash, data.Password); !isValid {
		return &models.ResponseObject{
			Code:    http.StatusBadRequest,
			Status:  "no-data-found",
			Message: messages.ErrWrongPassword.Error(),
		}
	}

	// generate jwt tokens
	token, err := c.middleware.Jwt.CreateAuthToken(user)
	if err != nil {
		return &models.ResponseObject{
			Code:    http.StatusInternalServerError,
			Status:  "server-erros",
			Message: err.Error(),
			Error:   err,
		}
	}
	authUser := &models.AuthenticatedUser{
		User:        user,
		AccessToken: token,
	}

	return &models.ResponseObject{Code: http.StatusCreated, Data: authUser, Status: "succes", Message: "user logged in successfully"}

}
