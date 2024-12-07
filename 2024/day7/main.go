package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type equation struct {
	testValue int
	numbers   []int
}

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var equations []equation
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ": ")
		numberStrings := strings.Split(s[1], " ")

		testValue, _ := strconv.Atoi(s[0])

		numbers := make([]int, len(numberStrings))
		for i, s := range numberStrings {
			n, _ := strconv.Atoi(s)
			numbers[i] = n
		}
		equations = append(equations, equation{testValue, numbers})
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanning input: %v", err)
		return
	}

	var count int
	for _, eq := range equations {
		candidates := f(eq.numbers)
		if slices.Contains(candidates, eq.testValue) {
			count += eq.testValue
		}
	}
	fmt.Println("Part 1:", count)
}

func f(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}

	tail := numbers[len(numbers)-1]
	rest := numbers[:len(numbers)-1]

	ret := make([]int, 0, len(numbers)*2)
	recurse := f(rest)
	for _, n := range recurse {
		ret = append(ret, tail+n)
	}
	for _, n := range recurse {
		ret = append(ret, tail*n)
	}
	return ret
}
