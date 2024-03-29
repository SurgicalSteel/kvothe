// Code generated by MockGen. DO NOT EDIT.
// Source: islack.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	slack "github.com/slack-go/slack"
)

// MockISlack is a mock of ISlack interface.
type MockISlack struct {
	ctrl     *gomock.Controller
	recorder *MockISlackMockRecorder
}

// MockISlackMockRecorder is the mock recorder for MockISlack.
type MockISlackMockRecorder struct {
	mock *MockISlack
}

// NewMockISlack creates a new mock instance.
func NewMockISlack(ctrl *gomock.Controller) *MockISlack {
	mock := &MockISlack{ctrl: ctrl}
	mock.recorder = &MockISlackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISlack) EXPECT() *MockISlackMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockISlack) New(token string, option ...slack.Option) *slack.Client {
	m.ctrl.T.Helper()
	varargs := []interface{}{token}
	for _, a := range option {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "New", varargs...)
	ret0, _ := ret[0].(*slack.Client)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockISlackMockRecorder) New(token interface{}, option ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{token}, option...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockISlack)(nil).New), varargs...)
}

// SendSlack mocks base method.
func (m *MockISlack) SendSlack(channelID string, options ...slack.MsgOption) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{channelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendSlack", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendSlack indicates an expected call of SendSlack.
func (mr *MockISlackMockRecorder) SendSlack(channelID interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSlack", reflect.TypeOf((*MockISlack)(nil).SendSlack), varargs...)
}

// UploadFile mocks base method.
func (m *MockISlack) UploadFile(params slack.FileUploadParameters) (*slack.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", params)
	ret0, _ := ret[0].(*slack.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockISlackMockRecorder) UploadFile(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockISlack)(nil).UploadFile), params)
}
