package main

import (
	"fmt"
)

func main() {
	// if true {
	// 	fmt.Println("001")
	// }

	// if false {
	// 	fmt.Println("002")
	// }

	// if !true {
	// 	fmt.Println("003")
	// }

	// if !false {
	// 	fmt.Println("004")
	// }

	// if 2 == 2 {
	// 	fmt.Println("005")
	// }

	// if 2 != 2 {
	// 	fmt.Println("006")
	// }

	// if !(2 == 2) {
	// 	fmt.Println("007")
	// }

	// if !(2 != 2) {
	// 	fmt.Println("008")
	// }

	x := 42

	if x == 40 {
		fmt.Println("Value equals: 40")
	} else if x == 41 {
		fmt.Println("Value equals: 41")
	} else {
		fmt.Println("Value does not equal 40 or 41")
	}

}