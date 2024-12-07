package main

import (
	"bufio"
	"fmt"
	"os"
)

type Guard struct {
	x   int
	y   int
	dir string
}

func (guard *Guard) turn() {
	guard.dir = guard.nextTurn()
}

func (guard Guard) nextTurn() string {
	switch guard.dir {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	panic("Invalid direction")
}

func (guard *Guard) move() {
	guard.x, guard.y = guard.nextMove()
	fmt.Printf("Guard at (%d, %d) facing %s\n", guard.x, guard.y, guard.dir)
}

func (guard Guard) nextMove() (int, int) {
	switch guard.dir {
	case "up":
		return guard.x, guard.y - 1
	case "right":
		return guard.x + 1, guard.y
	case "down":
		return guard.x, guard.y + 1
	case "left":
		return guard.x - 1, guard.y
	}
	panic("Invalid direction")
}

func textLineToMatrixRow(line string, guard *Guard, lineNumber int) []bool {
	row := make([]bool, len(line))
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '#':
			row[i] = true
		case '.':
			row[i] = false
		case '^':
			guard.x = i
			guard.y = lineNumber
			guard.dir = "up"
		}
	}
	return row
}

// ReadFileAsByteMatrix takes a file name as input and returns its content as a byte matrix (line by line).
func ReadInput(fileName string) ([][]bool, Guard, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, Guard{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var obstacleMatrix [][]bool
	var row []bool
	var guard Guard
	i := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		row = textLineToMatrixRow(scanner.Text(), &guard, i)
		obstacleMatrix = append(obstacleMatrix, row)
		i++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, Guard{}, fmt.Errorf("failed to read file: %w", err)
	}

	return obstacleMatrix, guard, nil
}
