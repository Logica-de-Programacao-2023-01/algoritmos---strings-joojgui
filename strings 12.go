package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	isPalindrome := checkPalindrome(inputString)

	if isPalindrome {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func checkPalindrome(input string) bool {
	input = strings.ToLower(input)
	input = removeSpaces(input)

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func removeSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}
