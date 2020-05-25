package graphical

import (
	"io/ioutil"
	"fmt"
	"Response"
	"in"
	"strings"
	"network"
	"create"
)

type CreateAl func(*network.Net)

type GoTo struct {

	ToDo 	int
	Create 	CreateAl
	Name	string
}

func ShowMain(D *GoTo) {

	var res string

	Choice := []string {
		"You wan't to learn or to predict ? [L/P] (L -> learn, P -> predict)",
		"You wan't wich neronal network ? [XOR]",
	}

	Responses := [][]string {
		[]string{ "L", "P" },
		[]string{ "XOR" },	
	}

	read, _ := ioutil.ReadFile("data/logo.txt")
	Response.PrintVerboseStep(string(read))

	fmt.Println("")

	for i := 0; i < len(Choice); i++ {

		_, res = in.ReadSTDIN(Choice[i], 1)
		for !In_array(Responses[i], res) {
			Response.Print("Response must be " + strings.Join(Responses[i], " or "))
			_, res = in.ReadSTDIN(Choice[i], 1)
		}
		if i == 0 {
			D.ToDo = GetIndex(Responses[i], res) 
		} else {
			D.Create = create.XOR
			D.Name = res
		}
	}
}

func GetIndex(data []string, search string) (int) {

	for i := 0; i < len(data); i++ {

		if data[i] == search {
			return (i)
		}
	} 
	return (0)
}

func In_array(data []string, search string) (bool) {

	for i := 0; i < len(data); i++ {

		if data[i] == search {
			return (true)
		}
	} 
	return (false)
}