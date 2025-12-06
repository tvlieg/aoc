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

	d := dial{50}
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.rotate(scanner.Text())
		if d.pos == 0 {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}

type dial struct {
	pos int
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
	d.pos = mod(d.pos-n, 100)
	return d.pos
}

func (d *dial) rotateRight(n int) int {
	d.pos = mod(d.pos+n, 100)
	return d.pos
}

func mod(a, b int) int {
	return (a%b + b) % b
}
