package main

import (
	"math/rand"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func render(rows []Row) {
	// rand.Seed(int64(0))
	for _, year := range uniqueYears {

		p, err := plot.New()
		if err != nil {
			panic(err)
		}

		p.Title.Text = "Plotutil example"
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"

		err = plotutil.AddLinePoints(p, "First", getPloterXYsOfYear(year, rows).XYs)

		if err != nil {
			panic(err)
		}

		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "build/"+strconv.Itoa(year)+".png"); err != nil {
			panic(err)
		}
	}
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

/*

	for _, year := range uniqueYears {

		// var graphValues = getGraphValuesOfYear(year, rows)

		var graph = chart.Chart{
			Series: []chart.Series{
				chart.ContinuousSeries{
					// XValues: graphValues.xValues,
					// YValues: graphValues.yValues,
					XValues: []float64{1.0, 2.0, 3.0, 4.0},
					YValues: []float64{1.0, 2.0, 3.0, 4.0},
				},
				chart.ContinuousSeries{
					// XValues: graphValues.xValues,
					// YValues: graphValues.yValues,
					XValues: []float64{2.0, 3.0, 5.0, 7.0},
					YValues: []float64{2.0, 7.0, 3.0, 4.0},
				},
			},
		}

		var fname = "build/" + strconv.Itoa(year) + ".png"

		var pngFile, err = os.Create(fname)

		if err != nil {
			panic("Error creating: " + fname)
		}

		err = graph.Render(chart.PNG, pngFile)

		if err != nil {
			panic("Error rendering: " + fname)
		}

		if err = pngFile.Close(); err == nil {
			hel.Pl("Created file:", fname)
		}

	}
*/
