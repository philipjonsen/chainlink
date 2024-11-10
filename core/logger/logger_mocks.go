// Code generated by mockery v2.46.3. DO NOT EDIT.

package logger

import (
	mock "github.com/stretchr/testify/mock"
	zapcore "go.uber.org/zap/zapcore"
)

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

type MockLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLogger) EXPECT() *MockLogger_Expecter {
	return &MockLogger_Expecter{mock: &_m.Mock}
}

// Critical provides a mock function with given fields: args
func (_m *MockLogger) Critical(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Critical_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Critical'
type MockLogger_Critical_Call struct {
	*mock.Call
}

// Critical is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Critical(args ...interface{}) *MockLogger_Critical_Call {
	return &MockLogger_Critical_Call{Call: _e.mock.On("Critical",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Critical_Call) Run(run func(args ...interface{})) *MockLogger_Critical_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Critical_Call) Return() *MockLogger_Critical_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Critical_Call) RunAndReturn(run func(...interface{})) *MockLogger_Critical_Call {
	_c.Call.Return(run)
	return _c
}

// Criticalf provides a mock function with given fields: format, values
func (_m *MockLogger) Criticalf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Criticalf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Criticalf'
type MockLogger_Criticalf_Call struct {
	*mock.Call
}

// Criticalf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Criticalf(format interface{}, values ...interface{}) *MockLogger_Criticalf_Call {
	return &MockLogger_Criticalf_Call{Call: _e.mock.On("Criticalf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Criticalf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Criticalf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Criticalf_Call) Return() *MockLogger_Criticalf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Criticalf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Criticalf_Call {
	_c.Call.Return(run)
	return _c
}

// Criticalw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Criticalw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Criticalw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Criticalw'
type MockLogger_Criticalw_Call struct {
	*mock.Call
}

// Criticalw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Criticalw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Criticalw_Call {
	return &MockLogger_Criticalw_Call{Call: _e.mock.On("Criticalw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Criticalw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Criticalw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Criticalw_Call) Return() *MockLogger_Criticalw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Criticalw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Criticalw_Call {
	_c.Call.Return(run)
	return _c
}

// Debug provides a mock function with given fields: args
func (_m *MockLogger) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type MockLogger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Debug(args ...interface{}) *MockLogger_Debug_Call {
	return &MockLogger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Debug_Call) Run(run func(args ...interface{})) *MockLogger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debug_Call) Return() *MockLogger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debug_Call) RunAndReturn(run func(...interface{})) *MockLogger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Debugf provides a mock function with given fields: format, values
func (_m *MockLogger) Debugf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type MockLogger_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Debugf(format interface{}, values ...interface{}) *MockLogger_Debugf_Call {
	return &MockLogger_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Debugf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debugf_Call) Return() *MockLogger_Debugf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debugf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Debugf_Call {
	_c.Call.Return(run)
	return _c
}

// Debugw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Debugw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Debugw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugw'
type MockLogger_Debugw_Call struct {
	*mock.Call
}

// Debugw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Debugw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Debugw_Call {
	return &MockLogger_Debugw_Call{Call: _e.mock.On("Debugw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Debugw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Debugw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debugw_Call) Return() *MockLogger_Debugw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debugw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Debugw_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: args
func (_m *MockLogger) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Error(args ...interface{}) *MockLogger_Error_Call {
	return &MockLogger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Error_Call) Run(run func(args ...interface{})) *MockLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Error_Call) Return() *MockLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Error_Call) RunAndReturn(run func(...interface{})) *MockLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Errorf provides a mock function with given fields: format, values
func (_m *MockLogger) Errorf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type MockLogger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Errorf(format interface{}, values ...interface{}) *MockLogger_Errorf_Call {
	return &MockLogger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Errorf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Errorf_Call) Return() *MockLogger_Errorf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Errorf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Errorf_Call {
	_c.Call.Return(run)
	return _c
}

// Errorw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Errorw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Errorw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorw'
type MockLogger_Errorw_Call struct {
	*mock.Call
}

