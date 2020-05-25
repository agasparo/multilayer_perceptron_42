package layer

import (
	"gonum.org/v1/gonum/mat"
)

type Activ func(*mat.Dense) (*mat.Dense)
type ActivP func(*mat.Dense) (*mat.Dense)

type Layers interface {

	Forward_propagation(*mat.Dense) (*mat.Dense)
	Backward_propagation(*mat.Dense, float64) (*mat.Dense)
	GetData() ([]float64, []float64)
}

type FC struct {

	Input 	*mat.Dense
	Output 	mat.Dense
	Weights mat.Dense
	Bias	mat.Dense
}

type AC struct {

	Activation 			Activ
	Activation_prime 	ActivP
	Input 				*mat.Dense
	Output 				*mat.Dense
}

// return Weigths and Bias for each c

func (Self *FC) GetData() ([]float64, []float64) {

	return Self.Weights.RawMatrix().Data, Self.Bias.RawMatrix().Data
}

func (Self *AC) GetData() ([]float64, []float64) {

	var a, b []float64

	return a, b
}

// returns output for a given input

func (Self *FC) Forward_propagation(input_data *mat.Dense) (*mat.Dense) {

	var res mat.Dense

	Self.Input = input_data
	res.Mul(Self.Input, &Self.Weights)
	Self.Output.Add(&res, &Self.Bias)
	return (&Self.Output)
}

func (Self *AC) Forward_propagation(input_data *mat.Dense) (*mat.Dense) {

	Self.Input = input_data
	Self.Output = Self.Activation(Self.Input)
	return (Self.Output)
}

// computes dE/dW, dE/dB for a given output_error=dE/dY. Returns input_error=dE/dX.

func (Self *FC) Backward_propagation(output_error *mat.Dense, learning_rate float64) (*mat.Dense) {

	var input_error, weights_error, res, res1 mat.Dense

	input_error.Mul(output_error, Self.Weights.T())
    weights_error.Mul(Self.Input.T(), output_error)

    // update parameters
    res.Scale(learning_rate, &weights_error)
    Self.Weights.Sub(&Self.Weights, &res)

    res1.Scale(learning_rate, output_error)
    Self.Bias.Sub(&Self.Bias, &res1)

    return (&input_error)
}

func (Self *AC) Backward_propagation(output_error *mat.Dense, learning_rate float64) (*mat.Dense) {

	var res mat.Dense

	res.Mul(Self.Activation_prime(Self.Input).T(), output_error)
	x, y := res.Dims()
	new := DiagToMat(res.RawMatrix().Data, x, y)
	return (new)
}

func DiagToMat(data []float64, line, col int) (*mat.Dense) {

	var tab []float64

	for i := 0; i < col; i++ {

		for a := 0; a < line; a++ {
			if a == i {
				tab = append(tab, data[a + (i * line)])
			}
		}
	}
	return (mat.NewDense(1, line, tab))
}

func CreateMat(line, cols int, nb float64) (*mat.Dense) {

	var data []float64

	for i := 0; i < (line * cols); i++ {

		data = append(data, nb)
	}
	return (mat.NewDense(line, cols, data))
}