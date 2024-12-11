package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	const (
		example = "125 17"
		input   = "554735 45401 8434 0 188 7487525 77 7"
	)
	fmt.Println("Example:", run(parse(example), 6))
	fmt.Println("Part 1:", run(parse(input), 25))
	fmt.Println("Part 2:", run(parse(input), 75))
}

func parse(s string) map[int]int {
	pebbles := make(map[int]int)
	for _, s := range strings.Split(s, " ") {
		n, _ := strconv.Atoi(s)
		pebbles[n] = 1 // Asume no duplicates in input
	}
	return pebbles
}

func run(pebbles map[int]int, n int) (count int) {
	for range n {
		pebbles = blink(pebbles)
	}

	for _, c := range pebbles {
		count += c
	}

	return
}

func blink(pebbles map[int]int) map[int]int {
	res := make(map[int]int, len(pebbles))
	for pebble, count := range pebbles {
		switch {
		case pebble == 0:
			res[1] += count
		case (numDigits(pebble) % 2) == 0:
			spl := split(pebble)
			res[spl[0]] += count
			res[spl[1]] += count
		default:
			res[pebble*2024] += count
		}
	}
	return res
}

func numDigits(n int) int {
	return int(math.Floor(math.Log10(float64(n)) + 1))
}

func split(pebble int) [2]int {
	d := numDigits(pebble)
	f := int(math.Pow10(d / 2))
	return [2]int{
		pebble / f,
		pebble % f,
	}
}
