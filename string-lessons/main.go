package main

import "fmt"

func main() {
	// STRINGS ARE IMMUTABLE
	// Any operation like toUpper or s+="es" all happen in a copy
	// Older string gets removed by the garbage collector
	s := "The quick brown fox"
	a := s[:3]
	c := s[4:9]
	d := s[:4] + "slow" + s[9:]
	// Copies contents of s to a diff location, adds es there. Others are still pointing to
	// original location
	s += "es"

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
}
