package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseSpaceSeparatedNumbers(input string) ([]int, error) {
	// Split the input string by spaces
	parts := strings.Split(input, ",")

	// Create an integer array to store the parsed numbers
	nums := make([]int, len(parts))

	// Parse each part into an integer
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse '%s' as an integer: %w", part, err)
		}
		nums[i] = num
	}

	return nums, nil
}

// ReadFileAsByteMatrix takes a file name as input and returns its content as a byte matrix (line by line).
func ReadInput(fileName string) (map[int]map[int]bool, [][]int, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	beforeMap := make(map[int]map[int]bool)
	var splitted []string
	var n1 int
	var n2 int
	var ok bool
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		splitted = strings.Split(scanner.Text(), "|")
		n1, err = strconv.Atoi(splitted[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert string to int: %w", err)
		}
		n2, err = strconv.Atoi(splitted[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert string to int: %w", err)
		}
		if _, ok = beforeMap[n2]; !ok {
			beforeMap[n2] = make(map[int]bool)
		}
		beforeMap[n2][n1] = true
	}

	var updatesMatrix [][]int
	var numbers []int
	for scanner.Scan() {
		numbers, err = ParseSpaceSeparatedNumbers(scanner.Text())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse numbers: %w", err)
		}
		updatesMatrix = append(updatesMatrix, numbers)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	return beforeMap, updatesMatrix, nil
}
