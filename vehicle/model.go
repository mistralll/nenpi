package vehicle

import "github.com/mistralll/nenpi/refueling"

type Vehicle struct {
	Title              string                `json:"title"`
	AvgFuelConsumption float64               `json:"avgFuelConsumption"`
	RefuelingRows      []refueling.Refueling `json:"refuelingRows"`
}

type Vehicles struct {
	Vehicles []Vehicle `json:"vehicles"`
}
