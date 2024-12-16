package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func (p *Position) getNextPosition(d Direction) Position {
	switch d {
	case NORTH:
		return p.up()
	case SOUTH:
		return p.down()
	case EAST:
		return p.right()
	default:
		return p.left()
	}
}

func (p *Position) up() Position {
	x, y := p.x, p.y
	return Position{x: x, y: y - 1}
}

func (p *Position) down() Position {
	x, y := p.x, p.y
	return Position{x: x, y: y + 1}
}

func (p *Position) left() Position {
	x, y := p.x, p.y
	return Position{x: x - 1, y: y}
}
func (p *Position) right() Position {
	x, y := p.x, p.y
	return Position{x: x + 1, y: y}
}

type Direction int

const (
	EAST Direction = iota
	SOUTH
	WEST
	NORTH
	MAX_DIRECTION
)

type Turn int

const (
	CLOCKWISE         Turn = 1
	COUNTER_CLOCKWISE Turn = -1
)

func (d *Direction) turn(t Turn) Direction {
	r := int(t) + int(*d)
	if r < 0 {
		return NORTH
	}
	return Direction(r % int(MAX_DIRECTION))
}

type MazeNodeType int

const (
	PATH MazeNodeType = iota
	WALL
	START
	END
)

type MazeNode struct {
	t       MazeNodeType
	out     map[Direction]*MazeNode
	visited int
}

var Maze = make(map[Position]*MazeNode)

type Reindeer struct {
	current *MazeNode
	score   int
	facing  Direction
	route   map[*MazeNode]struct{}
}

func (r *Reindeer) isAtEnd() bool {
	return r.current.t == END
}

var H int
var W int

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
	var start *MazeNode
	H = len(input)
	for y, row := range input {
		if W == 0 {
			W = len(row)
		}
		for x, ch := range row {
			p := Position{x: x, y: y}
			node := MazeNode{out: map[Direction]*MazeNode{}, visited: math.MaxInt32}

			if ch == '.' {
				node.t = PATH
			} else if ch == '#' {
				node.t = WALL
			} else if ch == 'S' {
				node.t = START
				start = &node
			} else if ch == 'E' {
				node.t = END
			}
			Maze[p] = &node
		}
	}

	for position, node := range Maze {
		if node.t == WALL {
			continue
		}

		for direction := EAST; direction < MAX_DIRECTION; direction++ {

			outNode := Maze[position.getNextPosition(direction)]
			if outNode.t == WALL || outNode.t == START {
				continue
			}
			node.out[direction] = outNode
		}
	}

	queue := []Reindeer{}

	queue = append(queue, Reindeer{current: start, facing: EAST, score: 0, route: map[*MazeNode]struct{}{}})

	low := math.MaxInt32
	best := map[*MazeNode]struct{}{}

	for len(queue) > 0 {
		reindeer := queue[0]
		queue = queue[1:]
		reindeer.route[reindeer.current] = struct{}{}

		if reindeer.score > low {
			continue
		}

		if reindeer.isAtEnd() {
			if reindeer.score < low {
				low = reindeer.score
				best = reindeer.route
				println(low)
			}

			if reindeer.score == low {
				for k := range reindeer.route {
					best[k] = struct{}{}
				}
			}
			continue
		}

		if reindeer.current.visited > reindeer.score {
			reindeer.current.visited = reindeer.score
		}
		enqueue(reindeer, &queue)
		//printDebug(reindeer)
		//time.Sleep(time.Millisecond * 10)
	}
	println(low)
	println(len(best))
}

func enqueue(head Reindeer, queue *[]Reindeer) {
	for direction, node := range head.current.out {
		if node.visited < head.score {
			continue
		}
		route := map[*MazeNode]struct{}{}

		for k, v := range head.route {
			route[k] = v
		}
		next := Reindeer{current: node, facing: direction, score: head.score, route: route}
		if head.facing == direction {
			next.score += 1
		} else {
			next.score += 1001
		}
		*queue = append(*queue, next)
	}
}

func printDebug(current Reindeer) {
	fmt.Printf("### SCORE: %d\n", current.score)
	debug := make([][]string, H)

	for y := 0; y < len(debug); y++ {
		debug[y] = make([]string, W)
	}

	for p, v := range Maze {
		x, y := p.x, p.y

		if v.visited != math.MaxInt32 {
			debug[y][x] = "@"
		} else if v.t == PATH {
			debug[y][x] = "."
		} else if v.t == WALL {
			debug[y][x] = "#"
		} else {
			debug[y][x] = "E"
		}
	}

	for _, row := range debug {
		println(strings.Join(row, ""))
	}
	println()
}
