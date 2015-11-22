package main

import (
	"fmt"
	"github.com/mitchellh/go-mruby"
)

func main() {
	mrb := mruby.NewMrb()
	defer mrb.Close()

	// Our custom function we'll expose to Ruby. The first return
	// value is what to return from the func and the second is an
	// exception to raise (if any).
	addFunc := func(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
		args := m.GetArgs()
		return mruby.Int(args[0].Fixnum() + args[1].Fixnum()), nil
	}

	// Lets define a custom class and a class method we can call.
	class := mrb.DefineClass("Example", nil)
	class.DefineClassMethod("add", addFunc, mruby.ArgsReq(2))

	// Let's call it and inspect the result
	result, err := mrb.LoadString(`Example.add(12, 30)`)
	if err != nil {
		panic(err.Error())
	}

	// This will output "Result: 42"
	fmt.Printf("Result: %s\n", result.String())
}
