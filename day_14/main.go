package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	WIDTH  = 101
	HEIGHT = 103
)

type Position struct {
	x int
	y int
}

func (p *Position) isEqual(other Position) bool {
	return p.x == other.x && p.y == other.y
}

func (p *Position) move(delta Position) {
	x, y := p.x, p.y
	dx, dy := delta.x, delta.y

	if x+dx < 0 {
		i := int(math.Abs(float64(dx)))

		for i > 0 {
			if x-1 < 0 {
				x = (WIDTH - 1)
			} else {
				x -= 1
			}
			i--
		}
	} else {
		x = (x + dx) % WIDTH
	}

	if y+dy < 0 {
		i := int(math.Abs(float64(dy)))

		for i > 0 {
			if y-1 < 0 {
				y = (HEIGHT - 1)
			} else {
				y -= 1
			}
			i--
		}
	} else {
		y = (y + dy) % HEIGHT
	}

	p.x, p.y = x, y
}

type Robot struct {
	p Position
	v Position
}

func (r *Robot) tick() {
	r.p.move(r.v)
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

	input := strings.Split(strings.TrimSpace(string(b)), "\n")
	robots := []*Robot{}

	for _, row := range input {
		position := row[strings.Index(row, "p=")+2 : strings.Index(row, " ")]
		velocity := row[strings.Index(row, "v=")+2:]

		x, y := strings.Split(position, ",")[0], strings.Split(position, ",")[1]
		dx, dy := strings.Split(velocity, ",")[0], strings.Split(velocity, ",")[1]

		px, _ := strconv.Atoi(x)
		py, _ := strconv.Atoi(y)
		vx, _ := strconv.Atoi(dx)
		vy, _ := strconv.Atoi(dy)

		robots = append(robots, &Robot{p: Position{x: px, y: py}, v: Position{x: vx, y: vy}})
	}

	seconds := 0
	for seconds < 100 {
		for _, r := range robots {
			r.tick()
		}
		seconds++
	}

	middleX := WIDTH / 2
	middleY := HEIGHT / 2

	first, second, third, fourth := 0, 0, 0, 0

	for _, r := range robots {
		x, y := r.p.x, r.p.y

		if x < middleX && y < middleY {
			first++
		} else if x > middleX && y < middleY {
			second++
		} else if x < middleX && y > middleY {
			third++
		} else if x > middleX && y > middleY {
			fourth++
		}
	}

	println(first * second * third * fourth)
}
