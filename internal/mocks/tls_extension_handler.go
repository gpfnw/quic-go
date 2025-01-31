// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gpfnw/quic-go/internal/handshake (interfaces: TLSExtensionHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	mint "github.com/bifurcation/mint"
	gomock "github.com/golang/mock/gomock"
	handshake "github.com/gpfnw/quic-go/internal/handshake"
)

// MockTLSExtensionHandler is a mock of TLSExtensionHandler interface
type MockTLSExtensionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTLSExtensionHandlerMockRecorder
}

// MockTLSExtensionHandlerMockRecorder is the mock recorder for MockTLSExtensionHandler
type MockTLSExtensionHandlerMockRecorder struct {
	mock *MockTLSExtensionHandler
}

// NewMockTLSExtensionHandler creates a new mock instance
func NewMockTLSExtensionHandler(ctrl *gomock.Controller) *MockTLSExtensionHandler {
	mock := &MockTLSExtensionHandler{ctrl: ctrl}
	mock.recorder = &MockTLSExtensionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTLSExtensionHandler) EXPECT() *MockTLSExtensionHandlerMockRecorder {
	return m.recorder
}

// GetPeerParams mocks base method
func (m *MockTLSExtensionHandler) GetPeerParams() <-chan handshake.TransportParameters {
	ret := m.ctrl.Call(m, "GetPeerParams")
	ret0, _ := ret[0].(<-chan handshake.TransportParameters)
	return ret0
}

// GetPeerParams indicates an expected call of GetPeerParams
func (mr *MockTLSExtensionHandlerMockRecorder) GetPeerParams() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerParams", reflect.TypeOf((*MockTLSExtensionHandler)(nil).GetPeerParams))
}

// Receive mocks base method
func (m *MockTLSExtensionHandler) Receive(arg0 mint.HandshakeType, arg1 *mint.ExtensionList) error {
	ret := m.ctrl.Call(m, "Receive", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive
func (mr *MockTLSExtensionHandlerMockRecorder) Receive(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockTLSExtensionHandler)(nil).Receive), arg0, arg1)
}

// Send mocks base method
func (m *MockTLSExtensionHandler) Send(arg0 mint.HandshakeType, arg1 *mint.ExtensionList) error {
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTLSExtensionHandlerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTLSExtensionHandler)(nil).Send), arg0, arg1)
}
