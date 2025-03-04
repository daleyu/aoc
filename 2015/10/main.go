package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("error")
		os.Exit(1)
	}
	return
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
}
