package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func checkAlphanum(inputStr string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(inputStr)
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	var i int
	var outputStr []rune
	for _, element := range s {
		capturedStr := checkAlphanum(string(element))
		if !capturedStr {
			outputStr = append(outputStr, element)
			continue
		}
		i++
		if i == 3 {

			changeCase := unicode.ToUpper(element)
			outputStr = append(outputStr, changeCase)
			i = 0
			continue
		}
		outputStr = append(outputStr, element)
	}
	output := string(outputStr)
	return output
}

func main() {

	var input string
	input = "divine oper..atio???n"

	outputStr := CapitalizeEveryThirdAlphanumericChar(input)
	fmt.Println(outputStr)
}
