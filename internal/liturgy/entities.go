package liturgy

import "time"

type Liturgy struct {
	Date             time.Time
	LiturgicalYear   string
	LiturgicalSeason string
	LiturgicalWeek   string
	LiturgicalColor  string
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
