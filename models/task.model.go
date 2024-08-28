package models

import "time"

type TaskStatus string

const (
	TODO        TaskStatus = "todo"
	IN_PROGRESS TaskStatus = "in-progress"
	COMPLETED   TaskStatus = "completed"
)

// Task is the task model
type Task struct {
	Id          string     `json:"id"`
	Slug        string     `json:"slug"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// CreateTaskDto is the data transfer object to create new task
type CreateTaskDto struct {
	Title       string `json:"title" validate:"required,min=4,max=30"`
	Description string `json:"description" validate:"required,min=4,max=100"`
}

// UpdateTaskDto is the data transfer object to update an existing task
type UpdateTaskDto struct {
	Title       *string     `json:"title" validate:"omitempty,min=4,max=30"`
	Description *string     `json:"description" validate:"omitempty,min=4,max=100"`
	Status      *TaskStatus `json:"status" validate:"omitempty,is_enum"`
}

// Tasks stores all task for all users
var AllTasks = make(map[string][]*Task)

// TasksResponse is the tasks data with pagination info
type TasksResponse struct {
	Tasks      []*Task     `json:"tasks"`
	PagingInfo *PagingInfo `json:"pagingInfo"`
}

// PagingInfo is the pagination info structure
type PagingInfo struct {
	TotalCount  int  `json:"totalCount"`
	Page        int  `json:"page"`
	HasNextPage bool `json:"hasNextPage"`
	Count       int  `json:"count"`
}

// APIPagingDto is the pagination data transfer object
type APIPagingDto struct {
	Limit  int    `json:"limit,omitempty"`
	Filter string `json:"filter,omitempty"`
	Page   int    `json:"page,omitempty"`
}

// IsValid checks if status is valid
func (t TaskStatus) IsValid() bool {
	switch t {
	case TODO, IN_PROGRESS:
		return true
	}
	return false
}
