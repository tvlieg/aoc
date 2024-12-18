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
		lines = append(lines, append([]byte{}, scanner.Bytes()...))
	}

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func part1(lines [][]byte) int {
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

func part2(lines [][]byte) int {
	var count int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			if lines[i][j] != 'A' {
				continue
			}

			topLeft := lines[i-1][j-1]
			topRight := lines[i-1][j+1]
			bottomLeft := lines[i+1][j-1]
			bottomRight := lines[i+1][j+1]

			one := (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')
			two := (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')

			if one && two {
				count++
			}
		}
	}
	return count
}
