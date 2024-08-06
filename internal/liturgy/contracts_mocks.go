package liturgy

import (
	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/mock"
)

type MockUsecases struct {
	mock.Mock
}

func (m *MockUsecases) GetCurrentLiturgicalInfo() (*Liturgy, *apperr.AppError) {
	args := m.Called()
	return args.Get(0).(*Liturgy), args.Get(1).(*apperr.AppError)
}
