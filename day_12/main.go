package main

import (
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

type PlotDimensions struct {
	area      int
	perimeter int
}

type Plant struct {
	p       Position
	t       string
	visited bool
}

func (p *Position) up() Position {
	return Position{x: p.x, y: p.y - 1}
}

func (p *Position) down() Position {
	return Position{x: p.x, y: p.y + 1}
}

func (p *Position) left() Position {
	return Position{x: p.x - 1, y: p.y}
}

func (p *Position) right() Position {
	return Position{x: p.x + 1, y: p.y}
}

func neighbouringPlants(plant *Plant, plantation map[Position]*Plant) []*Plant {
	position := plant.p
	above, below, left, right := position.up(), position.down(), position.left(), position.right()
	result := []*Plant{}

	if up, found := plantation[above]; found && up.t == plant.t && !up.visited {
		result = append(result, up)
	}

	if down, found := plantation[below]; found && down.t == plant.t && !down.visited {
		result = append(result, down)
	}

	if l, found := plantation[left]; found && l.t == plant.t && !l.visited {
		result = append(result, l)
	}

	if r, found := plantation[right]; found && r.t == plant.t && !r.visited {
		result = append(result, r)
	}

	return result
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

	rows := strings.Split(strings.TrimSpace(string(b)), "\n")
	plantation := map[Position]*Plant{}

	for y, row := range rows {
		for x, ch := range row {
			position := Position{x: x, y: y}
			plantation[position] = &Plant{t: string(ch), visited: false, p: position}
		}
	}

	plots := map[int][]*Plant{}
	id := 0

	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			current := Position{x: x, y: y}
			plant := plantation[current]
			if plant.visited {
				continue
			}
			id++

			stack := []*Plant{}
			neighbours := neighbouringPlants(plant, plantation)

			if len(neighbours) == 0 {
				plots[id] = append(plots[id], plant)
				continue
			}

			stack = append(stack, neighbours...)

			for len(stack) > 0 {
				current := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !current.visited {
					plots[id] = append(plots[id], current)
					stack = append(stack, neighbouringPlants(current, plantation)...)
					current.visited = true
				}
			}
		}
	}

	price := 0
	for _, plot := range plots {
		area := len(plot)
		perimeter := 0
		for _, plant := range plot {
			position := plant.p
			up, down, left, right := plantation[position.up()], plantation[position.down()], plantation[position.left()], plantation[position.right()]

			if up == nil || up.t != plant.t {
				perimeter++
			}

			if down == nil || down.t != plant.t {
				perimeter++
			}

			if left == nil || left.t != plant.t {
				perimeter++
			}

			if right == nil || right.t != plant.t {
				perimeter++
			}
		}
		price += area * perimeter
	}

	println(price)
}
