package row

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Row struct {
	Datetime  time.Time
	Odo       float64
	Refueling float64
	Total     int
	// memo      string

	// trip        float64
	// unitPrice   int
	// fuleMileage float64
}

var layout = "2006/01/02 15:04:05"

func (r *Row) SaveRow(title string) error {
	filename := title + ".csv"
	content := r.rowToString()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintln(f, content)

	return nil
}

func (r *Row) rowToString() string {
	d := r.Datetime
	date := d.Format(layout)
	odo := strconv.FormatFloat(r.Odo, 'f', 2, 64)
	fuel := strconv.FormatFloat(r.Refueling, 'f', 2, 64)
	total := strconv.Itoa(r.Total)
	rt := (date + "," + odo + "," + fuel + "," + total)
	return rt
}

func StringToRow(str string) *Row {
	arr := strings.Split(str, ",")

	date, _ := time.Parse(layout, arr[0])
	odo, _ := strconv.ParseFloat(arr[1], 64)
	fuel, _ := strconv.ParseFloat(arr[2], 64)
	total, _ := strconv.Atoi(arr[3])

	fmt.Println(date)
	fmt.Println(odo)
	fmt.Println(fuel)
	fmt.Println(total)

	rt := &Row{Datetime: date, Odo: odo, Refueling: fuel, Total: total}
	return rt
}

func HttprequesToRow(r *http.Request) *Row {
	date, _ := time.Parse("2006-01-02T15:04", r.FormValue("datetime"))
	odo, _ := strconv.ParseFloat(r.FormValue("odo"), 64)
	fuel, _ := strconv.ParseFloat(r.FormValue("refueling"), 64)
	total, _ := strconv.Atoi(r.FormValue("total"))

	return &Row{Datetime: date, Odo: odo, Refueling: fuel, Total: total}
}
