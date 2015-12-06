package main

import (
	"fmt"
)

func foo(i interface{}) {
	x := i.(int)
	fmt.Println(x)
}

func main() {
	foo("123")
}
