package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsStringArray(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, scanner.Text())
	}

	return stringArray, nil
}

func GetNumberInColumn(stringArray []string, index int) (int, error) {
	var stringBuilder strings.Builder
	for i := 0; i < len(stringArray)-1; i++ {
		if stringArray[i][index] != ' ' {
			stringBuilder.WriteByte(stringArray[i][index])
		}
	}
	intValue, err := strconv.Atoi(stringBuilder.String())
	if err != nil {
		return -1, err
	}
	return intValue, nil
}

func AddArray(array []int) int {
	total := 0
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	return total
}

func MultiplyArray(array []int) int {
	total := array[0]
	for i := 1; i < len(array); i++ {
		total *= array[i]
	}
	return total
}

func main() {
	stringArray, err := ReadFileAsStringArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	subTotal := 0
	for i := 0; i < len(stringArray[0]); {
		var numberArray []int
		number, err := GetNumberInColumn(stringArray, i)
		if err != nil {
			log.Fatal(err)
		}
		numberArray = append(numberArray, number)
		j := i + 1
		for ; j < len(stringArray[0]) && (j+1 >= len(stringArray[0]) || stringArray[len(stringArray)-1][j+1] == ' '); j++ {
			number, err = GetNumberInColumn(stringArray, j)
			if err != nil {
				log.Fatal(err)
			}
			numberArray = append(numberArray, number)
		}
		if stringArray[len(stringArray)-1][i] == '+' {
			subTotal = AddArray(numberArray)
		} else if stringArray[len(stringArray)-1][i] == '*' {
			subTotal = MultiplyArray(numberArray)
		} else {
			log.Fatalf("unexpected operator %c", stringArray[len(stringArray)-1][i])
		}
		total += subTotal
		i += j - i + 1
	}
	println(total)
}
