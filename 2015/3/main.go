package main

import (
	"fmt"
	"os"
	"strings"

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

	resultPart2 := part2(string(input))
	fmt.Printf("With Robo: %d\n", resultPart2)
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
	for _, c := range strings.Split(input, "") {
		fmt.Printf("%s\n", c)
		dx := directions[string(c)][0]
		dy := directions[string(c)][1]
		fmt.Printf(" dx: %d, dy %d\n", dx, dy)
		x += dx
		y += dy
		s.Add(fmt.Sprintf("(%d,%d)", x, y))
	}
	return s.Size()
}

func part2(input string) int {
	// robot santa and santa take turns
	s := set.NewStringSet()
	santa := [2]int{0, 0}
	robo := [2]int{0, 0}
	directions := make(map[string][]int)
	directions["^"] = []int{0, 1}
	directions["<"] = []int{-1, 0}
	directions[">"] = []int{1, 0}
	directions["v"] = []int{0, -1}

	for i, c := range strings.Split(input, "") {
		dx := directions[string(c)][0]
		dy := directions[string(c)][1]
		if i%2 == 0 {
			santa[0] += dx
			santa[1] += dy
			s.Add(fmt.Sprintf("(%d, %d)", santa[0], santa[1]))
		} else {
			robo[0] += dx
			robo[1] += dy
			s.Add(fmt.Sprintf("(%d, %d)", robo[0], robo[1]))
		}

	}
	return s.Size()
}
