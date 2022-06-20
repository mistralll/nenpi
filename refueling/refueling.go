package refueling

import (
	"bufio"
	"fmt"
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
		r.FuelMileage = r.Trip / r.Fuel
	}

	r.UnitPrice = int(float64(r.Total) / r.Fuel)

	return nil
}

func CalcAvgMileage(list []Refueling) int64 {
	var fuelSum int64 = 0
	var tripSum int64 = 0
	for _, row := range list {
		fuelSum += int64(row.Fuel)
		tripSum += int64(row.Trip)
	}
	var ans int64 = 0
	if(fuelSum > 0) {
		ans = tripSum / fuelSum
	}
	return ans
}
