package main

import "testing"

func TestSolution1(t *testing.T) {
	expected := 1
	input := "7 6 4 2 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d", input, expected)
	}
}

func TestSolution2(t *testing.T) {
	expected := 1
	input := "7 7 4 2 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}
func TestSolution3(t *testing.T) {
	expected := 1
	input := "7 4 2 1 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution4(t *testing.T) {
	expected := 0
	input := "7 7 4 2 1 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution5(t *testing.T) {
	expected := 1
	input := "7 6 5 5 3 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution6(t *testing.T) {
	expected := 0
	input := "7 6 5 5 3 1 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution7(t *testing.T) {
	expected := 0
	input := "7 7 6 5 5 3 1 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution8(t *testing.T) {
	expected := 0
	input := "7 6 9 5 5 3 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution9(t *testing.T) {
	expected := 0
	input := "1 2 3 2 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution10(t *testing.T) {
	expected := 0
	input := "1 2 3 3 2 1"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution11(t *testing.T) {
	expected := 1
	input := "1 4 7 10 11 5 14 16 17 20"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution12(t *testing.T) {
	expected := 1
	input := "1 4 4 7"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}

func TestSolution13(t *testing.T) {
	expected := 1
	input := "34 28 25 22"
	result := Solution(input)

	if result != expected {
		t.Errorf("\nInput:\n%s\nShould result in %d, got %d", input, expected, result)
	}
}
