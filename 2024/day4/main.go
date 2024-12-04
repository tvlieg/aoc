package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]byte
	for scanner.Scan() {
		lines = append(lines, append([]byte(nil), scanner.Bytes()...))
	}

	count := count(lines)

	fmt.Println(count)
}

func count(lines [][]byte) int {
	xmas := []byte("XMAS")
	reverse := []byte("SAMX")
	lenXmas := len(xmas)

	numRows := len(lines)
	numCols := len(lines[0])

	var count int
	for i, line := range lines {
		for j := range line {
			hasSpaceRight := j+lenXmas <= numCols
			hasSpaceLeft := j >= lenXmas-1
			hasSpaceDown := i+lenXmas <= numRows

			// check horizontal right
			// check horizontal left (== right reverse)
			if hasSpaceRight {
				horizontal := []byte{lines[i][j], lines[i][j+1], lines[i][j+2], lines[i][j+3]}
				if bytes.Equal(horizontal, xmas) || bytes.Equal(horizontal, reverse) {
					count++
				}
			}

			// check vertical down
			// check vertical up (== down reverse)
			if hasSpaceDown {
				vertical := []byte{lines[i][j], lines[i+1][j], lines[i+2][j], lines[i+3][j]}
				if slices.Equal(vertical, xmas) || slices.Equal(vertical, reverse) {
					count++
				}
			}

			if hasSpaceRight && hasSpaceDown {
				diagonalDownRight := []byte{lines[i][j], lines[i+1][j+1], lines[i+2][j+2], lines[i+3][j+3]}
				if slices.Equal(diagonalDownRight, xmas) || slices.Equal(diagonalDownRight, reverse) {
					count++
				}
			}

			if hasSpaceLeft && hasSpaceDown {
				diagonalDownLeft := []byte{lines[i][j], lines[i+1][j-1], lines[i+2][j-2], lines[i+3][j-3]}
				if slices.Equal(diagonalDownLeft, xmas) || slices.Equal(diagonalDownLeft, reverse) {
					count++
				}
			}
		}
	}
	return count
}
