// Code generated by protoc-gen-go.
// source: command.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	command.proto

It has these top-level messages:
	Command
	Commands
	Response
*/
package protos

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Type int32

const (
	Type_Get     Type = 0
	Type_Set     Type = 1
	Type_Add     Type = 2
	Type_Iterate Type = 3
	Type_Flush   Type = 4
)

var Type_name = map[int32]string{
	0: "Get",
	1: "Set",
	2: "Add",
	3: "Iterate",
	4: "Flush",
}
var Type_value = map[string]int32{
	"Get":     0,
	"Set":     1,
	"Add":     2,
	"Iterate": 3,
	"Flush":   4,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}

type Command struct {
	Type             *Type   `protobuf:"varint,1,req,name=type,enum=protos.Type" json:"type,omitempty"`
	Key              *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}

func (m *Command) GetType() Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Type_Get
}

func (m *Command) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Command) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Commands struct {
	Command          []*Command `protobuf:"bytes,1,rep,name=command" json:"command,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Commands) Reset()         { *m = Commands{} }
func (m *Commands) String() string { return proto.CompactTextString(m) }
func (*Commands) ProtoMessage()    {}

func (m *Commands) GetCommand() []*Command {
	if m != nil {
		return m.Command
	}
	return nil
}

type Response struct {
	Status           *bool    `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	Key              *string  `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            []string `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return false
}

func (m *Response) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Response) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.Type", Type_name, Type_value)
}