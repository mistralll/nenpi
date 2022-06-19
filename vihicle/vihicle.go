package vihicle

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mistralll/nenpi/refueling"
)

func LoadVihicle(title string) (*Vihicle, error) {
	filename := "data/csv/" + title + ".csv"
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

func SaveIcon(r *http.Request, title string) error {
	if r.Method != "POST" {
		return fmt.Errorf("method is not POST")
	}

	inFile, inFileHeader, err := r.FormFile("input_icon")
	if err != nil {
		return err
	}

	filetype := filepath.Ext(inFileHeader.Filename)

	filePath := "data/img/" + title + filetype

	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(newFile, inFile)
	if err != nil {
		return err
	}

	defer newFile.Close()
	defer inFile.Close()

	return nil
}
