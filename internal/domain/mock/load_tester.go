package mock

import (
	mock "github.com/stretchr/testify/mock"
	domain "gitlab.com/douglasschantz/stress-test/internal/domain"
)

type LoadTester struct {
	mock.Mock
}

type LoadTester_Expecter struct {
	mock *mock.Mock
}

func (_m *LoadTester) EXPECT() *LoadTester_Expecter {
	return &LoadTester_Expecter{mock: &_m.Mock}
}

func (_m *LoadTester) RunLoadTest(config domain.TestConfig) domain.TestResult {
	ret := _m.Called(config)

	if len(ret) == 0 {
		panic("no return value specified for RunLoadTest")
	}

	var r0 domain.TestResult
	if rf, ok := ret.Get(0).(func(domain.TestConfig) domain.TestResult); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Get(0).(domain.TestResult)
	}

	return r0
}

type LoadTester_RunLoadTest_Call struct {
	*mock.Call
}

func (_e *LoadTester_Expecter) RunLoadTest(config interface{}) *LoadTester_RunLoadTest_Call {
	return &LoadTester_RunLoadTest_Call{Call: _e.mock.On("RunLoadTest", config)}
}

func (_c *LoadTester_RunLoadTest_Call) Run(run func(config domain.TestConfig)) *LoadTester_RunLoadTest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.TestConfig))
	})
	return _c
}

func (_c *LoadTester_RunLoadTest_Call) Return(_a0 domain.TestResult) *LoadTester_RunLoadTest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LoadTester_RunLoadTest_Call) RunAndReturn(run func(domain.TestConfig) domain.TestResult) *LoadTester_RunLoadTest_Call {
	_c.Call.Return(run)
	return _c
}

func NewLoadTester(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoadTester {
	mock := &LoadTester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
