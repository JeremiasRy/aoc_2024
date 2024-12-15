package main

import (
	"os"
	"strings"
)

type DirectionToken int

const (
	T_UP DirectionToken = iota
	T_LEFT
	T_DOWN
	T_RIGHT
)

func parseDirection(ch rune) DirectionToken {
	switch ch {
	case '<':
		return T_LEFT
	case '>':
		return T_RIGHT
	case '^':
		return T_UP
	default:
		return T_DOWN
	}
}

type Position struct {
	x, y int
}

func (p *Position) getNextPosition(d DirectionToken) Position {
	x, y := p.x, p.y
	switch d {
	case T_DOWN:
		return Position{x: x, y: y + 1}
	case T_UP:
		return Position{x: x, y: y - 1}
	case T_LEFT:
		return Position{x: x - 1, y: y}
	case T_RIGHT:
		return Position{x: x + 1, y: y}
	}
	return Position{}
}

type MapTokenType int

const (
	T_ROBOT = iota
	T_BOX
	T_FREE
	T_WALL
)

func parseMapTokenType(ch rune) MapTokenType {
	if ch == '.' {
		return T_FREE
	} else if ch == 'O' {
		return T_BOX
	} else if ch == '#' {
		return T_WALL
	} else {
		return T_ROBOT
	}
}

var Map = map[Position]MapTokenType{}
var H, W int

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

	input := strings.Split(strings.TrimSpace(string(b)), "\n")
	var directions []string
	var current Position

	for y, row := range input {
		if W == 0 {
			W = len(row)
		}
		if len(row) == 0 {
			directions = input[y+1:]
			H = y
			break
		}
		for x, ch := range row {
			p := Position{x: x, y: y}
			t := parseMapTokenType(ch)

			Map[p] = t
			if t == T_ROBOT {
				current = p
			}
		}
	}

	for _, ch := range strings.Join(directions, "") {
		//debugPosition(Map)
		//println(string(ch))
		d := parseDirection(ch)

		next := current.getNextPosition(d)
		t, found := Map[next]

		if !found || t == T_WALL {
			continue
		}

		if t == T_FREE || (t == T_BOX && moveBox(next, d)) {
			Map[current] = T_FREE
			Map[next] = T_ROBOT
			current = next
		}
	}
	total := 0
	for p, t := range Map {
		if t == T_BOX {
			total += 100*p.y + p.x
		}
	}

	println(total)
}

func moveBox(p Position, direction DirectionToken) bool {
	next := p.getNextPosition(direction)
	t := Map[next]

	if t == T_FREE || (t == T_BOX && moveBox(next, direction)) {
		Map[p] = T_FREE
		Map[next] = T_BOX
		return true
	}
	return false
}

func debugPosition(m map[Position]MapTokenType) {
	debug := make([][]string, H)
	for i := 0; i < len(debug); i++ {
		debug[i] = make([]string, W)
		for j := 0; j < len(debug[i]); j++ {
			debug[i][j] = "."
		}
	}
	for p, t := range m {
		if t == T_BOX {
			debug[p.y][p.x] = "O"
		} else if t == T_WALL {
			debug[p.y][p.x] = "#"
		} else if t == T_ROBOT {
			debug[p.y][p.x] = "@"
		} else {
			debug[p.y][p.x] = "."
		}
	}
	for _, line := range debug {
		println(strings.Join(line, ""))
	}
	println()

}
