// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

type Test struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64 `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}

const Default_Test_Type int32 = 77

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}
