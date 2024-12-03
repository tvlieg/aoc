package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "example_input"
	// filePath := "input"

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
		fmt.Println(report)
		if isSafe(report) {
			fmt.Println("safe")
			safe++
			continue
		}
		for i := range report {
			tmp := make([]int, len(report)-1)
			s1 := report[0:i]
			s2 := report[i+1:]
			tmp := append(s1, s2...)
			fmt.Println("dampening: ", tmp)
			if isSafe(tmp) {
				fmt.Println("Safe when removing element: ", i)
				safe++
				continue
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
