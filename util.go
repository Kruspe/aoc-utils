package aoc_utils

import (
	"os"
	"strings"
)

func ReadInput(pathToFile string) []string {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func GetSeparators(input []string) []int {
	var separators []int
	for i, l := range input {
		if l == "" {
			separators = append(separators, i)
		}
	}
	return separators
}

func Pad2dArray(
	input []string,
	padding string,
) []string {
	result := make([]string, 0)
	emptyRow := strings.Repeat(padding, len(input[0])+2)

	result = append(result, emptyRow)
	for _, row := range input {
		result = append(result, padding+row+padding)
	}
	result = append(result, emptyRow)

	return result
}
