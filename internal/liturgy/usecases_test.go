package liturgy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLiturgyUsecases(t *testing.T) {
	t.Run("should return an liturgy usecases instance", func(t *testing.T) {
		liturgyUsecases := NewLiturgyUsecases()
		assert.IsType(t, &LiturgyUsecases{}, liturgyUsecases)
	})

	t.Run("GetCurrentLiturgicalInfo()", func(t *testing.T) {
		liturgyUsecases := NewLiturgyUsecases()

		t.Run("should return data in success of case", func(t *testing.T) {
			result, err := liturgyUsecases.GetCurrentLiturgicalInfo()
			assert.Nil(t, err)
			assert.NotNil(t, result)
		})
	})
}
