package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFileAsByteMatrix takes a file name as input and returns its content as a byte matrix (line by line).
func ReadFileAsByteMatrix(fileName string) ([][]byte, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var byteMatrix [][]byte
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line as a byte slice to the matrix
		byteMatrix = append(byteMatrix, []byte(scanner.Text()))
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return byteMatrix, nil
}
