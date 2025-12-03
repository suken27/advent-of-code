package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileAsStringArray(fileName string) ([]string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		stringArray = append(stringArray, text)
	}
	return stringArray, nil
}

func calculateHighestJoltage(battery string) (int, error) {
	i := 1
	maxFirstDigitIndex := 0
	maxFirstDigit := battery[maxFirstDigitIndex]
	for i < len(battery)-1 {
		if battery[i] > maxFirstDigit {
			maxFirstDigitIndex = i
			maxFirstDigit = battery[maxFirstDigitIndex]
		}
		i++
	}
	maxSecondDigit := battery[maxFirstDigitIndex+1]
	i = maxFirstDigitIndex + 1
	for i < len(battery) {
		if battery[i] > maxSecondDigit {
			maxSecondDigit = battery[i]
		}
		i++
	}
	result, err := strconv.Atoi(string(maxFirstDigit) + string(maxSecondDigit))
	if err != nil {
		return -1, err
	}
	return result, nil
}

func main() {
	batteryArray, error := ReadFileAsStringArray("input.txt")
	if error != nil {
		println(error)
		return
	}
	i := 0
	total := 0
	for i < len(batteryArray) {
		value, err := calculateHighestJoltage(batteryArray[i])
		if err != nil {
			println(err)
			return
		}
		total += value
		i++
	}
	println(total)
}
