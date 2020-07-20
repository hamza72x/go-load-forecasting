package main

import (
	hel "github.com/thejini3/go-helper"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func render() {
	// var arr []plotter.XYs
	var arr []interface{}

	for _, year := range uniqueYears {
		// get xys value for `year`
		var xys = getPloterXYsOfYear(year)

		arr = append(arr, timify(year), xys)

		// biuld each year png
		build("build/"+timify(year)+".png", []interface{}{
			timify(year), xys,
		})
	}

	build("build/2004vs2005.png", []interface{}{
		timify(2004), getPloterXYsOfYear(2004),
		timify(2005), getPloterXYsOfYear(2005),
	})

	build("build/2005vs2006.png", []interface{}{
		timify(2005), getPloterXYsOfYear(2005),
		timify(2006), getPloterXYsOfYear(2006),
	})

	// build vs-vs png
	build("build/all-year-vs-year.png", arr)

}

func build(fname string, v []interface{}) {
	p, err := plot.New()

	if err != nil {
		panic(err)
	}

	p.Title.Text = "Load Forecasting in Time Series"
	p.X.Label.Text = "Day"
	p.Y.Label.Text = "Load (KW)"

	err = plotutil.AddLinePoints(p, v...)

	if err := p.Save(10*vg.Inch, 10*vg.Inch, fname); err != nil {
		panic(err)
	} else {
		hel.Pl("Generated:", fname)
	}
}
