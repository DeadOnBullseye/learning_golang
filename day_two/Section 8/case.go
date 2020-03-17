package main

import (
	"fmt"
)

func main() {
	// switch {
	// case false:
	// 	fmt.Println("001")
	// case (2 == 4):
	// 	fmt.Println("002")
	// case (3 == 3):
	// 	fmt.Println("003")
	// 	fallthrough
	// case (4 == 4):
	// 	fmt.Println("004")
	// 	fallthrough
	// case (8 == 3):
	// 	fmt.Println("003")
	// 	fallthrough
	// case (7 == 3):
	// 	fmt.Println("003")
	// 	fallthrough
	// case (15 == 15):
	// 	fmt.Println("003")
	// default:
	// 	fmt.Println("Default")
	// }

	n := "Bond"

	switch n {
	case "Moneypenny", "bond":
		fmt.Println("Miss money")
	case "Bond":
		fmt.Println("Bond, James Bond.")
	case "Q":
		fmt.Println("This is q")
	default:
		fmt.Println("Default")
	}

}
