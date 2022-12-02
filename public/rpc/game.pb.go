// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	game.proto
	common.proto

It has these top-level messages:
	FailAck
	RoleLoginReq
	RoleLoginAck
*/
package rpc

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

type FailAck struct {
	ResultId ResultId `protobuf:"varint,1,opt,name=ResultId,enum=rpc.ResultId" json:"ResultId,omitempty"`
	ErrMsg   string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
}

func (m *FailAck) Reset()                    { *m = FailAck{} }
func (m *FailAck) String() string            { return proto.CompactTextString(m) }
func (*FailAck) ProtoMessage()               {}
func (*FailAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FailAck) GetResultId() ResultId {
	if m != nil {
		return m.ResultId
	}
	return ResultId_ResultIdInvalid
}

func (m *FailAck) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type RoleLoginReq struct {
	RoleId  int32  `protobuf:"varint,1,opt,name=RoleId" json:"RoleId,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=Account" json:"Account,omitempty"`
}

func (m *RoleLoginReq) Reset()                    { *m = RoleLoginReq{} }
func (m *RoleLoginReq) String() string            { return proto.CompactTextString(m) }
func (*RoleLoginReq) ProtoMessage()               {}
func (*RoleLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RoleLoginReq) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *RoleLoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type RoleLoginAck struct {
	ResultId      ResultId `protobuf:"varint,1,opt,name=ResultId,enum=rpc.ResultId" json:"ResultId,omitempty"`
	ErrMsg        string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	RoleId        int32    `protobuf:"varint,3,opt,name=RoleId" json:"RoleId,omitempty"`
	Account       string   `protobuf:"bytes,4,opt,name=Account" json:"Account,omitempty"`
	Name          string   `protobuf:"bytes,5,opt,name=Name" json:"Name,omitempty"`
	Level         int32    `protobuf:"varint,6,opt,name=Level" json:"Level,omitempty"`
	HasCreateRole bool     `protobuf:"varint,7,opt,name=HasCreateRole" json:"HasCreateRole,omitempty"`
}

func (m *RoleLoginAck) Reset()                    { *m = RoleLoginAck{} }
func (m *RoleLoginAck) String() string            { return proto.CompactTextString(m) }
func (*RoleLoginAck) ProtoMessage()               {}
func (*RoleLoginAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RoleLoginAck) GetResultId() ResultId {
	if m != nil {
		return m.ResultId
	}
	return ResultId_ResultIdInvalid
}

func (m *RoleLoginAck) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *RoleLoginAck) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *RoleLoginAck) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RoleLoginAck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleLoginAck) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *RoleLoginAck) GetHasCreateRole() bool {
	if m != nil {
		return m.HasCreateRole
	}
	return false
}

func init() {
	proto.RegisterType((*FailAck)(nil), "rpc.FailAck")
	proto.RegisterType((*RoleLoginReq)(nil), "rpc.RoleLoginReq")
	proto.RegisterType((*RoleLoginAck)(nil), "rpc.RoleLoginAck")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0x49, 0xce, 0xcf,
	0xcd, 0xcd, 0xcf, 0x83, 0x08, 0x29, 0xf9, 0x70, 0xb1, 0xbb, 0x25, 0x66, 0xe6, 0x38, 0x26, 0x67,
	0x0b, 0x69, 0x72, 0x71, 0x04, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x19, 0xf1, 0xea, 0x15, 0x15, 0x24, 0xeb, 0xc1, 0x04, 0x83, 0xe0, 0xd2, 0x42, 0x62,
	0x5c, 0x6c, 0xae, 0x45, 0x45, 0xbe, 0xc5, 0xe9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50,
	0x9e, 0x92, 0x03, 0x17, 0x4f, 0x50, 0x7e, 0x4e, 0xaa, 0x4f, 0x7e, 0x7a, 0x66, 0x5e, 0x50, 0x6a,
	0x21, 0x48, 0x1d, 0x88, 0x0f, 0x35, 0x90, 0x35, 0x08, 0xca, 0x13, 0x92, 0xe0, 0x62, 0x77, 0x4c,
	0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x81, 0x1a, 0x00, 0xe3, 0x2a, 0x5d, 0x66, 0x44, 0x32, 0x82, 0x3a,
	0xae, 0x42, 0x72, 0x05, 0x33, 0x2e, 0x57, 0xb0, 0xa0, 0xb8, 0x42, 0x48, 0x88, 0x8b, 0xc5, 0x2f,
	0x31, 0x37, 0x55, 0x82, 0x15, 0x2c, 0x0c, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xfa, 0xa4, 0x96, 0xa5,
	0xe6, 0x48, 0xb0, 0x81, 0x0d, 0x81, 0x70, 0x84, 0x54, 0xb8, 0x78, 0x3d, 0x12, 0x8b, 0x9d, 0x8b,
	0x52, 0x13, 0x4b, 0x52, 0x41, 0xc6, 0x4a, 0xb0, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0xa1, 0x0a, 0x3a,
	0xb1, 0x46, 0x81, 0x82, 0x3e, 0x89, 0x0d, 0x1c, 0xe6, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x41, 0xb3, 0x9d, 0x2c, 0x94, 0x01, 0x00, 0x00,
}
