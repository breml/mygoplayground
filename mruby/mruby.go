// Example based on https://github.com/mitchellh/go-mruby

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

	// ruby class
	result, err = mrb.LoadString(`
		module MyConfig
		        module DSL
		                def config(name, opts={})
		                        @config ||= Hash.new
		                        name = name.to_s if name.is_a?(Symbol)
		                        @config[name] = opts
		                        if name.is_a?(String)
		                                define_method(name) { instance_variable_get("@#{name}") }
		                                define_method("#{name}=") { |v| instance_variable_set("@#{name}", v) }
		                        end
		                end
		                def get_config
		                        return @config
		                end
		        end


		        def self.included(base)
		                # Add the DSL methods to the 'base' given.
		                base.extend(MyConfig::DSL)
		        end
		end

		class TestFilter
		        include MyConfig

		        config :message, :default => "defaultvalue"

		        public
		        def initialize()
		                @message = self.class.get_config["message"][:default]
		        end

		        public
		        def filter(event)
		                event["message"] = @message
		                event
		        end
		end
	`)
	if err != nil {
		panic(err.Error())
	}

	// instanciate and call method
	result, err = mrb.LoadString(`
		  tf = TestFilter.new
	      tf.filter({"message" => "input", "key" => "value"})
	    `)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Result: %s, Type: %d\n", result.String(), result.Type())
	if result.Type() == mruby.TypeHash {
		h := result.Hash()
		keys, err := h.Keys()
		if err != nil {
			panic(err.Error())
		}
		for i := 0; i < keys.Array().Len(); i++ {
			key, err := keys.Array().Get(i)
			if err != nil {
				panic(err.Error())
			}
			value, err := h.Get(key)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%s: %s\n", key.String(), value.String())
		}
	}
}
