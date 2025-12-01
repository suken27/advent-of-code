package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadFileAsString takes a file name as input and returns its content as a string.
func ReadFileAsIntArray(fileName string) ([]int, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var intArray []int
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text[1:])
		if err != nil {
			return nil, err
		}
		if text[0] == 'L' {
			number = number * -1
		}
		// Append each line as a byte slice to the matrix
		intArray = append(intArray, number)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return intArray, nil
}

func main() {
	array, error := ReadFileAsIntArray("input.txt")
	if error != nil {
		println(error)
		return
	}
	i := 0
	previousPos := 50
	zeroCount := 0
	nextPos := 0
	zeroClickTimes := 0
	for i < len(array) {
		nextPos = previousPos + array[i]
		nextPos = nextPos % 100
		if nextPos < 0 {
			nextPos = nextPos + 100
		}
		if array[i] < 0 {
			zeroClickTimes = int((array[i]*-1 + nextPos) / 100)
			if previousPos == 0 {
				zeroClickTimes--
			}
			if nextPos == 0 {
				zeroClickTimes++
			}
		} else {
			zeroClickTimes = int((array[i] + previousPos) / 100)
		}
		zeroCount += zeroClickTimes
		fmt.Printf("[%d]-->[%d]-->[%d] [%d]\n", previousPos, array[i], nextPos, zeroClickTimes)
		previousPos = nextPos
		i++
	}
	println(zeroCount)
}
