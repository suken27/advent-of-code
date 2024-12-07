package main

import "fmt"

func part1() {

	obstacleMatrix, guard, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultMatrix := make([][]bool, len(obstacleMatrix))
	for i := 0; i < len(obstacleMatrix); i++ {
		resultMatrix[i] = make([]bool, len(obstacleMatrix[i]))
	}
	result := 0
	exited := false
	var nextX int
	var nextY int
	for !exited {
		if !resultMatrix[guard.y][guard.x] {
			resultMatrix[guard.y][guard.x] = true
			result++
		}
		nextX, nextY = guard.nextMove()
		if nextY < 0 || nextY >= len(obstacleMatrix) || nextX < 0 || nextX >= len(obstacleMatrix[nextY]) {
			exited = true
		}
		for !exited && obstacleMatrix[nextY][nextX] {
			guard.turn()
			nextX, nextY = guard.nextMove()
			exited = nextY < 0 || nextY >= len(obstacleMatrix) || nextX < 0 || nextX >= len(obstacleMatrix[nextY])
		}
		guard.move()
	}

	fmt.Println(result)
}
