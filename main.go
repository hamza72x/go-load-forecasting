package main

import (
	"os"

	"github.com/gocarina/gocsv"
	hel "github.com/thejini3/go-helper"
	"gonum.org/v1/plot/plotter"
)

const csvLoadFile = "csv-files/load-history.csv"
const dateLayout1 = "2006-01-02"

var uniqueYears []int
var rows []Row

// h1, h2, h3, h4....h24
var hourKeys []string

const csvDailyData = "csv-files/daily data 18 19.csv"

func main() {
	// genereate hour keys
	// for i := 1; i <= 24; i++ {
	// 	hourKeys = append(hourKeys, "h"+strconv.Itoa(i))
	// }

	// setRowsFromCSV()
	// render()
	dailyData()

}

func dailyData() {
	type CSVDailyData struct {
		Date string  `json:"Date"`
		Load float64 `json:"Load"`
	}
	// open file
	var csvFile, err = os.OpenFile(csvDailyData, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	var csvRows = []*CSVDailyData{}
	// parse file to CSVRow
	if err := gocsv.UnmarshalFile(csvFile, &csvRows); err != nil { // Load clients from file
		panic(err)
	}

	var count float64 = 0
	var plotData plotter.XYs
	// convert []CSVRow to []Row
	for i := range csvRows {
		row := *csvRows[i]
		plotData = append(plotData, plotter.XY{
			X: count,
			Y: row.Load / 1000,
		})
		hel.Pl(row)
		count++
	}

	build("build/temp.png", []interface{}{
		"daily data", plotData,
	})
}
