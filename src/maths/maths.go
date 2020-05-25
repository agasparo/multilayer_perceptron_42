package maths

import (

)

func Count(data []float64) (float64) {

	c := 0

	for i := 0; i < len(data); i++ {
		c++
	}
	return (float64(c))
}

func Mean(data []float64) (float64) {

	var res float64
	c := Count(data)

	for i := 0; i < len(data); i++ {
		res += data[i]
	}
	return (res / c)
}