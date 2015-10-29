package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Foo string `val:"foo struct tag"`
	Bar string `val:"bar struct tag"`
}

func main() {
	test := Test{}
	st := reflect.TypeOf(test)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fmt.Println("test:", field.Tag)
		fmt.Println("test2:", field.Tag.Get("val"))
	}
}
