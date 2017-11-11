package main

import "fmt"

type test struct {
	value string
}

// Function with pointer receiver, which accesses the value without checking if the receiver points to an actual value
func (t *test) fail1() {
	fmt.Printf("in *t: %s\n", t.value)
}

func fail2(t *test) {
	fmt.Printf("t: %s", t.value)
}

func (t *test) ok1() {
	if t == nil {
		return
	}
	fmt.Printf("in *t: %s\n", t.value)
}

func ok2(t *test) {
	if t != nil {
		fmt.Printf("t: %s", t.value)
	}
}

func main() {
	var x *test
	x.ok1()
	ok2(x)

	// x.fail1()
	// fail2(x)
}
