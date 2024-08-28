package repo

import (
	"fmt"
	"strings"
	"task-app/common/messages"
	"task-app/helpers"
	"task-app/models"
	"time"

	"github.com/rs/zerolog/log"
)

// Task repo object
type Task struct {
}

// TaskRepo exposes task's methods to other packages
type TaskRepo interface {
	Create(task *models.Task, user *models.User) (*models.Task, error)
	GetOneTaskByField(userId string, field string, value interface{}) (*models.Task, *int, error)
	GetAllTasks(userId string, query *models.APIPagingDto) (*models.TasksResponse, error)
	UpdateTaskById(userId string, updateTask *models.Task, taskIndex int) (*models.Task, error)
	DeleteTask(userId string, index int) ([]*models.Task, error)
}

// NewTaskRepo instantiates the Task Repo object
func NewTaskRepo() *TaskRepo {
	task := &Task{}
	taskRepo := TaskRepo(task)
	return &taskRepo
}

// Create creates a new task for user
func (t *Task) Create(task *models.Task, user *models.User) (*models.Task, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// get users task
	tasks, err := t.getAll(user.Id)
	if err != nil && err != messages.ErrTaskNotFound {
		log.Logger.Debug().Msg(fmt.Sprintf("Create Task error : %s", err.Error()))
		return nil, err
	}

	// add to user's tasks
	tasks = append(tasks, task)
	// update all tasks on app
	models.AllTasks[user.Id] = tasks

	return task, nil
}

func (t *Task) getAll(userId string) ([]*models.Task, error) {
	tasks, exists := models.AllTasks[userId]
	if !exists {
		return nil, messages.ErrTaskNotFound
	}
	return tasks, nil
}

// GetOneTaskByField gets one task based on supplied field's value
func (t *Task) GetOneTaskByField(userId string, field string, value interface{}) (*models.Task, *int, error) {
	tasks, err := t.getAll(userId)
	if err != nil {
		return nil, nil, err
	}
	// get one from all user's tasks
	for i := 0; i < len(tasks); i++ {
		// convert struct to map
		taskMap := helpers.StructToMap(tasks[i])
		// check field with value
		if value == taskMap[field] {
			return tasks[i], &i, nil
		}
	}

	return nil, nil, messages.ErrTaskNotFound
}

// UpdateTaskById updates a task
func (t Task) UpdateTaskById(userId string, updateTask *models.Task, taskIndex int) (*models.Task, error) {
	updateTask.UpdatedAt = time.Now()

	// get all user tasks
	tasks, err := t.getAll(userId)
	if err != nil {
		return nil, err
	}

	// update user's tasks list
	tasks[taskIndex] = updateTask

	// update all tasks map
	models.AllTasks[userId] = tasks

	return updateTask, nil
}

// GetAllTasks gets all tasks created by a user
func (t *Task) GetAllTasks(userId string, query *models.APIPagingDto) (*models.TasksResponse, error) {
	allTasks, err := t.getAll(userId)
	if err != nil {
		return nil, err
	}
	// filter task by field
	filteredTasks := t.getTasksByField(allTasks, getWhereObject(query.Filter))
	// get offset
	tasks := t.getOffset(filteredTasks, query.Limit, query.Page)

	paingInfo := getPagingInfo(query, len(filteredTasks))

	paingInfo.TotalCount = len(filteredTasks)
	paingInfo.Count = len(tasks)
	return &models.TasksResponse{
		Tasks:      tasks,
		PagingInfo: paingInfo,
	}, nil
}

// DeleteTask deletes a task with a given task id
func (t *Task) DeleteTask(userId string, index int) ([]*models.Task, error) {
	var tasks []*models.Task
	allTasks, err := t.getAll(userId)
	if err != nil {
		return nil, err
	}

	tasks = append(allTasks[:index], allTasks[index+1:]...)
	// update all tasks map
	models.AllTasks[userId] = tasks
	return tasks, nil
}

// getTasksByField gets tasks by a provided field
func (t *Task) getTasksByField(tasks []*models.Task, where *WhereObj) []*models.Task {
	// do not filter if where object is nil
	if where == nil {
		return tasks
	}
	var filtered []*models.Task
	// get one from all user's tasks
	for i := 0; i < len(tasks); i++ {
		// convert struct to map
		taskMap := helpers.StructToMap(tasks[i])
		switch where.condition {
		case "eq":
			if taskMap[where.field] == where.value {
				filtered = append(filtered, tasks[i])
			}
		case "ne":
			if taskMap[where.field] != where.value {
				filtered = append(filtered, tasks[i])
			}
		case "like":
			if strings.Contains(taskMap[where.field].(string), where.value) {
				filtered = append(filtered, tasks[i])
			}
		}
	}
	return filtered
}

func (t *Task) getOffset(tasks []*models.Task, limit, page int) []*models.Task {
	if len(tasks) <= limit {
		return tasks
	}
	startIndex := (page * limit) - limit
	endIndex := page * limit

	if endIndex > len(tasks) {
		return tasks[startIndex:]
	}
	return tasks[startIndex:endIndex]
}
