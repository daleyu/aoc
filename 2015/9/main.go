package main

import (
	"fmt"
	"strings"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %v\n", e)
		panic(e)
	}
}

type Graph struct {
	NumNodes 	int
	Nodes		[][]Edges
}

type Edges struct {
	From 	string
	To		string
	weight 	int
}

func NewGraph() *Graph{

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

func part1(input string) int{
	result := 0 
	graph := map[
	for _, line := range strings.Split(input, "\n"){
		parts := strings.Split(line, " ")
		if len(parts) < 2{
			fmt.Println("Something is wrong with the input")
		}
		start, end := parts[0], parts[1]

	}

	return result
}

func part2(input string) int{
	result := 0 

	return result
}

