package main

import (
	"log"
	"regexp"

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

func getPloterXYsOfYear(year int, rows []Row) GraphValue {
	// var xValues []float64
	// var yValues []float64
	var XYs []plotter.XY
	for _, row := range rows {
		if year == row.Year {
			XYs = append(XYs, plotter.XY{X: row.getDayCountForYAxis(), Y: row.getDailyAverage()})
		}
	}
	return GraphValue{
		Year: year,
		XYs:  XYs,
	}
}
