package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	result := removeVowels(inputString)

	fmt.Println("Resultado:", result)
}

func removeVowels(input string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		input = strings.ReplaceAll(input, vowel, "")
	}

	return input
}
