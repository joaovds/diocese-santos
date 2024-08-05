package liturgy

import "time"

type GetCurrentLiturgicalInfoResponse struct {
	Date             time.Time `json:"date"`
	LiturgicalYear   string    `json:"liturgical_year"`
	LiturgicalSeason string    `json:"liturgical_season"`
	LiturgicalWeek   string    `json:"liturgical_week"`
	LiturgicalColor  string    `json:"liturgical_color"`
}

func NewGetCurrentLiturgicalInfoResponse(liturgy *Liturgy) *GetCurrentLiturgicalInfoResponse {
	return &GetCurrentLiturgicalInfoResponse{
		Date:             liturgy.Date,
		LiturgicalYear:   liturgy.LiturgicalYear,
		LiturgicalSeason: liturgy.LiturgicalSeason,
		LiturgicalWeek:   liturgy.LiturgicalWeek,
		LiturgicalColor:  liturgy.LiturgicalColor,
	}
}
