package hidden

import (
	"layer"
)

func Init(Self *layer.AC, activation layer.Activ, activation_prime layer.ActivP) {

	Self.Activation = activation
	Self.Activation_prime = activation_prime
}

func Create(length int) ([]layer.AC) {

	var AllInput []layer.AC

	for i := 0; i < length; i++ {
		AllInput = append(AllInput, layer.AC{})
	}
	return (AllInput)
}