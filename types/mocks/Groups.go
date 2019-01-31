// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import types "gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"

// Groups is an autogenerated mock type for the Groups type
type Groups struct {
	mock.Mock
}

// Create provides a mock function with given fields: name
func (_m *Groups) Create(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: name
func (_m *Groups) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *Groups) List() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDetails provides a mock function with given fields: search
func (_m *Groups) ListDetails(search string) ([]types.Group, error) {
	ret := _m.Called(search)

	var r0 []types.Group
	if rf, ok := ret.Get(0).(func(string) []types.Group); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: search
func (_m *Groups) Search(search string) ([]string, error) {
	ret := _m.Called(search)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubAdminList provides a mock function with given fields: name
func (_m *Groups) SubAdminList(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Users provides a mock function with given fields: name
func (_m *Groups) Users(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
