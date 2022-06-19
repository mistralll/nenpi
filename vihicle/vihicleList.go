package vihicle

import (
	"os"
	"path/filepath"
)

func getVihicleListStr() ([]string, error) {
	var files []string

	root := "data/csv/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fname := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
			files = append(files, fname)
		}
		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}

func GetVihicleList() (*Vihicles, error) {
	list, err := getVihicleListStr()
	if err != nil {
		return nil, err
	}

	var vihicles []Vihicle

	for _, v := range list {
		row, err := LoadVihicle(v)
		if err != nil {
			return nil, err
		}
		vihicles = append(vihicles, *row)
	}

	rt := &Vihicles{Vihicles: vihicles}

	return rt, nil
}
