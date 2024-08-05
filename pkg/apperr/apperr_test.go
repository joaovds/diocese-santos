package apperr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCode(t *testing.T) {
	t.Run("Test Error Code to string", func(t *testing.T) {
		errCode := ErrorCode("testing")
		assert.IsType(t, "string", errCode.String())
	})
}

func TestAppError(t *testing.T) {
	testErrCode := ErrorCode("testing")
	errMessageList := map[*ErrorCode]string{
		&testErrCode: "testing message",
	}

	t.Run("Test New App Error", func(t *testing.T) {
		appErr := NewAppError(&testErrCode, &errMessageList)
		assert.Equal(t, "testing message", appErr.Message)
		assert.Equal(t, 500, appErr.Status)
		assert.Equal(t, &testErrCode, appErr.ErrorCode)

		t.Run("should return default message if message not exists in list", func(t *testing.T) {
			notExistsErrCodeMessage := ErrorCode("not exists")
			appErr := NewAppError(&notExistsErrCodeMessage, &errMessageList)
			assert.Equal(t, "Unknown error", appErr.Message)
			assert.Equal(t, 500, appErr.Status)
		})
	})

	t.Run("Test Error()", func(t *testing.T) {
		appErr := NewAppError(nil, &errMessageList)
		assert.Equal(t, "Unknown error", appErr.Error())
	})

	t.Run("Test SetStatus()", func(t *testing.T) {
		appErr := NewAppError(&testErrCode, &errMessageList)
		assert.Equal(t, "testing message", appErr.Message)
		assert.Equal(t, 500, appErr.Status)
		appErr.SetStatus(201)
		assert.NotEqual(t, 500, appErr.Status)
		assert.Equal(t, 201, appErr.Status)
	})

	t.Run("Test SetMessage()", func(t *testing.T) {
		appErr := NewAppError(&testErrCode, &errMessageList)
		assert.Equal(t, "testing message", appErr.Message)
		appErr.SetMessage("new message")
		assert.NotEqual(t, "testing message", appErr.Message)
		assert.Equal(t, "new message", appErr.Message)
	})

	t.Run("Test IsNotError()", func(t *testing.T) {
		appErr := NewAppError(&testErrCode, &errMessageList)
		assert.False(t, appErr.IsNoError())
	})

	t.Run("Test IsError()", func(t *testing.T) {
		appErr := NewAppError(&testErrCode, &errMessageList)
		assert.True(t, appErr.IsError())
	})
}
