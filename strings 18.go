package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	isNumeric := checkNumeric(inputString)

	if isNumeric {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

func checkNumeric(input string) bool {
	_, err := strconv.Atoi(input)

	return err == nil
}
