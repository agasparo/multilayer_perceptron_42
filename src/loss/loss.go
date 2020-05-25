package loss

import (
	"maths"
	"gonum.org/v1/gonum/mat"
)

func Mse(y_true, y_pred *mat.Dense) (float64) {

	var res, carre mat.Dense

	res.Sub(y_true, y_pred)
	carre.Pow(&res, 2)
	return (maths.Mean(carre.RawMatrix().Data))
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