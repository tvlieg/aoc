package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	const input = "bgvyzdsv"

	var five, six int
	for i := 0; ; i++ {
		data := []byte(fmt.Sprintf("%s%d", input, i))
		sum := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(sum, "00000") && five == 0 {
			five = i
		}
		if strings.HasPrefix(sum, "000000") && six == 0 {
			six = i
			break
		}
	}
	fmt.Println("Part 1:", five)
	fmt.Println("Part 2:", six)
}
