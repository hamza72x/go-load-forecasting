package main

import (
	"strconv"
	"time"

	"gonum.org/v1/plot/plotter"
)

// GraphValue containing x and y values
type GraphValue struct {
	Year int
	XYs  plotter.XYs
}

// CSVRow a single row of the csv file
type CSVRow struct {
	ZoneID int    `csv:"zone_id"`
	Year   int    `csv:"year"`
	Month  int    `csv:"month"`
	Day    int    `csv:"day"`
	H1     string `csv:"h1"`
	H2     string `csv:"h2"`
	H3     string `csv:"h3"`
	H4     string `csv:"h4"`
	H5     string `csv:"h5"`
	H6     string `csv:"h6"`
	H7     string `csv:"h7"`
	H8     string `csv:"h8"`
	H9     string `csv:"h9"`
	H10    string `csv:"h10"`
	H11    string `csv:"h11"`
	H12    string `csv:"h12"`
	H13    string `csv:"h13"`
	H14    string `csv:"h14"`
	H15    string `csv:"h15"`
	H16    string `csv:"h16"`
	H17    string `csv:"h17"`
	H18    string `csv:"h18"`
	H19    string `csv:"h19"`
	H20    string `csv:"h20"`
	H21    string `csv:"h21"`
	H22    string `csv:"h22"`
	H23    string `csv:"h23"`
	H24    string `csv:"h24"`
}

// Row for usuability
type Row struct {
	ZoneID int
	Year   int
	Month  int
	Day    int
	Hours  map[string]string
}

// returns as 01
func (r *Row) monthStr() string {
	var str = strconv.Itoa(r.Month)
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}

// returns as 01
func (r *Row) dayStr() string {
	var str = strconv.Itoa(r.Day)
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}

// returns day count from 1 to 365
func (r *Row) getDayCountForYAxis() float64 {
	t1, err := time.Parse(dateLayout1, strconv.Itoa(r.Year)+"-01-01")
	if err != nil {
		panic("Error parsing t1: " + err.Error())
	}
	t2, err := time.Parse(dateLayout1, strconv.Itoa(r.Year)+"-"+r.monthStr()+"-"+r.dayStr())
	if err != nil {
		panic("Error parsing t2: " + err.Error())
	}
	return float64(t2.Sub(t1) / (24 * time.Hour))
}

// DailyAverage average of loads of 24 hours
func (r *Row) getDailyAverage() float64 {

	var total int = 0
	for i := range hourKeys {
		// loadValue => string
		if loadValue, ok := r.Hours[hourKeys[i]]; ok {
			if len(loadValue) > 0 {
				// since loadValue has comma (and/or if any other value)
				// so get only number characters
				filteredLoadValue := getNumbersOnly(loadValue)
				// since load value was parsed as string
				// convert this to integer
				// then if err == nil, add the value to total
				if i, err := strconv.Atoi(filteredLoadValue); err == nil {
					total += i
				}
			}
		}
	}
	return float64(total) / 24.0
}
