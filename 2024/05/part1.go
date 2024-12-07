package main

import "fmt"

func part1() {
	beforeMap, updates, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}
	result := 0
	var correctUpdate bool
	var middlePage int
	var exists bool
	for i := 0; i < len(updates); i++ {
		correctUpdate = true
		for j := 0; j < len(updates[i])-1 && correctUpdate; j++ {
			for k := j + 1; k < len(updates[i]) && correctUpdate; k++ {
				_, exists = beforeMap[updates[i][j]][updates[i][k]]
				if exists {
					fmt.Printf("Incorrect update: %d, %d needs to be before %d\n", updates[i], updates[i][k], updates[i][j])
					fmt.Printf("%d before map: %v\n", updates[i][j], beforeMap[updates[i][j]])
					correctUpdate = false
				}
			}
		}
		if correctUpdate {
			middlePage = updates[i][int(len(updates[i])/2)]
			result = result + middlePage
			fmt.Printf("Correct update: %d, adding %d\n", updates[i], middlePage)
		}
	}
	fmt.Println("Result:", result)
}
