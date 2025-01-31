// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gpfnw/quic-go/internal/crypto (interfaces: AEAD)

// Package mockcrypto is a generated GoMock package.
package mockcrypto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/gpfnw/quic-go/internal/protocol"
)

// MockAEAD is a mock of AEAD interface
type MockAEAD struct {
	ctrl     *gomock.Controller
	recorder *MockAEADMockRecorder
}

// MockAEADMockRecorder is the mock recorder for MockAEAD
type MockAEADMockRecorder struct {
	mock *MockAEAD
}

// NewMockAEAD creates a new mock instance
func NewMockAEAD(ctrl *gomock.Controller) *MockAEAD {
	mock := &MockAEAD{ctrl: ctrl}
	mock.recorder = &MockAEADMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAEAD) EXPECT() *MockAEADMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockAEAD) Open(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Open", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockAEADMockRecorder) Open(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockAEAD)(nil).Open), arg0, arg1, arg2, arg3)
}

// Overhead mocks base method
func (m *MockAEAD) Overhead() int {
	ret := m.ctrl.Call(m, "Overhead")
	ret0, _ := ret[0].(int)
	return ret0
}

// Overhead indicates an expected call of Overhead
func (mr *MockAEADMockRecorder) Overhead() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overhead", reflect.TypeOf((*MockAEAD)(nil).Overhead))
}

// Seal mocks base method
func (m *MockAEAD) Seal(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) []byte {
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockAEADMockRecorder) Seal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockAEAD)(nil).Seal), arg0, arg1, arg2, arg3)
}
