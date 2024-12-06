package main

import "strings"

const (
	none object = iota
	obstacle
)

type object int

func (o object) String() string {
	switch o {
	case none:
		return "."
	case obstacle:
		return "#"
	default:
		return "?"
	}
}

type grid [][]object

func (g grid) String() string {
	var b strings.Builder
	for _, line := range g {
		for _, v := range line {
			b.WriteString(v.String())
		}
		b.WriteString("\n")
	}
	return b.String()
}
