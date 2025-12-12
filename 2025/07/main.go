package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileAsBooleanMatrix(fileName string) ([][]bool, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var booleanMatrix [][]bool
	startingPoint := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var boolArray []bool
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				startingPoint = i
			}
			boolArray = append(boolArray, line[i] == '^')
		}
		booleanMatrix = append(booleanMatrix, boolArray)
	}

	return booleanMatrix, startingPoint, nil
}

func CalculateSplits(matrix [][]bool, beamMatrix [][]int, startingRow int, startingColumn int) int {
	if startingRow >= len(matrix)-1 || startingColumn < 0 || startingColumn > len(matrix[startingRow])-1 {
		return 0
	}
	if !matrix[startingRow+1][startingColumn] {
		if beamMatrix[startingRow+1][startingColumn] != -1 {
			return beamMatrix[startingRow+1][startingColumn]
		}
		beamMatrix[startingRow+1][startingColumn] = CalculateSplits(matrix, beamMatrix, startingRow+1, startingColumn)
		return beamMatrix[startingRow+1][startingColumn]
	}
	if beamMatrix[startingRow+1][startingColumn-1] == -1 {
		beamMatrix[startingRow+1][startingColumn-1] = CalculateSplits(matrix, beamMatrix, startingRow+1, startingColumn-1)
	}
	if beamMatrix[startingRow+1][startingColumn+1] == -1 {
		beamMatrix[startingRow+1][startingColumn+1] = CalculateSplits(matrix, beamMatrix, startingRow+1, startingColumn+1)
	}
	return beamMatrix[startingRow+1][startingColumn-1] + beamMatrix[startingRow+1][startingColumn+1] + 1
}

func main() {
	booleanMatrix, startingPoint, err := ReadFileAsBooleanMatrix("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	beamMatrix := make([][]int, len(booleanMatrix))
	for i := 0; i < len(booleanMatrix); i++ {
		beamMatrix[i] = make([]int, len(booleanMatrix[i]))
		for j := 0; j < len(booleanMatrix[i]); j++ {
			beamMatrix[i][j] = -1
		}
	}
	println(CalculateSplits(booleanMatrix, beamMatrix, 0, startingPoint) + 1)
}
