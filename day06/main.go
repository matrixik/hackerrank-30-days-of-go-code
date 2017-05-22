// https://www.hackerrank.com/challenges/30-review-loop

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
	a, _ := answer(os.Stdin)
	fmt.Println(a)
}

func answer(input io.Reader) (string, error) {
	read := func(i *bufio.Reader) string {
		line, err := i.ReadString('\n')
		if err != nil {
			panic(fmt.Errorf("ReadString error: %v", err))
		}
		return strings.TrimSpace(string(line))
	}

	in := bufio.NewReader(input)
	n, err := strconv.Atoi(read(in))
	if err != nil {
		return "", fmt.Errorf("convert: %v", err)
	}
	var out string
	for i := 0; i < n; i++ {
		s := read(in)
		even := ""
		odd := ""
		for j := 0; j < len(s); j++ {
			if j%2 == 0 {
				even += string(s[j])
			} else {
				odd += string(s[j])
			}
		}
		out += even + " " + odd + "\n"
	}
	return out, nil
}
