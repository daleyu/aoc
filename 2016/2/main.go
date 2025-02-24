package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <txt-input file>")
		os.Exit(1)
	}
	args := os.Args
	input, e := os.ReadFile(args[1])
	check(e)
	resultPart1 := part1(string(input))
	fmt.Printf("Part1 : %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("Part1 : %v\n", resultPart2)
}

func appendToNum(source, num int) int {
	source *= 10
	source += num
	return source
}

func calcNum(x, y int) int {
	return (x * 3) + (y + 1)
}

func moveButton(x, y *int, dx, dy int) {
	if *x+dx < 0 {
		*x = 0
	} else if *x+dx > 2 {
		*x = 2
	} else {
		*x += dx
	}

	if *y+dy < 0 {
		*y = 0
	} else if *y+dy > 2 {
		*y = 2
	} else {
		*y += dy
	}
}
func part1(input string) int {
	result := 0
	var mapping = map[string][]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	x := 1
	y := 1
	for _, line := range strings.Split(input, "\n") {
		print(line)
		if len(line) == 0 {
			break
		}
		for _, cha := range line {
			change := mapping[string(cha)]
			dx, dy := change[0], change[1]
			moveButton(&x, &y, dx, dy)
		}
		fmt.Printf("x: %d, y: %d \n", x, y)
		num := calcNum(x, y)
		fmt.Println(num)
		result = appendToNum(result, num)
	}

	return result
}

func part2(input string) []string {
	var result []string

	var mapping = map[string][]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	var phonePad = [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}
	x := 2
	y := 2
	rows := len(phonePad)
	cols := len(phonePad[0])
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}
		var num string
		for _, cha := range line {
			change, _ := mapping[string(cha)]
			dx, dy := change[0], change[1]
			new_x := x + dx
			new_y := y + dy
			if 0 <= new_x && new_x < rows && 0 <= new_y && new_y < cols {
				if phonePad[new_x][new_y] != "" {
					x = new_x
					y = new_y
				}
			}
		}
		num = phonePad[x][y]
		result = append(result, num)
	}

	return result
}
