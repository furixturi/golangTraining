package main

import (
	"fmt"
)

func main() {
	// Print 42 in decimal, binary, hexdecimal
	fmt.Printf("%d, %b, %x\n", 42, 42, 42)

	//%x is to print hexdecimal with 0x in front
	fmt.Printf("%d, %b, %#x\n", 42, 42, 42)

	//%X is to print hexdecimal with 0X in front and capital A-F
	fmt.Printf("%d, %b, %#X\n", 42, 42, 42)
}
