package vehicle

import "github.com/mistralll/nenpi/refueling"

type Vehicle struct {
	Title         string
	AvgMileage    float64
	RefuelingRows []refueling.Refueling
}

type Vehicles struct {
	Vehicles []Vehicle
}
