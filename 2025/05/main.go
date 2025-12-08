package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RemoveOverlap(range1 []int, range2 []int) []int {
	if (range1[0] <= range2[0] && range1[1] >= range2[0]) || (range1[0] <= range2[1] && range1[1] >= range2[1]) {
		//fmt.Printf("%d-%d overlaps with %d-%d\n", range1[0], range1[1], range2[0], range2[1])
		return []int{min(range1[0], range2[0]), max(range1[1], range2[1])}
	} else if (range2[0] <= range1[0] && range2[1] >= range1[0]) || (range2[0] <= range1[1] && range2[1] >= range1[1]) {
		//fmt.Printf("%d-%d overlaps with %d-%d\n", range1[0], range1[1], range2[0], range2[1])
		return []int{min(range1[0], range2[0]), max(range1[1], range2[1])}
	}
	return nil
}

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
	includedRanges := make([]bool, len(freshRanges))
	for i := 0; i < len(freshRanges); i++ {
		includedRanges[i] = false
	}
	for i := 0; i < len(freshRanges); i++ {
		if includedRanges[i] {
			continue
		}
		for j := 0; j < len(freshRanges); j++ {
			if i == j || includedRanges[j] {
				continue
			}
			resultingRange := RemoveOverlap(freshRanges[i], freshRanges[j])
			if resultingRange != nil {
				freshRanges[i][0] = resultingRange[0]
				freshRanges[i][1] = resultingRange[1]
				includedRanges[j] = true
				j = 0
			}
		}
	}
	totalFreshIds := 0
	for i := 0; i < len(freshRanges); i++ {
		if !includedRanges[i] {
			totalFreshIds += freshRanges[i][1] - freshRanges[i][0] + 1
		}
	}
	println(totalFreshIds)
}
