package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// filePath := "example_input_1"
	// filePath := "example_input_2"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	// Parse input
	var g grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for i, r := range line {
			row[i] = int(r - '0')
		}
		g = append(g, row)
	}

	var score1, score2 int
	for y, row := range g {
		for x, v := range row {
			if v != 0 {
				continue
			}
			score1 += part1(coord{x, y}, g)
			score2 += part2(coord{x, y}, g)
		}
	}
	fmt.Println("Part1:", score1)
	fmt.Println("Part2:", score2)
}

func part1(c coord, g grid) int {
	return len(uniq(g.recurse(0, []coord{c})))
}

func part2(c coord, g grid) int {
	return len(g.recurse(0, []coord{c}))
}

func uniq(cs []coord) []coord {
	m := make(map[coord]struct{})
	var res []coord
	for _, c := range cs {
		if _, ok := m[c]; ok {
			continue
		}
		m[c] = struct{}{}
		res = append(res, c)
	}
	return res
}

type coord struct {
	x, y int
}

type grid [][]int

func (g grid) get(c coord) (int, bool) {
	if c.x < 0 || c.x >= len(g[0]) || c.y < 0 || c.y >= len(g) {
		return 0, false
	}
	return g[c.y][c.x], true
}

func (g grid) recurse(h int, coords []coord) []coord {
	if h == 9 {
		return coords
	}
	var next []coord
	for _, c := range coords {
		next = append(next, g.move(h+1, c)...)
	}
	return g.recurse(h+1, next)
}

func (g grid) move(height int, c coord) []coord {
	nextCoords := make([]coord, 0, 4)
	for _, next := range []coord{
		up(c),
		down(c),
		left(c),
		right(c),
	} {
		if h, ok := g.get(next); ok && h == height {
			nextCoords = append(nextCoords, next)
		}
	}
	return nextCoords
}

func up(c coord) coord {
	return coord{c.x, c.y - 1}
}

func down(c coord) coord {
	return coord{c.x, c.y + 1}
}

func left(c coord) coord {
	return coord{c.x - 1, c.y}
}

func right(c coord) coord {
	return coord{c.x + 1, c.y}
}
