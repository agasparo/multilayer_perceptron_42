package maths

import (

)

func Max(data []float64) (float64) {

	max := data[0]

	for i := 0; i < len(data); i++ {
		
		if data[i] > max {
			max = data[i]
		}
	}
	return (max)
}

func Min(data []float64) (float64) {

	min := data[0]

	for i := 0; i < len(data); i++ {

		if data[i] < min {
			min = data[i]
		}
	}
	return (min)
}

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

func Abs(nb float64) (float64) {

	if nb < 0 {
		return (nb * -1)
	}
	return (nb)
}