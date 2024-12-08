package main

import (
	"fmt"
	"io"
	"os"
)

type coord [2]int

func main() {
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	moves, _ := io.ReadAll(file)

	pos := coord{0, 0}
	seen := map[coord]struct{}{
		pos: {},
	}
	for _, m := range moves {
		pos = move(pos, m)
		seen[pos] = struct{}{}
	}
	fmt.Println("Part 1:", len(seen))

	santa := coord{0, 0}
	robo := coord{0, 0}
	seen = map[coord]struct{}{
		santa: {},
	}
	for i, m := range moves {
		if i%2 == 0 {
			santa = move(santa, m)
			seen[santa] = struct{}{}
		} else {
			robo = move(robo, m)
			seen[robo] = struct{}{}
		}
	}
	fmt.Println("Part 2:", len(seen))
}

func move(c coord, dir byte) coord {
	switch dir {
	case '^':
		c[1]--
	case 'v':
		c[1]++
	case '>':
		c[0]++
	case '<':
		c[0]--
	}
	return c
}
