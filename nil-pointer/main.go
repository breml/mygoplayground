package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type myint struct {
	I int
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func saveCall(f interface{}, args interface{}) {
	vf := reflect.ValueOf(f)
	vargs := reflect.ValueOf(args)

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("failed to call: %s, signature %T, with values: %v, got err: %v\n", getFunctionName(f), f, args, err)
		}
	}()

	vf.Call([]reflect.Value{vargs})
}

func foo4(s *[1024]int) int {
	return s[0]
}

func foo3(s []int) int {
	return s[0]
}

func foo2(c chan int) int {
	return <-c
}

func foo(i *myint) int {
	return i.I
}

func main() {
	//var x *myint

	//saveCall(foo, nil)

	//saveCall(foo2, nil)
	
	//saveCall(foo3, nil)

	//fmt.Println(reflect.TypeOf(x))

	// fmt.Println(foo(x))
	
	//foo2(nil)
	
	//foo3(nil)
	
	//foo4(nil)
	
	
}