// Errorw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Errorw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Errorw_Call {
	return &MockLogger_Errorw_Call{Call: _e.mock.On("Errorw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Errorw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Errorw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Errorw_Call) Return() *MockLogger_Errorw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Errorw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Errorw_Call {
	_c.Call.Return(run)
	return _c
}

// Fatal provides a mock function with given fields: args
func (_m *MockLogger) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Fatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatal'
type MockLogger_Fatal_Call struct {
	*mock.Call
}

// Fatal is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Fatal(args ...interface{}) *MockLogger_Fatal_Call {
	return &MockLogger_Fatal_Call{Call: _e.mock.On("Fatal",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Fatal_Call) Run(run func(args ...interface{})) *MockLogger_Fatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Fatal_Call) Return() *MockLogger_Fatal_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatal_Call) RunAndReturn(run func(...interface{})) *MockLogger_Fatal_Call {
	_c.Call.Return(run)
	return _c
}

// Fatalf provides a mock function with given fields: format, values
func (_m *MockLogger) Fatalf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Fatalf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatalf'
type MockLogger_Fatalf_Call struct {
	*mock.Call
}

// Fatalf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Fatalf(format interface{}, values ...interface{}) *MockLogger_Fatalf_Call {
	return &MockLogger_Fatalf_Call{Call: _e.mock.On("Fatalf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Fatalf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Fatalf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Fatalf_Call) Return() *MockLogger_Fatalf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatalf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Fatalf_Call {
	_c.Call.Return(run)
	return _c
}

// Fatalw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Fatalw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatalw'
type MockLogger_Fatalw_Call struct {
	*mock.Call
}

// Fatalw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Fatalw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Fatalw_Call {
	return &MockLogger_Fatalw_Call{Call: _e.mock.On("Fatalw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Fatalw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Fatalw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Fatalw_Call) Return() *MockLogger_Fatalw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatalw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Fatalw_Call {
	_c.Call.Return(run)
	return _c
}

// Helper provides a mock function with given fields: skip
func (_m *MockLogger) Helper(skip int) Logger {
	ret := _m.Called(skip)

	if len(ret) == 0 {
		panic("no return value specified for Helper")
	}

	var r0 Logger
	if rf, ok := ret.Get(0).(func(int) Logger); ok {
		r0 = rf(skip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// MockLogger_Helper_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Helper'
type MockLogger_Helper_Call struct {
	*mock.Call
}

// Helper is a helper method to define mock.On call
//   - skip int
func (_e *MockLogger_Expecter) Helper(skip interface{}) *MockLogger_Helper_Call {
	return &MockLogger_Helper_Call{Call: _e.mock.On("Helper", skip)}
}

func (_c *MockLogger_Helper_Call) Run(run func(skip int)) *MockLogger_Helper_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockLogger_Helper_Call) Return(_a0 Logger) *MockLogger_Helper_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Helper_Call) RunAndReturn(run func(int) Logger) *MockLogger_Helper_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: args
func (_m *MockLogger) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type MockLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Info(args ...interface{}) *MockLogger_Info_Call {
	return &MockLogger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Info_Call) Run(run func(args ...interface{})) *MockLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Info_Call) Return() *MockLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Info_Call) RunAndReturn(run func(...interface{})) *MockLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Infof provides a mock function with given fields: format, values
func (_m *MockLogger) Infof(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type MockLogger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Infof(format interface{}, values ...interface{}) *MockLogger_Infof_Call {
	return &MockLogger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Infof_Call) Run(run func(format string, values ...interface{})) *MockLogger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Infof_Call) Return() *MockLogger_Infof_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Infof_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Infof_Call {
	_c.Call.Return(run)
	return _c
}

// Infow provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Infow(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Infow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infow'
type MockLogger_Infow_Call struct {
	*mock.Call
}

// Infow is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Infow(msg interface{}, keysAndValues ...interface{}) *MockLogger_Infow_Call {
	return &MockLogger_Infow_Call{Call: _e.mock.On("Infow",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Infow_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Infow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Infow_Call) Return() *MockLogger_Infow_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Infow_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Infow_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockLogger) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockLogger_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockLogger_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockLogger_Expecter) Name() *MockLogger_Name_Call {
	return &MockLogger_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockLogger_Name_Call) Run(run func()) *MockLogger_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLogger_Name_Call) Return(_a0 string) *MockLogger_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Name_Call) RunAndReturn(run func() string) *MockLogger_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Named provides a mock function with given fields: name
func (_m *MockLogger) Named(name string) Logger {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Named")
	}

	var r0 Logger
	if rf, ok := ret.Get(0).(func(string) Logger); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// MockLogger_Named_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Named'
type MockLogger_Named_Call struct {
	*mock.Call
}

// Named is a helper method to define mock.On call
//   - name string
func (_e *MockLogger_Expecter) Named(name interface{}) *MockLogger_Named_Call {
	return &MockLogger_Named_Call{Call: _e.mock.On("Named", name)}
}

