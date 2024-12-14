package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed example_input
	example string

	//go:embed input
	input string
)

type pos struct {
	x, y, vx, vy int
}

func main() {
	fmt.Println("part 1 (example):", part1(example, 11, 7))
	fmt.Println("part 1:", part1(input, 101, 103))
	// fmt.Println("part 2:", part2(input, 101, 103))
}

func part1(input string, width, height int) int {
	const pattern = `^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`
	re := regexp.MustCompile(pattern)

	const seconds = 100

	var tl, tr, bl, br int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		if len(m) != 5 {
			break
		}
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		vx, _ := strconv.Atoi(m[3])
		vy, _ := strconv.Atoi(m[4])

		x = mod(x+seconds*vx, width)
		y = mod(y+seconds*vy, height)

		// position on middle axes
		if x == width/2 || y == height/2 {
			continue
		}

		// fmt.Printf("after 100 seconds: %d,%d\n", px, py)
		isTop := y/(height/2) == 0
		isLeft := x/(width/2) == 0
		switch {
		case isTop && isLeft:
			tl++
		case isTop && !isLeft:
			tr++
		case !isTop && isLeft:
			bl++
		case !isTop && !isLeft:
			br++
		}
	}

	return tl * tr * bl * br
}

func part2(input string, width, height int) int {
	return 0
}

func mod(a, m int) int {
	return (a%m + m) % m
}
