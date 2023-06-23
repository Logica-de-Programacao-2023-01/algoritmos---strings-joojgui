package main

import (
	"fmt"
	"unicode"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	isCamelCase, wordCount := checkCamelCase(inputString)

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func checkCamelCase(input string) (bool, int) {
	if input == "" {
		return false, 0
	}

	wordCount := 1

	if !unicode.IsUpper(rune(input[0])) {
		return false, 0
	}

	for i := 1; i < len(input); i++ {

		if unicode.IsUpper(rune(input[i])) {

			if !unicode.IsUpper(rune(input[i-1])) {
				wordCount++
			}
		}
	}

	return true, wordCount
}