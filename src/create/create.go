package create

import (
	"loss"
	"activation"
	"hidden"
	"input"
	"network"
)

func XOR(Network *network.Net) {

	AllInput := input.Create(2)
	AllActive := hidden.Create(2)

	input.Init(&AllInput[0], 2, 3, 0)
	input.Init(&AllInput[1], 3, 1, 1)
	
	hidden.Init(&AllActive[0], activation.Tanh, activation.Tanh_prime)
	hidden.Init(&AllActive[1], activation.Tanh, activation.Tanh_prime)

	network.AddFc(Network, AllInput[0])
	network.AddAc(Network, AllActive[0])
	network.AddFc(Network, AllInput[1])
	network.AddAc(Network, AllActive[1])

	network.Use(Network, loss.Mse, loss.Mse_prime)
}