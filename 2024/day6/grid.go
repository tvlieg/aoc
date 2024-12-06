package main

import "strings"

const (
	none obstruction = iota
	obstacle
)

type obstruction int

func (o obstruction) String() string {
	switch o {
	case none:
		return "."
	case obstacle:
		return "#"
	default:
		return "?"
	}
}

type grid [][]obstruction

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
