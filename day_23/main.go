package main

import (
	"os"
	"slices"
	"strings"
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

	input := strings.TrimSpace(string(b))
	lan := map[string]map[string]struct{}{}

	for _, connection := range strings.Split(input, "\n") {
		split := strings.Split(connection, "-")
		left, right := split[0], split[len(split)-1]

		if lan[left] == nil {
			lan[left] = map[string]struct{}{}
		}

		if lan[right] == nil {
			lan[right] = map[string]struct{}{}
		}

		lan[left][right] = struct{}{}
		lan[right][left] = struct{}{}
	}
	high := 0
	result := ""
	for c := range lan {
		for c2 := range lan[c] {
			r([]string{c, c2}, lan, &high, &result)
		}

	}
	println(result)
}

func r(network []string, lanParty map[string]map[string]struct{}, high *int, result *string) {
	if len(network) > *high {
		slices.Sort(network)
		*result = strings.Join(network, ",")
		*high = len(network)
	}

	layer := lanParty[network[len(network)-1]]

	for c := range layer {
		for _, c2 := range network {
			if _, connected := lanParty[c][c2]; !connected {
				return
			}
		}
		r(append(network, c), lanParty, high, result)
	}
}
