package main

import (
	"regexp"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

type game struct {
	id   int
	sets []set
}

func newGame(str string) game {
	id, setStrs := parseGameString(str)

	sets := make([]set, 0)
	for _, s := range setStrs {
		sets = append(sets, newSet(s))
	}

	return game{
		id:   id,
		sets: sets,
	}
}

func (g game) isPossible() bool {
	for _, s := range g.sets {
		if !s.isPossible() {
			return false
		}
	}
	return true
}

func parseGameString(str string) (id int, sets []string) {
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	m := re.FindStringSubmatch(str)

	id, _ = strconv.Atoi(m[1])
	sets = strings.Split(m[2], "; ")

	return id, sets
}

type set struct {
	r, g, b int
}

func newSet(str string) set {
	s := set{}
	for _, sample := range strings.Split(str, ", ") {
		n, color := parseSample(sample)
		switch color {
		case "red":
			s.r = n
		case "green":
			s.g = n
		case "blue":
			s.b = n
		}
	}
	return s
}

func (s set) isPossible() bool {
	return s.r <= maxRed && s.g <= maxGreen && s.b <= maxBlue
}

func parseSample(s string) (n int, color string) {
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)
	m := re.FindStringSubmatch(s)
	n, _ = strconv.Atoi(m[1])
	color = m[2]

	return n, color
}
