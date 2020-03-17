package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb
	gb
	tb
	pb

	// kb = 1024
	// mb = 1024 * kb
	// gb = 1024 * mb
	// tb = 1024 * gb
	// pb = 1024 * tb
)

func main() {
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
	fmt.Printf("%d\t%b\n", tb, tb)
	fmt.Printf("%d\t%b\n", pb, pb)
}
