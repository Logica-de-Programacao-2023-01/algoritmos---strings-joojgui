package main

import "fmt"

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	reversedString := reverseString(inputString)

	fmt.Println("Resultado:", reversedString)
}

func reverseString(input string) string {
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
