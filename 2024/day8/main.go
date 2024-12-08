package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord [2]int

func (c coord) x() int {
	return c[0]
}

func (c coord) y() int {
	return c[1]
}

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var height, width int

	y := 0
	antennas := make(map[byte][]coord)
	for scanner.Scan() {
		line := scanner.Bytes()
		if y == 0 {
			width = len(line)
		}

		for x, b := range line {
			if b == '.' {
				continue
			}
			antennas[b] = append(antennas[b], coord{x, y})
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanning input: %v", err)
		return
	}
	height = y

	var antinodes []coord
	for _, as := range antennas {
		antinodes = append(antinodes, getAntinodes(height, width, as)...)

	}

	uniq := make(map[coord]struct{})
	for _, coord := range antinodes {
		uniq[coord] = struct{}{}
	}
	fmt.Println(len(uniq))
}

func getAntinodes(height, width int, cs []coord) []coord {
	calc := func(a1, a2 coord) (coord, bool) {
		x := 2*a1.x() - a2.x()
		if x < 0 || x >= width {
			return coord{}, false
		}
		y := 2*a1.y() - a2.y()
		if y < 0 || y >= height {
			return coord{}, false
		}
		return coord{x, y}, true
	}

	var res []coord
	for i := 0; i < len(cs)-1; i++ {
		for j := i + 1; j < len(cs); j++ {
			if an, ok := calc(cs[i], cs[j]); ok {
				res = append(res, an)
			}
			if an, ok := calc(cs[j], cs[i]); ok {
				res = append(res, an)
			}
		}
	}
	return res
}
