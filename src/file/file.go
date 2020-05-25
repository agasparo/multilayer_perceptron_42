package file

import (
	"encoding/json"
	"Response"
	"fmt"
	"os"
	"bytes"
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