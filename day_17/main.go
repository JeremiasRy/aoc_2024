package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type Register int
type LiteralOP int

const (
	A Register = iota
	B
	C
)

const (
	adv LiteralOP = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

var register = make(map[Register]int)
var combo = make(map[int]func() int)
var literal = make(map[LiteralOP]func(op int) bool)

var instructionPointer = 0
var output = []string{}

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Can't read file: ", os.Args[1])
		os.Exit(1)
	}

	populateComboOperations()
	populateLiteralOperations()

	input := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	a := input[0]
	b := input[1]
	c := input[2]

	valA, _ := strconv.Atoi(a[strings.Index(a, "Register A: ")+12:])
	valB, _ := strconv.Atoi(b[strings.Index(a, "Register B: ")+12:])
	valC, _ := strconv.Atoi(c[strings.Index(a, "Register C: ")+12:])

	register[A] = valA
	register[B] = valB
	register[C] = valC

	p := input[len(input)-1]

	p = p[strings.Index(p, "Program: ")+9:]

	program := []int{}

	for _, op := range strings.Split(p, ",") {
		val, _ := strconv.Atoi(op)
		program = append(program, val)
	}

	for instructionPointer < len(program)-1 {
		opCode := program[instructionPointer]
		operand := program[instructionPointer+1]
		//fmt.Printf("ip: %d, op: %d, register: %v, output: %v\n", instructionPointer, opCode, register, output)

		jump := literal[LiteralOP(opCode)](operand)

		if jump {
			instructionPointer += 2
		}
	}

	println(strings.Join(output, ","))
}

func populateComboOperations() {
	combo[0] = func() int { return 0 }
	combo[1] = func() int { return 1 }
	combo[2] = func() int { return 2 }
	combo[3] = func() int { return 3 }
	combo[4] = func() int { return register[A] }
	combo[5] = func() int { return register[B] }
	combo[6] = func() int { return register[C] }
	combo[7] = func() int {
		println("RESERVED")
		return -1
	}
}

func populateLiteralOperations() {
	literal[adv] = func(op int) bool {
		register[A] = register[A] / int(math.Pow(2, float64(combo[op]())))
		return true
	}

	literal[bxl] = func(op int) bool {
		register[B] = register[B] ^ op
		return true
	}

	literal[bst] = func(op int) bool {
		//println("bst, ", op)
		register[B] = combo[op]() % 8
		return true
	}

	literal[jnz] = func(op int) bool {
		if register[A] == 0 {
			return true
		}
		instructionPointer = op
		return false
	}

	literal[bxc] = func(op int) bool {
		register[B] = register[B] ^ register[C]
		return true
	}

	literal[out] = func(op int) bool {
		output = append(output, strconv.Itoa(combo[op]()%8))
		return true
	}

	literal[bdv] = func(op int) bool {
		register[B] = register[A] / int(math.Pow(2, float64(combo[op]())))
		return true
	}

	literal[cdv] = func(op int) bool {
		register[C] = register[A] / int(math.Pow(2, float64(combo[op]())))
		return true
	}
}
