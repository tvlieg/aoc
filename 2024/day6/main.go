package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "example_input"
	// filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	var (
		g    grid
		x, y int
		dir  direction
	)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		g = append(g, make([]object, len(line)))

		for j, b := range line {
			switch b {
			case '.':
				// fmt.Println(". at:", i, j)
				g[i][j] = none
			case '#':
				// fmt.Println("# at:", i, j)
				g[i][j] = obstacle
			case '>', '^', '<', 'v':
				x = j
				y = i
				dir = newDirection(b)
			}
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanning input: %v", err)
		return
	}

	// seen is the set of coordinates that have been visited.
	s := state{x, y, dir}
	coordsSeen := map[[2]int]struct{}{{s.x, s.y}: {}}
	for !s.leavesMappedArea(g) {
		s.move(g)
		coordsSeen[[2]int{s.x, s.y}] = struct{}{}
	}
	fmt.Println(len(coordsSeen))

	// s = state{x, y, dir}
	// statesSeen := map[[3]int]struct{}{{s.x, s.y, int(s.dir)}: {}}
}
