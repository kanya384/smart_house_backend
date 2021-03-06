// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "smart_house_backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// Devices is an autogenerated mock type for the Devices type
type Devices struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, deviceType
func (_m *Devices) Create(ctx context.Context, deviceType domain.Device) (string, error) {
	ret := _m.Called(ctx, deviceType)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, domain.Device) string); ok {
		r0 = rf(ctx, deviceType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Device) error); ok {
		r1 = rf(ctx, deviceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Devices) Delete(ctx context.Context, id string) error {
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
func (_m *Devices) Get(ctx context.Context, id string) (domain.Device, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Device
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Device); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, deviceType
func (_m *Devices) Update(ctx context.Context, deviceType domain.Device) error {
	ret := _m.Called(ctx, deviceType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Device) error); ok {
		r0 = rf(ctx, deviceType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
