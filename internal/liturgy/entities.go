package liturgy

import "time"

type Liturgy struct {
	Date             time.Time `json:"date"`
	LiturgicalYear   string    `json:"liturgical_year"`
	LiturgicalSeason string    `json:"liturgical_season"`
	LiturgicalWeek   string    `json:"liturgical_week"`
	LiturgicalColor  string    `json:"liturgical_color"`
}

func NewLiturgy(liturgicalYear, liturgicalSeason, liturgicalWeek, liturgicalColor string) *Liturgy {
	return &Liturgy{
		Date:             time.Now(),
		LiturgicalYear:   liturgicalYear,
		LiturgicalSeason: liturgicalSeason,
		LiturgicalWeek:   liturgicalWeek,
		LiturgicalColor:  liturgicalColor,
	}
}
