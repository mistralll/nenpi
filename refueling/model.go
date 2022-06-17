package refueling

import "time"

type Refueling struct {
	Datetime time.Time
	Odo      float64
	Fuel     float64
	Total    int
	// memo      string

	// trip        float64
	// unitPrice   int
	// fuleMileage float64
}
