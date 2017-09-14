// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package mypackage is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Test
*/
package mypackage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "example.Test.OptionalGroup")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x26, 0x32,
	0x72, 0xb1, 0x84, 0xa4, 0x16, 0x97, 0x08, 0xf1, 0x72, 0xb1, 0xe6, 0x24, 0x26, 0xa5, 0xe6, 0x48,
	0x30, 0x2a, 0x30, 0x69, 0x70, 0x0a, 0x09, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x29,
	0x30, 0x6a, 0xb0, 0x5a, 0x31, 0x99, 0x9b, 0x0b, 0xf1, 0x70, 0xb1, 0x14, 0xa5, 0x16, 0x14, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0x30, 0x0b, 0x19, 0x71, 0xf1, 0xe6, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x25,
	0xe6, 0xa4, 0x17, 0xe5, 0x97, 0x16, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x19, 0x49, 0xeb, 0xc1,
	0xec, 0x01, 0x19, 0xaa, 0xe7, 0x0f, 0x55, 0xe2, 0x0e, 0x52, 0x22, 0xa5, 0xc6, 0xc5, 0x8b, 0x22,
	0x20, 0x24, 0xca, 0xc5, 0x1b, 0x94, 0x5a, 0x58, 0x9a, 0x59, 0x94, 0x9a, 0xe2, 0x96, 0x99, 0x9a,
	0x93, 0x22, 0xc1, 0x0a, 0xb2, 0x5b, 0x8b, 0x87, 0x8b, 0xd9, 0xcd, 0xdf, 0x5f, 0x88, 0x95, 0x8b,
	0x31, 0x42, 0x40, 0x10, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf6, 0x43, 0x7a, 0xba, 0x00, 0x00,
	0x00,
}