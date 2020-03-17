package main

import (
	"fmt"
)

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done")
}
