package main

import "fmt"

type state struct {
	pos coord
	dir direction
}

func newState(x, y int, r byte) *state {
	var d direction
	switch r {
	case '^':
		d = north
	case '>':
		d = east
	case 'v':
		d = south
	case '<':
		d = east
	}
	return &state{
		pos: newCoord(x, y),
		dir: d,
	}
}

func (s *state) String() string {
	return fmt.Sprintf("x: %d y: %d dir: %v", s.pos.x, s.pos.y, s.dir)
}

func (s *state) leavesMappedArea(g grid) bool {
	if s.pos.x == 0 && s.dir == east {
		return true
	}
	if s.pos.x == len(g)-1 && s.dir == west {
		return true
	}
	if s.pos.y == 0 && s.dir == north {
		return true
	}
	if s.pos.y == len(g)-1 && s.dir == south {
		return true
	}
	return false
}

func (s *state) move(g grid) {
	next := coord{s.pos.x, s.pos.y}
	switch s.dir {
	case north:
		next.y = s.pos.y - 1
	case south:
		next.y = s.pos.y + 1
	case east:
		next.x = s.pos.x + 1
	case west:
		next.x = s.pos.x - 1
	}
	if g[next.y][next.x] == obstacle {
		s.turn()
		return
	}
	s.pos = next
}

func (s *state) turn() {
	switch s.dir {
	case north:
		s.dir = east
	case south:
		s.dir = west
	case east:
		s.dir = south
	case west:
		s.dir = north
	}
}

type direction int

const (
	north direction = iota + 1
	east
	south
	west
)

func (d direction) String() string {
	switch d {
	case north:
		return "^"
	case east:
		return ">"
	case south:
		return "v"
	case west:
		return "<"
	default:
		return "?"
	}
}

type coord struct {
	x, y int
}

func newCoord(x, y int) coord {
	return coord{x, y}
}
