package main

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
	"unicode"
)

var digitWords = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

//go:embed *.txt
var f embed.FS

func hasWordDigit(s string) (digit int, ok bool) {
	if unicode.IsDigit(rune(s[0])) {
		return int(s[0] - '0'), true
	}
	for i, w := range digitWords {
		if strings.HasPrefix(s, w) {
			return i, true
		}
	}
	return 0, false
}

func process(fileName string, calibrate func(string) int) (sum int, err error) {
	file, err := f.Open(fileName)
	if err != nil {
		fmt.Println("could not open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += calibrate(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w\n", err)
	}

	return sum, nil
}

func calibrateDigits(s string) int {
	i := strings.IndexFunc(s, unicode.IsDigit)
	first := int(s[i] - '0')

	i = strings.LastIndexFunc(s, unicode.IsDigit)
	last := int(s[i] - '0')

	return first*10 + last
}

func calibrateWordsAndDigits(s string) int {
	first := firstDigit(s)
	last := lastDigit(s)
	return first*10 + last
}

func firstDigit(s string) int {
	for len(s) != 0 {
		if d, ok := hasWordDigit(s); ok {
			return d
		}
		s = s[1:]
	}
	return 0
}

func lastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if d, ok := hasWordDigit(s[i:]); ok {
			return d
		}
	}
	return 0
}

func main() {
	// Part 1
	sum, err := process("example.txt", calibrateDigits)
	if err != nil {
		fmt.Println("Part 1 failed:", err)
	}
	fmt.Println(sum)

	// Part 2
	sum, err = process("input.txt", calibrateWordsAndDigits)
	if err != nil {
		fmt.Println("Part 2 failed:", err)
	}
	fmt.Println(sum)
}
