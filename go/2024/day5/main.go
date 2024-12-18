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

	scanner := bufio.NewScanner(file)

	rules := make(map[[2]int]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		nums := strings.Split(line, "|")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		rules[[2]int{a, b}] = struct{}{}
	}

	var updates [][]int
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
		if !isCorrectlyOrdered(update, rules) {
			continue
		}
		middle := len(update) / 2
		count += update[middle]
	}
	fmt.Println("part1:", count)

	count = 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			continue
		}
		slices.SortFunc(update, func(a, b int) int {
			if _, ok := rules[[2]int{a, b}]; ok {
				return -1
			}
			if _, ok := rules[[2]int{b, a}]; ok {
				return 1
			}
			return 0
		})
		middle := len(update) / 2
		count += update[middle]
	}
	fmt.Println("part2:", count)
}
func isCorrectlyOrdered(update []int, rules map[[2]int]struct{}) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if _, ok := rules[[2]int{update[j], update[i]}]; ok {
				return false
			}
		}
	}
	return true
}
