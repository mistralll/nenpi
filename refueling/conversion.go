package refueling

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (r *Refueling) refuelToStr() string {
	d := r.Datetime
	date := d.Format(layout)
	odo := strconv.FormatFloat(r.Odo, 'f', 2, 64)
	fuel := strconv.FormatFloat(r.Fuel, 'f', 2, 64)
	total := strconv.Itoa(r.Total)
	trip := strconv.FormatFloat(r.Trip, 'f', 2, 64)
	unit := strconv.Itoa(r.UnitPrice)
	mileage := strconv.FormatFloat(r.FuelMileage, 'f', 2, 64)

	rt := (date + "," + odo + "," + fuel + "," + total + "," + trip + "," + unit + "," + mileage)
	return rt
}

func StrToRefuel(str string) *Refueling {
	arr := strings.Split(str, ",")
	len := len(arr)

	var date time.Time
	var odo, fuel, trip, mileage float64
	var total, unit int

	if 0 <= len {
		date, _ = time.Parse(layout, arr[0])
	}
	if 1 <= len {
		odo, _ = strconv.ParseFloat(arr[1], 64)
	}
	if 2 <= len {
		fuel, _ = strconv.ParseFloat(arr[2], 64)
	}
	if 3 <= len {
		total, _ = strconv.Atoi(arr[3])
	}
	if 4 <= len {
		trip, _ = strconv.ParseFloat(arr[4], 64)
	}
	if 5 <= len {
		unit, _ = strconv.Atoi(arr[5])
	}
	if 6 <= len {
		mileage, _ = strconv.ParseFloat(arr[6], 64)
	}

	rt := &Refueling{Datetime: date, Odo: odo, Fuel: fuel, Total: total, Trip: trip, UnitPrice: unit, FuelMileage: mileage}
	return rt
}

func HttpReqToRefuel(r *http.Request) *Refueling {
	date, _ := time.Parse("2006-01-02T15:04", r.FormValue("datetime"))
	odo, _ := strconv.ParseFloat(r.FormValue("odo"), 64)
	fuel, _ := strconv.ParseFloat(r.FormValue("fuel"), 64)
	total, _ := strconv.Atoi(r.FormValue("total"))

	return &Refueling{Datetime: date, Odo: odo, Fuel: fuel, Total: total}
}
