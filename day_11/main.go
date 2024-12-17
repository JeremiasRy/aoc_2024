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

	input := string(b)
	counts := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(input), " ") {
		i, _ := strconv.Atoi(s)
		counts[i]++
	}

	blinks := 75

	for blinks > 0 {
		next := map[int]int{}
		for stone, count := range counts {
			first, second := blink(stone)
			next[first] += count
			if second != -1 {
				next[second] += count
			}
		}
		counts = next
		blinks--
	}

	total := 0
	for _, count := range counts {
		total += count
	}

	println(total)
}
