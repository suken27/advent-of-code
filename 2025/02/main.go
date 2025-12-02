package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsStringArray(fileName string) ([]string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		text := scanner.Text()
		stringArray := strings.Split(text, ",")
		return stringArray, nil
	} else {
		return nil, fmt.Errorf("file: %s had no text in it", fileName)
	}
}

func SplitRange(stringRange string) ([]int, error) {
	stringArray := strings.Split(stringRange, "-")
	if len(stringArray) != 2 {
		return nil, fmt.Errorf("incorrect format for range %s", stringRange)
	}
	var intArray []int
	i := 0
	for i < len(stringArray) {
		intValue, err := strconv.Atoi(stringArray[i])
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int: %w", err)
		}
		intArray = append(intArray, intValue)
		i++
	}
	return intArray, nil
}

func GetInvalidIds(idRange []int) []int {
	startId := idRange[0]
	endId := idRange[1]
	i := startId
	var invalidIds []int
	for i <= endId {
		if !IsValid(i) {
			invalidIds = append(invalidIds, i)
		}
		i++
	}
	return invalidIds
}

func IsValid(id int) bool {
	stringId := strconv.Itoa(id)
	i := 0
	var repeatingPattern string
	for i < int(len(stringId)/2) {
		repeatingPattern = stringId[:i+1]
		if ((len(stringId) - len(repeatingPattern)) % len(repeatingPattern)) == 0 {
			j := i + 1
			canBeInvalid := true
			for j < len(stringId) && canBeInvalid {
				if repeatingPattern[j%len(repeatingPattern)] != stringId[j] {
					canBeInvalid = false
				}
				j++
			}
			if canBeInvalid {
				return false
			}
		}
		i++
	}
	return true
}

func main() {
	array, error := ReadFileAsStringArray("input.txt")
	if error != nil {
		println(error)
		return
	}
	i := 0
	j := 0
	result := 0
	for i < len(array) {
		stringRange := array[i]
		intRange, err := SplitRange(stringRange)
		if err != nil {
			println(err)
		}
		invalidIds := GetInvalidIds(intRange)
		j = 0
		for j < len(invalidIds) {
			result += invalidIds[j]
			j++
		}
		i++
	}
	println(result)
}
