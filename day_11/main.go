package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func blink(i int) (int, int) {
	length := int(math.Log10(float64(i)) + 1)
	if i == 0 {
		return 1, -1
	} else if length&1 == 0 {
		split := int(math.Pow10(length - length/2))
		return i / split, i % split

	}
	return i * 2024, -1
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

	stones := []int{}
	input := string(b)
	for _, s := range strings.Split(strings.TrimSpace(input), " ") {
		i, _ := strconv.Atoi(s)
		stones = append(stones, i)
	}

	blinks := 25

	for blinks > 0 {
		new := []int{}
		for _, stone := range stones {
			first, second := blink(stone)
			new = append(new, first)
			if second >= 0 {
				new = append(new, second)
			}
		}
		stones = new
		blinks--
	}

	println(len(stones))
}
