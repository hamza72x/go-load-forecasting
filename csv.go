package main

import (
	"os"

	"github.com/gocarina/gocsv"
	hel "github.com/thejini3/go-helper"
)

// parse csv file to []Row array
func setRowsFromCSV() {
	// open file
	var csvFile, err = os.OpenFile(csvLoadFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	var csvRows = []*CSVRow{}
	// parse file to CSVRow
	if err := gocsv.UnmarshalFile(csvFile, &csvRows); err != nil { // Load clients from file
		panic(err)
	}

	var count = 0
	// convert []CSVRow to []Row
	for i := range csvRows {
		// * => converting pointer to local variable
		// since it was parsed as pointer
		// check top
		var csvRow = *csvRows[i]

		if !hel.ContainsInt(uniqueYears, csvRow.Year) {
			uniqueYears = append(uniqueYears, csvRow.Year)
		}
		if csvRow.ZoneID != 1 {
			continue
		}
		rows = append(rows, Row{
			ZoneID: csvRow.ZoneID,
			Year:   csvRow.Year,
			Month:  csvRow.Month,
			Day:    csvRow.Day,
			Hours: map[string]string{
				"h1": csvRow.H1, "h2": csvRow.H2, "h3": csvRow.H3, "h4": csvRow.H4,
				"h5": csvRow.H5, "h6": csvRow.H6, "h7": csvRow.H7, "h8": csvRow.H8,
				"h9": csvRow.H9, "h10": csvRow.H10, "h11": csvRow.H11, "h12": csvRow.H12,
				"h13": csvRow.H13, "h14": csvRow.H14, "h15": csvRow.H15, "h16": csvRow.H16,
				"h17": csvRow.H17, "h18": csvRow.H18, "h19": csvRow.H19, "h20": csvRow.H20,
				"h21": csvRow.H21, "h22": csvRow.H22, "h23": csvRow.H23, "h24": csvRow.H24,
			},
		})

		count++
	}

	hel.Pl("Unique years in csv file:", uniqueYears)
	hel.Pl("Total rows:", count)

}
