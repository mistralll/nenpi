package refueling

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var layout = "2006/01/02 15:04:05"

func (r *Refueling) SaveRefuel(title string) error {
	filename := "data/" + title + ".csv"
	content := r.refuelToStr()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
		// f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	}
	defer f.Close()
	fmt.Fprintln(f, content)

	return nil
}

func (r *Refueling) refuelToStr() string {
	d := r.Datetime
	date := d.Format(layout)
	odo := strconv.FormatFloat(r.Odo, 'f', 2, 64)
	fuel := strconv.FormatFloat(r.Fuel, 'f', 2, 64)
	total := strconv.Itoa(r.Total)
	rt := (date + "," + odo + "," + fuel + "," + total)
	return rt
}

func StrToRefuel(str string) *Refueling {
	arr := strings.Split(str, ",")

	date, _ := time.Parse(layout, arr[0])
	odo, _ := strconv.ParseFloat(arr[1], 64)
	fuel, _ := strconv.ParseFloat(arr[2], 64)
	total, _ := strconv.Atoi(arr[3])

	rt := &Refueling{Datetime: date, Odo: odo, Fuel: fuel, Total: total}
	return rt
}

func HttpReqToRefuel(r *http.Request) *Refueling {
	date, _ := time.Parse("2006-01-02T15:04", r.FormValue("datetime"))
	odo, _ := strconv.ParseFloat(r.FormValue("odo"), 64)
	fuel, _ := strconv.ParseFloat(r.FormValue("fuel"), 64)
	total, _ := strconv.Atoi(r.FormValue("total"))

	return &Refueling{Datetime: date, Odo: odo, Fuel: fuel, Total: total}
}
