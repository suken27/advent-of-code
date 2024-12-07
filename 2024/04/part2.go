package main

import "fmt"

func searchXRightDown(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i++
		j++
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchXLeftDown(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i++
		j--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchXRightUp(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i--
		j++
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchXLeftUp(matrix [][]byte, i, j int, word string) bool {
	wordIndex := 0
	for wordIndex < len(word) {
		if matrix[i][j] != word[wordIndex] {
			return false
		}
		i--
		j--
		wordIndex++
	}
	return wordIndex == len(word)
}

func searchX(matrix [][]byte, i, j int, word string) bool {
	result := 0
	if searchXRightDown(matrix, i-1, j-1, word) {
		result++
	}
	if searchXLeftDown(matrix, i-1, j+1, word) {
		result++
	}
	if result < 2 && searchXRightUp(matrix, i+1, j-1, word) {
		result++
	}
	if result < 2 && searchXLeftUp(matrix, i+1, j+1, word) {
		result++
	}
	return result == 2
}

func part2() {
	matrix, err := ReadFileAsByteMatrix("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'A' && i > 0 && i < len(matrix)-1 && j > 0 && j < len(matrix[i])-1 {
				if searchX(matrix, i, j, "MAS") {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}
