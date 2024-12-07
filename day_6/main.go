package main

import (
	"math"
	"os"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Position struct {
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

	w := len(rows[0]) + 1
	h := len(rows[:len(rows)-1])

	current := 0
	dir := UP

	obstacles := map[int]Position{}

	for i, ch := range input {
		if ch == '#' {
			obstacles[i] = getXY(i, w)
		} else if ch == '^' {
			current = i
		}
	}

	travel := map[int]struct{}{}

Loop:
	for {
		position := getXY(current, w)

		switch dir {
		case UP:
			{
				thereIs, at := isThereAnObstacleAbove(position, obstacles)
				if thereIs {
					diff := position.y - at - 1
					for i := 1; i <= diff; i++ {
						current -= w
						travel[current] = struct{}{}
					}
					dir = RIGHT
				} else {
					diff := h - (h - position.y)
					for i := 1; i <= diff; i++ {
						current -= w
						travel[current] = struct{}{}
					}
					break Loop
				}

			}
		case DOWN:
			{
				thereIs, at := isThereAnObstacleBelow(position, obstacles)
				if thereIs {
					diff := at - position.y - 1
					for i := 1; i <= diff; i++ {
						current += w
						travel[current] = struct{}{}
					}
					dir = LEFT
				} else {
					diff := h - position.y
					for i := 1; i <= diff; i++ {
						current += w
						travel[current] = struct{}{}
					}
					break Loop
				}
			}
		case LEFT:
			{
				thereIs, at := isThereAnObstacleLeft(position, obstacles)
				if thereIs {
					diff := position.x - at - 1
					for i := 1; i <= diff; i++ {
						current--
						travel[current] = struct{}{}
					}
					dir = UP
				} else {
					diff := w - (w - position.x)
					for i := 1; i <= diff; i++ {
						current--
						travel[current] = struct{}{}
					}
					break Loop
				}
			}
		case RIGHT:
			{
				thereIs, at := isThereAnObstacleRight(position, obstacles)

				if thereIs {
					diff := at - position.x - 1
					for i := 1; i <= diff; i++ {
						current++
						travel[current] = struct{}{}
					}
					dir = DOWN
				} else {

					diff := w - position.x
					for i := 1; i <= diff; i++ {
						current++
						travel[current] = struct{}{}
					}
					break Loop
				}
			}
		}
	}
	println(len(travel) - 1)
}

func isThereAnObstacleBelow(pos Position, obstacles map[int]Position) (bool, int) {
	x, y := pos.x, pos.y
	min := math.MaxInt32

	for _, o := range obstacles {
		if y < o.y && x == o.x {

			if o.y < min {
				min = o.y
			}
		}
	}
	return min != math.MaxInt32, min
}

func isThereAnObstacleAbove(pos Position, obstacles map[int]Position) (bool, int) {
	x, y := pos.x, pos.y
	min := -1

	for _, o := range obstacles {
		if y > o.y && x == o.x {
			if o.y > min {
				min = o.y
			}
		}
	}
	return min != -1, min
}

func isThereAnObstacleLeft(pos Position, obstacles map[int]Position) (bool, int) {
	x, y := pos.x, pos.y
	min := -1

	for _, o := range obstacles {
		if y == o.y && x > o.x {
			if o.x > min {
				min = o.x
			}
		}
	}
	return min != -1, min
}

func isThereAnObstacleRight(pos Position, obstacles map[int]Position) (bool, int) {
	x, y := pos.x, pos.y
	min := math.MaxInt32

	for _, o := range obstacles {
		if y == o.y && x < o.x {
			if o.x < min {
				min = o.x
			}
		}
	}
	return min != math.MaxInt32, min
}

func getXY(current int, w int) Position {
	return Position{x: current % w, y: current / w}
}
