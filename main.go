package main

import (
	"network"
	"gonum.org/v1/gonum/mat"
	"fmt"
	"file"
	"graphical"
	"in"
	"Response"
	"create"
)

func main() {

	Doing := graphical.GoTo{}
	Network := network.Net{}

	graphical.ShowMain(&Doing)

	final := Doing.Create(&Network)
	if Doing.ToDo == 1 {

		err, datas := file.GetDatas("data/" + Doing.Name + "/res.json")
		if err == 1 {
			Response.Print("Error with your file")
			return
		}
		create.ChangeDatas(&Network, datas)
		Predict(Network)
	} else {

		// lire le fichier de data
		ToLearn := file.Learn{}
		if file.ReadFile("data/data.csv", &ToLearn) == 0 {
			return
		}
		fmt.Println(final)
		// trier les datas
		//Train(Network, Doing, final)
	}
}

func Train(Network network.Net, Doing graphical.GoTo, final int) {

	var data [][]float64
	var savefile string

	x := []float64{ 0, 0, 0, 1, 1, 0, 1, 1 }
	y := []float64{ 0, 1, 1, 0 }

	x_train := mat.NewDense(2, 4, x)
	y_train := mat.NewDense(1, 4, y)

	epochs := 1000
	learning_rate := 0.1

	err := network.Train(x_train, y_train, epochs, learning_rate, Network, final)
	for i := 0; i < len(Network.Layer); i++ {

		tmp, tmp1 := Network.Layer[i].GetData()
		data = append(data, tmp, tmp1)
	}
	
	savefile = "data/" + Doing.Name + "/res.json" 
	if file.CompErr(err, savefile) == -1 {
		_, res := in.ReadSTDIN("Your error is more than your previous save file, would you like to save it ? [Y/N]", 1)
		for res != "Y" && res != "N" {
			Response.Print("Response must be Y or N")
			_, res = in.ReadSTDIN("Your error is more than your previous save file, would you like to save it ? [Y/N]", 1)
		}
		if res == "Y" {
			file.SaveFile(data, savefile, err)
		}
	} else {
		file.SaveFile(data, savefile, err)
	}

	// representer le reseau graphiquement
	


	// afficher le graph d'apprentissage
	

}

func Predict(Network network.Net) {

	x := []float64{ 0, 0, 0, 1, 1, 0, 1, 1 }
	x_train := mat.NewDense(2, 4, x)

	pred := network.Predict(&Network, x_train)
	fmt.Println("Prediction : ")
	fmt.Println(pred)
}