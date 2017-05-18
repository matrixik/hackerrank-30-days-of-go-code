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
	p, err := answer(os.Stdin)
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
	fmt.Println(p)
}

func answer(input io.Reader) (string, error) {
	in := bufio.NewReader(input)
	line, _, err := in.ReadLine()
	if err != nil {
		return "", fmt.Errorf("read error: %s", err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(line)))
	if err != nil {
		return "", fmt.Errorf("converting error: %s", err)
	}

	if n%2 == 1 || n >= 6 && n <= 20 {
		return "Weird", nil
	} else if n >= 2 && n <= 5 || n > 20 {
		return "Not Weird", nil
	}
	return "Not Weird", nil
}
