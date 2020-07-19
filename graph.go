package main

import (
	"os"

	hel "github.com/thejini3/go-helper"
	"github.com/wcharczuk/go-chart"
)

func renderGraph() {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	pngFile, err := os.Create("../build/output.png")
	c(err)

	err = graph.Render(chart.PNG, pngFile)
	c(err)

	pngFile.Close()
}

func c(err error) {
	if err != nil {
		hel.Pl("Error", err)
	}
}
