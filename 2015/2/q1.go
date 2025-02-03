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
		l, e := strconv.Atoi(dims[0])
		check(e)
		w, e := strconv.Atoi(dims[1])
		check(e)
		h, e := strconv.Atoi(dims[2])
		check(e)
		result += (2 * l * w) + (2 * w * h) + (2 * l * h)
	}

	fmt.Printf("Total amount of Paper: %d", result)
}
