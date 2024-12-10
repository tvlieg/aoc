package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	const free = -1

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
	fmt.Println("Part1:", sum)
}
