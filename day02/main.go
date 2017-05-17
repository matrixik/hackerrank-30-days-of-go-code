package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("The total meal cost is %d dollars.\n", response(os.Stdin))
}

func response(input io.Reader) int64 {
	in := bufio.NewReader(input)
	mealC, _ := in.ReadString('\n')
	tipP, _ := in.ReadString('\n')
	taxP, _ := in.ReadString('\n')

	getVal := func(x string) float64 {
		r, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
		return r
	}

	mealCost := getVal(mealC)
	tipPercent := getVal(tipP)
	taxPercent := getVal(taxP)
	return round(mealCost + (mealCost * ((tipPercent + taxPercent) / 100)))
}

func round(number float64) int64 {
	if number < 0 {
		return int64(math.Ceil(number - 0.5))
	}
	return int64(math.Floor(number + 0.5))
}
