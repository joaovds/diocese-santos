package parish

import (
	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/mock"
)

type MockUsecases struct {
	mock.Mock
}

func (m *MockUsecases) GetParishesByCity(citiesIDs []int) ([]*Parish, *apperr.AppError) {
	args := m.Called(citiesIDs)
	return args.Get(0).([]*Parish), args.Get(1).(*apperr.AppError)
}
