package main

import (
	"os"
	"strings"
)

const XMAS = "XMAS"
const SAMX = "SAMX"

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

	rows := strings.Split(string(b), "\n")

	rows = rows[:len(rows)-1]
	columns := []string{}
	diagonals := []string{}

	for r, row := range rows {
		c := []string{}
		for col := range row {
			c = append(c, string(rows[col][r]))
		}
		columns = append(columns, strings.Join(c, ""))
	}

	for r := len(rows) - 1; r >= 0; r-- {
		travel := len(rows) - 1 - r
		d := []string{}
		d = append(d, string(rows[r][0]))

		for c := 1; c <= travel; c++ {
			d = append(d, string(rows[r+c][c]))
		}
		diagonals = append(diagonals, strings.Join(d, ""))
	}

	for c := len(rows) - 1; c >= 1; c-- {
		travel := len(rows) - 1 - c
		d := []string{}
		d = append(d, string(rows[0][c]))

		for r := 1; r <= travel; r++ {
			d = append(d, string(rows[r][c+r]))
		}
		diagonals = append(diagonals, strings.Join(d, ""))
	}

	for r := len(rows) - 1; r >= 0; r-- {
		travel := len(rows) - 1 - r
		d := []string{}
		d = append(d, string(rows[r][len(rows)-1]))

		for c := 1; c <= travel; c++ {
			d = append(d, string(rows[r+c][len(rows)-1-c]))
		}
		diagonals = append(diagonals, strings.Join(d, ""))
	}

	for c := 0; c < len(rows)-1; c++ {
		travel := c
		d := []string{}
		d = append(d, string(rows[0][c]))

		for r := 1; r <= travel; r++ {
			d = append(d, string(rows[r][c-r]))
		}
		diagonals = append(diagonals, strings.Join(d, ""))
	}

	count := 0

	for i := 0; i < len(rows); i++ {
		count += strings.Count(rows[i], XMAS)
		count += strings.Count(rows[i], SAMX)
		count += strings.Count(columns[i], XMAS)
		count += strings.Count(columns[i], SAMX)
	}

	for _, diagonal := range diagonals {
		count += strings.Count(diagonal, XMAS)
		count += strings.Count(diagonal, SAMX)
	}

	println(count)
}
