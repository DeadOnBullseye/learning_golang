package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Outer: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tInner: %d\n", j)
		}
	}
}
