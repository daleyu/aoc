package main

import (
	"fmt"
	"os"

	"github.com/daleyu/aoc-practice/data-structures/set"
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
	resultPart1 := part1(string(input))

	fmt.Printf("Total houses: %d\n", resultPart1)
}

func part1(input string) int {
	// for santa to deliver presents we can just consider the starting point as origin (0,0)
	// This means we need like a set of some sort where we hash on the coordinates
	// making the set in data-structures now
	s := set.NewStringSet()

	x := 0
	y := 0
	directions := make(map[string][]int)
	directions["^"] = []int{0, 1}
	directions["<"] = []int{-1, 0}
	directions[">"] = []int{1, 0}
	directions["v"] = []int{0, -1}
	for _, c := range input {
		dx := directions[string(c)][0]
		dy := directions[string(c)][1]
		fmt.Printf(" dx: %d, dy %d\n", dx, dy)
		x += dx
		y += dy
		s.Add("(" + string(x) + "," + string(y) + ")")
	}
	fmt.Printf("keys: %v", s.Keys())
	return s.Size()
}
