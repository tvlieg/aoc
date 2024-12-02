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

	var safe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		level := make([]int, len(fields))
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			level[i] = n
		}
		if isSafe(level) {
			safe++
		}
	}

	fmt.Println(safe)
}

func isSafe(level []int) bool {
	if len(level) <= 1 {
		return true
	}

	var increase, decrease bool
	for i := 0; i < len(level)-1; i++ {
		a := level[i]
		b := level[i+1]
		if a == b {
			return false
		}

		diff := b - a
		if diff > 3 || diff < -3 {
			return false
		}

		if a > b {
			decrease = true
		} else {
			increase = true
		}

		if increase && decrease {
			return false
		}
	}

	return true
}
