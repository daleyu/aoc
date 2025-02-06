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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	args := os.Args
	println(args[1])
	input, err := os.ReadFile(args[1])
	check(err)
	resultPart1 := part1(string(input))

	fmt.Printf("Nice Strings: %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("Nice Strings Part 2: %d\n", resultPart2)
}

var vowels = set.NewStringSet("a", "e", "i", "o", "u")
var forbidden = set.NewStringSet("ab", "cd", "pq", "xy")

func part1(input string) int {
	result := 0
	for _, line := range strings.Split(input, "\n") {
		vowelCount := 0
		repeatCount := 0
		failed := false
		for i, c := range line {
			cha := string(c)
			if vowels.Has(cha) {
				vowelCount++
			}
			if i < len(line)-1 {
				subseq := line[i : i+2]
				if forbidden.Has(subseq) {
					failed = true
				}

				if line[i] == line[i+1] {
					repeatCount++
				}
			}
		}
		if failed {
			continue
		}

		if vowelCount >= 3 && repeatCount >= 1 {
			result++
		}
	}

	return result
}

func part2(input string) int {
	checkPair := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			temp := line[i : i+2]
			for j := i + 2; j < len(line)-1; j++ {
				if temp == line[j:j+2] {
					return true
				}
			}

		}
		return false
	}
	result := 0
	for _, line := range strings.Split(input, "\n") {
		pairsPossible := checkPair(line)
		if !pairsPossible {
			continue
		}
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				result++
				break
			}
		}
	}

	return result
}
