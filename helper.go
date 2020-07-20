package main

import (
	"log"
	"regexp"
	"strconv"
	"time"

	hel "github.com/thejini3/go-helper"
	"gonum.org/v1/plot/plotter"
)

func getNumbersOnly(str string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "")
}

// func getGraphValuesOfYear(year int, rows []Row) GraphValue {
// 	var xValues []float64
// 	var yValues []float64
// 	for _, row := range rows {
// 		if year == row.Year {
// 			xValues = append(xValues, row.getDailyAverage())
// 			yValues = append(yValues, row.getDayCountForYAxis())
// 		}
// 	}
// 	return GraphValue{
// 		Year:    year,
// 		xValues: xValues,
// 		yValues: yValues,
// 	}
// }

func getPloterXYsOfYear(year int) plotter.XYs {
	// var xValues []float64
	// var yValues []float64
	var XYs []plotter.XY
	var count = 0
	for month := 1; month <= 12; month++ {

		t := date(year, month, 0)

		for day := 1; day <= t.Day(); day++ {
			XYs = append(XYs, plotter.XY{X: float64(count), Y: getAverageLoadOfYMD(year, month, day)})
			count++
		}
	}
	// XYs = append(XYs, plotter.XY{X: 1, Y: 10})
	// XYs = append(XYs, plotter.XY{X: 2, Y: 11})
	// XYs = append(XYs, plotter.XY{X: 3, Y: 13})
	// XYs = append(XYs, plotter.XY{X: 4, Y: 14})
	hel.Pl("Generated plots for year:", year, "- total XYs:", count)
	return XYs
}

func getAverageLoadOfYMD(year, month, day int) float64 {
	var load float64 = 0
	for _, row := range rows {
		if year == row.Year && month == row.Month && day == row.Day {
			load = row.getDailyAverage()
			break
		}
	}
	return load
}

// dORm day or month
func timify(dORm int) string {
	var str = strconv.Itoa(dORm)
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
