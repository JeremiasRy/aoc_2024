package main

import (
	"os"
	"strconv"
	"strings"
)

type OP int

const (
	ADD OP = iota
	MULTIPLY
)

type StackItem struct {
	current  int
	equation []string
}

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
	result := 0

	for _, equation := range strings.Split(input, "\n") {
		if len(equation) == 0 {
			continue
		}

		nums := parseNums(strings.Split(strings.Join(strings.Split(equation, ":"), ""), " "))

		target, nums := nums[0], append(nums[1:], -1)
		stack := [][]int{}
		stack = append(stack, []int{nums[0]})

		for _, num := range nums[1:] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next := []int{}

			for _, current := range pop {
				if num == -1 && current == target {
					result += current
					break
				}

				sum := current + num
				product := current * num

				if sum <= target {
					next = append(next, sum)
				}

				if product <= target {
					next = append(next, product)
				}
			}
			stack = append(stack, next)
		}
	}
	println(result)

}

func parseNums(nums []string) []int {
	r := make([]int, len(nums))

	for i := range nums {
		num, _ := strconv.Atoi(nums[i])
		r[i] = num
	}
	return r
}
