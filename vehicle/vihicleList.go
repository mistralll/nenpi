package vehicle

import (
	"os"
	"path/filepath"
)

func getVehicleListStr() ([]string, error) {
	var files []string

	root := "data/csv/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fName := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
			files = append(files, fName)
		}
		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}

func GetVehicleList() (*Vehicles, error) {
	list, err := getVehicleListStr()
	if err != nil {
		return nil, err
	}

	var vihicles []Vehicle

	for _, v := range list {
		row, err := LoadVehicleInf(v)
		if err != nil {
			return nil, err
		}
		vihicles = append(vihicles, *row)
	}

	rt := &Vehicles{Vehicles: vihicles}

	return rt, nil
}
