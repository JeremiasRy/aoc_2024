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
	sets := map[string]struct{}{}
	for c1, connections := range lan {
		for c2 := range connections {
			layer1 := lan[c2]

			for c3 := range layer1 {
				layer2 := lan[c3]

				if _, found := layer2[c1]; found {
					set := []string{c1, c2, c3}
					slices.Sort(set)

					j := strings.Join(set, ",")
					sets[j] = struct{}{}
				}
			}
		}
	}

	count := 0
	for network := range sets {
		for _, computer := range strings.Split(network, ",") {
			if computer[0] == 't' {
				count++
				break
			}
		}
	}

	println(count)
}
