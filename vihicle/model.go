package vihicle

import "github.com/mistralll/goSrv/refueling"

type Vihicle struct {
	Title      string
	Refuelings []refueling.Refueling
}
