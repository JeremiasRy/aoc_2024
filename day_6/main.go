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

type Guard struct {
	current    int
	direction  Direction
	areaWidth  int
	areaHeight int
	area       string
	patrolled  map[int]int
}

func (g *Guard) peek() (bool, rune) {
	x, y := g.current%g.areaWidth, g.current/g.areaWidth

	switch g.direction {
	case UP:
		{
			if y-1 < 0 {
				return true, ' '
			}

			return false, rune(g.area[g.current-g.areaWidth])
		}
	case RIGHT:
		{
			if x+1 >= g.areaWidth {
				return true, ' '
			}

			return false, rune(g.area[g.current+1])
		}
	case DOWN:
		{
			if y+1 >= g.areaHeight {
				return true, ' '
			}

			return false, rune(g.area[g.current+g.areaWidth])
		}
	case LEFT:
		{
			if x-1 < 0 {
				return true, ' '
			}
			return false, rune(g.area[g.current-1])
		}
	}

	// unreachable
	return true, ' '
}

func (g *Guard) turn() {
	if g.direction+1 >= 4 {
		g.direction = 0
		return
	}

	g.direction++
}

func (g *Guard) move() {
	g.patrolled[g.current]++
	switch g.direction {
	case UP:
		{
			g.current = g.current - g.areaWidth
		}
	case RIGHT:
		{
			g.current = g.current + 1
		}
	case DOWN:
		{
			g.current = g.current + g.areaWidth
		}
	case LEFT:
		{
			g.current = g.current - 1
		}
	}
}

func (g *Guard) Patrol() {
	for {
		out, next := g.peek()

		if out {
			break
		}

		if next == '#' {
			g.turn()
		}

		g.move()
	}
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
	p := map[int]int{}
	p[current]++

	g := Guard{current: current, direction: UP, areaWidth: w, areaHeight: h, area: input, patrolled: p}
	g.Patrol()

	println(len(p))
}
