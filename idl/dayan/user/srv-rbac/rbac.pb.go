// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rbac.proto

package dayan_user_srv_rbac

import (
	fmt "fmt"
	_ "github.com/dy-dayan/user-srv-rbac/idl"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddRoleReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoleReq) Reset()         { *m = AddRoleReq{} }
func (m *AddRoleReq) String() string { return proto.CompactTextString(m) }
func (*AddRoleReq) ProtoMessage()    {}
func (*AddRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{0}
}

func (m *AddRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleReq.Unmarshal(m, b)
}
func (m *AddRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleReq.Marshal(b, m, deterministic)
}
func (m *AddRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleReq.Merge(m, src)
}
func (m *AddRoleReq) XXX_Size() int {
	return xxx_messageInfo_AddRoleReq.Size(m)
}
func (m *AddRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleReq proto.InternalMessageInfo

type AddRoleResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoleResp) Reset()         { *m = AddRoleResp{} }
func (m *AddRoleResp) String() string { return proto.CompactTextString(m) }
func (*AddRoleResp) ProtoMessage()    {}
func (*AddRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{1}
}

func (m *AddRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleResp.Unmarshal(m, b)
}
func (m *AddRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleResp.Marshal(b, m, deterministic)
}
func (m *AddRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleResp.Merge(m, src)
}
func (m *AddRoleResp) XXX_Size() int {
	return xxx_messageInfo_AddRoleResp.Size(m)
}
func (m *AddRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleResp proto.InternalMessageInfo

type DelRoleReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRoleReq) Reset()         { *m = DelRoleReq{} }
func (m *DelRoleReq) String() string { return proto.CompactTextString(m) }
func (*DelRoleReq) ProtoMessage()    {}
func (*DelRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{2}
}

func (m *DelRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRoleReq.Unmarshal(m, b)
}
func (m *DelRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRoleReq.Marshal(b, m, deterministic)
}
func (m *DelRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRoleReq.Merge(m, src)
}
func (m *DelRoleReq) XXX_Size() int {
	return xxx_messageInfo_DelRoleReq.Size(m)
}
func (m *DelRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelRoleReq proto.InternalMessageInfo

type DelRoleResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRoleResp) Reset()         { *m = DelRoleResp{} }
func (m *DelRoleResp) String() string { return proto.CompactTextString(m) }
func (*DelRoleResp) ProtoMessage()    {}
func (*DelRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{3}
}

func (m *DelRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRoleResp.Unmarshal(m, b)
}
func (m *DelRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRoleResp.Marshal(b, m, deterministic)
}
func (m *DelRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRoleResp.Merge(m, src)
}
func (m *DelRoleResp) XXX_Size() int {
	return xxx_messageInfo_DelRoleResp.Size(m)
}
func (m *DelRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelRoleResp proto.InternalMessageInfo

type AddAccessReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAccessReq) Reset()         { *m = AddAccessReq{} }
func (m *AddAccessReq) String() string { return proto.CompactTextString(m) }
func (*AddAccessReq) ProtoMessage()    {}
func (*AddAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{4}
}

func (m *AddAccessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAccessReq.Unmarshal(m, b)
}
func (m *AddAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAccessReq.Marshal(b, m, deterministic)
}
func (m *AddAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAccessReq.Merge(m, src)
}
func (m *AddAccessReq) XXX_Size() int {
	return xxx_messageInfo_AddAccessReq.Size(m)
}
func (m *AddAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddAccessReq proto.InternalMessageInfo

type AddAccessResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAccessResp) Reset()         { *m = AddAccessResp{} }
func (m *AddAccessResp) String() string { return proto.CompactTextString(m) }
func (*AddAccessResp) ProtoMessage()    {}
func (*AddAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{5}
}

func (m *AddAccessResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAccessResp.Unmarshal(m, b)
}
func (m *AddAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAccessResp.Marshal(b, m, deterministic)
}
func (m *AddAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAccessResp.Merge(m, src)
}
func (m *AddAccessResp) XXX_Size() int {
	return xxx_messageInfo_AddAccessResp.Size(m)
}
func (m *AddAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddAccessResp proto.InternalMessageInfo

type DelAccessReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelAccessReq) Reset()         { *m = DelAccessReq{} }
func (m *DelAccessReq) String() string { return proto.CompactTextString(m) }
func (*DelAccessReq) ProtoMessage()    {}
func (*DelAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{6}
}

