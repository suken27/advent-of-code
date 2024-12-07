package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// ReadFileAsString takes a file name as input and returns its content as a string.
func ReadFileAsString(fileName string) (string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file's content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

func checkEnabled(c byte, enableStr string, enableIndex *int) bool {
	if c != enableStr[*enableIndex] {
		*enableIndex = 0
	} else {
		*enableIndex++
		if *enableIndex == len(enableStr) {
			*enableIndex = 0
			return true
		}
	}
	return false
}

func checkDisabled(c byte, disableStr string, disableIndex *int) bool {
	if c != disableStr[*disableIndex] {
		*disableIndex = 0
	} else {
		*disableIndex++
		if *disableIndex == len(disableStr) {
			*disableIndex = 0
			return true
		}
	}
	return false
}

func enableOrDisable(c byte, enableStr string, disableStr string, enableIndex *int, disableIndex *int, enabled *bool) {
	if checkEnabled(c, enableStr, enableIndex) {
		*enabled = true
		fmt.Println("enabled")
	}
	if checkDisabled(c, disableStr, disableIndex) {
		*enabled = false
		fmt.Println("disabled")
	}
}

func parseOperator(content string, i *int) (int, error) {
	operator := ""
	for ; *i < len(content) && content[*i] != ',' && content[*i] != ')'; *i++ {
		if _, err := strconv.Atoi(string(content[*i])); err != nil {
			return 0, fmt.Errorf("failed to convert operator character to integer: %w", err)
		} else {
			operator += string(content[*i])
		}
	}
	if operator == "" {
		return 0, fmt.Errorf("expected operator but found no digits")
	}
	operatorInt, err := strconv.Atoi(operator)
	if err != nil {
		return 0, fmt.Errorf("failed to convert operator to integer: %w", err)
	}
	fmt.Printf("operator: %d\n", operatorInt)
	return operatorInt, nil
}

func mul(content string, i *int) (int, error) {
	operator1, error := parseOperator(content, i)
	if error != nil {
		return 0, fmt.Errorf("failed to parse operator1: %w", error)
	}
	if content[*i] != ',' {
		return 0, fmt.Errorf("expected ',' but found %c", content[*i])
	}
	*i++
	operator2, error := parseOperator(content, i)
	if error != nil {
		return 0, fmt.Errorf("failed to parse operator2: %w", error)
	}
	if content[*i] != ')' {
		return 0, fmt.Errorf("expected ')' but found %c", content[*i])
	}
	fmt.Printf("%d * %d\n", operator1, operator2)
	return operator1 * operator2, nil
}

func main() {

	fileName := "input.txt"
	content, err := ReadFileAsString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	enabled := true
	enableStr := "do()"
	disableStr := "don't()"
	mulStr := "mul("
	enableIndex := 0
	disableIndex := 0
	mulIndex := 0
	total := 0
	for i := 0; i < len(content); i++ {
		currentChar := content[i]
		enableOrDisable(currentChar, enableStr, disableStr, &enableIndex, &disableIndex, &enabled)
		if enabled {
			if currentChar == mulStr[mulIndex] {
				mulIndex++
				if mulIndex == len(mulStr) {
					enableIndex = 0
					disableIndex = 0
					mulIndex = 0
					i++
					result, error := mul(content, &i)
					if error == nil {
						total += result
					} else {
						fmt.Println(error)
					}
				}
			} else {
				mulIndex = 0
			}
		}
	}

	fmt.Println(total)

}
