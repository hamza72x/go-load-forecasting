package main

import (
	"strconv"
)

const csvLoadFile = "csv-files/load-history.csv"
const dateLayout1 = "2006-01-02"

var uniqueYears []int

// h1, h2, h3, h4....h24
var hourKeys []string

func main() {
	// genereate hour keys
	for i := 1; i <= 24; i++ {
		hourKeys = append(hourKeys, "h"+strconv.Itoa(i))
	}

	render(getRowsFromCSV())
}
