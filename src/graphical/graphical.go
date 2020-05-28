package graphical

import (
	"Response"
	"in"
	"strings"
	"network"
	"create"
	"strconv"
)

type CreateAl func(*network.Net) (int)

type GoTo struct {

	ToDo 	int
	Create 	CreateAl
	Name	string
	Decrease int
}

func ShowMain(D *GoTo) {

	var res string

	Choice := []string {
		"You wan't to learn or to predict ? [L, P, G] (L -> learn, P -> predict, G -> graph)",
		"You wan't wich neronal network ? [XOR, CUSTOM]",
		"You Want wich decrease for learning rate ? [0, 1, 2, 3] (exp, stair, linear, const)",
	}

	Responses := [][]string {
		[]string{ "L", "P", "G" },
		[]string{ "XOR", "CUSTOM" },
		[]string{ "0", "1", "2", "3" },
	}

	for i := 0; i < len(Choice); i++ {

		_, res = in.ReadSTDIN(Choice[i], 1)
		for !In_array(Responses[i], res) {
			Response.Print("Response must be " + strings.Join(Responses[i], " or "))
			_, res = in.ReadSTDIN(Choice[i], 1)
		}
		if i == 0 {
			D.ToDo = GetIndex(Responses[i], res) 
			if res == "G" {
				_, res = in.ReadSTDIN("wich network choose ? [XOR, CUSTOM]", 1)
				for res != "XOR" && res != "CUSTOM" {
					Response.Print("Response must be XOR or CUSTOM")
					_, res = in.ReadSTDIN(Choice[i], 1)
				}
				D.Name = res
				return
			}
		} else if i == 1 {
			if res == "XOR" {
				D.Create = create.XOR
			} else {
				D.Create = create.CUSTOM
			}
			D.Name = res
		} else {
			D.Decrease, _ = strconv.Atoi(res)
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