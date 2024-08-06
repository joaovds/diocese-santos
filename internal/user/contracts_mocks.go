package user

import (
	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/mock"
)

type MockUsecases struct {
	mock.Mock
}

func (m *MockUsecases) GetByID(id int) (*User, *apperr.AppError) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Get(1).(*apperr.AppError)
}
