// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gpfnw/quic-go/congestion (interfaces: SendAlgorithm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/gpfnw/quic-go/internal/protocol"
)

// MockSendAlgorithm is a mock of SendAlgorithm interface
type MockSendAlgorithm struct {
	ctrl     *gomock.Controller
	recorder *MockSendAlgorithmMockRecorder
}

// MockSendAlgorithmMockRecorder is the mock recorder for MockSendAlgorithm
type MockSendAlgorithmMockRecorder struct {
	mock *MockSendAlgorithm
}

// NewMockSendAlgorithm creates a new mock instance
func NewMockSendAlgorithm(ctrl *gomock.Controller) *MockSendAlgorithm {
	mock := &MockSendAlgorithm{ctrl: ctrl}
	mock.recorder = &MockSendAlgorithmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendAlgorithm) EXPECT() *MockSendAlgorithmMockRecorder {
	return m.recorder
}

// GetCongestionWindow mocks base method
func (m *MockSendAlgorithm) GetCongestionWindow() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetCongestionWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetCongestionWindow indicates an expected call of GetCongestionWindow
func (mr *MockSendAlgorithmMockRecorder) GetCongestionWindow() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCongestionWindow", reflect.TypeOf((*MockSendAlgorithm)(nil).GetCongestionWindow))
}

// MaybeExitSlowStart mocks base method
func (m *MockSendAlgorithm) MaybeExitSlowStart() {
	m.ctrl.Call(m, "MaybeExitSlowStart")
}

// MaybeExitSlowStart indicates an expected call of MaybeExitSlowStart
func (mr *MockSendAlgorithmMockRecorder) MaybeExitSlowStart() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeExitSlowStart", reflect.TypeOf((*MockSendAlgorithm)(nil).MaybeExitSlowStart))
}

// OnConnectionMigration mocks base method
func (m *MockSendAlgorithm) OnConnectionMigration() {
	m.ctrl.Call(m, "OnConnectionMigration")
}

// OnConnectionMigration indicates an expected call of OnConnectionMigration
func (mr *MockSendAlgorithmMockRecorder) OnConnectionMigration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnConnectionMigration", reflect.TypeOf((*MockSendAlgorithm)(nil).OnConnectionMigration))
}

// OnPacketAcked mocks base method
func (m *MockSendAlgorithm) OnPacketAcked(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount) {
	m.ctrl.Call(m, "OnPacketAcked", arg0, arg1, arg2)
}

// OnPacketAcked indicates an expected call of OnPacketAcked
func (mr *MockSendAlgorithmMockRecorder) OnPacketAcked(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketAcked", reflect.TypeOf((*MockSendAlgorithm)(nil).OnPacketAcked), arg0, arg1, arg2)
}

// OnPacketLost mocks base method
func (m *MockSendAlgorithm) OnPacketLost(arg0 protocol.PacketNumber, arg1, arg2 protocol.ByteCount) {
	m.ctrl.Call(m, "OnPacketLost", arg0, arg1, arg2)
}

// OnPacketLost indicates an expected call of OnPacketLost
func (mr *MockSendAlgorithmMockRecorder) OnPacketLost(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketLost", reflect.TypeOf((*MockSendAlgorithm)(nil).OnPacketLost), arg0, arg1, arg2)
}

// OnPacketSent mocks base method
func (m *MockSendAlgorithm) OnPacketSent(arg0 time.Time, arg1 protocol.ByteCount, arg2 protocol.PacketNumber, arg3 protocol.ByteCount, arg4 bool) bool {
	ret := m.ctrl.Call(m, "OnPacketSent", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	return ret0
}

// OnPacketSent indicates an expected call of OnPacketSent
func (mr *MockSendAlgorithmMockRecorder) OnPacketSent(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPacketSent", reflect.TypeOf((*MockSendAlgorithm)(nil).OnPacketSent), arg0, arg1, arg2, arg3, arg4)
}

// OnRetransmissionTimeout mocks base method
func (m *MockSendAlgorithm) OnRetransmissionTimeout(arg0 bool) {
	m.ctrl.Call(m, "OnRetransmissionTimeout", arg0)
}

// OnRetransmissionTimeout indicates an expected call of OnRetransmissionTimeout
func (mr *MockSendAlgorithmMockRecorder) OnRetransmissionTimeout(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRetransmissionTimeout", reflect.TypeOf((*MockSendAlgorithm)(nil).OnRetransmissionTimeout), arg0)
}

// RetransmissionDelay mocks base method
func (m *MockSendAlgorithm) RetransmissionDelay() time.Duration {
	ret := m.ctrl.Call(m, "RetransmissionDelay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetransmissionDelay indicates an expected call of RetransmissionDelay
func (mr *MockSendAlgorithmMockRecorder) RetransmissionDelay() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetransmissionDelay", reflect.TypeOf((*MockSendAlgorithm)(nil).RetransmissionDelay))
}

// SetNumEmulatedConnections mocks base method
func (m *MockSendAlgorithm) SetNumEmulatedConnections(arg0 int) {
	m.ctrl.Call(m, "SetNumEmulatedConnections", arg0)
}

// SetNumEmulatedConnections indicates an expected call of SetNumEmulatedConnections
func (mr *MockSendAlgorithmMockRecorder) SetNumEmulatedConnections(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNumEmulatedConnections", reflect.TypeOf((*MockSendAlgorithm)(nil).SetNumEmulatedConnections), arg0)
}

// SetSlowStartLargeReduction mocks base method
func (m *MockSendAlgorithm) SetSlowStartLargeReduction(arg0 bool) {
	m.ctrl.Call(m, "SetSlowStartLargeReduction", arg0)
}

// SetSlowStartLargeReduction indicates an expected call of SetSlowStartLargeReduction
func (mr *MockSendAlgorithmMockRecorder) SetSlowStartLargeReduction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSlowStartLargeReduction", reflect.TypeOf((*MockSendAlgorithm)(nil).SetSlowStartLargeReduction), arg0)
}

// TimeUntilSend mocks base method
func (m *MockSendAlgorithm) TimeUntilSend(arg0 time.Time, arg1 protocol.ByteCount) time.Duration {
	ret := m.ctrl.Call(m, "TimeUntilSend", arg0, arg1)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeUntilSend indicates an expected call of TimeUntilSend
func (mr *MockSendAlgorithmMockRecorder) TimeUntilSend(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilSend", reflect.TypeOf((*MockSendAlgorithm)(nil).TimeUntilSend), arg0, arg1)
}
