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
	outputTxt := "Date,Hour,Load\n"

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".csv" {
			return nil
		}

		var date = strings.ReplaceAll(info.Name(), ".csv", "")
		var rows []xTimeValue
		parseCsv(path, &rows)

		for _, row := range rows {

			for hour := 0; hour <= 23; hour++ {
				// row.Time 00:00
				// hourMin [01, 00]
				hourMin := strings.Split(row.Time, ":")
				if timify(hour) == hourMin[0] && hourMin[1] == "00" {
					outputTxt += fmt.Sprintf("%s,%d,%s\n", date, hour, row.Value)
				}
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	if err = hel.StrToFile(folder+"/processed.csv", outputTxt); err != nil {
		panic(err)
	}
}
