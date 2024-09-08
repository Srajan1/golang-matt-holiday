package main

import (
	"fmt"
	"os"
)

// run using cmd params
// go run .
// 4
// 4
// 6
// 9
// Average is  5.75
//
// Run using file
// go run . < nums.txt

func main() {
	var sum float64
	var n int

	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err == nil {
			sum += val
			n++
		} else {
			break
		}
	}
	if n == 0 {
		fmt.Fprintf(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("Average is ", sum/float64(n))
}
