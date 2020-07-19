package main

import (
	"strconv"

	hel "github.com/thejini3/go-helper"
)

const csvLoadFile = "csv-files/load-history.csv"

// h1, h2, h3, h4....h24
var hourKeys []string

func main() {
	// genereate hour keys
	for i := 1; i <= 24; i++ {
		hourKeys = append(hourKeys, "h"+strconv.Itoa(i))
	}
	// parse csv file (load)
	var rows = getRowsFromCSV()

	// int/key is the year
	// a year will generate one chart
	var graphs []Graph
	/* var graph = chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	} */

	for _, row := range rows {
		//hel.Pl(rows[i].DailyAverage())
		hel.Pl(row.getDailyAverage())
		// graphs
	}
	// csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(csvContent)
}
