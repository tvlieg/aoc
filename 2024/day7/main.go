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
		candidates := recurse1(eq.numbers)
		if slices.Contains(candidates, eq.testValue) {
			count += eq.testValue
		}
	}
	fmt.Println("Part 1:", count)

	count = 0
	for _, eq := range equations {
		candidates := recurse2(eq.numbers)
		if slices.Contains(candidates, eq.testValue) {
			count += eq.testValue
		}
	}
	fmt.Println("Part 2:", count)
}

func recurse1(ns []int) []int {
	if len(ns) == 1 {
		return ns
	}

	tail := ns[len(ns)-1]
	rest := ns[:len(ns)-1]
	results := recurse1(rest)

	var ret []int
	for _, n := range results {
		ret = append(ret, tail+n)
	}
	for _, n := range results {
		ret = append(ret, tail*n)
	}
	return ret
}

func recurse2(ns []int) []int {
	if len(ns) == 1 {
		return ns
	}

	tail := ns[len(ns)-1]
	rest := ns[:len(ns)-1]
	results := recurse2(rest)

	var ret []int
	for _, n := range results {
		ret = append(ret, tail+n)
	}
	for _, n := range results {
		ret = append(ret, tail*n)
	}
	for _, n := range results {
		ret = append(ret, concat(n, tail))
	}

	return ret
}

func concat(a, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	n, _ := strconv.Atoi(s)
	return n
}
