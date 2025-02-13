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

var directions = [][]int{
	{1,0},
	{0,1},
	{-1,0},
	{0,-1},
}
func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	args := os.Args
	println(args[1])
	input, err := os.ReadFile(args[1])
	check(err)
	input_str := string(input)

	resultPart1 := part1(input_str)
	fmt.Printf("Part 1: %d\n", resultPart1)


	resultPart2 := part2(input_str)
	fmt.Printf("Part 2: %d\n", resultPart2)
}

var index = 0


func mod(a, b int ) int {
	return (a % b + b) % b
}

func changeDirection(turn string) {
	if turn == "L"{
		index = ((index + 1) % 4)
	}else if turn == "R" {
		index = ((index + 3) % 4)
	}
}

func part1(input string) int{
	x := 0
	y := 0 
	for _, move := range strings.Split(input, ","){
		move = strings.TrimSpace(move)
		turn := string(move[0])
		magnitude, e := strconv.Atoi(string(move[1:]))
		check(e)
		fmt.Println(turn)
		changeDirection(turn)

		x += directions[index][0] * magnitude
		y += directions[index][1] * magnitude
		fmt.Println(magnitude)
	}
	return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
}

func part2(input string) int{
	visited := make(map[[2]int]bool)
	visited[[2]int{0,0}] = true
	x := 0
	y := 0 
	for _, move := range strings.Split(input, ","){
		move = strings.TrimSpace(move)
		turn := string(move[0])
		magnitude, e := strconv.Atoi(string(move[1:]))
		check(e)
		fmt.Println(turn)
		changeDirection(turn)

		for i := 0; i < magnitude; i++ {
			x += directions[index][0] 
			y += directions[index][1]
			curr := [2]int{x,y}
			if visited[curr] {
				return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
			}
			visited[curr] = true
		}
	}
	fmt.Println("We got to the end without ever getting to the easter bunny")
	return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
}