func (_c *MockLogger_Named_Call) Run(run func(name string)) *MockLogger_Named_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLogger_Named_Call) Return(_a0 Logger) *MockLogger_Named_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Named_Call) RunAndReturn(run func(string) Logger) *MockLogger_Named_Call {
	_c.Call.Return(run)
	return _c
}

// Panic provides a mock function with given fields: args
func (_m *MockLogger) Panic(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Panic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Panic'
type MockLogger_Panic_Call struct {
	*mock.Call
}

// Panic is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Panic(args ...interface{}) *MockLogger_Panic_Call {
	return &MockLogger_Panic_Call{Call: _e.mock.On("Panic",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Panic_Call) Run(run func(args ...interface{})) *MockLogger_Panic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Panic_Call) Return() *MockLogger_Panic_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Panic_Call) RunAndReturn(run func(...interface{})) *MockLogger_Panic_Call {
	_c.Call.Return(run)
	return _c
}

// Panicf provides a mock function with given fields: format, values
func (_m *MockLogger) Panicf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Panicf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Panicf'
type MockLogger_Panicf_Call struct {
	*mock.Call
}

// Panicf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Panicf(format interface{}, values ...interface{}) *MockLogger_Panicf_Call {
	return &MockLogger_Panicf_Call{Call: _e.mock.On("Panicf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Panicf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Panicf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Panicf_Call) Return() *MockLogger_Panicf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Panicf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Panicf_Call {
	_c.Call.Return(run)
	return _c
}

// Panicw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Panicw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Panicw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Panicw'
type MockLogger_Panicw_Call struct {
	*mock.Call
}

// Panicw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Panicw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Panicw_Call {
	return &MockLogger_Panicw_Call{Call: _e.mock.On("Panicw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Panicw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Panicw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Panicw_Call) Return() *MockLogger_Panicw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Panicw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Panicw_Call {
	_c.Call.Return(run)
	return _c
}

// Recover provides a mock function with given fields: panicErr
func (_m *MockLogger) Recover(panicErr interface{}) {
	_m.Called(panicErr)
}

// MockLogger_Recover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recover'
type MockLogger_Recover_Call struct {
	*mock.Call
}

// Recover is a helper method to define mock.On call
//   - panicErr interface{}
func (_e *MockLogger_Expecter) Recover(panicErr interface{}) *MockLogger_Recover_Call {
	return &MockLogger_Recover_Call{Call: _e.mock.On("Recover", panicErr)}
}

func (_c *MockLogger_Recover_Call) Run(run func(panicErr interface{})) *MockLogger_Recover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockLogger_Recover_Call) Return() *MockLogger_Recover_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Recover_Call) RunAndReturn(run func(interface{})) *MockLogger_Recover_Call {
	_c.Call.Return(run)
	return _c
}

// SetLogLevel provides a mock function with given fields: _a0
func (_m *MockLogger) SetLogLevel(_a0 zapcore.Level) {
	_m.Called(_a0)
}

// MockLogger_SetLogLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLogLevel'
type MockLogger_SetLogLevel_Call struct {
	*mock.Call
}

// SetLogLevel is a helper method to define mock.On call
//   - _a0 zapcore.Level
func (_e *MockLogger_Expecter) SetLogLevel(_a0 interface{}) *MockLogger_SetLogLevel_Call {
	return &MockLogger_SetLogLevel_Call{Call: _e.mock.On("SetLogLevel", _a0)}
}

func (_c *MockLogger_SetLogLevel_Call) Run(run func(_a0 zapcore.Level)) *MockLogger_SetLogLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(zapcore.Level))
	})
	return _c
}

func (_c *MockLogger_SetLogLevel_Call) Return() *MockLogger_SetLogLevel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_SetLogLevel_Call) RunAndReturn(run func(zapcore.Level)) *MockLogger_SetLogLevel_Call {
	_c.Call.Return(run)
	return _c
}

// Sync provides a mock function with given fields:
func (_m *MockLogger) Sync() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Sync")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLogger_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type MockLogger_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
func (_e *MockLogger_Expecter) Sync() *MockLogger_Sync_Call {
	return &MockLogger_Sync_Call{Call: _e.mock.On("Sync")}
}

func (_c *MockLogger_Sync_Call) Run(run func()) *MockLogger_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLogger_Sync_Call) Return(_a0 error) *MockLogger_Sync_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Sync_Call) RunAndReturn(run func() error) *MockLogger_Sync_Call {
	_c.Call.Return(run)
	return _c
}

