package http

import (
	"encoding/csv"
	"io"
)

type Field struct {
	Name      string `csv:"Field Name"`
	Template  string
	Status    string
	Reference string
	Comment   string `csv:"Comments"`
}

// TODO return an iterator over the CSV rows
func FieldsFromCsv(r io.Reader) ([]string, error) {
	csvReader := csv.NewReader(r)
	row, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	return row, nil
}
