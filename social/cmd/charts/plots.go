package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Latency"
	p.X.Label.Text = "connection"
	p.Y.Label.Text = "ms"

	err = plotutil.AddLinePoints(p,
		"missing-index", noindexLatency(),
		"present-index", indexLatency())
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "../../assets/img/latency.png"); err != nil {
		panic(err)
	}


	p, err = plot.New()
	if err != nil {
		panic(err)
	}


	p.Title.Text = "Throughput"
	p.X.Label.Text = "connection"
	p.Y.Label.Text = "req/s"

	err = plotutil.AddLinePoints(p,
		"missing-index", noindexReq(),
		"present-index", indexReq())
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "../../assets/img/throughput.png"); err != nil {
		panic(err)
	}


}

func noindexLatency() plotter.XYs {
	//pts := make(plotter.XYs, 4)
	//pts[0].X = 1
	//pts[0].Y =115.4
	pts := []plotter.XY{
		plotter.XY{X: 1, Y: 115.4},
		plotter.XY{X: 10, Y:  513.09},
		plotter.XY{X: 100, Y: 1020},
		plotter.XY{X: 1000, Y: 626.02},
	}
	return pts
}
func indexLatency() plotter.XYs {
	//pts := make(plotter.XYs, 4)
	//pts[0].X = 1
	//pts[0].Y =115.4
	pts := []plotter.XY{
		plotter.XY{X: 1, Y: 20.82},
		plotter.XY{X: 10, Y:   75.40},
		plotter.XY{X: 100, Y: 710.20},
		plotter.XY{X: 1000, Y: 605.95},
	}
	return pts
}
func noindexReq() plotter.XYs {
	//pts := make(plotter.XYs, 4)
	//pts[0].X = 1
	//pts[0].Y =115.4
	pts := []plotter.XY{
		plotter.XY{X: 1, Y: 8.70},
		plotter.XY{X: 10, Y: 19.37},
		plotter.XY{X: 100, Y:  18.25},
		plotter.XY{X: 1000, Y: 143.62},
	}
	return pts
}
func indexReq() plotter.XYs {
	//pts := make(plotter.XYs, 4)
	//pts[0].X = 1
	//pts[0].Y =115.4
	pts := []plotter.XY{
		plotter.XY{X: 1, Y: 54.13},
		plotter.XY{X: 10, Y:   134.06},
		plotter.XY{X: 100, Y:139.17},
		plotter.XY{X: 1000, Y:  544.76},
	}
	return pts
}
