package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	count := make(map[int]int)
	factor := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])
		count[l]++
		factor[r] += r
	}

	var similarity int
	for n := range count {
		similarity += count[n] * factor[n]
	}

	fmt.Printf("Similarity: %d\n", similarity)
}
