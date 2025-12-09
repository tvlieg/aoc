package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// f, err := os.Open("./example.txt")
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var coords []coord
	for scanner.Scan() {
		xStr, yStr, ok := strings.Cut(scanner.Text(), ",")
		if !ok {
			continue
		}
		x, err := strconv.Atoi(xStr)
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(yStr)
		if err != nil {
			panic(err)
		}
		coords = append(coords, coord{x, y})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := run(coords)
	fmt.Println("1:", result)
}

func run(coords []coord) int {
	var maxArea int
	for i := range coords {
		for j := range coords[i:] {
			if a := area(coords[i], coords[j]); a > maxArea {
				maxArea = a
			}
		}

	}
	return maxArea
}

type coord struct {
	x, y int
}

func area(a, b coord) int {
	x := int(math.Abs(float64(a.x-b.x))) + 1
	y := int(math.Abs(float64(a.y-b.y))) + 1
	return x * y
}
