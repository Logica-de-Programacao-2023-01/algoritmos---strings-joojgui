package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputString string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	isIncreasing := checkIncreasingSequence(inputString)

	if isIncreasing {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}

func checkIncreasingSequence(input string) bool {
	if len(input) <= 1 {
		return false
	}

	for i := 0; i < len(input)-1; i++ {
		currNum, err1 := strconv.Atoi(string(input[i]))
		nextNum, err2 := strconv.Atoi(string(input[i+1]))

		if err1 != nil || err2 != nil || nextNum != currNum+1 {
			return false
		}
	}

	return true
}
