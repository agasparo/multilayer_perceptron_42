package graphs

import (
	"github.com/wcharczuk/go-chart"
	"os"
)

func LearmingRateTrace() {

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}
	f, _ := os.Create("view/output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}