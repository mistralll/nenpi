package vihicle

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mistralll/nenpi/refueling"
)

func LoadVihicle(title string) (*Vihicle, error) {
	filename := "data/" + title + ".csv"
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
	fmt.Println("vihicle.SaveIcon が呼ばれました")

	if r.Method != "POST" {
		return fmt.Errorf("method is not POST")
	}

	file, fileHeader, err := r.FormFile("input_icon")
	if err != nil {
		return err
	}

	uploadFileName := fileHeader.Filename

	imagePath := "data/" + uploadFileName
	// ここ後でtitle + 拡張子に変更する

	saveImg, err := os.Create(imagePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(saveImg, file)
	if err != nil {
		return err
	}

	defer saveImg.Close()
	defer file.Close()

	fmt.Println("vihicle.SaveIcon が終了しました")

	return nil
}
