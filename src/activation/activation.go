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

func MaxFloatInSlice(fls []float64) (m float64) {

    m = fls[len(fls)-1]
    for _, e := range fls {
        if m <= e {
            m = e
        }
    }
    return m
}

func SumExpC(fls []float64) (float64) {
    var s float64 = 0
    c := MaxFloatInSlice(fls)
    for _, e := range fls {
        s += math.Exp(e - c)
    }
    return (s)
}

func Softmax(matr *mat.Dense) (*mat.Dense) {

    var sm []float64

    fls := matr.RawMatrix().Data
    fx, fy := matr.Dims()

    c := MaxFloatInSlice(fls)
    sum_exp_c := SumExpC(fls)
    sm = make([]float64, len(fls))

    for i, v := range fls {
        sm[i] = math.Exp(v - c) / sum_exp_c
    }
    return (mat.NewDense(fx, fy, sm))
}
