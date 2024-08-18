// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/grpc-course/07-todo-service/internal/todo (interfaces: TaskStore)

// Package store_mock is a generated GoMock package.
package store_mock

import (
	reflect "reflect"

	store "github.com/cshep4/grpc-course/07-todo-service/internal/store"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskStore is a mock of TaskStore interface.
type MockTaskStore struct {
	ctrl     *gomock.Controller
	recorder *MockTaskStoreMockRecorder
}

// MockTaskStoreMockRecorder is the mock recorder for MockTaskStore.
type MockTaskStoreMockRecorder struct {
	mock *MockTaskStore
}

// NewMockTaskStore creates a new mock instance.
func NewMockTaskStore(ctrl *gomock.Controller) *MockTaskStore {
	mock := &MockTaskStore{ctrl: ctrl}
	mock.recorder = &MockTaskStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskStore) EXPECT() *MockTaskStoreMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockTaskStore) AddTask(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTask indicates an expected call of AddTask.
func (mr *MockTaskStoreMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTaskStore)(nil).AddTask), arg0)
}

// CompleteTask mocks base method.
func (m *MockTaskStore) CompleteTask(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteTask indicates an expected call of CompleteTask.
func (mr *MockTaskStoreMockRecorder) CompleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTask", reflect.TypeOf((*MockTaskStore)(nil).CompleteTask), arg0)
}

// ListTasks mocks base method.
func (m *MockTaskStore) ListTasks() ([]store.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].([]store.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockTaskStoreMockRecorder) ListTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskStore)(nil).ListTasks))
}
