package main

import (
	"conveyorbelt"
	"fmt"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func main() {
	var completed, unusedA, unusedB []float64

	for i := 0; i < 1000; i++ {
		belt := conveyorbelt.NewConveyorBelt(conveyorbelt.FinishedProduct, 3)

		for j := 0; j < 100; j++ {
			belt.Run()
		}

		if i == 0 {
			fmt.Println()
			fmt.Println("Single conveyor belt run (100 steps):")
			fmt.Println(belt)
			fmt.Println()
		}

		output := belt.OutputCount()
		completed = append(completed, float64(output[conveyorbelt.FinishedProduct]))
		unusedA = append(unusedA, float64(output[conveyorbelt.ComponentA]))
		unusedB = append(unusedB, float64(output[conveyorbelt.ComponentB]))
	}

	fmt.Println("Averages over 1000 trials (100 steps per trial):")

	mean, _ := stats.Mean(completed)
	sd, _ := stats.StandardDeviation(completed)
	fmt.Printf("Finished Product: mean %.2f, standard deviation %.2f\n", mean, sd)
	plotHistogram(completed, "Finished Products")

	mean, _ = stats.Mean(unusedA)
	sd, _ = stats.StandardDeviation(unusedA)
	fmt.Printf("Unused Component A: mean %.2f, standard deviation %.2f\n", mean, sd)
	plotHistogram(unusedA, "Unused Component A")

	mean, _ = stats.Mean(unusedB)
	sd, _ = stats.StandardDeviation(unusedB)
	fmt.Printf("Unused Component B: mean %.2f, standard deviation %.2f\n", mean, sd)
	plotHistogram(unusedB, "Unused Component B")
}

func plotHistogram(data []float64, title string) {
	p, _ := plot.New()
	p.Title.Text = title
	v := make(plotter.Values, len(data))
	for i := range v {
		v[i] = data[i]
	}

	h, _ := plotter.NewHist(v, 16)
	h.Normalize(1)
	p.Add(h)

	mean, _ := stats.Mean(data)
	_, _, _, yMax := h.DataRange()
	meanLine, _ := plotter.NewLine(plotter.XYs{plotter.XY{mean, 0}, plotter.XY{mean, yMax}})
	meanLine.Width = vg.Points(2)
	meanLine.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	meanLine.Color = color.RGBA{R: 255, A: 255}
	p.Add(meanLine)

	p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("img/%s Histogram.png", title))
}
