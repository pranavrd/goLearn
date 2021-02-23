package main

import (
	"fmt"
)

const (
	_ = 1 << iota
	a
	b
	c
	d
)

func main() {
	// i := 42
	// j := strconv.Itoa(i)
	// fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
}
