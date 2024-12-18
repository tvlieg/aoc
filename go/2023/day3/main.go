package main

import (
	"bufio"
	"embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input example
var f embed.FS

var re = regexp.MustCompile(`\d+`)

func process(fileName string, fn func(prev, cur, next string) int) (sum int, err error) {
	file, err := f.Open(fileName)
	if err != nil {
		fmt.Println("could not open file:", err)
	}
	defer file.Close()

	var prev, cur, next string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prev = cur
		cur = next
		next = scanner.Text()
		if cur != "" {
			sum += fn(prev, cur, next)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	// Last line
	prev = cur
	cur = next
	sum += part1(prev, cur, "")

	return sum, nil
}

func part1(before, line, after string) (sum int) {
	locs := re.FindAllStringIndex(line, -1)
	for _, l := range locs {
		start := l[0]
		end := l[1]
		n, _ := strconv.Atoi(line[start:end])
		surround := getSurround(before, line, after, start, end)
		if strings.ContainsFunc(surround, isSymbol) {
			sum += n
		}
	}
	return sum
}

func isSymbol(r rune) bool {
	if r == '.' {
		return false
	}
	return unicode.IsSymbol(r) || unicode.IsPunct(r)
}

func getSurround(before, line, after string, start, end int) string {
	a := start
	if start > 0 {
		a--
	}
	z := end
	if end < len(line)-1 {
		z++
	}

	surround := make([]string, 0)
	surround = append(surround, line[a:start])
	surround = append(surround, line[end:z])
	if before != "" {
		surround = append(surround, before[a:z])
	}
	if after != "" {
		surround = append(surround, after[a:z])
	}

	return strings.Join(surround, "")
}

func part2(before, line, after string) (sum int) {
	locs := re.FindAllStringIndex(line, -1)
	for _, l := range locs {
		start := l[0]
		end := l[1]
		n, _ := strconv.Atoi(line[start:end])
		surround := getSurround(before, line, after, start, end)
		if strings.ContainsFunc(surround, isSymbol) {
			sum += n
		}
	}
	return sum
}

func main() {
	// Part 1
	sum, err := process("input", part1)
	if err != nil {
		fmt.Println("Part 1 failed:", err)
	}
	fmt.Println(sum)

	// Part 2
	sum, err = process("example", part2)
	if err != nil {
		fmt.Println("Part 2 failed:", err)
	}
	fmt.Println(sum)
}
