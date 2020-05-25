package main

import (
	"network"
	"gonum.org/v1/gonum/mat"
	"fmt"
	"file"
	"graphical"
)

func main() {

	Doing := graphical.GoTo{}
	Network := network.Net{}

	graphical.ShowMain(&Doing)

	Doing.Create(&Network)
	if Doing.ToDo == 1 {
		// set les poids et le bias sur le reseau
		Predict(Network)
	} else {
		Train(Network, Doing)
	}
}

func Train(Network network.Net, Doing graphical.GoTo) {

	x := []float64{ 0, 0, 0, 1, 1, 0, 1, 1 }
	y := []float64{ 0, 1, 1, 0 }

	x_train := mat.NewDense(2, 4, x)
	y_train := mat.NewDense(1, 4, y)

	epochs := 10000
	learning_rate := 0.1

	err := network.Train(x_train, y_train, epochs, learning_rate, Network)

	// save bias, weigth
	var data [][]float64

	for i := 0; i < len(Network.Layer); i++ {

		tmp, tmp1 := Network.Layer[i].GetData()
		data = append(data, tmp, tmp1)
	}
	// demander stocker le reseau si le poid est moins bon
	file.SaveFile(data, "data/" + Doing.Name + "/res.json", err)

	// repreenter le reseau graphiquement
	// afficher le graph d'apprentissage
}

func Predict(Network network.Net) {

	x := []float64{ 0, 0, 0, 1, 1, 0, 1, 1 }
	x_train := mat.NewDense(2, 4, x)

	pred := network.Predict(&Network, x_train)
	fmt.Println("Prediction : ")
	fmt.Println(pred)
}