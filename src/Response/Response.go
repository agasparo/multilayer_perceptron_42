package Response

import (
	"github.com/fatih/color"
)

func Print(str string) {

	color.Red(str)
}

func Sucess(str string) {
	
	color.Green(str)
}

func PrintVerboseStep(str string) {

	color.Magenta(str)
}

func PrintVerbose(str string) {

	color.White(str)
}