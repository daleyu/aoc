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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	args := os.Args
	println(args[1])
	input, err := os.ReadFile(args[1])
	check(err)
	resultPart1 := part1(string(input))

	fmt.Printf("Value on a: %d\n", resultPart1)

	resultPart2 := part2(string(input))
	fmt.Printf("New value: %d\n", resultPart2)
}

func part1(input string) int {
	var result int

	dp := make(map[string]int)
	graph := make(map[string]string)

	var dfs func(currWire string) int
	dfs = func(currWire string) int {
		var wireVal int
		// if it is just a num
		if val, err := strconv.Atoi(currWire); err == nil {
			return val
		}

		// if we have seen this wire in our dp
		if dpVal, ok := dp[currWire]; ok {
			return dpVal
		}

		// run dfs
		children := graph[currWire]
		parts := strings.Split(children, " ")

		switch {
		case len(parts) == 1:
			wireVal = dfs(parts[0])
		case parts[0] == "NOT":
			wireVal = (math.MaxUint16) ^ dfs(parts[1])

		case parts[1] == "AND":
			wireVal = dfs(parts[0]) & dfs(parts[2])
		case parts[1] == "OR":
			wireVal = dfs(parts[0]) | dfs(parts[2])
		case parts[1] == "RSHIFT":
			wireVal = dfs(parts[0]) >> dfs(parts[2])
		case parts[1] == "LSHIFT":
			wireVal = dfs(parts[0]) << dfs(parts[2])
		default:
			fmt.Println("Something must've went wrong with the string")
		}

		// init dp mapping
		dp[currWire] = wireVal
		return wireVal
	}

	// create graph
	for _, line := range strings.Split(input, "\n") {
		nodes := strings.Split(line, "->")
		for i := range nodes {
			nodes[i] = strings.TrimSpace(nodes[i])
		}
		graph[nodes[1]] = nodes[0]
	}

	result = dfs("a")

	return result
}

func part2(input string) int {
	var result int

	dp := make(map[string]int)
	graph := make(map[string]string)

	var dfs func(currWire string) int
	dfs = func(currWire string) int {
		var wireVal int
		// if it is just a num
		if val, err := strconv.Atoi(currWire); err == nil {
			return val
		}

		// if we have seen this wire in our dp
		if dpVal, ok := dp[currWire]; ok {
			return dpVal
		}

		// run dfs
		children := graph[currWire]
		parts := strings.Split(children, " ")

		switch {
		case len(parts) == 1:
			wireVal = dfs(parts[0])
		case parts[0] == "NOT":
			wireVal = (math.MaxUint16) ^ dfs(parts[1])

		case parts[1] == "AND":
			wireVal = dfs(parts[0]) & dfs(parts[2])
		case parts[1] == "OR":
			wireVal = dfs(parts[0]) | dfs(parts[2])
		case parts[1] == "RSHIFT":
			wireVal = dfs(parts[0]) >> dfs(parts[2])
		case parts[1] == "LSHIFT":
			wireVal = dfs(parts[0]) << dfs(parts[2])
		default:
			fmt.Println("Something must've went wrong with the string")
		}

		// init dp mapping
		dp[currWire] = wireVal
		return wireVal
	}

	// create graph
	for _, line := range strings.Split(input, "\n") {
		nodes := strings.Split(line, "->")
		for i := range nodes {
			nodes[i] = strings.TrimSpace(nodes[i])
		}
		graph[nodes[1]] = nodes[0]
	}

	result = dfs("a")

	graph["b"] = strconv.Itoa(result)
	dp = make(map[string]int)
	result = dfs("a")

	return result
}
