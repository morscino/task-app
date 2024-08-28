package controllers

import (
	"task-app/common/middleware"
	"task-app/config"
	"task-app/models"
	"task-app/repo"
)

// Cobtroller is the controller object
type Controller struct {
	middleware *middleware.Middleware
	Config     *config.ConfigType

	taskRepo repo.TaskRepo
	userRepo repo.UserRepo
}

// Operations registers all controllers method
type Operations interface {
	Middleware() *middleware.Middleware
	// tasks
	CreateTask(data *models.CreateTaskDto, user *models.User) *models.ResponseObject
	GetAllTasks(user *models.User, query *models.APIPagingDto) *models.ResponseObject
	GetTaskById(user *models.User, taskId string) *models.ResponseObject
	UpdateTaskbyId(user *models.User, taskId string, data *models.UpdateTaskDto) *models.ResponseObject
	DeleteTask(user *models.User, taskId string) *models.ResponseObject
	MarkTaskAsCompleted(user *models.User, taskId string) *models.ResponseObject

	// users
	RegisterUser(data *models.SignUpDto) *models.ResponseObject
	Login(data *models.SignInDto) *models.ResponseObject
}

// NewController loads all controllers resources
func NewController(middleware *middleware.Middleware) *Operations {
	c := &Controller{
		middleware: middleware,
		taskRepo:   *repo.NewTaskRepo(),
		userRepo:   *repo.NewUserRepo(),
	}
	op := Operations(c)

	return &op
}
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}
