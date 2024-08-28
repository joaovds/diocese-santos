package usecases

import (
	"context"

	"github.com/joaovds/diocese-santos/pkg/apperr"
)

type UserUsecases struct{}

func NewUserUsecases() *UserUsecases {
	return &UserUsecases{}
}

type UserUsecasesContract interface {
	SignIn(ctx context.Context, params *SignInUsecaseParams) (*SignInUsecaseResult, *apperr.AppError)
}
