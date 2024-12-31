package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"
)

//go:embed example_input
var example string

//go:embed input
var input string

func main() {
	fmt.Println("Part1 (example):", part1(example))
	fmt.Println("Part1:", part1(input))

	fmt.Println("Part2: (example)", part2(example))
	fmt.Println("Part2:", part2(input))
}

func part1(input string) int {
	neighbors := parse(input)

	lans := make(map[[3]string]struct{})
	for host := range neighbors {
		if host[0] != 't' {
			continue
		}
		for _, n := range neighbors[host] {
			for _, m := range neighbors[n] {
				if !slices.Contains(neighbors[m], host) {
					continue
				}
				lan := []string{host, n, m}
				slices.Sort(lan)
				lans[[3]string{lan[0], lan[1], lan[2]}] = struct{}{}
			}
		}
	}

	return len(lans)
}

func part2(input string) string {
	g := parse(input)

	remaining := make(map[string]struct{}, len(g))
	for node := range g {
		remaining[node] = struct{}{}
	}
	pwds := bronKerbosch(g, make(map[string]struct{}), remaining, make(map[string]struct{}))
	var longest string
	for _, pwd := range pwds {
		if len(pwd) <= len(longest) {
			continue
		}
		longest = pwd
	}
	return longest
}

func bronKerbosch(graph map[string][]string, potential, remaining, skip map[string]struct{}) (pwds []string) {
	if len(remaining) == 0 && len(skip) == 0 {
		pwd := make([]string, 0, len(potential))
		for host := range potential {
			pwd = append(pwd, host)
		}
		slices.Sort(pwd)
		return []string{strings.Join(pwd, ",")}
	}

	for node := range remaining {
		newPotential := make(map[string]struct{})
		maps.Copy(newPotential, potential)
		newPotential[node] = struct{}{}

		newRemaining := make(map[string]struct{})
		for n := range remaining {
			if slices.Contains(graph[node], n) {
				newRemaining[n] = struct{}{}
			}
		}

		newSkip := make(map[string]struct{})
		for n := range skip {
			if slices.Contains(graph[node], n) {
				newSkip[n] = struct{}{}
			}
		}

		pwds = append(pwds, bronKerbosch(graph, newPotential, newRemaining, newSkip)...)
		delete(remaining, node)
		delete(skip, node)
	}

	return
}

func parse(input string) map[string][]string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	g := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		h1, h2, _ := strings.Cut(line, "-")

		g[h1] = append(g[h1], h2)
		g[h2] = append(g[h2], h1)
	}

	return g
}
