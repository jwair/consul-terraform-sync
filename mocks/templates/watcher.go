// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	hcat "github.com/hashicorp/hcat"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"
)

// Watcher is an autogenerated mock type for the Watcher type
type Watcher struct {
	mock.Mock
}

// BufferReset provides a mock function with given fields: _a0
func (_m *Watcher) BufferReset(_a0 hcat.Notifier) {
	_m.Called(_a0)
}

// Buffering provides a mock function with given fields: _a0
func (_m *Watcher) Buffering(_a0 hcat.Notifier) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(hcat.Notifier) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Clients provides a mock function with given fields:
func (_m *Watcher) Clients() hcat.Looker {
	ret := _m.Called()

	var r0 hcat.Looker
	if rf, ok := ret.Get(0).(func() hcat.Looker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hcat.Looker)
		}
	}

	return r0
}

// Complete provides a mock function with given fields: _a0
func (_m *Watcher) Complete(_a0 hcat.Notifier) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(hcat.Notifier) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Deregister provides a mock function with given fields: ns
func (_m *Watcher) Deregister(ns ...hcat.Notifier) {
	_va := make([]interface{}, len(ns))
	for _i := range ns {
		_va[_i] = ns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MarkForSweep provides a mock function with given fields: notifier
func (_m *Watcher) MarkForSweep(notifier hcat.IDer) {
	_m.Called(notifier)
}

// Recaller provides a mock function with given fields: _a0
func (_m *Watcher) Recaller(_a0 hcat.Notifier) hcat.Recaller {
	ret := _m.Called(_a0)

	var r0 hcat.Recaller
	if rf, ok := ret.Get(0).(func(hcat.Notifier) hcat.Recaller); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hcat.Recaller)
		}
	}

	return r0
}

// Register provides a mock function with given fields: ns
func (_m *Watcher) Register(ns ...hcat.Notifier) error {
	_va := make([]interface{}, len(ns))
	for _i := range ns {
		_va[_i] = ns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...hcat.Notifier) error); ok {
		r0 = rf(ns...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBufferPeriod provides a mock function with given fields: min, max, tmplIDs
func (_m *Watcher) SetBufferPeriod(min time.Duration, max time.Duration, tmplIDs ...string) {
	_va := make([]interface{}, len(tmplIDs))
	for _i := range tmplIDs {
		_va[_i] = tmplIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, min, max)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Size provides a mock function with given fields:
func (_m *Watcher) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Watcher) Stop() {
	_m.Called()
}

// Sweep provides a mock function with given fields: notifier
func (_m *Watcher) Sweep(notifier hcat.IDer) {
	_m.Called(notifier)
}

// WaitCh provides a mock function with given fields: _a0
func (_m *Watcher) WaitCh(_a0 context.Context) <-chan error {
	ret := _m.Called(_a0)

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context) <-chan error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// Watch provides a mock function with given fields: _a0, _a1
func (_m *Watcher) Watch(_a0 context.Context, _a1 chan string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, chan string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWatcher creates a new instance of Watcher. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewWatcher(t testing.TB) *Watcher {
	mock := &Watcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
