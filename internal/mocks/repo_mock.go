// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-suggestion-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-suggestion-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddSuggestions mocks base method.
func (m *MockRepo) AddSuggestions(arg0 context.Context, arg1 []models.Suggestion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSuggestions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSuggestions indicates an expected call of AddSuggestions.
func (mr *MockRepoMockRecorder) AddSuggestions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSuggestions", reflect.TypeOf((*MockRepo)(nil).AddSuggestions), arg0, arg1)
}

// CreateSuggestion mocks base method.
func (m *MockRepo) CreateSuggestion(arg0 context.Context, arg1 models.Suggestion) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSuggestion", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSuggestion indicates an expected call of CreateSuggestion.
func (mr *MockRepoMockRecorder) CreateSuggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSuggestion", reflect.TypeOf((*MockRepo)(nil).CreateSuggestion), arg0, arg1)
}

// DescribeSuggestion mocks base method.
func (m *MockRepo) DescribeSuggestion(arg0 context.Context, arg1 uint64) (*models.Suggestion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSuggestion", arg0, arg1)
	ret0, _ := ret[0].(*models.Suggestion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSuggestion indicates an expected call of DescribeSuggestion.
func (mr *MockRepoMockRecorder) DescribeSuggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSuggestion", reflect.TypeOf((*MockRepo)(nil).DescribeSuggestion), arg0, arg1)
}

// ListSuggestions mocks base method.
func (m *MockRepo) ListSuggestions(arg0 context.Context, arg1, arg2 uint64) ([]models.Suggestion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSuggestions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Suggestion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSuggestions indicates an expected call of ListSuggestions.
func (mr *MockRepoMockRecorder) ListSuggestions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuggestions", reflect.TypeOf((*MockRepo)(nil).ListSuggestions), arg0, arg1, arg2)
}

// RemoveSuggestion mocks base method.
func (m *MockRepo) RemoveSuggestion(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSuggestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSuggestion indicates an expected call of RemoveSuggestion.
func (mr *MockRepoMockRecorder) RemoveSuggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSuggestion", reflect.TypeOf((*MockRepo)(nil).RemoveSuggestion), arg0, arg1)
}

// UpdateSuggestion mocks base method.
func (m *MockRepo) UpdateSuggestion(arg0 context.Context, arg1 models.Suggestion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSuggestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSuggestion indicates an expected call of UpdateSuggestion.
func (mr *MockRepoMockRecorder) UpdateSuggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuggestion", reflect.TypeOf((*MockRepo)(nil).UpdateSuggestion), arg0, arg1)
}
