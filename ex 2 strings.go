package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase:")
	scanner.Scan()
	frase1 := scanner.Text()

	frase1SE := strings.ReplaceAll(frase1, " ", "")

	fmt.Println(" A frase removendo os espaços fica :", frase1SE)

}
]