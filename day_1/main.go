package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Can't read file: ", os.Args[1])
		os.Exit(1)
	}

	input := string(b)

	leftIds := []int{}
	rightIds := map[int]int{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lr := strings.Split(line, " ")
		left, right := lr[0], lr[len(lr)-1]

		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)

		leftIds = append(leftIds, leftInt)
		rightIds[rightInt] += 1
	}

	similarity := 0
	for _, id := range leftIds {
		multiply := rightIds[id]
		similarity += id * multiply
	}

	println(similarity)
}
