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
	result := 0
	input, err := os.ReadFile("input.txt")
	check(err)
	for _, line := range strings.Split(string(input), "\n") {
		dims := strings.Split(line, "x")
		extraSlack, e := strconv.Atoi(dims[0])
		check(e)
		for _, side := range dims {
			if side <= extraSlack {
				extraSlack = side
			}
		}
		l, e := strconv.Atoi(dims[0])
		check(e)
		w, e := strconv.Atoi(dims[1])
		check(e)
		h, e := strconv.Atoi(dims[2])
		check(e)
		fmt.Printf("%d * %d * %d = %d\n", l, w, h, result)
		result += (2 * l * w) + (2 * w * h) + (2 * l * h) + extraSlack
	}

	fmt.Printf("Total amount of Paper: %d\n", result)
}
