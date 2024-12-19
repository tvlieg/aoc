package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("Part 1 (example):", part1(example, 7, 12))
	fmt.Println("Part 1 :", part1(input, 71, 1024))
	fmt.Println("Part 2 (example):", part2(example, 7, 12))
	fmt.Println("Part 2:", part2(input, 71, 1024))
}

func part1(input string, dim, byteCount int) int {
	bytes := parse(input)

	g := newGrid(dim, dim, bytes)
	for range byteCount {
		g.fall()
	}

	return bfc(g)
}

func part2(input string, dim, byteCount int) string {
	bytes := parse(input)

	g := newGrid(dim, dim, bytes)
	for range byteCount {
		g.fall()
	}

	var last coord
	for bfc(g) > 0 {
		last = g.fall()
	}

	return fmt.Sprintf("%d,%d", last.x, last.y)
}

func parse(input string) []coord {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var coords []coord
	for scanner.Scan() {
		line := scanner.Text()
		xs, ys, _ := strings.Cut(line, ",")
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		coords = append(coords, coord{x, y})
	}

	return coords
}

func bfc(g *grid) int {
	type node struct {
		coord
		parent *node
	}
	start := coord{0, 0}
	goal := coord{g.cols - 1, g.rows - 1}

	queue := []node{{coord: start}}
	seen := map[coord]struct{}{start: {}}

	var end node
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v.coord == goal {
			end = v
			break
		}
		for _, n := range g.neighbors(v.coord) {
			if _, ok := seen[n]; ok {
				continue
			}
			seen[n] = struct{}{}
			queue = append(queue, node{coord: n, parent: &v})
		}
	}

	n := end
	var count int
	for n.parent != nil {
		// fmt.Println(n.coord)
		n = *n.parent
		count++
	}
	return count
}

type coord struct {
	x, y int
}

type grid struct {
	data  []byte
	cols  int
	rows  int
	bytes []coord
}

func newGrid(cols, rows int, bytes []coord) *grid {
	g := grid{
		data:  make([]byte, cols*rows),
		cols:  cols,
		rows:  rows,
		bytes: bytes,
	}
	for i := range cols * rows {
		g.data[i] = '.'
	}
	return &g
}

func (g *grid) neighbors(c coord) []coord {
	dirs := []struct{ dx, dy int }{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var res []coord
	for _, d := range dirs {
		nx, ny := c.x+d.dx, c.y+d.dy
		if nx < 0 || nx >= g.cols || ny < 0 || ny >= g.rows {
			continue
		}
		if g.get(nx, ny) == '#' {
			continue
		}
		res = append(res, coord{x: nx, y: ny})
	}
	return res
}

func (g *grid) get(x, y int) byte {
	return g.data[y*g.rows+x]
}

func (g *grid) fall() coord {
	if len(g.bytes) == 0 {
		panic("no more bytes")
	}
	b := g.bytes[0]
	g.bytes = g.bytes[1:]

	g.data[b.y*g.rows+b.x] = '#'

	return b
}

func (g *grid) String() string {
	var b strings.Builder
	for i := 0; i < g.rows; i++ {
		b.WriteString(string(g.data[i*g.rows : (i+1)*g.rows]))
		if i < g.rows-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
