#!/bin/bash

mkdir day_$1
cd day_$1

go mod init day_$1
touch main.go
touch input

echo "package main 
import (\"os\")
func main(){if len(os.Args)!= 2{println(\"Usage: go run main.go <input>\")
os.Exit(1)}

b, err := os.ReadFile(os.Args[1])
if err != nil {println(\"Can't read file: \", os.Args[1])
os.Exit(1)}

//input := string(b)
}" > main.go

go fmt main.go

SESSION=`cat ../session`
URL="https://adventofcode.com/2024/day/$1/input"

curl $URL -X GET --cookie "session=$SESSION" > input