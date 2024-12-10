package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x, y int
}

func main() {
	// filePath := "example_input_1"
	// filePath := "example_input_2"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for i, r := range line {
			row[i] = int(r - '0')
		}
		grid = append(grid, row)
	}

	var trails []map[coord]struct{}
	for y, row := range grid {
		for x, v := range row {
			if v == 0 {
				trails = append(trails, map[coord]struct{}{{x, y}: {}})
			}
		}
	}

	for i := 0; i < 9; i++ {
		for i, trail := range trails {
			trails[i] = move(trail, grid)
		}
	}

	var score int
	for _, trail := range trails {
		score += len(trail)
	}
	fmt.Println("Part1:", score)
}

func move(trail map[coord]struct{}, grid [][]int) map[coord]struct{} {
	height := len(grid)
	width := len(grid[0])

	newTrail := make(map[coord]struct{})
	for pos := range trail {
		next := grid[pos.y][pos.x] + 1
		if pos.y < height-1 && grid[pos.y+1][pos.x] == next {
			newTrail[coord{pos.x, pos.y + 1}] = struct{}{}
		}
		if pos.x < width-1 && grid[pos.y][pos.x+1] == next {
			newTrail[coord{pos.x + 1, pos.y}] = struct{}{}
		}
		if pos.y > 0 && grid[pos.y-1][pos.x] == next {
			newTrail[coord{pos.x, pos.y - 1}] = struct{}{}
		}
		if pos.x > 0 && grid[pos.y][pos.x-1] == next {
			newTrail[coord{pos.x - 1, pos.y}] = struct{}{}
		}
	}

	return newTrail
}
