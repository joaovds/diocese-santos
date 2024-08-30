package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMissingFieldErr(t *testing.T) {
	t.Run("should return missing field error", func(t *testing.T) {
		field := "testing"
		expectedMessage := "Missing field: testing"
		result := NewMissingFieldErr(field)
		assert.Equal(t, expectedMessage, result.Error())
		assert.Equal(t, MISSING_FIELD.String(), result.ErrorCode.String())
	})
}
