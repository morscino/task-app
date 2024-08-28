package repo

import (
	"task-app/common/messages"
	"task-app/fake"
	"task-app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *Suite) TestCreateUser(t *testing.T) {
	userObj := &User{}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		shouldError bool
		err         error
	}{
		{
			name:        "should create user successfully",
			user:        testUser1,
			shouldError: false,
			err:         nil,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			newUser := userObj.CreateUser(testCase.user)

			if !testCase.shouldError {
				assert.NoError(t, nil)
			}
			if testCase.shouldError {
				assert.Error(t, nil)
			}

			assert.Equal(t, newUser.Email, testCase.user.Email, "the two emails are the same")

		})
	}
}

func (s *Suite) TestGetOneUserByField(t *testing.T) {
	userObj := &User{}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		shouldError bool
		field       string
		value       interface{}
		err         error
	}{
		{
			name:        "should fetch user by field successfully",
			user:        testUser1,
			shouldError: false,
			err:         nil,
			field:       "email",
			value:       "abc123@ymail.com",
		},
		{
			name:        "should not fetch user by field successfully",
			user:        testUser1,
			shouldError: true,
			err:         messages.ErrUserNotFound,
			field:       "emails",
			value:       "abc123@ymail.com",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			newUser, err := userObj.GetOneUserByField(testCase.field, testCase.value)

			if !testCase.shouldError {
				assert.NoError(t, err)
				assert.Equal(t, newUser.Email, testCase.user.Email, "the two emails are the same")
			}
			if testCase.shouldError {
				if assert.Error(t, err) {
					assert.Equal(t, err, testCase.err)
				}
			}

		})
	}
}
