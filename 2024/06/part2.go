package main

import "fmt"

func isLoop(obstacleMatrix [][]bool, guard Guard, pathMatrix [][]map[string]bool) bool {
	var nextX int
	var nextY int
	pathCopy := make([][]map[string]bool, len(pathMatrix))
	for i := 0; i < len(pathMatrix); i++ {
		pathCopy[i] = make([]map[string]bool, len(pathMatrix[i]))
		for j := 0; j < len(pathMatrix[i]); j++ {
			pathCopy[i][j] = make(map[string]bool)
			for k, v := range pathMatrix[i][j] {
				pathCopy[i][j][k] = v
			}
		}
	}
	guardCopy := guard
	// Turns as if it would have an obstacle in front of it
	guardCopy.turn()
	var exists bool
	var walkedPath bool
	for {
		nextX, nextY = guardCopy.nextMove()
		if nextY < 0 || nextY >= len(obstacleMatrix) || nextX < 0 || nextX >= len(obstacleMatrix[nextY]) {
			return false
		}
		for obstacleMatrix[nextY][nextX] {
			guardCopy.turn()
			nextX, nextY = guardCopy.nextMove()
			walkedPath, exists = pathCopy[guardCopy.y][guardCopy.x][guardCopy.dir]
			if nextY < 0 || nextY >= len(obstacleMatrix) || nextX < 0 || nextX >= len(obstacleMatrix[nextY]) {
				return false
			} else if exists && walkedPath {
				return true
			}
		}
		walkedPath, exists = pathCopy[guardCopy.y][guardCopy.x][guardCopy.dir]
		if !exists {
			pathCopy[guardCopy.y][guardCopy.x][guardCopy.dir] = true
		} else if walkedPath {
			return true
		}
		guardCopy.move()
	}
}

func part2() {
	obstacleMatrix, guard, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	pathMatrix := make([][]map[string]bool, len(obstacleMatrix))
	for i := 0; i < len(obstacleMatrix); i++ {
		pathMatrix[i] = make([]map[string]bool, len(obstacleMatrix[i]))
		for j := 0; j < len(obstacleMatrix[i]); j++ {
			pathMatrix[i][j] = make(map[string]bool)
		}
	}
	newObstaclesMatrix := make([][]bool, len(obstacleMatrix))
	for i := 0; i < len(obstacleMatrix); i++ {
		newObstaclesMatrix[i] = make([]bool, len(obstacleMatrix[i]))
	}
	result := 0
	exited := false
	var nextX int
	var nextY int
	var exists bool
	for !exited {
		if !newObstaclesMatrix[guard.y][guard.x] && isLoop(obstacleMatrix, guard, pathMatrix) {
			result++
			newObstaclesMatrix[guard.y][guard.x] = true
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
		_, exists = pathMatrix[guard.y][guard.x][guard.dir]
		if !exists {
			pathMatrix[guard.y][guard.x][guard.dir] = true
		}
		guard.move()
	}

	fmt.Println(result)
}
