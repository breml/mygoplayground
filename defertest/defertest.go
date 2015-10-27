package main

import (
	"fmt"
)

func function() {
	defer fmt.Println("first defer in function")
	fmt.Println("in function")
	defer fmt.Println("last defer in function")
}

func main() {
	fmt.Println("start of main")
	defer fmt.Println("first defer at start of main")
	if true {
		defer fmt.Println("defer in if")
	}
	if false {
		defer fmt.Println("defer in if false")
	} else {
		defer fmt.Println("defer in else false")
	}
	fmt.Println("in main before function")
	function()
	fmt.Println("in main after function")
	defer fmt.Println("last defer at end of main")
	fmt.Println("end of main")
}
