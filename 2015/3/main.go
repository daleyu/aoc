package main

import (
	"fmt"
	"math"
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
	input, err := os.ReadFile("input.txt")
	check(err)
	result := part1(string(input))

	fmt.Printf("Total amount of Paper: %d\n", resultPart1)
}

func part1(input string) int {
	// for santa to deliver presents we can just consider the starting point as origin (0,0)
	// This means we need like a set of some sort where we hash on the coordinates
	// making the set in data-structures now
	return 0
}
