// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "belajar-go-echo/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *UserRepository) Create(_a0 models.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *UserRepository) GetAll() (*[]models.User, error) {
	ret := _m.Called()

	var r0 *[]models.User
	if rf, ok := ret.Get(0).(func() *[]models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.User)
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

// Login provides a mock function with given fields: _a0
func (_m *UserRepository) Login(_a0 models.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
