package main

import (
	"os"
	"slices"
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

	current := strings.Index(input, "^")
	start := current
	dir := UP

	route := map[int][]Direction{}
	print := strings.Split(input, "")

	print[current] = "0"

	for {
		out, next := peek(input, current, dir, w, h)

		if out {
			break
		}

		route[current] = append(route[current], dir)

		if next == '#' {
			if print[current] != "0" {
				print[current] = "+"
			}
			if dir+1 >= 4 {
				dir = UP
			} else {
				dir++
			}
		}
		current = move(current, dir, w)
	}
	route[current] = append(route[current], dir)

	for k, v := range route {
		//fmt.Printf("%v\n", getXY(k, w))
		if len(v) >= 2 && print[k] != "0" {
			print[k] = "+"
		} else if print[k] != "+" && print[k] != "0" {
			switch v[0] {
			case UP:
				print[k] = "^"
			case DOWN:
				print[k] = "v"
			case LEFT:
				print[k] = "<"
			case RIGHT:
				print[k] = ">"
			}
		}

	}

	println(strings.Join(print, ""))
	count := 0
	infiniteLoopObstaclePositions := map[int]int{}
	for obstacle, directions := range route {
		possible := addObstacle(input, obstacle)
		current = start

		loop, _ := isInfiniteLoop(possible, current, w, h)
		if loop {
			count++
			infiniteLoopObstaclePositions[obstacle]++
		}

		for _, dir := range directions {
			obstacle = move(obstacle, dir, w)
			possible := addObstacle(input, obstacle)
			current = start

			loop, _ := isInfiniteLoop(possible, current, w, h)
			if loop {
				infiniteLoopObstaclePositions[obstacle]++
			}
		}

	}
	println(len(infiniteLoopObstaclePositions))
	println(count)

	hailMary := 0
	for _, v := range infiniteLoopObstaclePositions {
		hailMary += v
	}

	println(hailMary)
}

func isInfiniteLoop(input string, current int, w int, h int) (bool, string) {
	route := map[int][]Direction{}
	debug := strings.Split(input, "")
	dir := UP

	isInfinite := false
	for {
		out, next := peek(input, current, dir, w, h)

		if out {
			break
		}

		if slices.Contains(route[current], dir) {
			isInfinite = true
			break
		}

		route[current] = append(route[current], dir)

		if next == '#' || next == 'O' {
			if debug[current] != "^" {
				debug[current] = "+"
			}

			if dir+1 >= 4 {
				dir = UP
			} else {
				dir++
			}
		}
		current = move(current, dir, w)
	}

	for k, v := range route {
		if len(v) >= 2 && debug[k] != "^" {
			debug[k] = "+"
		} else if debug[k] != "+" && debug[k] != "^" && debug[k] != "0" {
			switch v[0] {
			case UP:
				fallthrough
			case DOWN:
				debug[k] = "|"
			case LEFT:
				fallthrough
			case RIGHT:
				debug[k] = "-"
			}
		}
	}

	return isInfinite, strings.Join(debug, "")
}

func peek(input string, current int, dir Direction, w int, h int) (bool, byte) {
	next := getXY(current, w)
	x, y := next.x, next.y
	switch dir {
	case UP:
		{
			if y-1 < 0 {
				return true, ' '
			}
			return false, input[current-w]
		}
	case RIGHT:
		{
			if x+1 >= w-1 {
				return true, ' '
			}
			return false, input[current+1]
		}
	case DOWN:
		{
			if y+1 >= h {
				return true, ' '
			}
			return false, input[current+w]
		}
	case LEFT:
		{
			if x-1 < 0 {
				return true, ' '
			}
			return false, input[current-1]
		}
	}
	// unreachable
	return true, ' '
}

func move(current int, dir Direction, w int) int {
	switch dir {
	case UP:
		return current - w
	case RIGHT:
		return current + 1
	case DOWN:
		return current + w
	case LEFT:
		return current - 1
	}

	return -1
}

func addObstacle(input string, obstacle int) string {
	old := strings.Split(input, "")
	new := make([]string, len(old))
	copy(new, old)

	new[obstacle] = "O"
	return strings.Join(new, "")
}

func getXY(current int, w int) Position {
	return Position{x: current % w, y: current / w}
}
