package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	var (
		g grid
		s *state
	)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		g = append(g, make([]obstruction, len(line)))

		for j, b := range line {
			switch b {
			case '.':
				// fmt.Println(". at:", i, j)
				g[i][j] = none
			case '#':
				// fmt.Println("# at:", i, j)
				g[i][j] = obstacle
			case '>', '^', '<', 'v':
				s = newState(j, i, b)
			}
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanning input: %v", err)
		return
	}

	// seen is the set of coordinates that have been visited.
	seen := map[[2]int]struct{}{{s.pos.x, s.pos.y}: {}}
	for !s.leavesMappedArea(g) {
		s.move(g)
		seen[[2]int{s.pos.x, s.pos.y}] = struct{}{}
	}
	fmt.Println(len(seen))
}
