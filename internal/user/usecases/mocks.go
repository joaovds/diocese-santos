package usecases

import (
	"context"

	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecases struct {
	mock.Mock
}

func (m *MockUserUsecases) SignIn(ctx context.Context, params *SignInUsecaseParams) (*SignInUsecaseResult, *apperr.AppError) {
	args := m.Called(ctx, params)
	if result, ok := args.Get(0).(*SignInUsecaseResult); ok {
		return result, nil
	}
	if err, ok := args.Get(1).(*apperr.AppError); ok {
		return nil, err
	}
	return nil, nil
}