func (m *DelAccessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelAccessReq.Unmarshal(m, b)
}
func (m *DelAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelAccessReq.Marshal(b, m, deterministic)
}
func (m *DelAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelAccessReq.Merge(m, src)
}
func (m *DelAccessReq) XXX_Size() int {
	return xxx_messageInfo_DelAccessReq.Size(m)
}
func (m *DelAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelAccessReq proto.InternalMessageInfo

type DelAccessResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelAccessResp) Reset()         { *m = DelAccessResp{} }
func (m *DelAccessResp) String() string { return proto.CompactTextString(m) }
func (*DelAccessResp) ProtoMessage()    {}
func (*DelAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{7}
}

func (m *DelAccessResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelAccessResp.Unmarshal(m, b)
}
func (m *DelAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelAccessResp.Marshal(b, m, deterministic)
}
func (m *DelAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelAccessResp.Merge(m, src)
}
func (m *DelAccessResp) XXX_Size() int {
	return xxx_messageInfo_DelAccessResp.Size(m)
}
func (m *DelAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelAccessResp proto.InternalMessageInfo

type AssignRoleAccessReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignRoleAccessReq) Reset()         { *m = AssignRoleAccessReq{} }
func (m *AssignRoleAccessReq) String() string { return proto.CompactTextString(m) }
func (*AssignRoleAccessReq) ProtoMessage()    {}
func (*AssignRoleAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{8}
}

func (m *AssignRoleAccessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignRoleAccessReq.Unmarshal(m, b)
}
func (m *AssignRoleAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignRoleAccessReq.Marshal(b, m, deterministic)
}
func (m *AssignRoleAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignRoleAccessReq.Merge(m, src)
}
func (m *AssignRoleAccessReq) XXX_Size() int {
	return xxx_messageInfo_AssignRoleAccessReq.Size(m)
}
func (m *AssignRoleAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignRoleAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_AssignRoleAccessReq proto.InternalMessageInfo

type AssignRoleAccessResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignRoleAccessResp) Reset()         { *m = AssignRoleAccessResp{} }
func (m *AssignRoleAccessResp) String() string { return proto.CompactTextString(m) }
func (*AssignRoleAccessResp) ProtoMessage()    {}
func (*AssignRoleAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{9}
}

func (m *AssignRoleAccessResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignRoleAccessResp.Unmarshal(m, b)
}
func (m *AssignRoleAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignRoleAccessResp.Marshal(b, m, deterministic)
}
func (m *AssignRoleAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignRoleAccessResp.Merge(m, src)
}
func (m *AssignRoleAccessResp) XXX_Size() int {
	return xxx_messageInfo_AssignRoleAccessResp.Size(m)
}
func (m *AssignRoleAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignRoleAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_AssignRoleAccessResp proto.InternalMessageInfo

type RemoveRoleAccessReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRoleAccessReq) Reset()         { *m = RemoveRoleAccessReq{} }
func (m *RemoveRoleAccessReq) String() string { return proto.CompactTextString(m) }
func (*RemoveRoleAccessReq) ProtoMessage()    {}
func (*RemoveRoleAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{10}
}

func (m *RemoveRoleAccessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoleAccessReq.Unmarshal(m, b)
}
func (m *RemoveRoleAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoleAccessReq.Marshal(b, m, deterministic)
}
func (m *RemoveRoleAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoleAccessReq.Merge(m, src)
}
func (m *RemoveRoleAccessReq) XXX_Size() int {
	return xxx_messageInfo_RemoveRoleAccessReq.Size(m)
}
func (m *RemoveRoleAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoleAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoleAccessReq proto.InternalMessageInfo

type RemoveRoleAccessResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRoleAccessResp) Reset()         { *m = RemoveRoleAccessResp{} }
func (m *RemoveRoleAccessResp) String() string { return proto.CompactTextString(m) }
func (*RemoveRoleAccessResp) ProtoMessage()    {}
func (*RemoveRoleAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f88ffdd966c9c7ed, []int{11}
}

