package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

var rnd = rand.New(rand.NewSource(42))

func TestFoo(t *testing.T) {
	var x *myint
	val, ok := quick.Value(reflect.TypeOf(x), rnd)
	if !ok {
		t.Log("Error creating value")
	}
	t.Log(val)

	foo(nil)
}
