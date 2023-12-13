package aoc_utils

import (
	"fmt"
	"os"
	"slices"
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

func Pad2dStringArray(
	input [][]string,
	padding string) [][]string {
	result := make([][]string, 0)
	emptyRow := strings.Repeat(padding, len(input[0])+2)

	result = append(result, strings.Split(emptyRow, ""))
	for _, row := range input {
		r := make([]string, 0)
		r = append(r, padding)
		r = append(r, row...)
		r = append(r, padding)
		result = append(result, r)
	}
	result = append(result, strings.Split(emptyRow, ""))

	return result
}

func PadStringArray(
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

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(values []int) (int, error) {
	if len(values) < 2 {
		return 0, fmt.Errorf("values must contain at least 2 elements")
	}
	result := values[0] * values[1] / GCD(values[0], values[1])
	for i := 0; i < len(values)-2; i++ {
		result, _ = LCM([]int{result, values[i+2]})
	}
	return result, nil
}

func FloodFill(m [][]string, border []string, paint string, x, y int) {
	if m[y][x] != paint && !slices.Contains(border, m[y][x]) {
		m[y][x] = paint
		if x+1 != len(m[0]) {
			FloodFill(m, border, paint, x+1, y)
		}
		if x-1 != -1 {
			FloodFill(m, border, paint, x-1, y)
		}
		if y+1 != len(m) {
			FloodFill(m, border, paint, x, y+1)
		}
		if y-1 != -1 {
			FloodFill(m, border, paint, x, y-1)
		}
	}
}
