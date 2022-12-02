// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MsgId int32

const (
	MsgId_Invalid MsgId = 0
	// client <> game
	MsgId_CToG_RoleLogin MsgId = 1
)

var MsgId_name = map[int32]string{
	0: "Invalid",
	1: "CToG_RoleLogin",
}
var MsgId_value = map[string]int32{
	"Invalid":        0,
	"CToG_RoleLogin": 1,
}

func (x MsgId) String() string {
	return proto.EnumName(MsgId_name, int32(x))
}
func (MsgId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ResultId int32

const (
	ResultId_ResultIdInvalid ResultId = 0
	ResultId_Success         ResultId = 1
	ResultId_RoleNotExist    ResultId = 2
	ResultId_System          ResultId = 3
	ResultId_InvalidParams   ResultId = 4
	ResultId_Fail            ResultId = 5
	ResultId_ViolateWord     ResultId = 6
	ResultId_SystemBusy      ResultId = 7
)

var ResultId_name = map[int32]string{
	0: "ResultIdInvalid",
	1: "Success",
	2: "RoleNotExist",
	3: "System",
	4: "InvalidParams",
	5: "Fail",
	6: "ViolateWord",
	7: "SystemBusy",
}
var ResultId_value = map[string]int32{
	"ResultIdInvalid": 0,
	"Success":         1,
	"RoleNotExist":    2,
	"System":          3,
	"InvalidParams":   4,
	"Fail":            5,
	"ViolateWord":     6,
	"SystemBusy":      7,
}

func (x ResultId) String() string {
	return proto.EnumName(ResultId_name, int32(x))
}
func (ResultId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterEnum("rpc.MsgId", MsgId_name, MsgId_value)
	proto.RegisterEnum("rpc.ResultId", ResultId_name, ResultId_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x71, 0x63, 0x9b, 0xa4, 0x4c, 0x6a, 0x3b, 0x8e, 0x2f, 0x21, 0x3d, 0x78, 0xf1, 0x0d,
	0x2a, 0x2a, 0x01, 0x15, 0x49, 0x44, 0xc1, 0x8b, 0xac, 0xbb, 0x4b, 0x58, 0xd8, 0xcd, 0x84, 0x9d,
	0x8d, 0x98, 0x37, 0xf0, 0xb1, 0x25, 0xfe, 0xa1, 0xb7, 0xdf, 0xe5, 0x03, 0xbf, 0x2f, 0xac, 0x35,
	0x87, 0xc0, 0xfd, 0xc5, 0x10, 0x39, 0x31, 0x2d, 0xe2, 0xa0, 0x77, 0xe7, 0x90, 0xdf, 0x4b, 0x57,
	0x1b, 0xaa, 0xa0, 0xac, 0xfb, 0x0f, 0xe5, 0x9d, 0xc1, 0x23, 0x22, 0xd8, 0x5c, 0x3d, 0xf1, 0xed,
	0x5b, 0xc3, 0xde, 0xde, 0x71, 0xe7, 0x7a, 0xcc, 0x76, 0x5f, 0x19, 0xac, 0x1a, 0x2b, 0xa3, 0x4f,
	0xb5, 0xa1, 0x33, 0xd8, 0xfe, 0xef, 0x83, 0xaa, 0xa0, 0x6c, 0x47, 0xad, 0xad, 0x08, 0x66, 0x84,
	0xb0, 0x9e, 0xf5, 0x03, 0xa7, 0xeb, 0x4f, 0x27, 0x09, 0x8f, 0x09, 0xa0, 0x68, 0x27, 0x49, 0x36,
	0xe0, 0x82, 0x4e, 0xe1, 0xe4, 0xcf, 0x3d, 0xaa, 0xa8, 0x82, 0xe0, 0x92, 0x56, 0xb0, 0xbc, 0x51,
	0xce, 0x63, 0x4e, 0x5b, 0xa8, 0x9e, 0x1d, 0x7b, 0x95, 0xec, 0x0b, 0x47, 0x83, 0x05, 0x6d, 0x00,
	0x7e, 0xe5, 0x7e, 0x94, 0x09, 0xcb, 0x7d, 0xfe, 0x3a, 0x7f, 0x7f, 0x2f, 0x7e, 0x3a, 0x2e, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xe2, 0x3b, 0x22, 0xd7, 0x00, 0x00, 0x00,
}
