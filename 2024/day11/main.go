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
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var pebbles []int
	for _, s := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(s)
		pebbles = append(pebbles, n)
	}

	for range 25 {
		var nextPebbles []int
		for _, p := range pebbles {
			switch {
			case p == 0:
				nextPebbles = append(nextPebbles, 1)
			case (numDigits(p) % 2) == 0:
				nextPebbles = append(nextPebbles, split(p)...)
			default:
				nextPebbles = append(nextPebbles, p*2024)
			}
		}
		pebbles = nextPebbles
	}
	fmt.Println("Part 1:", len(pebbles))
}

func numDigits(n int) int {
	return int(math.Floor(math.Log10(float64(n)) + 1))
}

func split(p int) []int {
	d := numDigits(p)
	f := int(math.Pow10(d / 2))
	return []int{
		p / f,
		p % f,
	}
}
