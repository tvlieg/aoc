package main

import (
	"bufio"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func process(fileName string) (sum int, err error) {
	file, err := f.Open(fileName)
	if err != nil {
		fmt.Println("could not open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		g := newGame(scanner.Text())
		if g.isPossible() {
			sum += g.id
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w\n", err)
	}

	return sum, nil
}

func main() {
	// Part 1
	sum, err := process("input.txt")
	if err != nil {
		fmt.Println("Part 1 failed:", err)
	}
	fmt.Println(sum)
}
