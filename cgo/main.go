package main

/*

void cCall();

*/
import "C"

import "fmt"

//export goCall
func goCall() {
	fmt.Println("test from goCall")
}

func main() {
	fmt.Println("test from go")

	C.cCall()
}
