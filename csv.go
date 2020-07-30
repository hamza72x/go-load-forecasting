package main

import (
	"os"

	"github.com/gocarina/gocsv"
	hel "github.com/thejini3/go-helper"
)

// parse csv file to []Row array
func setRowsFromCSV() {

	var csvRows []xCSVRow
	parseCsv(csvLoadFile, &csvRows)

	var count = 0
	// convert []CSVRow to []Row
	for _, row := range csvRows {
		// * => converting pointer to local variable
		// since it was parsed as pointer
		// check top

		if !hel.ContainsInt(uniqueYears, row.Year) {
			uniqueYears = append(uniqueYears, row.Year)
		}
		if row.ZoneID != 1 {
			continue
		}
		rows = append(rows, xRow{
			ZoneID: row.ZoneID,
			Year:   row.Year,
			Month:  row.Month,
			Day:    row.Day,
			Hours: map[string]string{
				"h1": row.H1, "h2": row.H2, "h3": row.H3, "h4": row.H4,
				"h5": row.H5, "h6": row.H6, "h7": row.H7, "h8": row.H8,
				"h9": row.H9, "h10": row.H10, "h11": row.H11, "h12": row.H12,
				"h13": row.H13, "h14": row.H14, "h15": row.H15, "h16": row.H16,
				"h17": row.H17, "h18": row.H18, "h19": row.H19, "h20": row.H20,
				"h21": row.H21, "h22": row.H22, "h23": row.H23, "h24": row.H24,
			},
		})

		count++
	}

	hel.Pl("Unique years in csv file:", uniqueYears)
	hel.Pl("Total rows:", count)

}

func parseCsv(filename string, v interface{}) {

	// hel.Pl("Parsing", filename)

	var csvFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	// parse file to CSVRow
	if err := gocsv.UnmarshalFile(csvFile, v); err != nil {
		panic(err)
	}

	// hel.Pl("Parsed", filename)
}
