package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type ParseState int
type Machine struct {
	a      XY
	b      XY
	target XY
}

type Game struct {
	m        *Machine
	spend    int
	aPresses int
	bPresses int
	current  XY
}

func (g *Game) isOver() bool {
	return (g.current.x > g.m.target.x || g.current.y > g.m.target.y) || (g.aPresses >= MAX_PLAY || g.bPresses >= MAX_PLAY)
}

func (g *Game) targetHit() bool {
	return g.current.x == g.m.target.x && g.current.y == g.m.target.y
}

func (g *Game) pressA() {
	delta := g.m.a
	g.current.x += delta.x
	g.current.y += delta.y

	g.aPresses++
	g.spend += A_PRICE
}

func (g *Game) dePressA() {
	delta := g.m.a
	g.current.x -= delta.x
	g.current.y -= delta.y

	g.aPresses--
	g.spend -= A_PRICE
}

func (g *Game) pressB() {
	delta := g.m.b
	g.current.x += delta.x
	g.current.y += delta.y

	g.bPresses++
	g.spend += B_PRICE
}

func (g *Game) dePressB() {
	delta := g.m.b
	g.current.x -= delta.x
	g.current.y -= delta.y

	g.bPresses--
	g.spend -= B_PRICE
}

type XY struct {
	x int
	y int
}

const MAX_PLAY = 100
const A_PRICE = 3
const B_PRICE = 1

const (
	BUTTON_A ParseState = iota
	BUTTON_B
	PRIZE
	MAX_PARSE_STATE
)

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

	p := BUTTON_A
	machines := []*Machine{}
	currentMachine := Machine{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		switch p {
		case BUTTON_A:
			currentMachine.a = parseXY(row)
		case BUTTON_B:
			currentMachine.b = parseXY(row)
		case PRIZE:
			{
				x := row[strings.Index(row, "X=")+2 : strings.Index(row, ",")]
				y := row[strings.Index(row, "Y=")+2:]

				prizeX, _ := strconv.Atoi(x)
				prizeY, _ := strconv.Atoi(y)
				currentMachine.target = XY{x: prizeX, y: prizeY}

				c := currentMachine
				machines = append(machines, &c)
			}
		}
		p = (p + 1) % MAX_PARSE_STATE
	}

	result := 0
	for _, machine := range machines {
		gameA := Game{m: machine, current: XY{x: 0, y: 0}, spend: 0, aPresses: 0, bPresses: 0}
		gameB := Game{m: machine, current: XY{x: 0, y: 0}, spend: 0, aPresses: 0, bPresses: 0}
		low := math.MaxInt32
		for !gameA.isOver() {
			gameA.pressA()
		}

		l := backTrackGameA(gameA)

		if l != math.MaxInt32 {
			low = l
		}

		for !gameB.isOver() {
			gameB.pressB()
		}

		l = backTrackGameB(gameB)

		if l < low {
			low = l
		}

		if low != math.MaxInt32 {
			result += low
		}

	}
	println(result)
}

func backTrackGameA(game Game) int {
	low := math.MaxInt32

	for i := 0; i < game.aPresses; i++ {
		game.dePressA()

		bees := 0
		for !game.isOver() {
			game.pressB()
			bees++

			if game.targetHit() {
				if game.spend < low {
					low = game.spend
				}
				for bees >= 0 {
					game.dePressB()
					bees--
				}
				break
			}
		}
		for bees >= 0 {
			game.dePressB()
			bees--
		}
	}
	return low
}

func backTrackGameB(game Game) int {
	low := math.MaxInt32

	for i := 0; i < game.bPresses; i++ {
		game.dePressB()

		aas := 0
		for !game.isOver() {
			game.pressA()
			aas++

			if game.targetHit() {
				if game.spend < low {
					low = game.spend
				}
				for aas > 0 {
					game.dePressA()
					aas--
				}
				break
			}
		}
		for aas > 0 {
			game.dePressA()
			aas--
		}
	}
	return low
}

func parseXY(row string) XY {
	x := row[strings.Index(row, "X+")+2 : strings.Index(row, ",")]
	y := row[strings.Index(row, "Y+")+2:]
	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)

	return XY{x: xInt, y: yInt}
}
