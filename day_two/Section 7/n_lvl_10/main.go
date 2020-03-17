package main

import (
	"fmt"
)

func main() {
	num := 42
	fmt.Printf("NUMBER ONE\n----------\nDec: %d\nBin: %b\nHex: %#x\n\n", num, num, num)
	num2 := num << 1
	fmt.Printf("NUMBER TWO\n----------\nDec: %d\nBin: %b\nHex: %#x\n", num2, num2, num2)
}
