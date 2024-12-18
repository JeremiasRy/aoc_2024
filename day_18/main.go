package main

import (
	"os"
	"strconv"
	"strings"
)

type XY struct {
	x, y int
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
	MAX_DIRECTION
)

func (p *XY) getNextPosition(d Direction) XY {
	x, y := p.x, p.y
	switch d {
	case UP:
		return XY{x: x, y: y - 1}
	case DOWN:
		return XY{x: x, y: y + 1}
	case RIGHT:
		return XY{x: x + 1, y: y}
	case LEFT:
		return XY{x: x - 1, y: y}
	default:
		println("Invalid direction")
		os.Exit(1)
		return XY{}
	}
}

var MemorySpace = make(map[XY]bool)

const WIDTH = 71
const HEIGHT = 71

type Historians struct {
	p     XY
	steps int
}

type Queue struct {
	vals []Historians
}

func (q *Queue) enqueue(val Historians) {
	q.vals = append(q.vals, val)
}

func (q *Queue) dequeue() Historians {
	r := q.vals[0]
	q.vals = q.vals[1:]
	return r
}

func (q *Queue) isEmpty() bool {
	return len(q.vals) == 0
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

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			p := XY{x: x, y: y}
			MemorySpace[p] = false
		}
	}

	input := strings.Split(strings.TrimSpace(string(b)), "\n")

	for _, row := range input[:1024] {
		position := strings.Split(row, ",")

		x, y := position[0], position[len(position)-1]
		p := parsePosition(x, y)

		MemorySpace[p] = true
	}

	for _, pixel := range input[1024:] {
		position := strings.Split(pixel, ",")

		x, y := position[0], position[len(position)-1]
		p := parsePosition(x, y)

		MemorySpace[p] = true

		if !canReachExit() {
			println(pixel)
			break
		}
	}
}

func canReachExit() bool {
	init := Historians{p: XY{x: 0, y: 0}, steps: 0}
	queue := Queue{vals: []Historians{}}
	best := map[XY]int{}

	queue.enqueue(init)

	for !queue.isEmpty() {
		current := queue.dequeue()

		low, visited := best[current.p]

		if !visited {
			best[current.p] = current.steps
		}

		if visited && current.steps > low {
			continue
		}

		if current.steps < low {
			best[current.p] = current.steps
		}

		//printDebug(MemorySpace, BestRoute)

		for direction := UP; direction < MAX_DIRECTION; direction++ {
			next := current.p.getNextPosition(direction)
			corrupted, inBounds := MemorySpace[next]

			if !inBounds || corrupted || visited {
				continue
			}

			clone := Historians{p: next, steps: current.steps + 1}
			queue.enqueue(clone)
		}
	}
	return best[XY{x: 70, y: 70}] != 0
}

func printDebug(m map[XY]bool, visited map[XY]int) {
	s := make([][]string, HEIGHT)
	for y := 0; y < len(s); y++ {
		s[y] = make([]string, WIDTH)
	}

	for position, corrupted := range m {
		x, y := position.x, position.y

		_, onRoute := visited[position]
		if onRoute {
			s[y][x] = "O"
		} else if corrupted {
			s[y][x] = "#"
		} else {
			s[y][x] = "."
		}
	}

	for _, row := range s {
		println(strings.Join(row, ""))
	}
	println()
}

func parsePosition(x string, y string) XY {
	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)

	return XY{x: xInt, y: yInt}
}
