package controllers

import (
	"testing"

	"task-app/common/messages"
	mock_controllers "task-app/controllers/mock"
	"task-app/fake"
	"task-app/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testResponse1 := &models.ResponseObject{
		Code:   201,
		Status: "success",
		Data:   fake.User("abc1@ymail.com", "ade", "yemi"),
	}

	testResponse2 := &models.ResponseObject{
		Code:   400,
		Status: "bad-request",
		Data:   fake.User("abc1@ymail.com", "ade", "yemi"),
		Error:  messages.ErrUserWithEmailAlreadyExists,
	}

	testSingupDto := &models.SignUpDto{
		Email:     "abc1@ymail.com",
		Password:  "Y65far$$nhdyi",
		FirstName: "ade",
		LastName:  "yemi",
	}

	cases := []struct {
		name        string
		dto         *models.SignUpDto
		shouldError bool
		response    *models.ResponseObject
		err         error
	}{
		{
			name:        "should create user successfully",
			dto:         testSingupDto,
			shouldError: false,
			response:    testResponse1,
			err:         nil,
		},
		{
			name:        "user email lready exists",
			dto:         testSingupDto,
			shouldError: true,
			response:    testResponse2,
			err:         messages.ErrUserWithEmailAlreadyExists,
		},
	}

	for _, testCase := range cases {
		operations := mock_controllers.NewMockOperations(ctrl)
		operations.EXPECT().RegisterUser(testCase.dto).Return(testCase.response)

		t.Run(testCase.name, func(t *testing.T) {
			result := operations.RegisterUser(testCase.dto)

			if !testCase.shouldError {
				user := result.Data.(*models.User)
				assert.NoError(t, nil)
				assert.Equal(t, user.Email, testCase.dto.Email, "the two emails are the same")
			}
			if testCase.shouldError {
				assert.Error(t, result.Error.(error), testCase.err)
			}

			assert.Equal(t, result.Code, testCase.response.Code)

		})
	}

}
