package main

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Printf("Part 1: %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("Part 2: %d\n", resultPart2)
}

func part1(input string) int{
	result := 0 
	for _, line := range strings.Split(input, "\n"){
		parts := strings.Split(line, " ")
		start, end := parts[0], parts[1]
		fmt.Printf("start: %s\n", start)
		fmt.Printf("end: %s\n", end)
	}

	return result
}

func part2(input string) int{
	result := 0
	return result
}
