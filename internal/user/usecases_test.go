package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserUsecases(t *testing.T) {
	t.Run("should return an user usecases instance", func(t *testing.T) {
		userUsecases := NewUserUsecases()
		assert.IsType(t, &UserUsecases{}, userUsecases)
	})

	t.Run("GetByID", func(t *testing.T) {
		userUsecases := NewUserUsecases()

		t.Run("should return data in success of case", func(t *testing.T) {
			result, err := userUsecases.GetByID(1)
			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, 1, result.ID)
			assert.Equal(t, "Carol", result.FirstName)
			assert.Equal(t, "thebestpope", result.Password)
		})
	})
}
