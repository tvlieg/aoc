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

	var updates [][]int
	before := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		nums := strings.Split(line, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])

		before[first] = append(before[first], second)

	}
	for scanner.Scan() {
		line := scanner.Text()
		strings := strings.Split(line, ",")
		nums := make([]int, len(strings))
		for i, s := range strings {
			n, _ := strconv.Atoi(s)
			nums[i] = n
		}

		updates = append(updates, nums)

	}

	var count int
	for _, update := range updates {
		if isCorrectlyOrdered(update, before) {
			middle := len(update) / 2
			count += update[middle]
		}
	}
	fmt.Println(count)
}

func isCorrectlyOrdered(update []int, rules map[int][]int) bool {
	// Loop over pages in update
	for i := len(update) - 1; i >= 0; i-- {
		page := update[i]
		mustAfter, ok := rules[page]
		if !ok {
			// no rules for page
			continue
		}

		before := update[0:i]
		for _, pageBefore := range before {
			if slices.Contains(mustAfter, pageBefore) {
				return false
			}
		}
	}
	return true
}
