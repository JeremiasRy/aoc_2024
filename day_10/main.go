package main

import (
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func (p *Position) up() Position {
	return Position{y: p.y - 1, x: p.x}
}
func (p *Position) down() Position {
	return Position{y: p.y + 1, x: p.x}
}
func (p *Position) left() Position {
	return Position{y: p.y, x: p.x - 1}
}
func (p *Position) right() Position {
	return Position{y: p.y, x: p.x + 1}
}

type Node struct {
	h        int
	w        int
	val      int
	position Position
	next     []*Node
	visited  map[*Node]struct{}
}

func (n *Node) isTrailHead() bool {
	return n.val == 0
}

func (n *Node) isTrailEnd() bool {
	return n.val == 9
}

func (n *Node) mapNeighbourNodes(topoMap map[Position]*Node) {
	for _, position := range []Position{n.position.up(), n.position.down(), n.position.left(), n.position.right()} {
		if !outOfBounds(position, n.h, n.w) {
			trail, found := topoMap[position]
			if found && trail.val-n.val == 1 {
				n.next = append(n.next, trail)
			}
		}
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

	rows := strings.Split(strings.TrimSpace(input), "\n")

	h := len(rows)
	w := len(rows[0])

	topoMap := map[Position]*Node{}

	for y, row := range rows {
		for x, ch := range row {
			val, err := strconv.Atoi(string(ch))
			if err != nil {
				continue
			}
			position := Position{y, x}
			topoMap[position] = &Node{h: h, w: w, val: val, position: position, next: []*Node{}, visited: map[*Node]struct{}{}}
		}
	}

	trailHeads := []*Node{}
	for _, node := range topoMap {
		node.mapNeighbourNodes(topoMap)
		if node.isTrailHead() {
			trailHeads = append(trailHeads, node)
		}
	}

	result := 0
	for _, head := range trailHeads {
		start := head
		current := head
		ends := 0

		stack := current.next

		for len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if current.isTrailEnd() {
				if _, visited := current.visited[start]; !visited {
					current.visited[start] = struct{}{}
					ends++
				}
				continue
			}

			stack = append(stack, current.next...)
		}
		result += ends
	}
	println(result)

}

func outOfBounds(p Position, w int, h int) bool {
	x, y := p.x, p.y
	return x < 0 || x >= w || y < 0 || y >= h
}
