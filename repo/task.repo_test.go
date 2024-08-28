package repo

import (
	"task-app/common/messages"
	"task-app/fake"
	"task-app/helpers"
	"task-app/models"

	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TaskCreate(t *testing.T) {
	tasObj := &Task{}
	testTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       "test task",
		Slug:        helpers.ToSlug("test task"),
		Description: "just for test",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		task        *models.Task
		shouldError bool
		err         error
	}{
		{
			name:        "should create task successfully",
			user:        testUser1,
			task:        testTask,
			shouldError: false,
			err:         nil,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			newTask, err := tasObj.Create(testCase.task, testCase.user)

			if !testCase.shouldError {
				assert.NoError(t, err)
			}
			if testCase.shouldError {
				assert.Error(t, err)
			}

			assert.Equal(t, newTask.Slug, testCase.task.Slug, "the two slugs are the same")

		})
	}
}

func TestGetOneTaskByField(t *testing.T) {
	taskObj := &Task{}
	testTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       "test task",
		Slug:        "test-task",
		Description: "just for test",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		task        *models.Task
		shouldError bool
		field       string
		value       interface{}
		err         error
	}{

		{
			name:        "should not successfully get task by field",
			user:        testUser1,
			task:        testTask,
			field:       "slug",
			value:       "test-taskuuu",
			shouldError: true,
			err:         messages.ErrTaskNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, err := taskObj.GetOneTaskByField(testCase.user.Id, testCase.field, testCase.value)

			if !testCase.shouldError {
				assert.NoError(t, err)
			}
			if testCase.shouldError {
				assert.Error(t, err)
			}

		})
	}
}

func TestUpdateTaskById(t *testing.T) {
	taskObj := &Task{}
	testUpdateTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       "test task",
		Slug:        "test-task",
		Description: "just for test",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		task        *models.Task
		shouldError bool
		taskIndex   int
		err         error
	}{

		{
			name:        "should not successfully update task",
			user:        testUser1,
			task:        testUpdateTask,
			taskIndex:   1,
			shouldError: true,
			err:         messages.ErrTaskNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := taskObj.UpdateTaskById(testCase.user.Id, testCase.task, testCase.taskIndex)

			if !testCase.shouldError {
				assert.NoError(t, err)
			}
			if testCase.shouldError {
				assert.Error(t, err)
			}

		})
	}
}

func TestGetAllTasks(t *testing.T) {
	taskObj := &Task{}
	testTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       "test task",
		Slug:        "test-task",
		Description: "just for test",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		task        *models.Task
		shouldError bool
		taskIndex   int
		err         error
	}{

		{
			name:        "should not successfully fetch tasks",
			user:        testUser1,
			task:        testTask,
			taskIndex:   1,
			shouldError: true,
			err:         messages.ErrTaskNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := taskObj.GetAllTasks(testCase.user.Id, nil)

			if !testCase.shouldError {
				assert.NoError(t, err)
			}
			if testCase.shouldError {
				assert.Error(t, err)
			}

		})
	}
}

func TestDeleteTask(t *testing.T) {
	taskObj := &Task{}
	testTask := &models.Task{
		Id:          uuid.New().String(),
		Title:       "test task",
		Slug:        "test-task",
		Description: "just for test",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testUser1 := fake.User("abc123@ymail.com", "john", "doe")

	cases := []struct {
		name        string
		user        *models.User
		task        *models.Task
		shouldError bool
		taskIndex   int
		err         error
	}{

		{
			name:        "should not successfully delete task",
			user:        testUser1,
			task:        testTask,
			taskIndex:   1,
			shouldError: true,
			err:         messages.ErrTaskNotFound,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := taskObj.DeleteTask(testCase.user.Id, testCase.taskIndex)

			if !testCase.shouldError {
				assert.NoError(t, err)
			}
			if testCase.shouldError {
				assert.Error(t, err)
			}

		})
	}
}
