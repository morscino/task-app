package controllers

import (
	mock_controllers "task-app/controllers/mock"
	"task-app/fake"
	"task-app/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	response := &models.ResponseObject{
		Code:   201,
		Status: "success",
		Data:   fake.User("abc1@ymail.com", "ade", "yemi"),
	}

	testSingupDto := &models.SignUpDto{
		Email:     "abc1@ymail.com",
		Password:  "Y65far$$nhdyi",
		FirstName: "ade",
		LastName:  "yemi",
	}

	operations := mock_controllers.NewMockOperations(ctrl)
	operations.EXPECT().RegisterUser(testSingupDto).Return(response)

	result := operations.RegisterUser(testSingupDto)
	assert.Equal(t, result.Code, response.Code)

}
