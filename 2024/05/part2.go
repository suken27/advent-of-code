package main

import "fmt"

func part2() {
	beforeMap, updates, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}
	result := 0
	var correctUpdate bool
	var middlePage int
	var exists bool
	var swap int
	for i := 0; i < len(updates); i++ {
		correctUpdate = true
		for j := 0; j < len(updates[i])-1; j++ {
			// This only works because there is only one incorrect page in the update (bad code)
			for k := j + 1; k < len(updates[i]); k++ {
				_, exists = beforeMap[updates[i][j]][updates[i][k]]
				if exists {
					swap = updates[i][j]
					updates[i][j] = updates[i][k]
					updates[i][k] = swap
					correctUpdate = false
				}
			}
		}
		if !correctUpdate {
			middlePage = updates[i][int(len(updates[i])/2)]
			result = result + middlePage
			fmt.Printf("Incorrect update (corrected): %d, adding %d\n", updates[i], middlePage)
		}
	}
	fmt.Println("Result:", result)
}
