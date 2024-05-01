// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cshep4/grpc-course/module7/internal/transport/grpc (interfaces: ChatService)

// Package chat_mock is a generated GoMock package.
package chat_mock

import (
	context "context"
	reflect "reflect"

	chat "github.com/cshep4/grpc-course/module7/internal/chat"
	gomock "github.com/golang/mock/gomock"
)

// MockChatService is a mock of ChatService interface.
type MockChatService struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceMockRecorder
}

// MockChatServiceMockRecorder is the mock recorder for MockChatService.
type MockChatServiceMockRecorder struct {
	mock *MockChatService
}

// NewMockChatService creates a new mock instance.
func NewMockChatService(ctrl *gomock.Controller) *MockChatService {
	mock := &MockChatService{ctrl: ctrl}
	mock.recorder = &MockChatServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService) EXPECT() *MockChatServiceMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockChatService) SendMessage(arg0 context.Context, arg1 string, arg2 chat.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockChatServiceMockRecorder) SendMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockChatService)(nil).SendMessage), arg0, arg1, arg2)
}

// Subscribe mocks base method.
func (m *MockChatService) Subscribe(arg0 context.Context, arg1, arg2 string) (chan chat.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1, arg2)
	ret0, _ := ret[0].(chan chat.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockChatServiceMockRecorder) Subscribe(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockChatService)(nil).Subscribe), arg0, arg1, arg2)
}

// Unsubscribe mocks base method.
func (m *MockChatService) Unsubscribe(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockChatServiceMockRecorder) Unsubscribe(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockChatService)(nil).Unsubscribe), arg0, arg1, arg2)
}