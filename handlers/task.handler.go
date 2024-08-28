package handlers

import (
	"net/http"
	"strings"
	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"

	"github.com/gin-gonic/gin"
)

// @Tags Task
// @Summary Create new Task
// @Schemes
// @Description Creates a new task
// @Param   request   body     models.CreateTaskDto   true  "data to create new task"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseObject "desc"
// @Router /tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var input models.CreateTaskDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {

		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	user := c.MustGet("authUser").(*models.User) // auth user
	// send to controller
	result := h.controller.CreateTask(&input, user)
	c.JSON(result.Code, result)
}

// @Tags Task
// @Summary Get All Tasks
// @Description Gets All tasks
// @Accept  json
// @Produce  json
// @Param   request   body     models.APIPagingDto   true  "data to query for all "
// @Success 200 {string} {object} models.ResponseObject{data=models.TasksResponse} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks [get]
func (h *Handler) GetTasks(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.GetAllTasks(user, query)
	c.JSON(result.Code, result)
}

// @Tags Task
// @Summary Get Single task
// @Description Get Single task by id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Task Id"
// @Success 200 {string} {object} models.ResponseObject{data=models.Task} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks/{id} [get]
func (h *Handler) GetTaskById(c *gin.Context) {
	id := strings.Trim(c.Param("id"), " ")

	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.GetTaskById(user, id)
	c.JSON(result.Code, result)
}

// @Tags Task
// @Summary Update Task
// @Description Upadates Task with a guven Id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Task Id"
// @Param   request   body     models.UpdateTaskDto   true  "data to update task with"
// @Success 200 {string} {object} models.ResponseObject{data=models.Task} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	id := strings.Trim(c.Param("id"), " ")
	var input models.UpdateTaskDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {

		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.UpdateTaskbyId(user, id, &input)
	c.JSON(result.Code, result)
}

// @Tags Task
// @Summary Delete task
// @Description Delete task by id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Task Id"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	id := strings.Trim(c.Param("id"), " ")

	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.DeleteTask(user, id)
	c.JSON(result.Code, result)
}

// @Tags Task
// @Summary Complete task
// @Description Marks task as completed
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Task Id"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /tasks/{id}/complete [put]
func (h *Handler) CompletesTask(c *gin.Context) {
	id := strings.Trim(c.Param("id"), " ")

	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.MarkTaskAsCompleted(user, id)
	c.JSON(result.Code, result)
}
