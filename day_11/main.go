package main

import (
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	val string
}

func (s *Stone) blink() []Stone {
	r := []Stone{}
	if s.isZero() {
		r = append(r, Stone{val: "1"})
	} else if s.hasEvenNumberOfDigits() {
		r = append(r, s.split()...)
	} else {
		i, _ := strconv.Atoi(s.val)
		v := strconv.Itoa(i * 2024)
		r = append(r, Stone{val: v})
	}

	return r
}

func (s *Stone) isZero() bool {
	return s.val == "0"
}

func (s *Stone) hasEvenNumberOfDigits() bool {
	return len(s.val)&1 == 0
}

func (s *Stone) split() []Stone {
	arr := strings.Split(s.val, "")

	half := len(arr) / 2

	first, second := strings.Join(arr[0:half], ""), strings.Join(arr[half:], "")

	if allZerosAndLongerThanOne(second) {
		second = second[len(second)-2 : len(second)-1]
	} else if len(second) > 1 {
		second = strings.TrimPrefix(second, "0")
	}

	return []Stone{{val: first}, {val: second}}
}

func allZerosAndLongerThanOne(s string) bool {
	if len(s) == 1 {
		return false
	}
	for _, ch := range s {
		if ch != '0' {
			return false
		}
	}

	return true
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

	stones := []Stone{}
	input := string(b)
	for _, s := range strings.Split(strings.TrimSpace(input), " ") {
		stones = append(stones, Stone{val: s})
	}

	blinks := 25

	for blinks > 0 {
		new := []Stone{}
		for _, stone := range stones {
			new = append(new, stone.blink()...)
		}
		stones = new
		blinks--
	}

	println(len(stones))
}
