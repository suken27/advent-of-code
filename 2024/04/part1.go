package main

import (
	"fmt"
)

func searchRight(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for j < len(matrix[i]) && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		j++
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchDown(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i < len(matrix) && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i++
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchRigthDown(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i < len(matrix) && j < len(matrix[i]) && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i++
		j++
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchLeftDown(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i < len(matrix) && j >= 0 && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i++
		j--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchUp(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i >= 0 && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchLeft(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for j >= 0 && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		j--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchLeftUp(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i >= 0 && j >= 0 && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i--
		j--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchRightUp(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for i >= 0 && j < len(matrix[i]) && wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i--
		j++
		wordIndex++
	}
	return wordIndex == len(word)
}

func search(matrix [][]byte, i, j int, word string) int {
	total := 0
	if searchRight(matrix, i, j, word) {
		total++
	}
	if searchDown(matrix, i, j, word) {
		total++
	}
	if searchRigthDown(matrix, i, j, word) {
		total++
	}
	if searchLeftDown(matrix, i, j, word) {
		total++
	}
	if searchUp(matrix, i, j, word) {
		total++
	}
	if searchLeft(matrix, i, j, word) {
		total++
	}
	if searchLeftUp(matrix, i, j, word) {
		total++
	}
	if searchRightUp(matrix, i, j, word) {
		total++
	}
	return total
}

func part1() {
	byteMatrix, err := ReadFileAsByteMatrix("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	result := 0
	for i := 0; i < len(byteMatrix); i++ {
		for j := 0; j < len(byteMatrix[i]); j++ {
			if byteMatrix[i][j] == 'X' {
				result += search(byteMatrix, i, j, "XMAS")
			}
		}
	}
	fmt.Println(result)
}
