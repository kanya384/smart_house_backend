// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "smart_house_backend/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// HouseParts is an autogenerated mock type for the HouseParts type
type HouseParts struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, housePart
func (_m *HouseParts) Create(ctx context.Context, housePart domain.HousePart) (string, error) {
	ret := _m.Called(ctx, housePart)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, domain.HousePart) string); ok {
		r0 = rf(ctx, housePart)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.HousePart) error); ok {
		r1 = rf(ctx, housePart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *HouseParts) Delete(ctx context.Context, id string) error {
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
func (_m *HouseParts) Get(ctx context.Context, id string) (domain.HousePart, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.HousePart
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.HousePart); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.HousePart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, housePart
func (_m *HouseParts) Update(ctx context.Context, housePart domain.HousePart) error {
	ret := _m.Called(ctx, housePart)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.HousePart) error); ok {
		r0 = rf(ctx, housePart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
