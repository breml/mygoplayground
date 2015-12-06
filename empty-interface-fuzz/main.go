package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type myint int

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

func fooVarBad(i ...interface{}) {
	x := i[0].(int)
	fmt.Println(x)
}

func fooVarBad2(i ...interface{}) {
	for _, ii := range i {
		x := ii.(int)
		fmt.Println(x)
	}
}

func fooMyInt(i interface{}) {
	x := i.(myint)
	fmt.Println(x)
}

func save(f func(x interface{}), arg interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("work failed: %s %T %v %v\n", GetFunctionName(f), f, arg, err)
		}
	}()
	f(arg)
}

func saveVar(f func(x ...interface{}), arg ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("work failed: %s %T %v %v\n", GetFunctionName(f), f, arg, err)
		}
	}()
	f(arg...)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	fooCorrect("123")
	fooCorrect2("123")
	save(fooBad, "123")
	save(fooBad, nil)
	saveVar(fooVarBad, "123", "123")
	saveVar(fooVarBad2, "123", "123")
	saveVar(fooVarBad2, 123, nil)
	saveVar(fooVarBad2, 123, 123)
	save(fooMyInt, 123)
	save(fooMyInt, nil)

	var i myint = 124
	save(fooMyInt, i)
}
