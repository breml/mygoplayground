package main

import (
	"fmt"
)

func main() {
	var base rune
	base = '\u2500'
	for i := 0; i < 255; i++ {
		// base := int('\u2500')
		// fmt.Print(string(rune(base+i)) + " ")
		fmt.Printf("%c = %x\n", base+rune(i), base+rune(i))
	}
}
