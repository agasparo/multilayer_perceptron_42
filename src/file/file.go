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

	Datas 	map[int][]float64
	Ret 	int
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
			
			if record[i] == "M" {
				record[i] = "1"
			} else if record[i] == "B" {
				record[i] = "0"
			}
			a, _ := strconv.ParseFloat(record[i], 64)
			Add = append(Add, a)
		}
		TL.Datas[len(TL.Datas)] = Add
	}
	return (1)
}

func SaveFile(data [][]float64, name string, ner float64) {

	user := TransformData(data, ner)
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

func TransformData(x [][]float64, ner float64) (Base) {

	Datas := Base{}
	Datas.Name = "XOR"
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