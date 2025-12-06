package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("./example.txt")
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := dial{50, 100}
	var (
		count1 int
		count2 int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count2 += d.rotate(scanner.Text())
		if d.pos == 0 {
			count1++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("1:", count1)
	fmt.Println("2:", count2)
}

type dial struct {
	pos int
	max int
}

func (d *dial) rotate(s string) int {
	if len(s) == 0 {
		panic("empty string")
	}

	dir := s[0]
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		panic("strconv")
	}

	switch dir {
	case 'L':
		return d.rotateLeft(n)
	case 'R':
		return d.rotateRight(n)
	}

	return 0
}

func (d *dial) rotateLeft(n int) int {
	zeros := (n + (d.max - d.pos)) / d.max
	if d.pos == 0 {
		zeros--
	}
	d.pos = mod(d.pos-n, d.max)

	return zeros
}

func (d *dial) rotateRight(n int) int {
	zeros := (n + d.pos) / d.max
	d.pos = mod(d.pos+n, d.max)

	return zeros
}

func mod(a, b int) int {
	return (a%b + b) % b
}
