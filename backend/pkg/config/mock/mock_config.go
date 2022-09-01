// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	viper "github.com/spf13/viper"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AllSettings mocks base method.
func (m *MockInterface) AllSettings() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSettings")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// AllSettings indicates an expected call of AllSettings.
func (mr *MockInterfaceMockRecorder) AllSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSettings", reflect.TypeOf((*MockInterface)(nil).AllSettings))
}

// Get mocks base method.
func (m *MockInterface) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), key)
}

// GetBool mocks base method.
func (m *MockInterface) GetBool(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBool", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool.
func (mr *MockInterfaceMockRecorder) GetBool(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockInterface)(nil).GetBool), key)
}

// GetInt mocks base method.
func (m *MockInterface) GetInt(key string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt", key)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt.
func (mr *MockInterfaceMockRecorder) GetInt(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockInterface)(nil).GetInt), key)
}

// GetString mocks base method.
func (m *MockInterface) GetString(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString.
func (mr *MockInterfaceMockRecorder) GetString(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockInterface)(nil).GetString), key)
}

// GetStringSlice mocks base method.
func (m *MockInterface) GetStringSlice(key string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringSlice", key)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringSlice indicates an expected call of GetStringSlice.
func (mr *MockInterfaceMockRecorder) GetStringSlice(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringSlice", reflect.TypeOf((*MockInterface)(nil).GetStringSlice), key)
}

// Init mocks base method.
func (m *MockInterface) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockInterfaceMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockInterface)(nil).Init))
}

// IsSet mocks base method.
func (m *MockInterface) IsSet(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSet", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSet indicates an expected call of IsSet.
func (mr *MockInterfaceMockRecorder) IsSet(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSet", reflect.TypeOf((*MockInterface)(nil).IsSet), key)
}

// ReadConfig mocks base method.
func (m *MockInterface) ReadConfig(configFilePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConfig", configFilePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadConfig indicates an expected call of ReadConfig.
func (mr *MockInterfaceMockRecorder) ReadConfig(configFilePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfig", reflect.TypeOf((*MockInterface)(nil).ReadConfig), configFilePath)
}

// Set mocks base method.
func (m *MockInterface) Set(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockInterfaceMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInterface)(nil).Set), key, value)
}

// SetDefault mocks base method.
func (m *MockInterface) SetDefault(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDefault", key, value)
}

// SetDefault indicates an expected call of SetDefault.
func (mr *MockInterfaceMockRecorder) SetDefault(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefault", reflect.TypeOf((*MockInterface)(nil).SetDefault), key, value)
}

// UnmarshalKey mocks base method.
func (m *MockInterface) UnmarshalKey(key string, rawVal interface{}, decoder ...viper.DecoderConfigOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{key, rawVal}
	for _, a := range decoder {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnmarshalKey", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalKey indicates an expected call of UnmarshalKey.
func (mr *MockInterfaceMockRecorder) UnmarshalKey(key, rawVal interface{}, decoder ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key, rawVal}, decoder...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalKey", reflect.TypeOf((*MockInterface)(nil).UnmarshalKey), varargs...)
}
