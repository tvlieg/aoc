package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// filePath := "example_input"
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	input, _ := io.ReadAll(file)

	const expr = `mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`
	re := regexp.MustCompile(expr)
	matches := re.FindAllStringSubmatch(string(input), -1)

	do := true
	var sum int
	for _, m := range matches {
		switch m[0] {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				a, _ := strconv.Atoi(m[1])
				b, _ := strconv.Atoi(m[2])
				sum += a * b
			}
		}
	}
	fmt.Println(sum)
}
