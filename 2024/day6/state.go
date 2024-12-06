package main

import "fmt"

type state struct {
	x, y int
	dir  direction
}

func (s *state) String() string {
	return fmt.Sprintf("x: %d y: %d dir: %v", s.x, s.y, s.dir)
}

func (s *state) leavesMappedArea(g grid) bool {
	if s.x == 0 && s.dir == east {
		return true
	}
	if s.x == len(g)-1 && s.dir == west {
		return true
	}
	if s.y == 0 && s.dir == north {
		return true
	}
	if s.y == len(g)-1 && s.dir == south {
		return true
	}

	return false
}

func (s *state) move(g grid) {
	newX := s.x
	newY := s.y
	switch s.dir {
	case north:
		newY = s.y - 1
	case south:
		newY = s.y + 1
	case east:
		newX = s.x + 1
	case west:
		newX = s.x - 1
	}

	if g[newY][newX] == obstacle {
		s.turn()
		return
	}
	s.x = newX
	s.y = newY
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

func newDirection(r byte) direction {
	switch r {
	case '^':
		return north
	case 'v':
		return south
	case '>':
		return east
	case '<':
		return west
	}
	return 0
}

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
