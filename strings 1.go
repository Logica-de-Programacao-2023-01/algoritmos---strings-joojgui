package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e converta todas as letras minúsculas para maiúsculas. Informe ao usuário o resultado.

func main() {
	var s string

	fmt.Print("digite uma frase em com letras minusculas :")
	fmt.Scan(&s)

	s = strings.ToUpper(s)

	fmt.Println("A string em maisuculo fica:", s)

}
