package main

import (
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
	rows := strings.Split(strings.TrimSpace(input), "\n")

	var start Position

Loop:
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {

			if rows[i][j] == '^' {
				start = Position{x: j, y: i}
				break Loop
			}
		}
	}

	current := start
	route := map[Position]struct{}{}

	w := len(rows[0])
	h := len(rows)

	dir := UP
	for {
		out, next, position := peek(rows, current, dir, w, h)
		if out {
			route[current] = struct{}{}
			break
		}

		route[current] = struct{}{}

		if next == '#' {
			dir = (dir + 1) % 4
		} else {
			current = position
		}
	}
	println(len(route))
	loops := 0

	for position := range route {
		if isInfiniteLoop(rows, start, w, h, position) {
			loops++
		}
	}
	println(loops)
}

func isInfiniteLoop(rows []string, current Position, w int, h int, newObstacle Position) bool {
	direction := UP
	obstacles := map[Position]map[Direction]bool{}
	for {
		out, next, position := peek(rows, current, direction, w, h)

		if out {
			return false
		}

		if next == '#' || newObstacle == position {
			if _, exists := obstacles[position]; !exists {
				obstacles[position] = map[Direction]bool{}
			}
			if obstacles[position][direction] {
				return true
			}
			obstacles[position][direction] = true
			direction = (direction + 1) % 4
		} else {
			current = position
		}
	}
}

func peek(rows []string, p Position, dir Direction, w int, h int) (bool, byte, Position) {
	x, y := p.x, p.y
	switch dir {
	case UP:
		{
			if isOutOfBounds(w, h, x, y-1) {
				return true, 0, Position{}
			}

			return false, rows[y-1][x], Position{x, y - 1}

		}
	case DOWN:
		{
			if isOutOfBounds(w, h, x, y+1) {
				return true, 0, Position{}
			}
			return false, rows[y+1][x], Position{x, y + 1}

		}
	case LEFT:
		{
			if isOutOfBounds(w, h, x-1, y) {
				return true, 0, Position{}
			}
			return false, rows[y][x-1], Position{x - 1, y}
		}
	case RIGHT:
		{
			if isOutOfBounds(w, h, x+1, y) {
				return true, 0, Position{}
			}
			return false, rows[y][x+1], Position{x + 1, y}
		}
	default:
		{
			os.Exit(1)
		}
	}

	return false, 0, Position{}
}

func isOutOfBounds(w int, h int, x int, y int) bool {
	return x < 0 || x >= w || y < 0 || y >= h
}
