// https://www.hackerrank.com/challenges/30-2d-arrays

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
	var inArray [6][6]int
	var out int
	for i := 0; i < 6; i++ {
		line, _, err := in.ReadLine()
		if err != nil {
			fmt.Println(err)
			return -1
		}
		numbers := strings.Split(string(line), " ")
		for j := 0; j < 6; j++ {
			inArray[i][j], err = strconv.Atoi(strings.TrimSpace(numbers[j]))
			if err != nil {
				fmt.Println(err)
				return -2
			}
		}

	}
	for k := 1; k < 5; k++ {
		for l := 1; l < 5; l++ {
			sum := inArray[l-1][k-1] + inArray[l-1][k] + inArray[l-1][k+1] +
				inArray[l][k] +
				inArray[l+1][k-1] + inArray[l+1][k] + inArray[l+1][k+1]
			if l == 1 && k == 1 {
				out = sum
			}
			if sum > out {
				out = sum
			}
		}
	}
	return out
}
