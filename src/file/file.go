package file

import (
	"encoding/json"
	"Response"
	"fmt"
	"os"
	"bytes"
	"io/ioutil"
	"encoding/csv"
	"strconv"
	"io"
	"network"
)

type Base struct {

	Name string
	Error 	float64
	Layers []Layers_json
}

type Layers_json struct {

	Type 	string
	Weigths []float64
	Bias 	[]float64
}

type Learn struct {

	Datas 		map[int][]float64
	Response	[]float64
	Ret 		int
}

func ReadFile(file_name string, TL *Learn) (int) {

	var Add, got []float64

	TL.Datas = make(map[int][]float64)
	TL.Ret = 1

	csvfile, err := os.Open(file_name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return (0)
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Response.Print(fmt.Sprintf("%s\n", err))
			return (0)
		}
		Add = got
		for i := 1; i < len(record); i++ {
			
			if record[i] == "M" || record[i] == "B" {
				b := 0.0
				a := 1.0
				if record[i] == "M" {
					b = 1.0
					a = 0.0
				} 
				TL.Response = append(TL.Response, b, a)
			} else {
				a, _ := strconv.ParseFloat(record[i], 64)
				Add = append(Add, a)
			}
		}
		TL.Datas[len(TL.Datas)] = Add
	}
	return (1)
}

func SaveFile(data [][]float64, name string, ner float64, name_aolgo string) {

	user := TransformData(data, ner, name_aolgo)
    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent("", "\t")

    err := encoder.Encode(user)
    if err != nil {
        check(err, name, 0)
        return
    }
    file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        check(err, name, 0)
        return
    }
    defer file.Close()
    file.Truncate(0)
	file.Seek(0,0)
    _, err = file.Write(buffer.Bytes())
    if err != nil {
        check(err, name, 0)
        return
    }
    check(err, name, 1)
}

func SaveGraph(SaveData network.Save, name string) {

	user := SaveData
    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent("", "\t")

    err := encoder.Encode(user)
    if err != nil {
        check(err, name, 0)
        return
    }
    file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        check(err, name, 0)
        return
    }
    defer file.Close()
    file.Truncate(0)
	file.Seek(0,0)
    _, err = file.Write(buffer.Bytes())
    if err != nil {
        check(err, name, 0)
        return
    }
    check(err, name, 1)
}

func TransformData(x [][]float64, ner float64, name_aolgo string) (Base) {

	Datas := Base{}
	Datas.Name = name_aolgo
	Datas.Error = ner
	a := 0

	for i := 0; i < len(x); i += 2 {

		Add := Layers_json{}
		if a % 2 == 0 {
			Add.Type = "Fully Connected Layer"
		} else {
			Add.Type = "Activation layer"
		}
		Add.Weigths = x[i]
		Add.Bias = x[i + 1]
		Datas.Layers = append(Datas.Layers, Add)
		a++
	}
	return (Datas)
}

func check(e error, name string, v int) {
    
    if e != nil {
        Response.Print(fmt.Sprintf("%s\n", e))
    } else {
    	if v == 1 {
    		Response.Sucess(fmt.Sprintf("File %s created", name))
    	}
    }
}

func CompErr(err float64, file_name string) (int) {

	file, e := ioutil.ReadFile(file_name)
	if e != nil {
		check(e, file_name, 0)
		return (0)
	}

	data := Base{}
	_ = json.Unmarshal([]byte(file), &data)
	if data.Error < err {
		return (-1)
	}
	return (1)
}

func GetDatas(file_name string) (int, Base) {

	file, e := ioutil.ReadFile(file_name)
	if e != nil {
		check(e, file_name, 0)
		return 1, Base{}

	}
	data := Base{}
	_ = json.Unmarshal([]byte(file), &data)
	return 0, data
}

func ReadGraph(path string) (int, []network.Save) {

	var donne []network.Save
	tab := [4]string{ "exponnential", "stair", "linear", "constant" }

	for i := 0; i < len(tab); i++ {

		file_name := path + tab[i] + ".json"
		file, e := ioutil.ReadFile(file_name)
		if e != nil {
			check(e, file_name, 0)
			return 1, donne
		}
		Datas := network.Save{}
		_ = json.Unmarshal([]byte(file), &Datas)
		donne = append(donne, Datas)
	}
	return 0, donne
}