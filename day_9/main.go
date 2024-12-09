package main

import (
	"os"
	"strconv"
	"strings"
)

type BlockType int

const (
	FREE BlockType = iota
	FILE
)

type Block struct {
	id int
	t  BlockType
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

	input := strings.TrimSpace(string(b))
	t := FILE
	disk := []Block{}
	id := 0

	for _, ch := range input {
		num, _ := strconv.Atoi(string(ch))

		for num > 0 {
			disk = append(disk, Block{id, t})
			num--
		}
		if t == FILE {
			id++
		}
		t = (t + 1) % 2
	}

	lo, hi := 0, len(disk)-1

	for lo < hi {
		if disk[hi].t != FILE {
			hi--
			continue
		}

		if disk[lo].t != FREE {
			lo++
			continue
		}

		if disk[hi].t == FILE && disk[lo].t == FREE {
			disk[lo], disk[hi] = disk[hi], disk[lo]
		}
	}

	checksum := 0

	for pos, block := range disk {
		if block.t == FREE {
			break
		}
		checksum += pos * block.id
	}
	println(checksum)
}
