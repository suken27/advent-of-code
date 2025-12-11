package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileAsStringMatrix(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var stringMatrix [][]string
	scanner := bufio.NewScanner(file)
	regexExpresion := regexp.MustCompile(` +`)
	for scanner.Scan() {
		line := scanner.Text()
		line = regexExpresion.ReplaceAllString(line, " ")
		slices := strings.Split(line, " ")
		stringMatrix = append(stringMatrix, slices)
	}

	return stringMatrix, nil
}

func TransposeMatrix(matrix [][]string) [][]string {
	transposed := make([][]string, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposed[j] = append(transposed[j], matrix[i][j])
		}
	}
	return transposed
}

func CalculateColumn(column []string) (int, error) {
	intValue, err := strconv.Atoi(column[0])
	if err != nil {
		return -1, err
	}
	total := intValue
	columnLength := len(column)
	switch column[columnLength-1] {
	case "+":
		for i := 1; i < columnLength-1; i++ {
			intValue, err = strconv.Atoi(column[i])
			if err != nil {
				return -1, err
			}
			total += intValue
		}
	case "*":
		for i := 1; i < columnLength-1; i++ {
			intValue, err = strconv.Atoi(column[i])
			if err != nil {
				return -1, err
			}
			total *= intValue
		}
	default:
		return -1, fmt.Errorf("operand %s is not supported", column[columnLength-1])
	}
	return total, nil
}

func main() {
	stringMatrix, err := ReadFileAsStringMatrix("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stringMatrix = TransposeMatrix(stringMatrix)
	grandTotal := 0
	for i := 0; i < len(stringMatrix); i++ {
		subTotal, calcErr := CalculateColumn(stringMatrix[i])
		if calcErr != nil {
			log.Fatal(calcErr)
		}
		grandTotal += subTotal
	}
	println(grandTotal)
}
