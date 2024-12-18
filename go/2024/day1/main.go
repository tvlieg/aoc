package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filePath := "example_input"
	// filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		l, _ := strconv.Atoi(fields[0])
		left = append(left, l)

		r, _ := strconv.Atoi(fields[1])
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int
	for i := range left {
		if left[i] >= right[i] {
			distance += (left[i] - right[i])
			continue
		}
		distance += (right[i] - left[i])
	}
	fmt.Printf("Distance: %d", distance)
}