func (m *RemoveRoleAccessResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoleAccessResp.Unmarshal(m, b)
}
func (m *RemoveRoleAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoleAccessResp.Marshal(b, m, deterministic)
}
func (m *RemoveRoleAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoleAccessResp.Merge(m, src)
}
func (m *RemoveRoleAccessResp) XXX_Size() int {
	return xxx_messageInfo_RemoveRoleAccessResp.Size(m)
}
func (m *RemoveRoleAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoleAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoleAccessResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddRoleReq)(nil), "dayan.user.srv.rbac.AddRoleReq")
	proto.RegisterType((*AddRoleResp)(nil), "dayan.user.srv.rbac.AddRoleResp")
	proto.RegisterType((*DelRoleReq)(nil), "dayan.user.srv.rbac.DelRoleReq")
	proto.RegisterType((*DelRoleResp)(nil), "dayan.user.srv.rbac.DelRoleResp")
	proto.RegisterType((*AddAccessReq)(nil), "dayan.user.srv.rbac.AddAccessReq")
	proto.RegisterType((*AddAccessResp)(nil), "dayan.user.srv.rbac.AddAccessResp")
	proto.RegisterType((*DelAccessReq)(nil), "dayan.user.srv.rbac.DelAccessReq")
	proto.RegisterType((*DelAccessResp)(nil), "dayan.user.srv.rbac.DelAccessResp")
	proto.RegisterType((*AssignRoleAccessReq)(nil), "dayan.user.srv.rbac.AssignRoleAccessReq")
	proto.RegisterType((*AssignRoleAccessResp)(nil), "dayan.user.srv.rbac.AssignRoleAccessResp")
	proto.RegisterType((*RemoveRoleAccessReq)(nil), "dayan.user.srv.rbac.RemoveRoleAccessReq")
	proto.RegisterType((*RemoveRoleAccessResp)(nil), "dayan.user.srv.rbac.RemoveRoleAccessResp")
}

func init() { proto.RegisterFile("rbac.proto", fileDescriptor_f88ffdd966c9c7ed) }

var fileDescriptor_f88ffdd966c9c7ed = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x03, 0x21,
	0x10, 0x86, 0x2f, 0x8d, 0xc6, 0xb1, 0x55, 0x43, 0xd5, 0xc3, 0x5e, 0x54, 0x4e, 0x7a, 0xe1, 0xa0,
	0x4f, 0x80, 0xf6, 0xe8, 0x89, 0x37, 0xd8, 0x5d, 0x26, 0x8d, 0xc9, 0x5a, 0x90, 0xd1, 0x26, 0xbe,
	0x8b, 0x0f, 0x6b, 0x90, 0x86, 0xa5, 0x0b, 0xa8, 0xc7, 0x61, 0xbe, 0xfc, 0xc3, 0xfc, 0xff, 0x00,
	0xb8, 0xae, 0xed, 0x85, 0x75, 0xe6, 0xdd, 0xb0, 0xa5, 0x6e, 0x3f, 0xdb, 0x8d, 0xf8, 0x20, 0x74,
	0x82, 0xdc, 0x56, 0xf8, 0x56, 0x03, 0x5d, 0x4b, 0x18, 0x00, 0x3e, 0x07, 0x90, 0x5a, 0x2b, 0x33,
	0xa0, 0xc2, 0x37, 0xbe, 0x80, 0xe3, 0x58, 0x91, 0xf5, 0xcd, 0x15, 0x0e, 0x49, 0x33, 0x56, 0x64,
	0xf9, 0x09, 0xcc, 0xa5, 0xd6, 0xb2, 0xef, 0x91, 0xc8, 0xb7, 0x4f, 0x61, 0x91, 0xd4, 0x01, 0x58,
	0xe1, 0xb0, 0x07, 0x24, 0x35, 0x59, 0x7e, 0x01, 0x4b, 0x49, 0xf4, 0xb2, 0xde, 0x78, 0xcd, 0x91,
	0xbb, 0x84, 0xf3, 0xfc, 0x39, 0xe0, 0x0a, 0x5f, 0xcd, 0x16, 0x33, 0x3c, 0x7f, 0x26, 0x7b, 0xff,
	0x35, 0x83, 0x99, 0x7a, 0x94, 0x4f, 0xec, 0x19, 0x0e, 0x77, 0x4b, 0xb1, 0x2b, 0x51, 0xf0, 0x43,
	0x8c, 0x06, 0x34, 0xd7, 0xbf, 0x03, 0x64, 0xbd, 0xda, 0xce, 0x85, 0x8a, 0xda, 0xe8, 0x58, 0x45,
	0x2d, 0x31, 0x91, 0x29, 0x38, 0x8a, 0xa6, 0xb1, 0x9b, 0xda, 0xf0, 0xb8, 0x6c, 0xc3, 0xff, 0x42,
	0x82, 0x66, 0xf4, 0xb9, 0xa2, 0x99, 0xe6, 0x52, 0xd1, 0xdc, 0x8b, 0x8a, 0xad, 0xe1, 0x6c, 0x9a,
	0x09, 0xbb, 0x2d, 0xff, 0x25, 0x4f, 0xb4, 0xb9, 0xfb, 0x27, 0x19, 0x06, 0x4d, 0xd3, 0xac, 0x0c,
	0x2a, 0xdc, 0x42, 0x65, 0x50, 0xe9, 0x3c, 0xba, 0x83, 0x9f, 0xfb, 0x7f, 0xf8, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0xc6, 0x7d, 0xa2, 0x2e, 0x03, 0x00, 0x00,
}
