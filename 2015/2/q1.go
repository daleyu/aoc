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
	resultPart1 := part1(string(input))

	fmt.Printf("Total amount of Paper: %d\n", resultPart1)

	//part2
	resultPart2 := part2(string(input))
	fmt.Printf("Material for ribbons: %d\n", resultPart2)
}

func part1(input string) int {
	result := 0
	for _, line := range strings.Split(string(input), "\n") {
		dims := strings.Split(line, "x")
		l, e := strconv.Atoi(dims[0])
		check(e)
		w, e := strconv.Atoi(dims[1])
		check(e)
		h, e := strconv.Atoi(dims[2])
		check(e)
		extraSlack := l * w
		if extraSlack > l*h {
			extraSlack = l * h
		}
		if extraSlack > w*h {
			extraSlack = w * h
		}
		result += (2 * l * w) + (2 * w * h) + (2 * l * h) + extraSlack
	}
	return result
}

func part2(input string) int {
	result := 0
	//we need to find the smallest perimeter and the cubic volume
	for _, line := range strings.Split(input, "\n") {
		dims := strings.Split(line, "x")
		l, e := strconv.Atoi(dims[0])
		check(e)
		w, e := strconv.Atoi(dims[1])
		check(e)
		h, e := strconv.Atoi(dims[2])
		check(e)
		bow := l * w * h
		result += bow

		//smallest perimeter
		smallest := math.Inf(1)
		for i := 0; i < 3; i++ {
			for j := i; j < 3; j++ {
				if i == j {
					fmt.Println("Exiting here")
					continue
				}
				side1, e := strconv.Atoi(dims[i])
				check(e)
				side2, e := strconv.Atoi(dims[j])
				check(e)
				fmt.Printf("%d,%d\n", side1, side2)
				perim := side1*2 + side2*2
				if perim < int(smallest) {
					smallest = float64(perim)
				}
			}
		}
		result += int(smallest)
	}
	return result
}
