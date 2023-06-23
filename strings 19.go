package main

import "fmt"

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	reversedString := reversesString(inputString)

	fmt.Println("String invertida:", reversedString)
}

func reversesString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
