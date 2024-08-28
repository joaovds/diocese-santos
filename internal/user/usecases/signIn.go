package usecases

import (
	"context"
	"net/http"

	"github.com/joaovds/diocese-santos/internal/user/errs"
	"github.com/joaovds/diocese-santos/pkg/apperr"
)

func (u *UserUsecases) SignIn(ctx context.Context, params *SignInUsecaseParams) (*SignInUsecaseResult, *apperr.AppError) {
	if err := params.validate(); err != nil {
		return nil, err.SetStatus(http.StatusBadRequest)
	}
	println("created")
	return NewSignInUsecaseResult(), nil
}

type SignInUsecaseParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewSignInUsecaseParams() *SignInUsecaseParams {
	return &SignInUsecaseParams{}
}

func (sp *SignInUsecaseParams) validate() *apperr.AppError {
	if len(sp.FirstName) == 0 {
		return errs.NewMissingFieldErr("first_name")
	}
	if len(sp.LastName) == 0 {
		return errs.NewMissingFieldErr("last_name")
	}
	if len(sp.Email) == 0 {
		return errs.NewMissingFieldErr("email")
	}
	if len(sp.Password) == 0 {
		return errs.NewMissingFieldErr("password")
	}

	return nil
}

// ----- ... -----

type SignInUsecaseResult struct{}

func NewSignInUsecaseResult() *SignInUsecaseResult {
	return &SignInUsecaseResult{}
}
