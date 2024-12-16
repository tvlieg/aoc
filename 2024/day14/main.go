package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("part 1 (example):", part1(parse(example), 11, 7))
	fmt.Println("part 1:", part1(parse(input), 101, 103))
	fmt.Println("part 2:", part2(parse(input), 101, 103))
}

func parse(input string) []robot {
	const pattern = `^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`
	re := regexp.MustCompile(pattern)

	var robots []robot
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

		robots = append(robots, robot{x, y, vx, vy})
	}
	return robots
}

func part1(robots []robot, w, h int) int {
	xAxis := w / 2
	yAxis := h / 2

	var tl, tr, bl, br int
	for _, r := range robots {
		r = r.moveN(100, w, h)

		if r.x == xAxis || r.y == yAxis {
			continue
		}

		top := r.y < yAxis
		left := r.x < xAxis
		switch {
		case top && left:
			tl++
		case top && !left:
			tr++
		case !top && left:
			bl++
		case !top && !left:
			br++
		}
	}

	return tl * tr * bl * br
}

func part2(robots []robot, w, h int) int {
	for i := 1; ; i++ {
		for j, r := range robots {
			robots[j] = r.move(w, h)
		}
		if checkTree(robots, w) {
			return i
		}
	}
}

func mod(a, m int) int {
	return (a%m + m) % m
}

type robot struct {
	x, y, vx, vy int
}

func (r robot) moveN(count, w, h int) robot {
	r.x = mod(r.x+r.vx*count, w)
	r.y = mod(r.y+r.vy*count, h)
	return r
}

func (r robot) move(w, h int) robot {
	return r.moveN(1, w, h)
}

func checkTree(robots []robot, w int) bool {
	xAxis := w / 2

	m := make(map[[2]int]struct{}, len(robots))
	for _, r := range robots {
		m[[2]int{r.x, r.y}] = struct{}{}
	}

	// check for symmetry
	var score int
	for coord := range m {
		if coord[0] <= xAxis {
			mirror := [2]int{w - coord[0], coord[1]}
			if _, ok := m[mirror]; !ok {
				continue
			}
			score++

		}
	}
	// Try increasing scores
	if score > 50 {
		fmt.Println("score:", score)
		return true
	}

	return false
}
