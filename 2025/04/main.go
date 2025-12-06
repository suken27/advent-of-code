package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileAsBooleanMatrix(fileName string) ([][]bool, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var booleanMatrix [][]bool
	var booleanArray []bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		booleanArray = []bool{}
		for i := 0; i < len(text); i++ {
			if text[i] == '@' {
				booleanArray = append(booleanArray, true)
			} else {
				booleanArray = append(booleanArray, false)
			}
		}
		booleanMatrix = append(booleanMatrix, booleanArray)
	}
	return booleanMatrix, nil
}

func IsAccesible(matrix [][]bool, x int, y int) (bool, error) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) {
		return false, fmt.Errorf("coordinates [%d, %d] are out of bounds of the matrix [0-%d, 0-%d]", x, y, len(matrix)-1, len(matrix[x])-1)
	}
	if !matrix[x][y] {
		return false, nil
	}

	surroundingRolls := -1
	for i := max(0, x-1); i <= min(len(matrix)-1, x+1); i++ {
		for j := max(0, y-1); j <= min(len(matrix[i])-1, y+1); j++ {
			if matrix[i][j] {
				surroundingRolls++
				if surroundingRolls >= 4 {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func RemoveAccesibles(matrix [][]bool) (int, [][]bool, error) {
	updatedMatrix := [][]bool{}
	totalAccesibleRolls := 0
	for i := 0; i < len(matrix); i++ {
		updatedRow := []bool{}
		for j := 0; j < len(matrix[i]); j++ {
			isAccessible, err := IsAccesible(matrix, i, j)
			if err != nil {
				return -1, nil, err
			}
			if isAccessible {
				updatedRow = append(updatedRow, false)
				totalAccesibleRolls++
			} else {
				updatedRow = append(updatedRow, matrix[i][j])
			}
		}
		updatedMatrix = append(updatedMatrix, updatedRow)
	}
	return totalAccesibleRolls, updatedMatrix, nil
}

func main() {
	rollsMatrix, error := ReadFileAsBooleanMatrix("input.txt")
	if error != nil {
		println(error)
		return
	}
	totalAccesibleRolls := 0
	accessedRolls := 1
	for accessedRolls > 0 {
		newAccessedRolls, updatedMatrix, err := RemoveAccesibles(rollsMatrix)
		accessedRolls = newAccessedRolls
		if err != nil {
			println(err)
			return
		}
		totalAccesibleRolls += accessedRolls
		rollsMatrix = updatedMatrix
	}
	println(totalAccesibleRolls)
}
