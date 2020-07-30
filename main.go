package main

import (
	"time"

	hel "github.com/thejini3/go-helper"
)

const csvLoadFile = "csv-files/load-history.csv"
const dateLayout1 = "2006-01-02"

var uniqueYears []int
var rows []xRow

// h1, h2, h3, h4....h24
var hourKeys []string

const csvDailyData = "csv-files/daily data 18 19.csv"

var started time.Time

func main() {
	started = time.Now()
	// genereate hour keys
	// for i := 1; i <= 24; i++ {
	// 	hourKeys = append(hourKeys, "h"+strconv.Itoa(i))
	// }

	// setRowsFromCSV()
	// render()
	// dailyData1()
	sldcToDailyData()
	done()
}

func done() {
	hel.Pl("Execution completed within", time.Since(started))
}
