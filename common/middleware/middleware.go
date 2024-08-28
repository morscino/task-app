package middleware

import (
	"task-app/common/messages"
	"task-app/config"
	"task-app/models"
	"task-app/repo"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	authorizationHeader = "Authorization"
	authorizationBearer = "Bearer"
)

type TokenMaker interface {
	CreateAuthToken(user *models.User) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Middleware struct {
	Jwt      TokenMaker
	logger   *zerolog.Logger
	userRepo repo.UserRepo
	config   *config.ConfigType
}

func NewMiddleware(config *config.ConfigType) (*Middleware, error) {
	l := log.With().Str("middleware", "api").Logger()

	jwt, err := NewJwtMaker(config)
	if err != nil {
		return nil, err
	}

	m := &Middleware{
		Jwt:      jwt,
		logger:   &l,
		config:   config,
		userRepo: *repo.NewUserRepo(),
	}

	return m, nil
}

// JwtUserAuth hybrid middleware returns an authorized user
func (m *Middleware) JwtUserAuth(c *gin.Context) (*models.User, error) {
	authorization := c.GetHeader(authorizationHeader)
	if len(authorization) < 1 {
		return nil, messages.ErrInvalidToken
	}

	fields := strings.Fields(authorization)
	if len(fields) != 2 {
		return nil, messages.ErrInvalidToken
	}

	return m.getUserFromToken(fields[1])
}

func (m *Middleware) getUserFromToken(token string) (*models.User, error) {
	verified, err := m.Jwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	return m.userRepo.GetOneUserByField("id", verified.UserId)
}
