// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	storage "github.com/klevtcov/makemenu_go/storage"
	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// PickRandomDish provides a mock function with given fields: ctx, quantity
func (_m *Storage) PickRandomDish(ctx context.Context, quantity int) (*storage.Dishes, error) {
	ret := _m.Called(ctx, quantity)

	var r0 *storage.Dishes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*storage.Dishes, error)); ok {
		return rf(ctx, quantity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *storage.Dishes); ok {
		r0 = rf(ctx, quantity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.Dishes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PickRandomIngridient provides a mock function with given fields: ctx, ingridient, quantity
func (_m *Storage) PickRandomIngridient(ctx context.Context, ingridient string, quantity int) ([]string, error) {
	ret := _m.Called(ctx, ingridient, quantity)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) ([]string, error)); ok {
		return rf(ctx, ingridient, quantity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []string); ok {
		r0 = rf(ctx, ingridient, quantity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, ingridient, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAllIngridients provides a mock function with given fields: ctx, ingridient
func (_m *Storage) ShowAllIngridients(ctx context.Context, ingridient string) ([]string, error) {
	ret := _m.Called(ctx, ingridient)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, ingridient)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, ingridient)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ingridient)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
