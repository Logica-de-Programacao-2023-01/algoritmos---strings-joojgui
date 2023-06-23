package main

import (
	"fmt"
	"regexp"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	containsNumber := containsNumber(inputString)

	if containsNumber {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}

func containsNumber(input string) bool {
	match, _ := regexp.MatchString("[0-9]", input)
	return match
}
