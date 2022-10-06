package refueling

import "time"

type Refueling struct {
	DateTime        time.Time `json:"dateTime"`
	FuelConsumption float64   `json:"fuelConsumption"`
	RefuelAmount    float64   `json:"refuelAmount"`
	TotalCost       int       `json:"totalCost"`
	UnitPrice       int       `json:"unitPrice"`
	Odo             float64   `json:"odo"`
	Trip            float64   `json:"trip"`
	// memo      string
}
