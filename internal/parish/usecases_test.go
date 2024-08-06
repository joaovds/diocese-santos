package parish

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParishUsecases(t *testing.T) {
	t.Run("should return an parish usecases instance", func(t *testing.T) {
		parishUsecases := NewParishUsecases()
		assert.IsType(t, &ParishUsecases{}, parishUsecases)
	})

	t.Run("GetParishesByCity", func(t *testing.T) {
		parishUsecases := NewParishUsecases()

		t.Run("should return data in success of case", func(t *testing.T) {
			result, err := parishUsecases.GetParishesByCity([]int{1})
			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Len(t, result, 1)
		})
	})
}
