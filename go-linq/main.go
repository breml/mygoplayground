package main

import (
	"fmt"
	linq "github.com/ahmetalpbalkan/go-linq"
)

func main() {
	var rs []linq.T
	var ri int
	var err error

	s := []string{"A", "B", "C", "B"}

	ri, err = linq.From(s).Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("res:", ri)

	rs, err = linq.From(s).Distinct().Results()
	if err != nil {
		panic(err)
	}
	fmt.Println("res:", rs)

	rs, err = linq.From(s).Reverse().Results()
	if err != nil {
		panic(err)
	}
	fmt.Println("res:", rs)

	rs, err = linq.From(s).Where(func(a linq.T) (bool, error) {
		return a.(string) == "A", nil
	}).Results()
	if err != nil {
		panic(err)
	}
	fmt.Println("res:", rs)
}
