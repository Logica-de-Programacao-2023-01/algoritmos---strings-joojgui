package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	var oldChar, newChar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&newChar)

	result := strings.ReplaceAll(inputString, oldChar, newChar)

	fmt.Println("Resultado:", result)
}
