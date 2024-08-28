// Code generated by MockGen. DO NOT EDIT.
// Source: ./repo/task.repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	reflect "reflect"
	models "task-app/models"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskRepo is a mock of TaskRepo interface.
type MockTaskRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepoMockRecorder
}

// MockTaskRepoMockRecorder is the mock recorder for MockTaskRepo.
type MockTaskRepoMockRecorder struct {
	mock *MockTaskRepo
}

// NewMockTaskRepo creates a new mock instance.
func NewMockTaskRepo(ctrl *gomock.Controller) *MockTaskRepo {
	mock := &MockTaskRepo{ctrl: ctrl}
	mock.recorder = &MockTaskRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepo) EXPECT() *MockTaskRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskRepo) Create(task *models.Task, user *models.User) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", task, user)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskRepoMockRecorder) Create(task, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepo)(nil).Create), task, user)
}

// DeleteTask mocks base method.
func (m *MockTaskRepo) DeleteTask(userId string, index int) ([]*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", userId, index)
	ret0, _ := ret[0].([]*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskRepoMockRecorder) DeleteTask(userId, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskRepo)(nil).DeleteTask), userId, index)
}

// GetAllTasks mocks base method.
func (m *MockTaskRepo) GetAllTasks(userId string, query *models.APIPagingDto) (*models.TasksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTasks", userId, query)
	ret0, _ := ret[0].(*models.TasksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTasks indicates an expected call of GetAllTasks.
func (mr *MockTaskRepoMockRecorder) GetAllTasks(userId, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTasks", reflect.TypeOf((*MockTaskRepo)(nil).GetAllTasks), userId, query)
}

// GetOneTaskByField mocks base method.
func (m *MockTaskRepo) GetOneTaskByField(userId, field string, value interface{}) (*models.Task, *int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneTaskByField", userId, field, value)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(*int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOneTaskByField indicates an expected call of GetOneTaskByField.
func (mr *MockTaskRepoMockRecorder) GetOneTaskByField(userId, field, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneTaskByField", reflect.TypeOf((*MockTaskRepo)(nil).GetOneTaskByField), userId, field, value)
}

// UpdateTaskById mocks base method.
func (m *MockTaskRepo) UpdateTaskById(userId string, updateTask *models.Task, taskIndex int) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskById", userId, updateTask, taskIndex)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskById indicates an expected call of UpdateTaskById.
func (mr *MockTaskRepoMockRecorder) UpdateTaskById(userId, updateTask, taskIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskById", reflect.TypeOf((*MockTaskRepo)(nil).UpdateTaskById), userId, updateTask, taskIndex)
}
