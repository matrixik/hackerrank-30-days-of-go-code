// https://www.hackerrank.com/challenges/30-recursion

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
	n := answer(os.Stdin)
	fmt.Println(n)
}

func answer(input io.Reader) int {
	in := bufio.NewReader(input)
	line, err := in.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			return -1
		}
	}
	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return -2
	}
	return factorial(n)
}

func factorial(n int) int {
	if n > 1 {
		return n * factorial(n-1)
	}
	return 1
}
