package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var freshRanges [][]int
	readingRanges := true
	for readingRanges && scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			splittedRange := strings.Split(text, "-")
			startValue, startValueErr := strconv.Atoi(splittedRange[0])
			if startValueErr != nil {
				println(startValueErr)
				return
			}
			endValue, endValueErr := strconv.Atoi(splittedRange[1])
			if endValueErr != nil {
				println(endValueErr)
				return
			}
			freshRanges = append(freshRanges, []int{startValue, endValue})
		} else {
			readingRanges = false
		}
	}
	totalFreshItems := 0
	for scanner.Scan() {
		text := scanner.Text()
		foodId, foodIdErr := strconv.Atoi(text)
		if foodIdErr != nil {
			println(foodIdErr)
			return
		}
		itemFound := false
		for i := 0; i < len(freshRanges) && !itemFound; i++ {
			if foodId >= freshRanges[i][0] && foodId <= freshRanges[i][1] {
				itemFound = true
				totalFreshItems++
			}
		}
	}
	println(totalFreshItems)
}
