package main

import (
	"os"
	"strings"
)

var Available = make(map[string]struct{})
var Possible = make(map[string]struct{})

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
	for _, pattern := range strings.Split(input[0], ",") {
		Available[strings.TrimSpace(pattern)] = struct{}{}
	}

	for _, pattern := range input[2:] {
		r(0, pattern, map[string]map[string]struct{}{})
	}
	println(len(Possible))
}

func r(current int, pattern string, visited map[string]map[string]struct{}) {
	remainder := pattern[current:]

	if _, exists := visited[remainder]; !exists {
		visited[remainder] = map[string]struct{}{}
	}

	if _, solved := Possible[pattern]; solved {
		return
	}

	if current == len(pattern) {
		Possible[pattern] = struct{}{}
		return
	}

	for i := current + 1; i <= len(pattern); i++ {
		_, match := Available[pattern[current:i]]

		if _, explored := visited[remainder][pattern[current:i]]; explored {
			continue
		}

		if match {
			visited[remainder][pattern[current:i]] = struct{}{}
			r(i, pattern, visited)
		}
	}
}
