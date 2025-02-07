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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	args := os.Args
	println(args[1])
	input, err := os.ReadFile(args[1])
	check(err)
	resultPart1 := part1(string(input))

	fmt.Printf("Value on a: %d\n", resultPart1)

	// resultPart2 := part2(string(input))
	// fmt.Printf("New value: %d\n", resultPart2)
}

func part1(input string) int {
	var result int

	return result
}
