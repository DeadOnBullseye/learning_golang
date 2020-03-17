package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	foo()

	fmt.Println("Something More")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in the foo func!")
}

func bar() {
	fmt.Println("And then we exited!")
}
