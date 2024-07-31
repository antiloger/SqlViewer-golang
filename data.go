package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ImportData() (header []Col, body [][]string, err error) {
	file, err := os.Open("simple.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rec, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	h := rec[0]
	header = make([]Col, len(h))
	for i, v := range h {
		header[i] = Col{
			Name: v,
		}
	}

	body = rec[1:]

	return header, body, nil
}
