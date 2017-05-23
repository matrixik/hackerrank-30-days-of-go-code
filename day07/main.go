// https://www.hackerrank.com/challenges/30-arrays

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	a, err := answer(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func answer(input io.Reader) (string, error) {

	in := bufio.NewReader(input)
	// Skip number
	_ = readLine(in)
	ar := readLine(in)
	a := strings.Split(ar, " ")
	var out string
	for i := range a {
		out += fmt.Sprintf("%s ", a[len(a)-1-i])
	}
	return strings.TrimSpace(out), nil
}

func readLine(i *bufio.Reader) string {
	line, err := i.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			panic(fmt.Errorf("ReadString: %v", err))
		}
	}
	return strings.TrimSpace(line)
}
