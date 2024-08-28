package middleware

import (
	"fmt"

	"task-app/common/messages"
	"task-app/config"
	"task-app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Payload struct {
		UserId string `json:"sub"`
		Email  string `json:"email"`
		jwt.RegisteredClaims
	}

	JwtMaker struct {
		secretKey []byte
		config    *config.ConfigType
	}
)

func NewPayload(user *models.User, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		user.Id,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return payload, nil
}

func (p Payload) Valid() error {
	if time.Now().After(p.RegisteredClaims.ExpiresAt.Time) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func NewJwtMaker(config *config.ConfigType) (TokenMaker, error) {
	return &JwtMaker{
		secretKey: []byte(config.JwtSecret),
		config:    config,
	}, nil
}

func (j JwtMaker) CreateAuthToken(user *models.User) (string, error) {
	duration, _ := time.ParseDuration(j.config.JwtSecretExpiry)
	payload, _ := NewPayload(user, duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Create the JWT string
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JwtMaker) VerifyToken(tokenString string) (*Payload, error) {
	payload := Payload{}

	tkn, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, messages.ErrInvalidToken
	}
	return &payload, nil
}
