package main

import (
	"bufio"
	"embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input example
var f embed.FS

var re = regexp.MustCompile(`\d+`)

func process(fileName string, fn func(a, b, c string) int) (sum int, err error) {
	file, err := f.Open(fileName)
	if err != nil {
		fmt.Println("could not open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prev, cur, next string
	for scanner.Scan() {
		sum += fn(prev, cur, next)
		prev = cur
		cur = next
		next = scanner.Text()
	}
	sum += fn(prev, cur, next)
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return sum, nil
}

func part1(prev, cur, next string) int {
	locs := re.FindAllStringIndex(cur, -1)
	for _, l := range locs {
		n, _ := strconv.Atoi(cur[l[0]:l[1]])
		fmt.Println(n)
	}
	return 0
}

func main() {
	// Part 1
	sum, err := process("example", part1)
	if err != nil {
		fmt.Println("Part 1 failed:", err)
	}
	fmt.Println(sum)

	//// Part 2
	//sum, err = process("input.txt", part2)
	//if err != nil {
	//	fmt.Println("Part 2 failed:", err)
	//}
	//fmt.Println(sum)
}
