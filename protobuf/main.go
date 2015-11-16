package main

import (
	"fmt"
	"log"

	"github.com/breml/mygoplayground/protobuf/example"
	"github.com/golang/protobuf/proto"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// etc.
	fmt.Println(test.GetLabel(), "=", newTest.GetLabel(), test.GetType(), "=", newTest.GetType())
}
