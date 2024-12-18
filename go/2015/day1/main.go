package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "input"

	file, _ := os.Open(filePath)
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var floor, pos int
	for i, b := range bytes {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 && pos == 0 {
			fmt.Println("basement")
			pos = i + 1
		}
	}
	fmt.Println("Part 1:", floor)
	fmt.Println("Part 2:", pos)
}
