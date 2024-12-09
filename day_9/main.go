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

type Chunk struct {
	lo int
	hi int
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

	fileChunks := getFileChunks(disk)
	freeChunks := getFreeChunks(disk)
	for _, file := range fileChunks {
		for _, free := range freeChunks {
			fileSize := file.hi - file.lo
			freeSize := free.hi - free.lo

			if freeSize >= fileSize && free.lo < file.lo {
				i := 0

				for i < fileSize {
					disk[free.lo+i], disk[file.hi-i] = disk[file.hi-i], disk[free.lo+i]
					i++
				}
				break
			}
		}
		freeChunks = getFreeChunks(disk)
	}

	checksum := 0

	for pos, block := range disk {

		if block.t == FREE {
			continue
		}
		checksum += pos * block.id
	}
	println(checksum)
}

func getFileChunks(disk []Block) []Chunk {
	result := []Chunk{}

Loop:
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i].t == FREE {
			continue
		}
		hi := i
		id := disk[i].id
		for disk[i].t == FILE && disk[i].id == id {
			i--
			if i < 0 {
				break Loop
			}
		}
		lo := i
		i++
		result = append(result, Chunk{lo, hi})
	}
	return result
}

func getFreeChunks(disk []Block) []Chunk {
	result := []Chunk{}

	for i := 0; i < len(disk); i++ {
		if disk[i].t == FILE {
			continue
		}

		lo := i
		for disk[i].t == FREE {
			i++
			if i >= len(disk) {
				break
			}
		}

		hi := i
		result = append(result, Chunk{lo, hi})
	}
	return result
}
