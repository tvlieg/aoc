package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	fmt.Println("Part 1:", part1(filePath))
	fmt.Println("Part 2:", part2(filePath))
}

func part1(fp string) int {
	file, _ := os.Open(fp)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const free = -1
	var blocks []int
	scanner.Scan()
	for i, r := range scanner.Text() {
		n := int(r - '0')

		if i%2 == 0 {
			id := i / 2
			for range n {
				blocks = append(blocks, id)
			}
			continue
		}

		for range n {
			blocks = append(blocks, free)
		}
	}

	var i, sum int
	for len(blocks) > 0 {
		if blocks[0] == free {
			// strip trailing free space
			for blocks[len(blocks)-1] == free {
				blocks = blocks[:len(blocks)-1]
			}
			blocks[0] = blocks[len(blocks)-1]
			blocks = blocks[:len(blocks)-1]
		}
		sum += blocks[0] * i
		i++
		blocks = blocks[1:]
	}
	return sum
}

func part2(fp string) int {
	f, _ := os.Open(fp)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var (
		files  []*file
		spaces []*space
		cursor int
	)
	for i, r := range scanner.Text() {
		n := int(r - '0')

		if i%2 == 0 {
			id := i / 2
			files = append(files, &file{id: id, pos: cursor, size: n})
		} else {
			spaces = append(spaces, &space{pos: cursor, size: n})
		}

		cursor += n
	}

	compacted := make([]*file, 0, len(files))
	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		for _, s := range spaces {
			if f.size <= s.size && s.pos < f.pos {
				f.pos = s.pos
				s.pos += f.size
				s.size -= f.size
				break
			}
		}
		compacted = append(compacted, f)
	}

	var sum int
	for _, f := range compacted {
		for i := range f.size {
			sum += f.id * (f.pos + i)
		}
	}
	return sum
}

type file struct {
	id, pos, size int
}

type space struct {
	pos, size int
}
