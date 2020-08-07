package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	hel "github.com/thejini3/go-helper"
	"gonum.org/v1/plot/plotter"
)

func dailyData1() {

	var csvRows []xDailyData

	parseCsv(csvDailyData, &csvRows)

	var count float64 = 0
	var plotData plotter.XYs

	for _, row := range csvRows {
		plotData = append(plotData, plotter.XY{
			X: count,
			Y: row.Load / 1000,
		})
		count++
	}

	build("build/temp.png", []interface{}{
		"daily data", plotData,
	})
}

func sldcToDailyData() {

	folder := "SLDC_Data"
	oDateHourMinLoad := "Date,Hour,Min,Load\n"
	oDateHourLoad := "Date,Hour,Load\n"
	oDateAvgPeak := "Date,Avg,Peak\n"

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".csv" {
			return nil
		}

		var date = strings.ReplaceAll(info.Name(), ".csv", "")
		var rows []xTimeValue
		parseCsv(path, &rows)
		var total float64 = 0
		var count float64 = 0
		var peak float64 = 0
		for _, row := range rows {

			if len(row.Time) == 0 || row.Value <= 1 {
				continue
			}
			count++
			total += row.Value
			if row.Value > peak {
				peak = row.Value
			}
			for hour := 0; hour <= 23; hour++ {
				// row.Time 00:00
				// hourMin [01, 00]
				hourMin := strings.Split(row.Time, ":")
				// 01 or 02 or 03
				hourStr := hourMin[0]
				// minute string : 00 or 30
				minStr := hourMin[1]

				if len(hourMin) != 2 {
					hel.Pl(hourMin)
					panic("Invalid hourMin")
				}

				if timify(hour) == hourStr && minStr == "00" {
					oDateHourLoad += fmt.Sprintf("%s,%d,%0.2f\n", date, hour, row.Value)
				}

				if timify(hour) == hourStr && (minStr == "00" || minStr == "30") {

					var minute int

					if minStr == "00" {
						minute = 0
					} else if minStr == "30" {
						minute = 30
					} else {
						panic("Unknown minute: " + err.Error())
					}

					oDateHourMinLoad += fmt.Sprintf("%s,%d,%d,%0.2f\n", date, hour, minute, row.Value)

				}
			}
		}

		oDateAvgPeak += fmt.Sprintf("%s,%.2f,%.2f\n", date, total/count, peak)

		return err
	})

	if err != nil {
		panic(err)
	}

	if err = hel.StrToFile("build/processed-date-hour-load.csv", oDateHourLoad); err != nil {
		panic(err)
	}

	if err = hel.StrToFile("build/processed-date-hour-minute-load.csv", oDateHourMinLoad); err != nil {
		panic(err)
	}

	if err = hel.StrToFile("build/processed-avg-peak.csv", oDateAvgPeak); err != nil {
		panic(err)
	}
}
