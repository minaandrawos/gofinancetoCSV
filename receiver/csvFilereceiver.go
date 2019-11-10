package receiver

import (
	"bufio"
	"encoding/csv"
	"gofinancetoCSV/dataconverter"
	"os"
	"strings"
)

type CSVRcv struct {
	headerIndexMap map[string]int
	*csv.Writer
	file *os.File
}

func NewCSVRcv(filename string, headers []string) (*CSVRcv, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	w := bufio.NewWriter(f)
	writer := csv.NewWriter(w)

	csvRcv := &CSVRcv{
		Writer:         writer,
		file:           f,
		headerIndexMap: make(map[string]int),
	}
	err = csvRcv.Write(headers)
	if err != nil {
		return nil, err
	}
	for i, h := range headers {
		csvRcv.headerIndexMap[h] = i
	}
	return csvRcv, nil
}

func (csvRcv *CSVRcv) Receive(p dataconverter.Params) error {
	header, row, err := p.ToCSV()
	if err != nil {
		return err
	}

	orderedRows := make([]string, len(row))

	for i, h := range header {
		if v, ok := csvRcv.headerIndexMap[strings.ToLower(h)]; ok {
			orderedRows[v] = row[i]
		}
	}
	return csvRcv.Write(orderedRows)
}

func (csvRcv *CSVRcv) Stop() error {
	csvRcv.Flush()
	return csvRcv.file.Close()
}
