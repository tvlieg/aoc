package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("./example.txt")
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var count int
	for r := range strings.SplitSeq(input, `,`) {
		before, after, ok := strings.Cut(r, `-`)
		if !ok {
			panic(r)
		}
		begin, err := strconv.Atoi(before)
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(after)
		if err != nil {
			panic(err)
		}
		count += getRange(begin, end)
	}
	fmt.Println("1:", count)
}

func getRange(begin, end int) int {
	var sum int
	for i := first(begin); double(i) <= end; i++ {
		sum += double(i)
	}
	return sum
}

func double(n int) int {
	if n == 0 {
		return 0
	}

	exp := int(math.Log10(float64(n))) + 1
	return n*int(math.Pow10(exp)) + n
}

// first gets the first invalid number greater than or equal to n.
func first(n int) int {
	digits := int(math.Log10(float64(n))) + 1

	// odd number of digits
	if digits%2 == 1 {
		return int(math.Pow10(digits / 2))
	}

	// even number of digits
	base := n / int(math.Pow10(digits/2))
	if n > double(base) {
		base++
	}

	return base
}
