package main

import (
	"fmt"
	"os"
	"strconv"
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

func unquote(text string) string {
	s, err := strconv.Unquote(text)
	check(err)

	return s
}

func quote(text string) string {
	s := strconv.Quote(text)
	return s
}

// Disregarding the whitespace in the file, what is the number of characters of
// code for string literals minus the number of characters in memory for the
// values of the strings in total for the entire file?
func part1(input string) int {
	var result int
	for _, line := range strings.Split(input, "\n") {
		result += len(line) - len(unquote(line))
	}

	return result
}

func part2(input string) int {
	var result int
	for _, line := range strings.Split(input, "\n") {
		result += len(quote(line)) - len(line)
	}

	return result
}
