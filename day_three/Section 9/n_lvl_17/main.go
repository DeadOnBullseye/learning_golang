package main

import (
	"fmt"
)

func main() {
	x := 1999
	for {
		x++
		if x > 2020 {
			break
		}

		fmt.Println(x)
	}
}
