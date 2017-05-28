// https://www.hackerrank.com/challenges/30-binary-numbers

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := answer(os.Stdin)
	fmt.Println(a)
}

func answer(input io.Reader) int {
	in := bufio.NewReader(input)
	line, _, err := in.ReadLine()
	if err != nil {
		panic(err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(line)))
	if err != nil {
		panic(err)
	}
	binary := strconv.FormatInt(int64(n), 2)
	var count, maximum int
	for _, l := range binary {
		if l%2 == 1 {
			count++
			if count > maximum {
				maximum = count
			}
		} else {
			count = 0
		}
	}
	return maximum

	// Or with string split and sort but slower in this cases
	// out := strings.Split(binary, "0")
	// sort.Slice(out, func(i int, j int) bool {
	// 	return len(out[i]) > len(out[j])
	// })
	// return len(out[0])
}
