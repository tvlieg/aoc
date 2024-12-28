package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("Part 1 (example):", part1(example))
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2 (example):", part2(example))
	fmt.Println("Part 2:", part2(input))
}

func part1(input string) int {
	r := newRaceTrack(input)

	// distance to end
	dte := r.distanceToEndMap()
	var sum int
	for c := range dte {
		sum += r.cheat(c, 2, 100, dte)
	}
	return sum
}

func part2(input string) int {
	r := newRaceTrack(input)

	// distance to end
	dte := r.distanceToEndMap()
	var sum int
	for c := range dte {
		for i := 2; i <= 20; i++ {
			sum += r.cheat(c, i, 100, dte)
		}
	}
	return sum
}

func (r *racetrack) distanceToEndMap() map[coord]int {
	queue := []coord{r.end}
	seen := map[coord]int{r.end: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, n := range r.neighbors(current) {
			if _, ok := seen[n]; ok {
				continue
			}
			seen[n] = seen[current] + 1
			queue = append(queue, n)
		}
	}

	return seen
}

const (
	track = 0
	wall  = -1
)

type coord struct {
	x, y int
}

type racetrack struct {
	data  []int
	cols  int
	rows  int
	start coord
	end   coord
}

func newRaceTrack(input string) *racetrack {
	r := racetrack{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Bytes()
		for x, b := range line {
			var v int
			switch b {
			case 'S':
				r.start = coord{x, y}
				v = track
			case 'E':
				r.end = coord{x, y}
				v = track
			case '.':
				v = track
			case '#':
				v = wall
			default:
				panic("unknown object")
			}
			r.data = append(r.data, v)
		}
		r.cols = len(line)
		r.rows++
	}
	return &r
}

func (r *racetrack) neighbors(c coord) []coord {
	dirs := []struct{ dx, dy int }{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var res []coord
	for _, d := range dirs {
		nx, ny := c.x+d.dx, c.y+d.dy
		if nx < 0 || nx >= r.cols || ny < 0 || ny >= r.rows {
			continue
		}
		if r.get(coord{nx, ny}) == wall {
			continue
		}
		res = append(res, coord{x: nx, y: ny})
	}
	return res
}

// cheat finds the best cheat for a given coordinate and returns the ???.
func (r *racetrack) cheat(c coord, ps, thres int, dte map[coord]int) int {
	dirs := func(r int) (dirs []struct{ dx, dy int }) {
		for i := 0; i < r; i++ {
			dirs = append(dirs, []struct{ dx, dy int }{
				{i, -r + i},
				{r - i, i},
				{-i, r - i},
				{-r + i, -i},
			}...)
		}
		return
	}

	var count int
	for _, d := range dirs(ps) {
		nc := coord{c.x + d.dx, c.y + d.dy}

		// off grid
		if nc.x < 0 || nc.x >= r.cols || nc.y < 0 || nc.y >= r.rows {
			continue
		}
		if r.get(nc) == wall {
			continue
		}

		saved := dte[c] - dte[nc] - ps
		if saved >= thres {
			count++
		}
	}
	return count
}

func (r *racetrack) get(c coord) int {
	return r.data[c.y*r.rows+c.x]
}

func (r *racetrack) String() string {
	var b strings.Builder
	for y := range r.rows {
		for x := range r.cols {
			if x == r.start.x && y == r.start.y {
				b.WriteByte('S')
				continue
			}
			if x == r.end.x && y == r.end.y {
				b.WriteByte('E')
				continue
			}
			switch r.data[y*r.rows+x] {
			case wall:
				b.WriteByte('#')
			case track:
				b.WriteByte('.')
			}
		}
		if y < r.rows-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
