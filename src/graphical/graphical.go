package graphical

import (
	"Response"
	"in"
	"strings"
	"network"
	"create"
	"strconv"
	"fmt"
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

	defs()
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

func defs() {

	Response.PrintVerboseStep("Feedforward :")
	fmt.Println("L'information ne se déplace que dans une seule direction, vers l'avant, à partir des nœuds d'entrée, en passant par les couches cachées (le cas échéant) et vers les noeuds de sortie. Il n'y a pas de cycles ou de boucles dans le réseau")
	Response.PrintVerboseStep("\nBackpropagation :")
	fmt.Println("la rétropropagation calcule le gradient de la fonction de perte par rapport aux poids du réseau pour un exemple d'entrée-sortie unique, et le fait efficacement, contrairement à un calcul direct naïf du gradient par rapport à chaque poids individuellement")
	Response.PrintVerboseStep("\nGradient descent :")
	fmt.Println("est une méthode pour calculer le gradient de l'erreur pour chaque neurone")
	Response.PrintVerboseStep("\nOverfitting : (surapprentissage)")
	fmt.Println("est une analyse statistique qui correspond trop étroitement ou exactement à un ensemble particulier de données. Ainsi, cette analyse peut ne pas correspondre à des données supplémentaires ou ne pas prévoir de manière fiable les observations futures")
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