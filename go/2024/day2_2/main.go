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
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	var safe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			report[i] = n
		}
		if isSafe(report) {
			safe++
			continue
		}
		for i := range report {
			tmp := make([]int, len(report))
			copy(tmp, report)
			tmp = slices.Delete(tmp, i, i+1)
			if isSafe(tmp) {
				safe++
				break
			}
		}
	}

	fmt.Println(safe)
}

func isSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	var increase, decrease bool
	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]
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
