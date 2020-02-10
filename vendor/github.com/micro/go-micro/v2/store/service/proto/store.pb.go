// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/go-micro/store/service/proto/store.proto

package go_micro_store

import (
	fmt "fmt"
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

type Record struct {
	// key of the record
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value in the record
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp in unix seconds
	Expiry               int64    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type ReadOptions struct {
	Prefix               bool     `protobuf:"varint,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{1}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

func (m *ReadOptions) GetPrefix() bool {
	if m != nil {
		return m.Prefix
	}
	return false
}

type ReadRequest struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Options              *ReadOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{2}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReadRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{3}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type WriteRequest struct {
	Record               *Record  `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{4}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{5}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84ccc98e143ed3e, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "go.micro.store.Record")
	proto.RegisterType((*ReadOptions)(nil), "go.micro.store.ReadOptions")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.store.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.store.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "go.micro.store.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "go.micro.store.WriteResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.store.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.store.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.store.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.store.ListResponse")
}

func init() {
	proto.RegisterFile("micro/go-micro/store/service/proto/store.proto", fileDescriptor_f84ccc98e143ed3e)
}

var fileDescriptor_f84ccc98e143ed3e = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x9a, 0x36, 0xd5, 0x49, 0x5b, 0xcb, 0x22, 0x25, 0xd4, 0x0f, 0xe2, 0x82, 0x90, 0x8b,
	0x69, 0xa9, 0x78, 0x15, 0xc1, 0x0f, 0x14, 0x04, 0x61, 0x05, 0x3d, 0xd7, 0x76, 0x2c, 0xc1, 0xda,
	0x8d, 0xbb, 0x69, 0x69, 0xff, 0x90, 0xbf, 0x53, 0xb2, 0xbb, 0xd1, 0x94, 0x34, 0x17, 0x6f, 0x33,
	0xfb, 0xde, 0xbc, 0x99, 0x79, 0xc3, 0x42, 0xf8, 0x19, 0x8d, 0x05, 0xef, 0x4f, 0xf9, 0x99, 0x0e,
	0x64, 0xc2, 0x05, 0xf6, 0x25, 0x8a, 0x65, 0x34, 0xc6, 0x7e, 0x2c, 0x78, 0x62, 0xde, 0x42, 0x15,
	0x93, 0xf6, 0x94, 0xeb, 0x92, 0x50, 0xbd, 0xd2, 0x7b, 0x70, 0x18, 0x8e, 0xb9, 0x98, 0x90, 0x0e,
	0xd8, 0x1f, 0xb8, 0xf6, 0x2c, 0xdf, 0x0a, 0x76, 0x59, 0x1a, 0x92, 0x7d, 0xa8, 0x2f, 0x47, 0xb3,
	0x05, 0x7a, 0x55, 0xdf, 0x0a, 0x9a, 0x4c, 0x27, 0xa4, 0x0b, 0x0e, 0xae, 0xe2, 0x48, 0xac, 0x3d,
	0xdb, 0xb7, 0x02, 0x9b, 0x99, 0x8c, 0x9e, 0x82, 0xcb, 0x70, 0x34, 0x79, 0x8a, 0x93, 0x88, 0xcf,
	0x65, 0x4a, 0x8b, 0x05, 0xbe, 0x47, 0x2b, 0xa5, 0xb8, 0xc3, 0x4c, 0x46, 0x5f, 0x34, 0x8d, 0xe1,
	0xd7, 0x02, 0x65, 0xb2, 0xa5, 0xeb, 0x05, 0x34, 0xb8, 0xd6, 0x50, 0x7d, 0xdd, 0xe1, 0x41, 0xb8,
	0x39, 0x73, 0x98, 0x6b, 0xc3, 0x32, 0x2e, 0xbd, 0x82, 0xa6, 0xd6, 0x95, 0x31, 0x9f, 0x4b, 0x24,
	0x03, 0x68, 0x08, 0xb5, 0x98, 0xf4, 0x2c, 0xdf, 0x0e, 0xdc, 0x61, 0xb7, 0x28, 0x93, 0xc2, 0x2c,
	0xa3, 0xd1, 0x4b, 0x68, 0xbe, 0x8a, 0x28, 0xc1, 0x6c, 0xb4, 0x10, 0x1c, 0x0d, 0xa9, 0xe9, 0xca,
	0x05, 0x0c, 0x8b, 0xee, 0x41, 0xcb, 0xd4, 0xeb, 0x11, 0xe8, 0x09, 0xb4, 0x6e, 0x70, 0x86, 0x7f,
	0x8a, 0x85, 0x65, 0x69, 0x07, 0xda, 0x19, 0xc5, 0x14, 0xb5, 0xc0, 0x7d, 0x8c, 0x64, 0x62, 0x4a,
	0xd2, 0xb5, 0x74, 0xfa, 0xdf, 0xb5, 0x86, 0xdf, 0x55, 0xa8, 0x3f, 0xa7, 0x08, 0xb9, 0x85, 0x5a,
	0xaa, 0x45, 0x0a, 0x86, 0xe6, 0x1a, 0xf6, 0x0e, 0xb7, 0x83, 0x66, 0xba, 0xca, 0xc0, 0x22, 0xd7,
	0x50, 0x4b, 0x9d, 0x26, 0x5b, 0xef, 0x52, 0x2a, 0x93, 0x3f, 0x0e, 0xad, 0x90, 0x3b, 0xa8, 0x2b,
	0xb3, 0x48, 0x81, 0x98, 0xbf, 0x41, 0xef, 0xa8, 0x04, 0xfd, 0xd5, 0x79, 0x00, 0x47, 0x1b, 0x48,
	0x0a, 0xd4, 0x0d, 0xef, 0x7b, 0xc7, 0x65, 0x70, 0x26, 0xf5, 0xe6, 0xa8, 0x1f, 0x72, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0x20, 0x54, 0x71, 0x53, 0x03, 0x00, 0x00,
}