package main

import (
	"strconv"
	"strings"
)

func Solution(input string) int {
	safe := 0
	for _, report := range strings.Split(input, "\n") {
		if len(report) == 0 {
			continue
		}

		nums := parseSequence(strings.Split(report, " "))
		isSafe := isValidSequence(nums)

		if !isSafe {
			for i := range nums {
				sub := append(make([]int, 0), nums[:i]...)
				sub = append(sub, nums[i+1:]...)

				if isValidSequence(sub) {
					safe++
					break
				}
			}
		} else {
			safe++
		}
	}
	return safe
}

func parseSequence(nums []string) []int {
	r := make([]int, len(nums))
	for i, str := range nums {
		num, _ := strconv.Atoi(str)
		r[i] = num
	}
	return r
}

func isValidSequence(nums []int) bool {
	prev := nums[0]
	valid := true
	dir := NONE

	for _, current := range nums[1:] {
		diff := prev - current

		if dir == NONE {
			if diff > 0 {
				dir = Decreasing
			} else if diff < 0 {
				dir = Increasing
			} else {
				valid = false
				break
			}
		}

		if dir == Increasing && diff >= -3 && diff <= -1 {
			prev = current
			continue
		}

		if dir == Decreasing && diff <= 3 && diff >= 1 {
			prev = current
			continue
		}
		valid = false
		break
	}
	/*
		if dir == Increasing {
			println("Ascending")
		} else if dir == Decreasing {
			println("Descending")
		}
	*/
	return valid
}
