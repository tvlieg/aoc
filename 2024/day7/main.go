package main

import (
	"bufio"
	"fmt"
	"math"
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
		candidates := recurse1(eq.testValue, eq.numbers)
		if slices.Contains(candidates, eq.testValue) {
			count += eq.testValue
		}
	}
	fmt.Println("Part 1:", count)

	count = 0
	for _, eq := range equations {
		candidates := recurse2(eq.testValue, eq.numbers)
		if slices.Contains(candidates, eq.testValue) {
			count += eq.testValue
		}
	}
	fmt.Println("Part 2:", count)
}

func recurse1(goal int, ns []int) []int {
	if len(ns) == 1 {
		return ns
	}

	tail := ns[len(ns)-1]
	rest := ns[:len(ns)-1]

	var ret []int
	for _, r := range recurse1(goal, rest) {
		if sum := tail + r; sum <= goal {
			ret = append(ret, sum)
		}
		if product := tail * r; product <= goal {
			ret = append(ret, product)
		}
	}

	return ret
}

func recurse2(goal int, ns []int) []int {
	if len(ns) == 1 {
		return ns
	}

	tail := ns[len(ns)-1]
	rest := ns[:len(ns)-1]

	var ret []int
	for _, r := range recurse2(goal, rest) {
		if sum := tail + r; sum <= goal {
			ret = append(ret, sum)
		}
		if product := tail * r; product <= goal {
			ret = append(ret, product)
		}
		if conc := concat(r, tail); conc <= goal {
			ret = append(ret, conc)
		}
	}

	return ret
}

func concat(a, b int) int {
	digits := int(math.Floor(math.Log10(float64(b)) + 1))
	return a*int(math.Pow10(digits)) + b
}
