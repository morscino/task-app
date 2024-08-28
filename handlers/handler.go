package handlers

import (
	"fmt"

	"strconv"
	"task-app/common/middleware"

	"task-app/config"
	"task-app/controllers"
	"task-app/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	controller controllers.Operations
}

type Operations interface {
	// middleware
	AuthenticatedUserMiddleware() gin.HandlerFunc

	// tasks
	CreateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	CompletesTask(c *gin.Context)

	// users
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

func NewHandler(config *config.ConfigType) Operations {
	middleware, err := middleware.NewMiddleware(config)
	if err != nil {
		log.Logger.Fatal().Msg(fmt.Sprintf("Create middleware error : %s", err.Error()))
	}
	h := &Handler{
		controller: *controllers.NewController(middleware),
	}
	return Operations(h)
}

func getPagingInfo(c *gin.Context) *models.APIPagingDto {
	var paging models.APIPagingDto

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	paging.Filter = c.Query("filter")

	// default limit is 10
	if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}
	paging.Limit = limit
	paging.Page = page
	return &paging
}
