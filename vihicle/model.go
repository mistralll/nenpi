package vihicle

import "github.com/mistralll/nenpi/refueling"

type Vihicle struct {
	Title      string
	AvgMileage int64
	Refuelings []refueling.Refueling
}

type Vihicles struct {
	Vihicles []Vihicle
}
