package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	var (
		inputInt   uint64
		inputFloat float64
		// str        string
	)

	fmt.Scanf("%d", &inputInt)
	fmt.Scanf("%f", &inputFloat)

	for scanner.Scan() {
		fmt.Printf("%d\n", i+inputInt)
		fmt.Printf("%.1f\n", d+inputFloat)
		fmt.Println(s + scanner.Text())
		break
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}
}
