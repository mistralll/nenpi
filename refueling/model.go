package refueling

import "time"

type Refueling struct {
	Datetime time.Time
	Odo      float64
	Fuel     float64
	Total    int
	// memo      string

	Trip        float64
	UnitPrice   int
	FuelMileage float64
}
