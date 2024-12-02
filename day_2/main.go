package main

import (
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Increasing Direction = iota
	Decreasing
	NONE
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
	safe := 0
	for _, report := range strings.Split(input, "\n") {
		if len(report) == 0 {
			continue
		}

		nums := strings.Split(report, " ")
		prev, _ := strconv.Atoi(nums[0])
		dir := NONE
		isSafe := true

		for _, str := range nums[1:] {
			current, _ := strconv.Atoi(str)
			diff := prev - current

			if dir == NONE {
				if diff == 0 {
					isSafe = false
					break
				}

				if diff <= -1 && diff >= -3 {
					dir = Increasing
					prev = current
					continue
				}

				if diff >= 1 && diff <= 3 {
					dir = Decreasing
					prev = current
					continue
				}
				isSafe = false
				break
			}

			if dir == Increasing {
				if !(diff <= -1 && diff >= -3) {
					isSafe = false
					break
				}
			}

			if dir == Decreasing {
				if !(diff >= 1 && diff <= 3) {
					isSafe = false
					break
				}
			}

			prev = current
		}
		if isSafe {
			safe += 1
		}
	}
	println(safe)
}
