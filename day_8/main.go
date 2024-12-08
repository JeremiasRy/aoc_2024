package main

import (
	"os"
	"strings"
)

type Antenna struct {
	x int
	y int
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
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]

	antennas := map[string][]Antenna{}
	antinodes := map[int]struct{}{}

	w := len(rows[0])
	h := len(rows)

	for y, row := range rows {
		for x, ch := range row {
			frequency := string(ch)
			if frequency != "." && frequency != "\n" {
				antennas[frequency] = append(antennas[frequency], Antenna{x, y})
			}
		}
	}

	for _, nodes := range antennas {
		for _, a := range nodes {
			for _, b := range nodes {
				if a == b {
					continue
				}
				deltaX, deltaY := distanceBetween(a, b)
				antinodeX, antinodeY := a.x+deltaX, a.y+deltaY

				if !isOutOfBounds(antinodeX, antinodeY, w, h) {
					index := XYToPosition(antinodeX, antinodeY, w+1)
					antinodes[index] = struct{}{}
				}
			}
		}
	}
	println(len(antinodes))
}

func XYToPosition(x int, y int, w int) int {
	return w*y + x
}

func distanceBetween(a Antenna, b Antenna) (int, int) {
	return a.x - b.x, a.y - b.y
}

func isOutOfBounds(x int, y int, w int, h int) bool {
	return x < 0 || x >= w || y < 0 || y >= h
}
