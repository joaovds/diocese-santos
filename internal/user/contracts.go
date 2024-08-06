package user

import "github.com/joaovds/diocese-santos/pkg/apperr"

type Usecases interface {
	GetByID(id int) (*User, *apperr.AppError)
}
