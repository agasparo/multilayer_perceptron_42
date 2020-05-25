package activation

import (
	"math"
    "gonum.org/v1/gonum/mat"
)

func Tanh(matr *mat.Dense) (*mat.Dense) {

    x := matr.RawMatrix().Data
    line, col := matr.Dims()

    for i := 0; i < len(x); i++ {

    	x[i] = math.Tanh(x[i])
    }
    return (mat.NewDense(line, col, x))
}

func Tanh_prime(matr *mat.Dense) (*mat.Dense) { //peut etre a changer
    
    var res mat.Dense

	matr1 := Tanh(matr)
    carre := Pow(matr1)

    line, col := carre.Dims()
    tmpOne := CreateMat(line, col, 1.0)
    res.Sub(tmpOne, carre)
    return (&res)
}

func Pow(matr *mat.Dense) (*mat.Dense) {

    x := matr.RawMatrix().Data
    line, col := matr.Dims()

    for i := 0; i < len(x); i++ {

        x[i] = math.Pow(x[i], 2)
    }
    return (mat.NewDense(line, col, x))
}

func CreateMat(line, cols int, nb float64) (*mat.Dense) {

    var data []float64

    for i := 0; i < (line * cols); i++ {

        data = append(data, nb)
    }
    return (mat.NewDense(line, cols, data))
}
