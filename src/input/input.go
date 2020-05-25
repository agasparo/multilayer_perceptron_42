package input

import (
	"math/rand"
	"time"
	"layer"
	"gonum.org/v1/gonum/mat"
)

// input_size = number of input neurons
// output_size = number of output neurons

func Init(Self *layer.FC, input_size, output_size, t int) {

	var tmp, local *mat.Dense

	tmp = mat.NewDense(input_size, output_size, random(input_size, output_size))
	local = mat.NewDense(input_size, output_size, def(input_size, output_size, 0.5))
	Self.Weights.Sub(tmp, local)

	tmp = mat.NewDense(1, output_size, random(1, output_size))
	local = mat.NewDense(1, output_size, def(1, output_size, 0.5))
	Self.Bias.Sub(tmp, local)
}

func random(input_size, output_size int) (data []float64) {

	for i := 0; i < (input_size * output_size); i++ {

		rand.Seed(time.Now().UnixNano())
		data = append(data, rand.Float64())
	}
	return (data)
}

func def(input_size, output_size int, nb float64) (data []float64) {

	for i := 0; i < (input_size * output_size); i++ {

		data = append(data, nb)
	}
	return (data)
}

func Create(length int) ([]layer.FC) {

	var AllInput []layer.FC

	for i := 0; i < length; i++ {
		AllInput = append(AllInput, layer.FC{})
	}
	return (AllInput)
}