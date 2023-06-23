package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString, oldLetter, newLetter string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&oldLetter)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&newLetter)

	replacedString := strings.ReplaceAll(inputString, oldLetter, newLetter)

	fmt.Println("Resultado:", replacedString)
}
