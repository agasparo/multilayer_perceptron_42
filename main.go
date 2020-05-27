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
	"norm"
	"text/tabwriter"
	"os"
	"maths"
)

func main() {

	Doing := graphical.GoTo{}
	Network := network.Net{}

	graphical.ShowMain(&Doing)

	final := Doing.Create(&Network)
	ToLearn := file.Learn{}
	if file.ReadFile("data/data.csv", &ToLearn) == 0 {
		return
	}
	norm.Normalize(ToLearn.Datas)

	if Doing.ToDo == 1 {

		err, datas := file.GetDatas("data/" + Doing.Name + "/res.json")
		if err == 1 {
			Response.Print("Error with your file")
			return
		}
		create.ChangeDatas(&Network, datas)
		Predict(Network, ToLearn)
	} else {

		Train(Network, Doing, final, ToLearn)
	}
}

func Train(Network network.Net, Doing graphical.GoTo, final int, TL file.Learn) {

	var data [][]float64
	var savefile string
	var x []float64

	for i := 0; i < len(TL.Datas); i++ {

		for e := 0; e < len(TL.Datas[i]); e ++ {
			x = append(x, TL.Datas[i][e])
		}
	}

	x_train := mat.NewDense(len(TL.Datas[0]), len(TL.Datas), x)
	y_train := mat.NewDense(1, len(TL.Response), TL.Response)

	epochs := 1000
	learning_rate := 0.1

	fmt.Println(x_train.Dims())

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
			file.SaveFile(data, savefile, err, Doing.Name)
		}
	} else {
		file.SaveFile(data, savefile, err, Doing.Name)
	}

	// representer le reseau graphiquement
	


	// afficher le graph d'apprentissage
	

}

func Predict(Network network.Net, TL file.Learn) {

	var x []float64

	for i := 0; i < len(TL.Datas); i++ {

		for e := 0; e < len(TL.Datas[i]); e ++ {
			x = append(x, TL.Datas[i][e])
		}
	}
	x_train := mat.NewDense(len(TL.Datas[0]), len(TL.Datas), x)
	real_data := TL.Response

	pred := network.Predict(&Network, x_train)
	x_data := pred.RawMatrix().Data

	fmt.Println("Prediction : ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "index\tNeuronal Network response\treal response\tdiff\t")
	for i := 0; i < len(x_data); i++ {
			fmt.Fprintln(w, toString(float64(i), 0) + "\t" + toString(x_data[i], 0) + "\t" + toString(real_data[i], 0) + "\t" + toString(x_data[i] - real_data[i], 1)  + "\t")
	}
    fmt.Fprintln(w)
    w.Flush()
}

func toString(nb float64, t int) (string) {

	if t == 1 {

		nbc := maths.Abs(nb)
		if nbc <= 0.30 {
			return (fmt.Sprintf("\033[1;32m%f \033[0m", nb))
		} else if nbc > 0.30 && nbc <= 0.60 {
			return (fmt.Sprintf("\033[1;33m%f \033[0m", nb))
		} else {
			return (fmt.Sprintf("\033[1;31m%f \033[0m", nb))
		}
	}
	return (fmt.Sprintf("%f", nb))
}