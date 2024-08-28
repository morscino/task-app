package repo

import (
	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"
	"time"
)

// User repo object
type User struct{}

// TaskRepo exposes user's methods to other packages
type UserRepo interface {
	CreateUser(user *models.User) *models.User
	GetOneUserByField(field string, value interface{}) (*models.User, error)
}

// NewUserRepo instantiates the User Repo object
func NewUserRepo() *UserRepo {
	user := &User{}
	userRepo := UserRepo(user)
	return &userRepo
}

// Create stores a new user
func (u *User) CreateUser(user *models.User) *models.User {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// store user
	models.AllUsers[user.Id] = user

	return user
}

// GetOneUserByField gets one user based on the supplied field's value
func (u *User) GetOneUserByField(field string, value interface{}) (*models.User, error) {
	// get one user from all users
	for _, user := range models.AllUsers {
		// convert struct to map
		userMap := helpers.StructToMap(user)
		// check field with value
		if value == userMap[field] {
			return user, nil
		}
	}
	return nil, messages.ErrUserNotFound
}

func (u *User) getUser(userId string) (*models.User, error) {
	user, exists := models.AllUsers[userId]
	if !exists {
		return nil, messages.ErrUserNotFound
	}
	return user, nil
}
