package refueling

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var layout = "2006/01/02 15:04:05"

func (r *Refueling) SaveRefuel(title string) error {
	filename := "data/csv/" + title + ".csv"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	fmt.Println(r.Datetime.Format(layout))

	r.calcRefuel(title)
	content := r.refuelToStr()
	fmt.Fprintln(f, content)

	return nil
}

func (r *Refueling) calcRefuel(title string) error {
	filename := "data/csv/" + title + ".csv"
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastline := ""
	for scanner.Scan() {
		lastline = scanner.Text()
	}

	if lastline != "" {
		prev := StrToRefuel(lastline)
		r.Trip = r.Odo - prev.Odo
		if r.Fuel > 0 {
			r.FuelMileage = r.Trip / r.Fuel
		} else {
			r.FuelMileage = 0
		}
	}

	if r.Fuel > 0 {
		r.UnitPrice = int(float64(r.Total) / r.Fuel)
	} else {
		r.UnitPrice = 0
	}

	return nil
}

func CalcAvgMileage(list []Refueling) float64 {
	var fuelSum float64 = 0
	var tripSum float64 = 0
	for _, row := range list {
		if row.Trip != 0 {
			fuelSum += float64(row.Fuel)
			tripSum += float64(row.Trip)
		}
	}
	var ans float64 = 0
	if fuelSum > 0 {
		ans = tripSum / fuelSum
		ans = math.Round(ans*100) / 100
	}
	return ans
}
