package refueling

import (
	"bufio"
	"fmt"
	"os"
)

var layout = "2006/01/02 15:04:05"

func (r *Refueling) SaveRefuel(title string) error {
	filename := "data/" + title + ".csv"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		f, err = os.Create(filename)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	
	r.calcRefuel(title)
	content := r.refuelToStr()
	fmt.Fprintln(f, content)

	return nil
}

func (r *Refueling) calcRefuel(title string) error {
	filename := "data/" + title + ".csv"
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
