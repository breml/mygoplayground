package main

type fooer interface {
	foo()
}

func main() {
	var f fooer

	f.foo()
}
