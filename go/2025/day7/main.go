package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./example.txt")
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var current string
	var sum int
	for scanner.Scan() {
		next, splits := apply([]byte(current), scanner.Bytes())
		sum += splits
		fmt.Println(string(next))
		current = string(next)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Splits:", sum)
}

func apply(current, next []byte) ([]byte, int) {
	if len(current) == 0 {
		return bytes.Replace(next, []byte{'S'}, []byte{'|'}, 1), 0
	}

	var splits int
	for {
		i := bytes.IndexRune(current, '|')
		if i == -1 {
			break
		}

		switch next[i] {
		case '.':
			next[i] = '|'
		case '^':
			next[i-1] = '|'
			next[i+1] = '|'
			splits++
		}

		current[i] = '.'
	}

	return next, splits
}
