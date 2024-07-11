package file

import (
	"os"
)

type CSV struct {
	File     *os.File
	LineRead uint8
}

func LoadCSV(p string) (*CSV, error) {
	var pathErr error

	if p == "" {
		p, pathErr = os.Getwd()
		if pathErr != nil {
			return nil, pathErr
		}
	}

	f, err := os.Open(p + "/questions.cvs")
	if err != nil {
		return nil, err
	}
	csv := CSV{
		File: f,
	}
	return &csv, nil
}
