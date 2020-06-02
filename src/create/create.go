package create

import (
	"loss"
	"activation"
	"hidden"
	"input"
	"network"
	"file"
)

func ChangeDatas(Network *network.Net, data file.Base) {

	for i := 0; i < len(data.Layers); i++ {
		Network.Layer[i].ModifiData(data.Layers[i].Weigths, data.Layers[i].Bias)
	}
}

func XOR(Network *network.Net) (int) {

	AllInput := input.Create(2)
	AllActive := hidden.Create(2)

	input.Init(&AllInput[0], 30, 60)
	input.Init(&AllInput[1], 60, 1)
	
	hidden.Init(&AllActive[0], activation.Tanh, activation.Tanh_prime)
	hidden.Init(&AllActive[1], activation.Tanh, activation.Tanh_prime)

	network.AddFc(Network, AllInput[0])
	network.AddAc(Network, AllActive[0])
	network.AddFc(Network, AllInput[1])
	network.AddAc(Network, AllActive[1])

	network.Use(Network, loss.Mse, loss.Mse_prime)

	return 1
}

func CUSTOM(Network *network.Net) (int) {

	AllInput := input.Create(4)
	AllActive := hidden.Create(4)

	input.Init(&AllInput[0], 30, 50)
	input.Init(&AllInput[1], 50, 60)
	input.Init(&AllInput[2], 60, 70)
	input.Init(&AllInput[3], 70, 1)

	hidden.Init(&AllActive[0], activation.Tanh, activation.Tanh_prime)
	hidden.Init(&AllActive[1], activation.Tanh, activation.Tanh_prime)
	hidden.Init(&AllActive[2], activation.Tanh, activation.Tanh_prime)
	hidden.Init(&AllActive[3], activation.Tanh, activation.Softmax)

	network.AddFc(Network, AllInput[0])
	network.AddAc(Network, AllActive[0])
	network.AddFc(Network, AllInput[1])
	network.AddAc(Network, AllActive[1])
	network.AddFc(Network, AllInput[2])
	network.AddAc(Network, AllActive[2])
	network.AddFc(Network, AllInput[3])
	network.AddAc(Network, AllActive[3])

	network.Use(Network, loss.Mse, loss.Mse_prime)

	return 1
}