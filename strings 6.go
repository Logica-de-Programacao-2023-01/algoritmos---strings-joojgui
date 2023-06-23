package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	palavras := strings.Fields(inputString)
	numPalavras := len(palavras)

	fmt.Printf("A string contém %d palavra(s).\n", numPalavras)
}
