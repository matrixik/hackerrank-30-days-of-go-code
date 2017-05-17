package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var in int
	var do float64
	var st string
	// Read and save an integer, double, and String to your variables.
	num := 0
	for num != 3 {
		scanner.Scan()
		v := scanner.Text()
		switch num {
		case 0:
			in, _ = strconv.Atoi(v)
			num += 1
		case 1:
			do, _ = strconv.ParseFloat(v, 64)
			num += 1
		case 2:
			st = v
			num += 1
		}
	}

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + uint64(in))
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+do)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + st)
}
