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
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wrapping, ribbon int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "x")

		l, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])
		h, _ := strconv.Atoi(nums[2])
		wrapping += calcWrapping(l, w, h)
		ribbon += calcRibbon(l, w, h)
	}
	fmt.Println("Part 1:", wrapping)
	fmt.Println("Part 2:", ribbon)
}

func calcWrapping(l, w, h int) int {
	a := l * w
	b := l * h
	c := w * h
	extra := slices.Min([]int{a, b, c})
	surface := 2*l*w + 2*w*h + 2*h*l + extra
	return surface
}

func calcRibbon(l, w, h int) int {
	a := 2*l + 2*w
	b := 2*l + 2*h
	c := 2*w + 2*h
	wrap := slices.Min([]int{a, b, c})
	bow := l * w * h
	return wrap + bow
}
