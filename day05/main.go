package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/30-loops
func main() {
	ans, err := answer(os.Stdin)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(ans)
	}
}

func answer(input io.Reader) (string, error) {
	in := bufio.NewReader(input)
	line, _, err := in.ReadLine()
	if err != nil {
		return "", fmt.Errorf("ReadLine: %v", err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(line)))
	if err != nil {
		return "", fmt.Errorf("convert: %v", err)
	}
	var out string
	for i := 1; i <= 10; i++ {
		out += fmt.Sprintf("%d x %d = %d\n", n, i, n*i)
	}
	return out, nil
}
