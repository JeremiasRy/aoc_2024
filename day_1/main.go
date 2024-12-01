package main

import (
	"math"
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
	rightIds := []int{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lr := strings.Split(line, " ")
		left, right := lr[0], lr[len(lr)-1]

		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)

		leftIds = insert(leftInt, leftIds)
		rightIds = insert(rightInt, rightIds)
	}
	distance := 0
	i := 0

	for i < len(leftIds) {
		distance += int(math.Abs(float64(leftIds[i] - rightIds[i])))
		i++
	}
	println(distance)
}

func insert(num int, nums []int) []int {
	nums = append(nums, num)
	i := 0

	for i < len(nums) {
		j := i
		for j > 0 && nums[j-1] > nums[j] {
			nums[j-1], nums[j] = nums[j], nums[j-1]
			j = j - 1
		}
		i++
	}
	return nums
}
