package norm

import (
	"maths"
)

func Normalize(Data map[int][]float64) {

	for i := 0; i < len(Data); i++ {

		minK := maths.Min(Data[i])
		maxK := maths.Max(Data[i])

		for j := 0; j < len(Data[i]); j++ {
			Data[i][j] = (Data[i][j] - minK) / (maxK - minK)
		}
	}
}