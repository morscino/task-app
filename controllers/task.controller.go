package controllers

import (
	"net/http"
	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"

	"github.com/google/uuid"
)

// CreateTask creates a new task
func (c *Controller) CreateTask(data *models.CreateTaskDto, user *models.User) *models.ResponseObject {
	// check if task with same slug exists
	slug := helpers.ToSlug(data.Title)
	existingTask, _, err := c.taskRepo.GetOneTaskByField(user.Id, "slug", slug)
	if err != nil && err != messages.ErrTaskNotFound {
		return &models.ResponseObject{
			Code:    http.StatusInternalServerError,
			Status:  "server-error",
			Error:   messages.ErrServerError,
			Message: messages.ErrServerError.Error(),
		}
	}
	if existingTask != nil {
		return &models.ResponseObject{
			Code:    http.StatusBadRequest,
			Status:  "bad-request",
			Error:   messages.ErrTaskWithSlugAlreadyExists,
			Message: messages.ErrTaskWithSlugAlreadyExists.Error(),
		}
	}

	newTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       data.Title,
		Slug:        slug,
		Description: data.Description,
		Status:      models.TODO,
	}

	task, err := c.taskRepo.Create(newTask, user)
	if err != nil {
		return &models.ResponseObject{
			Code:    http.StatusInternalServerError,
			Status:  "server-error",
			Error:   messages.ErrServerError,
			Message: messages.ErrServerError.Error(),
		}
	}
	return &models.ResponseObject{Code: http.StatusCreated, Data: task, Status: "succes", Message: "task created successfully"}
}

// GetAllTasks gets alltask for a user
func (c *Controller) GetAllTasks(user *models.User, query *models.APIPagingDto) *models.ResponseObject {
	result, err := c.taskRepo.GetAllTasks(user.Id, query)
	if err != nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}
	return &models.ResponseObject{Code: http.StatusOK, Data: result, Status: "succes", Message: "tasks fetched successfully"}
}

// GetTaskById gets a single task by Id
func (c *Controller) GetTaskById(user *models.User, taskId string) *models.ResponseObject {
	task, _, err := c.taskRepo.GetOneTaskByField(user.Id, "id", taskId)
	if err != nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	return &models.ResponseObject{Code: http.StatusOK, Data: task, Status: "succes", Message: "task fetched successfully"}
}

// DeleteTask deletes a single task
func (c *Controller) DeleteTask(user *models.User, taskId string) *models.ResponseObject {
	// get task
	_, index, err := c.taskRepo.GetOneTaskByField(user.Id, "id", taskId)
	if err != nil || index == nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	_, err = c.taskRepo.DeleteTask(user.Id, *index)
	if err != nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	return &models.ResponseObject{Code: http.StatusOK, Status: "succes", Message: "task deleted successfully"}

}

// UpdateTaskbyId updates a task
func (c *Controller) UpdateTaskbyId(user *models.User, taskId string, data *models.UpdateTaskDto) *models.ResponseObject {
	// get task
	task, index, err := c.taskRepo.GetOneTaskByField(user.Id, "id", taskId)
	if err != nil || index == nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	if data.Title != nil {
		// ensure title slug stays unique
		slug := helpers.ToSlug(*data.Title)
		existingTask, _, err := c.taskRepo.GetOneTaskByField("test-user", "slug", slug)
		if err != nil && err != messages.ErrTaskNotFound {
			return &models.ResponseObject{Code: http.StatusInternalServerError, Status: "server-error", Error: messages.ErrServerError, Message: messages.ErrServerError.Error()}
		}
		if existingTask != nil {
			return &models.ResponseObject{Code: http.StatusBadRequest, Status: "bad-request", Error: messages.ErrTaskWithSlugAlreadyExists, Message: messages.ErrTaskWithSlugAlreadyExists.Error()}
		}

		task.Title = *data.Title
		task.Slug = slug
	}

	if data.Description != nil {
		task.Description = *data.Description
	}

	if data.Status != nil {
		task.Status = *data.Status
	}

	task, err = c.taskRepo.UpdateTaskById(user.Id, task, *index)
	if err != nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	return &models.ResponseObject{Code: http.StatusOK, Data: task, Status: "succes", Message: "task updated successfully"}

}

// MarkTaskAsCompleted marks task as completed
func (c *Controller) MarkTaskAsCompleted(user *models.User, taskId string) *models.ResponseObject {
	// get task
	task, index, err := c.taskRepo.GetOneTaskByField(user.Id, "id", taskId)
	if err != nil || index == nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	task.Status = models.COMPLETED
	task.Completed = true

	task, err = c.taskRepo.UpdateTaskById(user.Id, task, *index)
	if err != nil {
		return &models.ResponseObject{Code: http.StatusOK, Status: "no-data-found", Message: err.Error()}
	}

	return &models.ResponseObject{Code: http.StatusOK, Data: task, Status: "succes", Message: "task marked as completed successfully"}

}
