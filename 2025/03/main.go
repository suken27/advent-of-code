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

func calculateHighestJoltage(battery string, batteriesToEnable int) string {

	i := 0
	maxFirstDigitIndex := 0
	maxFirstDigit := battery[maxFirstDigitIndex]
	for i < len(battery)-batteriesToEnable+1 {
		if battery[i] > maxFirstDigit {
			maxFirstDigitIndex = i
			maxFirstDigit = battery[maxFirstDigitIndex]
		}
		i++
	}
	if batteriesToEnable == 1 {
		return string(maxFirstDigit)
	} else {
		value := calculateHighestJoltage(battery[maxFirstDigitIndex+1:], batteriesToEnable-1)
		return string(maxFirstDigit) + value
	}
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
		value := calculateHighestJoltage(batteryArray[i], 12)
		intValue, err := strconv.Atoi(value)
		if err != nil {
			println(err)
			return
		}
		println(intValue)
		total += intValue
		i++
	}
	println(total)
}
