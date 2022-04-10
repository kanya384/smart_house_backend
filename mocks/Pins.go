// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "smart_house_backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// Pins is an autogenerated mock type for the Pins type
type Pins struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, pin
func (_m *Pins) Create(ctx context.Context, pin domain.Pin) (string, error) {
	ret := _m.Called(ctx, pin)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, domain.Pin) string); ok {
		r0 = rf(ctx, pin)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Pin) error); ok {
		r1 = rf(ctx, pin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Pins) Delete(ctx context.Context, id string) error {
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
func (_m *Pins) Get(ctx context.Context, id string) (domain.Pin, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Pin
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Pin); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Pin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, pin
func (_m *Pins) Update(ctx context.Context, pin domain.Pin) error {
	ret := _m.Called(ctx, pin)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Pin) error); ok {
		r0 = rf(ctx, pin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}