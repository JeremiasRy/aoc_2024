package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	input := string(file)

	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	count := 0

	for r, row := range rows {
		for c, ch := range row {
			if r-1 >= 0 && r+1 < len(row) && c-1 >= 0 && c+1 < len(row) && ch == 'A' {
				tl, tr, bl, br := rows[r-1][c-1], rows[r-1][c+1], rows[r+1][c-1], rows[r+1][c+1]

				d1 := strings.Join([]string{string(tl), string(ch), string(br)}, "")
				d2 := strings.Join([]string{string(tr), string(ch), string(bl)}, "")

				if (d1 == "MAS" || d1 == "SAM") && (d2 == "MAS" || d2 == "SAM") {
					count++
				}
			}
		}
	}
	println(count)
}
