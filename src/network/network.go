package network

import (
	"gonum.org/v1/gonum/mat"
	"fmt"
	"layer"
	"math"
)

type floss func(*mat.Dense, *mat.Dense) (float64)
type flossP func(*mat.Dense, *mat.Dense) (*mat.Dense)

type Save struct {

	Errors []float64
	Epochs int
	Lr 	   []float64
	Lr_t    string
}

type Net struct {

	Layer		[]layer.Layers
	Loss		floss
	LossPrime	flossP
}

func AddAc(Self *Net, ac layer.AC) {

	Self.Layer = append(Self.Layer, &ac)
}

func AddFc(Self *Net, fc layer.FC) {

	Self.Layer = append(Self.Layer, &fc)
}

func Use(Self *Net, loss floss, loss_prime flossP) {

	Self.Loss = loss
    Self.LossPrime = loss_prime
}

func Predict(Self *Net, x *mat.Dense) (*mat.Dense) {

	var res []float64
	var outpout *mat.Dense

	lines, samples := x.Dims()
	data_x := x.RawMatrix().Data

	for i := 0; i < samples; i++ {

		outpout = mat.NewDense(1, lines, transform(data_x, i, lines))
		for z := 0; z < len(Self.Layer); z++ {
			outpout = Self.Layer[z].Forward_propagation(outpout)
		}
		res = append(res, outpout.RawMatrix().Data[0])
	}
	return (mat.NewDense(len(res), 1, res))
}

func Train(x, y *mat.Dense, epochs int, learning_rate float64, Self Net, outpout_s int, S *Save, lr_algo int) (float64) {

	var err float64
	var outpout, error *mat.Dense

	data_x := x.RawMatrix().Data
	data_y := y.RawMatrix().Data
	lines, samples := x.Dims()
	y_lines, _ := y.Dims()
	lr_base := learning_rate
	type_lr := [2]string{ "exponnential", "constant" }
	S.Lr_t = type_lr[lr_algo]

	for i := 0; i < epochs; i++ {
		
		err = 0.0
		for j := 0; j < samples; j++ {

			// forward propagation
			
			outpout = mat.NewDense(outpout_s, lines, transform(data_x, j, lines))
			for z := 0; z < len(Self.Layer); z++ {
				outpout = Self.Layer[z].Forward_propagation(outpout)
			}

			// compute loss (for display purpose only)
	
			real := mat.NewDense(outpout_s, y_lines, transform(data_y, j, y_lines))
			err += Self.Loss(real, outpout)

			// backward propagation
			
			error = Self.LossPrime(real, outpout)
			for k := len(Self.Layer) - 1; k >= 0; k-- {
				error = Self.Layer[k].Backward_propagation(error, learning_rate)
			}
		}
		err /= float64(samples) //save data for graph
		S.Errors = append(S.Errors, err)
		S.Lr = append(S.Lr, learning_rate)
		fmt.Printf("epoch %d / %d error = %f, learning rate : %f\n", i + 1, epochs, err, learning_rate)
		learning_rate = LearningRate(lr_base, float64(i), lr_algo)
	}
	S.Epochs = epochs
	return (err)
}

func LearningRate(lr_init, epoch float64, lr_algo int) (float64) {

	var lrate float64

	if lr_algo == 0 {
		k := 0.09
		lrate = lr_init * math.Exp(-k * epoch)
	} else {
		lrate = lr_init
	}
    return (lrate)
}

func transform(data []float64, j, mul int) ([]float64) {

	var dat []float64

	deb := j * mul
	fin := (j + 1) * mul

	for i := deb; i < fin; i++ {
		dat = append(dat, data[i])
	} 

	return (dat)
}