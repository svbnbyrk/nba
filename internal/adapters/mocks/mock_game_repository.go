// Code generated by mockery v2.20.0. DO NOT EDIT.

package adapters_mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/svbnbyrk/nba/internal/core/domain"
)

// GameRepositoryInterface is an autogenerated mock type for the GameRepositoryInterface type
type GameRepositoryInterface struct {
	mock.Mock
}

// GetGamesByFilter provides a mock function with given fields: ctx, filter
func (_m *GameRepositoryInterface) GetGamesByFilter(ctx context.Context, filter domain.GameFilter) ([]domain.Game, error) {
	ret := _m.Called(ctx, filter)

	var r0 []domain.Game
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GameFilter) ([]domain.Game, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GameFilter) []domain.Game); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Game)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GameFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertGame provides a mock function with given fields: ctx, game
func (_m *GameRepositoryInterface) UpsertGame(ctx context.Context, game domain.Game) error {
	ret := _m.Called(ctx, game)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Game) error); ok {
		r0 = rf(ctx, game)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewGameRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameRepositoryInterface creates a new instance of GameRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameRepositoryInterface(t mockConstructorTestingTNewGameRepositoryInterface) *GameRepositoryInterface {
	mock := &GameRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
