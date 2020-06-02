package loss

import (
	"maths"
	"gonum.org/v1/gonum/mat"
)

func Mse(y_true, y_pred *mat.Dense) (float64) {

	var res mat.Dense

	res.Sub(y_true, y_pred)
	return (maths.Mean(calcCarre(res.RawMatrix().Data)))
}

func calcCarre(data []float64) ([]float64) {

	var tab []float64

	for i := 0; i < len(data); i++{

		tab = append(tab, data[i] * data[i])
	}
	return (tab)
}

func Mse_prime(y_true, y_pred *mat.Dense) (*mat.Dense) {
    
    var res, two, divi mat.Dense

    res.Sub(y_pred, y_true)
    two.Scale(2, &res)
   
    line, col := two.Dims() 
    tmpSize := CreateMat(line, col, float64(len(y_true.RawMatrix().Data)))
    divi.DivElem(&two, tmpSize)
    return (&divi)
}

func CreateMat(line, cols int, nb float64) (*mat.Dense) {

	var data []float64

	for i := 0; i < (line * cols); i++ {

		data = append(data, nb)
	}
	return (mat.NewDense(line, cols, data))
}