package refueling

import (
	"fmt"
	"os"
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
	}
	defer f.Close()
	fmt.Fprintln(f, content)

	return nil
}