// Trace provides a mock function with given fields: args
func (_m *MockLogger) Trace(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Trace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trace'
type MockLogger_Trace_Call struct {
	*mock.Call
}

// Trace is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Trace(args ...interface{}) *MockLogger_Trace_Call {
	return &MockLogger_Trace_Call{Call: _e.mock.On("Trace",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Trace_Call) Run(run func(args ...interface{})) *MockLogger_Trace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Trace_Call) Return() *MockLogger_Trace_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Trace_Call) RunAndReturn(run func(...interface{})) *MockLogger_Trace_Call {
	_c.Call.Return(run)
	return _c
}

// Tracef provides a mock function with given fields: format, values
func (_m *MockLogger) Tracef(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Tracef_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tracef'
type MockLogger_Tracef_Call struct {
	*mock.Call
}

// Tracef is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Tracef(format interface{}, values ...interface{}) *MockLogger_Tracef_Call {
	return &MockLogger_Tracef_Call{Call: _e.mock.On("Tracef",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Tracef_Call) Run(run func(format string, values ...interface{})) *MockLogger_Tracef_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Tracef_Call) Return() *MockLogger_Tracef_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Tracef_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Tracef_Call {
	_c.Call.Return(run)
	return _c
}

// Tracew provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Tracew(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Tracew_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tracew'
type MockLogger_Tracew_Call struct {
	*mock.Call
}

// Tracew is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Tracew(msg interface{}, keysAndValues ...interface{}) *MockLogger_Tracew_Call {
	return &MockLogger_Tracew_Call{Call: _e.mock.On("Tracew",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Tracew_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Tracew_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Tracew_Call) Return() *MockLogger_Tracew_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Tracew_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Tracew_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: args
func (_m *MockLogger) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type MockLogger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) Warn(args ...interface{}) *MockLogger_Warn_Call {
	return &MockLogger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_Warn_Call) Run(run func(args ...interface{})) *MockLogger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warn_Call) Return() *MockLogger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warn_Call) RunAndReturn(run func(...interface{})) *MockLogger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

// Warnf provides a mock function with given fields: format, values
func (_m *MockLogger) Warnf(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// MockLogger_Warnf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnf'
type MockLogger_Warnf_Call struct {
	*mock.Call
}

// Warnf is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *MockLogger_Expecter) Warnf(format interface{}, values ...interface{}) *MockLogger_Warnf_Call {
	return &MockLogger_Warnf_Call{Call: _e.mock.On("Warnf",
		append([]interface{}{format}, values...)...)}
}

func (_c *MockLogger_Warnf_Call) Run(run func(format string, values ...interface{})) *MockLogger_Warnf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warnf_Call) Return() *MockLogger_Warnf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warnf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Warnf_Call {
	_c.Call.Return(run)
	return _c
}

// Warnw provides a mock function with given fields: msg, keysAndValues
func (_m *MockLogger) Warnw(msg string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLogger_Warnw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnw'
type MockLogger_Warnw_Call struct {
	*mock.Call
}

// Warnw is a helper method to define mock.On call
//   - msg string
//   - keysAndValues ...interface{}
func (_e *MockLogger_Expecter) Warnw(msg interface{}, keysAndValues ...interface{}) *MockLogger_Warnw_Call {
	return &MockLogger_Warnw_Call{Call: _e.mock.On("Warnw",
		append([]interface{}{msg}, keysAndValues...)...)}
}

func (_c *MockLogger_Warnw_Call) Run(run func(msg string, keysAndValues ...interface{})) *MockLogger_Warnw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warnw_Call) Return() *MockLogger_Warnw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warnw_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Warnw_Call {
	_c.Call.Return(run)
	return _c
}

// With provides a mock function with given fields: args
func (_m *MockLogger) With(args ...interface{}) Logger {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for With")
	}

	var r0 Logger
	if rf, ok := ret.Get(0).(func(...interface{}) Logger); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// MockLogger_With_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'With'
type MockLogger_With_Call struct {
	*mock.Call
}

// With is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLogger_Expecter) With(args ...interface{}) *MockLogger_With_Call {
	return &MockLogger_With_Call{Call: _e.mock.On("With",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLogger_With_Call) Run(run func(args ...interface{})) *MockLogger_With_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_With_Call) Return(_a0 Logger) *MockLogger_With_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_With_Call) RunAndReturn(run func(...interface{}) Logger) *MockLogger_With_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
