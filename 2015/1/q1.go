package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %v\n", e)
		panic(e)
	}
}

func main() {
	result := 0
	input, err := os.ReadFile("q1.sample.txt")
	enteredBasement := 0
	check(err)
	for i, c := range input {
		currChar := string(c)
		if currChar == "(" {
			result += 1
		} else {
			result -= 1
		}
		if result < 0 && enteredBasement < 1 {
			fmt.Printf("Entered Basement at position: %d\n", i+1)
			enteredBasement += 1
		}
	}

	fmt.Printf("Floor: %d", result)
}
