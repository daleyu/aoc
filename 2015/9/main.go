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

type Graph map[string]map[string]int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	args := os.Args
	println(args[1])
	input, err := os.ReadFile(args[1])
	check(err)
	input_str := string(input)

	resultPart1, resultPart2 := part1(input_str)
	fmt.Printf("Part 1: %d\nPart 2: %d \n", resultPart1, resultPart2)
}

func part1(input string) (int, int) {
	graph := make(Graph)
	// create the graph
	for _, line := range strings.Split(input, "\n") {
		if len(line) <= 0 {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			fmt.Println("Something is wrong with the input")
			os.Exit(1)
		}
		start, end := parts[0], parts[2]
		weight, e := strconv.Atoi(parts[4])
		check(e)
		if graph[start] == nil {
			graph[start] = make(map[string]int)
		}
		if graph[end] == nil {
			graph[end] = make(map[string]int)
		}
		graph[start][end] = weight
		graph[end][start] = weight
	}

	min := math.MaxInt32
	max := 0

	// run traveling statesman
	for node := range graph {
		visited := map[string]bool{node: true}
		min_dist, max_dist := dfs(graph, node, visited)
		if min_dist < min {
			min = min_dist
		}
		if max_dist > max {
			max = max_dist
		}
	}
	return min, max
}

func dfs(graph Graph, node string, visited map[string]bool) (int, int) {
	if len(visited) == len(graph) {
		return 0, 0
	}
	min := math.MaxInt32
	max := 0
	for neigh := range graph {
		if visited[neigh] {
			continue
		}

		visited[neigh] = true
		weight := graph[node][neigh]
		min_dist, max_dist := dfs(graph, neigh, visited)
		if weight+min_dist < min {
			min = weight + min_dist
		}
		if weight+max_dist > max {
			max = weight + max_dist
		}

		delete(visited, neigh)
	}

	return min, max
}
