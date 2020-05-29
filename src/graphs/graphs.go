package graphs

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"network"
)

func Draw(data []network.Save) {

	for i := 0; i < len(data); i++ {

		LearningRate(Prepare(len(data[i].Lr)), data[i].Lr, data[i].Lr_t)
		Train(Prepare(len(data[i].Errors)), data[i].Errors, data[i].ValLoss, data[i].Lr_t)
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

func Train(x, y, z []float64, name string) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "epochs",
		},
		YAxis: chart.YAxis{
			Name: "loss, val_loss",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "loss",
				XValues: x,
				YValues: y,
			},
			chart.ContinuousSeries{
				Name:    "val loss",
				XValues: x,
				YValues: z,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}
	f, _ := os.Create("data/view/error_" + name + ".png")
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