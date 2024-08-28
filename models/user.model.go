package models

import "time"

// User the user object model
type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// SignUpDto the sign up data transfer object
type SignUpDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,is_password"`
	LastName  string `json:"lastName" validate:"required,min=2,max=25"`
	FirstName string `json:"firstName" validate:"required,min=2,max=25"`
}

// SignInDto the sign in data transfer object
type SignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,is_password"`
}

// AuthenticatedUser the authenticated user object
type AuthenticatedUser struct {
	User        *User  `json:"user"`
	AccessToken string `json:"accessToken"`
}

// AllUsers stores all users
var AllUsers = make(map[string]*User)
