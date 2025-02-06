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

	fmt.Printf("Lights on: %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("Light level = %d\n", resultPart2)
}

func part1(input string) int {
	result := 0
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		toggle := false
		on := false
		startX := 0
		startY := 0
		endX := 0
		endY := 0
		line := string(line)

		switch {
		case strings.HasPrefix(line, "toggle"):
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			toggle = true
		case strings.HasPrefix(line, "turn on"):
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			on = true
		case strings.HasPrefix(line, "turn off"):
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			on = false
		default:
			fmt.Println(line)
			fmt.Println("Something must be wrong with input")
			continue
		}

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				if toggle {
					grid[i][j] = !grid[i][j]
					continue
				}
				if on {
					grid[i][j] = true
				} else {
					grid[i][j] = false
				}
			}
		}

	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				fmt.Printf("x: %d, y: %d\n", i, j)
				result++
			}
		}
	}
	return result
}

func part2(input string) int {
	result := 0
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		toggle := false
		on := false
		startX := 0
		startY := 0
		endX := 0
		endY := 0
		line := string(line)

		switch {
		case strings.HasPrefix(line, "toggle"):
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			toggle = true
		case strings.HasPrefix(line, "turn on"):
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			on = true
		case strings.HasPrefix(line, "turn off"):
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &startX, &startY, &endX, &endY)
			on = false
		default:
			fmt.Println(line)
			fmt.Println("Something must be wrong with input")
			continue
		}

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				if toggle {
					grid[i][j] += 2
					continue
				}
				if on {
					grid[i][j]++
				} else {
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				}
			}
		}

	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			result += grid[i][j]
		}
	}
	return result
}
