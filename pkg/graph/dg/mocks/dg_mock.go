// Code generated by MockGen. DO NOT EDIT.
// Source: dg.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dg "github.com/XDTD/Dragonfly2/pkg/graph/dg"
	gomock "github.com/golang/mock/gomock"
)

// MockDG is a mock of DG interface.
type MockDG[T comparable] struct {
	ctrl     *gomock.Controller
	recorder *MockDGMockRecorder[T]
}

// MockDGMockRecorder is the mock recorder for MockDG.
type MockDGMockRecorder[T comparable] struct {
	mock *MockDG[T]
}

// NewMockDG creates a new mock instance.
func NewMockDG[T comparable](ctrl *gomock.Controller) *MockDG[T] {
	mock := &MockDG[T]{ctrl: ctrl}
	mock.recorder = &MockDGMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDG[T]) EXPECT() *MockDGMockRecorder[T] {
	return m.recorder
}

// AddEdge mocks base method.
func (m *MockDG[T]) AddEdge(fromVertexID, toVertexID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEdge", fromVertexID, toVertexID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEdge indicates an expected call of AddEdge.
func (mr *MockDGMockRecorder[T]) AddEdge(fromVertexID, toVertexID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEdge", reflect.TypeOf((*MockDG[T])(nil).AddEdge), fromVertexID, toVertexID)
}

// AddVertex mocks base method.
func (m *MockDG[T]) AddVertex(id string, value T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVertex", id, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVertex indicates an expected call of AddVertex.
func (mr *MockDGMockRecorder[T]) AddVertex(id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVertex", reflect.TypeOf((*MockDG[T])(nil).AddVertex), id, value)
}

// CanAddEdge mocks base method.
func (m *MockDG[T]) CanAddEdge(fromVertexID, toVertexID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanAddEdge", fromVertexID, toVertexID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanAddEdge indicates an expected call of CanAddEdge.
func (mr *MockDGMockRecorder[T]) CanAddEdge(fromVertexID, toVertexID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanAddEdge", reflect.TypeOf((*MockDG[T])(nil).CanAddEdge), fromVertexID, toVertexID)
}

// DeleteEdge mocks base method.
func (m *MockDG[T]) DeleteEdge(fromVertexID, toVertexID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEdge", fromVertexID, toVertexID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEdge indicates an expected call of DeleteEdge.
func (mr *MockDGMockRecorder[T]) DeleteEdge(fromVertexID, toVertexID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEdge", reflect.TypeOf((*MockDG[T])(nil).DeleteEdge), fromVertexID, toVertexID)
}

// DeleteVertex mocks base method.
func (m *MockDG[T]) DeleteVertex(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteVertex", id)
}

// DeleteVertex indicates an expected call of DeleteVertex.
func (mr *MockDGMockRecorder[T]) DeleteVertex(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVertex", reflect.TypeOf((*MockDG[T])(nil).DeleteVertex), id)
}

// GetRandomVertices mocks base method.
func (m *MockDG[T]) GetRandomVertices(n uint) []*dg.Vertex[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomVertices", n)
	ret0, _ := ret[0].([]*dg.Vertex[T])
	return ret0
}

// GetRandomVertices indicates an expected call of GetRandomVertices.
func (mr *MockDGMockRecorder[T]) GetRandomVertices(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomVertices", reflect.TypeOf((*MockDG[T])(nil).GetRandomVertices), n)
}

// GetSinkVertices mocks base method.
func (m *MockDG[T]) GetSinkVertices() []*dg.Vertex[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSinkVertices")
	ret0, _ := ret[0].([]*dg.Vertex[T])
	return ret0
}

// GetSinkVertices indicates an expected call of GetSinkVertices.
func (mr *MockDGMockRecorder[T]) GetSinkVertices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSinkVertices", reflect.TypeOf((*MockDG[T])(nil).GetSinkVertices))
}

// GetSourceVertices mocks base method.
func (m *MockDG[T]) GetSourceVertices() []*dg.Vertex[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceVertices")
	ret0, _ := ret[0].([]*dg.Vertex[T])
	return ret0
}

// GetSourceVertices indicates an expected call of GetSourceVertices.
func (mr *MockDGMockRecorder[T]) GetSourceVertices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceVertices", reflect.TypeOf((*MockDG[T])(nil).GetSourceVertices))
}

// GetVertex mocks base method.
func (m *MockDG[T]) GetVertex(id string) (*dg.Vertex[T], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVertex", id)
	ret0, _ := ret[0].(*dg.Vertex[T])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVertex indicates an expected call of GetVertex.
func (mr *MockDGMockRecorder[T]) GetVertex(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVertex", reflect.TypeOf((*MockDG[T])(nil).GetVertex), id)
}

// GetVertexKeys mocks base method.
func (m *MockDG[T]) GetVertexKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVertexKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetVertexKeys indicates an expected call of GetVertexKeys.
func (mr *MockDGMockRecorder[T]) GetVertexKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVertexKeys", reflect.TypeOf((*MockDG[T])(nil).GetVertexKeys))
}

// GetVertices mocks base method.
func (m *MockDG[T]) GetVertices() map[string]*dg.Vertex[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVertices")
	ret0, _ := ret[0].(map[string]*dg.Vertex[T])
	return ret0
}

// GetVertices indicates an expected call of GetVertices.
func (mr *MockDGMockRecorder[T]) GetVertices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVertices", reflect.TypeOf((*MockDG[T])(nil).GetVertices))
}

// VertexCount mocks base method.
func (m *MockDG[T]) VertexCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VertexCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// VertexCount indicates an expected call of VertexCount.
func (mr *MockDGMockRecorder[T]) VertexCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VertexCount", reflect.TypeOf((*MockDG[T])(nil).VertexCount))
}
