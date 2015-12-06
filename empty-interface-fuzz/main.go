package main

import (
	"fmt"
)

func fooCorrect(i interface{}) {
	x, ok := i.(int)
	if ok {
		fmt.Printf("int value is: %q\n", x)
	} else {
		fmt.Printf("value is not an int\n")
	}
}

func fooCorrect2(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x)
	default:
		fmt.Printf("value is not an int\n")
	}
}

func fooBad(i interface{}) {
	x := i.(int)
	fmt.Println(x)
}

func main() {
	fooCorrect("123")
	fooCorrect2("123")
	fooBad("123")
}
