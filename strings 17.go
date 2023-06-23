package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	uniqueLetters := findUniqueLetters(inputString)

	fmt.Println("Letras únicas:", uniqueLetters)
}

func findUniqueLetters(input string) string {
	letterCount := make(map[rune]int)

	for _, char := range input {
		letterCount[char]++
	}

	var uniqueLetters []string

	for _, char := range input {
		if letterCount[char] == 1 {
			uniqueLetters = append(uniqueLetters, string(char))
		}
	}

	return strings.Join(uniqueLetters, "")
}
