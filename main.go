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
	"graphs"
)

func main() {

	Doing := graphical.GoTo{}
	Network := network.Net{}

	graphical.ShowMain(&Doing)

	if Doing.ToDo == 2 {
		errs, datas := file.ReadGraph("data/graph/" + Doing.Name + "_")
		if errs == 1 {
			return
		}
		graphs.Draw(datas)
		
		return
	}

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

		SaveData := network.Save{}
		Train(Network, Doing, final, ToLearn, SaveData)
	}
}

func Train(Network network.Net, Doing graphical.GoTo, final int, TL file.Learn, SaveData network.Save) {

	var data [][]float64
	var savefile string
	var x []float64

	for i := 0; i < len(TL.Datas); i++ {

		for e := 0; e < len(TL.Datas[i]); e++ {
			x = append(x, TL.Datas[i][e])
		}
	}

	x_train := mat.NewDense(len(TL.Datas[0]), len(TL.Datas), x)
	y_train := mat.NewDense(1, len(TL.Response), TL.Response)

	epochs := 250
	learning_rate := 0.1

	err := network.Train(x_train, y_train, epochs, learning_rate, Network, final, &SaveData, Doing.Decrease)
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
			file.SaveGraph(SaveData, "data/graph/" + Doing.Name + "_" + SaveData.Lr_t + ".json")
		}
	} else {
		file.SaveFile(data, savefile, err, Doing.Name)
		file.SaveGraph(SaveData, "data/graph/" + Doing.Name + "_" + SaveData.Lr_t + ".json")
	}
}

func Predict(Network network.Net, TL file.Learn) {

	var x []float64

	for i := 0; i < len(TL.Datas); i++ {

		for e := 0; e < len(TL.Datas[i]); e++ {
			x = append(x, TL.Datas[i][e])
		}
	}
	x_train := mat.NewDense(len(TL.Datas[0]), len(TL.Datas), x)
	real_data := TL.Response

	pred := network.Predict(&Network, x_train)
	x_data := pred.RawMatrix().Data

	fmt.Println("Prediction : ")

	var percent[]int
	percent = append(percent, 0, 0, 0)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "index\tNeuronal Network response\treal response\tM / B real\tM / B NN\tcompare\tdiff")
	for i := 0; i < len(x_data); i++ {
			fmt.Fprintln(w, toString(float64(i), 0, percent) + "\t" + toString(x_data[i], 0, percent) + "\t" + toString(real_data[i], 0, percent) + "\t" + IS(x_data[i]) + "\t" + IS(real_data[i]) + Cmp(x_data[i], real_data[i]) + "\t" + toString(x_data[i] - real_data[i], 1, percent)  + "\t")
	}
    fmt.Fprintln(w)
    w.Flush()

    Response.PrintVerboseStep("Accuracy of neuronal network :")
    tab := []string{ "Green", "Orange", "Red" }
    w = new(tabwriter.Writer)
    w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Color\tNumber\tNumber of data\tpercent\t")
	for i := 0; i < len(percent); i++ {
			fmt.Fprintln(w, tab[i] + "\t" + fmt.Sprintf("%d", percent[i]) + "\t" + fmt.Sprintf("%d", len(x_data)) + "\t" + fmt.Sprintf("%f", float64(percent[i]) / float64(len(x_data)) * 100) + " %\t")
	}
    fmt.Fprintln(w)
    w.Flush()
}

func Cmp(nb, nb1 float64) (string) {

	if IS(nb) == IS(nb1) {
		return("\t\033[1;32m[OK]\033[0m")
	}
	return ("\t\033[1;31m[KO]\033[0m")
}

func IS(nb float64) (string) {

	if nb >= 0.5 {
		return ("M")
	}
	return ("B")
}

func toString(nb float64, t int, percent []int) (string) {

	if t == 1 {

		nbc := maths.Abs(nb)
		if nbc <= 0.33 {
			percent[0] = percent[0] + 1
			return (fmt.Sprintf("\033[1;32m%f \033[0m", nb))
		} else if nbc > 0.33 && nbc <= 0.66 {
			percent[1] = percent[1] + 1
			return (fmt.Sprintf("\033[1;33m%f \033[0m", nb))
		} else {
			percent[2] = percent[2] + 1
			return (fmt.Sprintf("\033[1;31m%f \033[0m", nb))
		}
	}
	return (fmt.Sprintf("%f", nb))
}