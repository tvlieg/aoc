package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const precision = 1e-4

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	var (
		matrices               [][][]float64
		ax, ay, bx, by, px, py float64
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		prefix, toParse, ok := strings.Cut(line, ": ")
		if !ok {
			continue
		}

		parseButton := func(s string) (float64, float64) {
			parse := func(s string) float64 {
				_, fStr, _ := strings.Cut(s, "+")
				f, _ := strconv.ParseFloat(fStr, 64)
				return f
			}
			x, y, _ := strings.Cut(s, ", ")
			return parse(x), parse(y)
		}
		parsePrize := func(s string) (float64, float64) {
			parse := func(s string) float64 {
				_, fStr, _ := strings.Cut(s, "=")
				f, _ := strconv.ParseFloat(fStr, 64)
				return f
			}
			x, y, _ := strings.Cut(s, ", ")
			return parse(x), parse(y)
		}

		switch prefix {
		case "Button A":
			ax, ay = parseButton(toParse)
			continue
		case "Button B":
			bx, by = parseButton(toParse)
			continue
		case "Prize":
			px, py = parsePrize(toParse)
			matrices = append(matrices, [][]float64{{ax, bx, px}, {ay, by, py}})
			// matrices = append(matrices, [][]float64{{ax, bx, px + 10000000000000}, {ay, by, py + 10000000000000}})
		}

	}

	var count int
	for _, m := range matrices {
		result, ok := gaussJordan(m)
		if !ok {
			continue
		}
		a := result[0][2]
		b := result[1][2]

		if a < 0 || b < 0 {
			continue
		}

		if math.IsNaN(a) || math.IsNaN(b) {
			continue
		}

		if !isInteger(a) || !isInteger(b) {
			continue
		}

		count += 3*int(math.Round(a)) + int(math.Round(b))
	}
	fmt.Println("Part1:", count)
}

func gaussJordan(matrix [][]float64) ([][]float64, bool) {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {

		pivot := i
		for k := i + 1; k < rows; k++ {
			if math.Abs(matrix[k][i]) > math.Abs(matrix[pivot][i]) {
				pivot = k
			}
		}

		if i != pivot {
			matrix[i], matrix[pivot] = matrix[pivot], matrix[i]
		}

		if math.Abs(matrix[i][i]) < precision {
			// Return matrix unchanged to indicate no solution
			return nil, false
		}

		div := matrix[i][i]
		for j := 0; j < cols; j++ {
			matrix[i][j] /= div
		}

		for k := 0; k < rows; k++ {
			if k != i {
				factor := matrix[k][i]
				for j := 0; j < cols; j++ {
					matrix[k][j] -= factor * matrix[i][j]
				}
			}
		}
	}

	return matrix, true
}

func isInteger(value float64) bool {
	return math.Abs(value-math.Round(value)) < precision

}
