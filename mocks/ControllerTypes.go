// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "smart_house_backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ControllerTypes is an autogenerated mock type for the ControllerTypes type
type ControllerTypes struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, controllerType
func (_m *ControllerTypes) Create(ctx context.Context, controllerType domain.ControllerType) (string, error) {
	ret := _m.Called(ctx, controllerType)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, domain.ControllerType) string); ok {
		r0 = rf(ctx, controllerType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ControllerType) error); ok {
		r1 = rf(ctx, controllerType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ControllerTypes) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *ControllerTypes) Get(ctx context.Context, id string) (domain.Controller, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Controller
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Controller); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Controller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, controllerType
func (_m *ControllerTypes) Update(ctx context.Context, controllerType domain.ControllerType) error {
	ret := _m.Called(ctx, controllerType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ControllerType) error); ok {
		r0 = rf(ctx, controllerType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
