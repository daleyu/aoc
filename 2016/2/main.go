package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error){
	if e != nil{
		fmt.Printf("Error: %s\n", e)
		os.Exit(1)
	}
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <txt-input file>\n")
		os.Exit(1)
	}
	args := os.Args
	input, e := os.ReadFile(args[1])
	check(e)
	resultPart1 := part1(string(input))
	fmt.Printf("Part1 : %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("Part1 : %d\n", resultPart2)
}

func appendToNum(source, num int) int {
	source *= 10
	source += num
	return source
}

func calcNum(x, y int) int{
	return x * 3 + y
}

func moveButton(

func part1(input string) int{
	result := 0 
	var mapping = map[string][]int{
		"U": {1,0},
		"D": {0,1},
		"L": {-1,0},
		"R": {0,-1},
	}
	x := 1
	y := 1
	for _, line := range strings.Split(input, "\n") {
		for _, cha := range strings.Split(line, ""){
			change := mapping[cha]
			dx, dy := change[0], change[1]
			x += dx 
			y += dy 
			num := calcNum(x,y)
			result = appendToNum(result, num)
		}
	}

	return result
}

func part2(input string) int{
	result := 0 

	return result
}
