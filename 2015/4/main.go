package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
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

	fmt.Printf("Hash Solutions 5 zeros: %d\n", resultPart1)
	resultPart2 := part2(string(input))

	fmt.Printf("Hash Solutions 6 zeros: %d\n", resultPart2)
}

func part1(input string) int {
	result := 0
	numZeros := 5
	for i := 0; i < math.MaxInt32; i++ {
		entry := fmt.Sprintf("%s%d", input, i)
		md5Hash := md5.Sum([]byte(entry))
		hashed := fmt.Sprintf("%x", md5Hash)
		if hashed[:numZeros] == "00000" {
			return i
		}

	}
	return result
}

func part2(input string) int {
	result := 0
	numZeros := 6
	for i := 0; i < math.MaxInt32; i++ {
		entry := fmt.Sprintf("%s%d", input, i)
		md5Hash := md5.Sum([]byte(entry))
		hashed := fmt.Sprintf("%x", md5Hash)
		if hashed[:numZeros] == "000000" {
			return i
		}

	}
	return result
}
