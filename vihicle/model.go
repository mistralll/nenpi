package vihicle

import "github.com/mistralll/nenpi/refueling"

type Vihicle struct {
	Title      string
	Refuelings []refueling.Refueling
}

type Vihicles struct {
	Vihicles []Vihicle
}
