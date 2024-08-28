package messages

import "errors"

var (
	ErrNoDataFound                = errors.New("no data found")
	ErrTaskNotFound               = errors.New("task not found")
	ErrUserNotFound               = errors.New("user not found")
	ErrUserWithEmailAlreadyExists = errors.New("user with email already exists")
	ErrInvalidToken               = errors.New("invalid token")
	ErrWrongPassword              = errors.New("wrong password")
	ErrCouldNotGenerateToken      = errors.New("could not generate user token")
	ErrTaskWithSlugAlreadyExists  = errors.New("task with slug already exists")
	ErrInvalidInput               = errors.New("invalid input")
	ErrServerError                = errors.New("server error")
)
