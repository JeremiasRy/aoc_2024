package main

import (
	"os"
	"slices"
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
	ruleMap := map[int][]int{}

	for _, rule := range strings.Split(input, "\n") {
		if rule == "\n" {
			break
		}

		nums := strings.Split(rule, "|")

		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[len(nums)-1])

		ruleMap[first] = append(ruleMap[first], second)
	}

	updates := strings.Split(input, "\n\n")[1]
	result := 0

	for _, update := range strings.Split(updates, "\n") {
		if len(update) == 0 {
			continue
		}
		nums := parseNums(strings.Split(update, ","))
		corrected := false

		i := 0
		for i < len(nums) {
			current := nums[i]
			rules := ruleMap[current]

			j := i + 1

			for j < len(nums) {
				next := nums[j]
				if !slices.Contains(rules, next) {
					corrected = true
					nums[i], nums[j] = nums[j], nums[i]
					i -= 1
					break
				}
				j++
			}
			i++
		}

		if corrected {
			middle := nums[len(nums)/2]
			result += middle
		}
	}

	println(result)
}

func parseNums(s []string) []int {
	r := make([]int, len(s))

	for i := range s {
		num, _ := strconv.Atoi(s[i])
		r[i] = num
	}
	return r
}
