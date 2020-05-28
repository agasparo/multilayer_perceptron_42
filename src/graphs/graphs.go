package graphs

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"network"
)

func Draw(data []network.Save) {

	for i := 0; i < len(data); i++ {

		LearningRate(Prepare(len(data[i].Lr)), data[i].Lr, data[i].Lr_t)
	}
}

func LearningRate(x, y []float64, name string) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "epochs",
		},
		YAxis: chart.YAxis{
			Name: "Learning rate",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}
	f, _ := os.Create("data/view/learningRate_" + name + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func Prepare(size int) ([]float64) {

	var table []float64

	for i := 0; i < size; i++ {
		table = append(table, float64(i))
	}
	return (table)
}