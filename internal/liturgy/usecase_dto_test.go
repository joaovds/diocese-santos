package liturgy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGetCurrentLiturgicalInfoResponse(t *testing.T) {
	liturgy := NewLiturgy("year", "season", "week", "color")
	response := NewGetCurrentLiturgicalInfoResponse(liturgy)
	assert.IsType(t, &GetCurrentLiturgicalInfoResponse{}, response)
	assert.NotNil(t, response)
	assert.Equal(t, liturgy.LiturgicalYear, response.LiturgicalYear)
	assert.Equal(t, liturgy.LiturgicalSeason, response.LiturgicalSeason)
	assert.Equal(t, liturgy.LiturgicalWeek, response.LiturgicalWeek)
	assert.Equal(t, liturgy.LiturgicalColor, response.LiturgicalColor)
}
