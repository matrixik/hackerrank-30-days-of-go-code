package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Your text: ")
	in := bufio.NewReader(os.Stdin)
	var input string
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, World.")
	fmt.Println(input)
}
