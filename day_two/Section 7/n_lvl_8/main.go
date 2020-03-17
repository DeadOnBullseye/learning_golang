package main

import (
	"fmt"
)

const (
	a = 42 == 42
	b = 42 <= 42
	c = 42 >= 42
	d = 42 != 42
	e = 42 < 42
	f = 42 > 42
)

func main() {

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

}
