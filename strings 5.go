package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputString string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&inputString)

	_, err := strconv.ParseFloat(inputString, 64)
	if err == nil {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
