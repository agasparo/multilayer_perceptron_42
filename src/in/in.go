package in

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func ReadSTDIN(str string, t int) (float64, string) {

	fmt.Println(str)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input := ReplaceWhiteSpace(text)
	if t == 0 {
		f, _ := strconv.ParseFloat(input, 64)
		return f, ""
	}
	return 0, input
}

func ReplaceWhiteSpace(text string) (string) {

	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\r\n", "")
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\v", "")
	return (text)
}