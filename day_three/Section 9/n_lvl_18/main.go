package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		r := i % 4
		fmt.Println(r)
	}
}
