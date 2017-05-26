// https://www.hackerrank.com/challenges/30-dictionaries-and-maps

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	err := answer(os.Stdin)
	if err != nil {
		panic(err)
	}
}

func answer(input io.Reader) error {
	var phoneNum int
	fmt.Scan(&phoneNum)

	phoneBook := make(map[string]string)
	for i := 0; i < phoneNum; i++ {
		var name, phone string
		fmt.Scan(&name, &phone)
		phoneBook[name] = phone
	}

	// Must output straight to os.Stdout instead getting timeouts on Hackerrank
	// on test with 50000 names to print
	s := bufio.NewScanner(input)
	for s.Scan() {
		nameToPrint := s.Text()
		if p, ok := phoneBook[nameToPrint]; ok {
			fmt.Printf("%s=%s\n", nameToPrint, p)
		} else {
			fmt.Println("Not found")
		}
	}

	return nil
}
