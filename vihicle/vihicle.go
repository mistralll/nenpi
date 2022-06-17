package vihicle

import (
	"bufio"
	"os"

	"github.com/mistralll/goSrv/refueling"
)

func LoadVihicle(title string) (*Vihicle, error) {
	filename := title + ".csv"
	fp, err := os.Open(filename)
	if err != nil {
		fp, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var rows []refueling.Refueling
	for scanner.Scan() {
		line := scanner.Text()
		row := refueling.StrToRefuel(line)
		rows = append(rows, *row)
	}

	p := &Vihicle{Title: title, Refuelings: rows}

	return p, nil
}
