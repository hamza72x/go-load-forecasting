package main

import (
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

		if !hel.IntContains(uniqueYears, row.Year) {
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

// last lines empty new line isn't checked
// so handle that
/*
if len(row.Time) == 0 || len(row.Value) == 0 {
	continue
}
*/
func parseCsv(filename string, v interface{}) {
	str, err := hel.FileStr(filename)
	if err != nil {
		panic("error hel.FileStr: " + filename)
	}
	err = gocsv.UnmarshalString(str, v)
	if err != nil {
		panic("Error UnmarshalString parseCsv" + err.Error())
	}
}
