// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package nano_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// API version numbers. As proto3 doesn't have constants, so we use an enum
// which allows aliases. VERSION_INVALID is a 0-value placeholder, which
// protobuf requires.
type APIVersion int32

const (
	APIVersion_VERSION_INVALID APIVersion = 0
	APIVersion_VERSION_MAJOR   APIVersion = 1
	APIVersion_VERSION_MINOR   APIVersion = 0
)

var APIVersion_name = map[int32]string{
	0: "VERSION_INVALID",
	1: "VERSION_MAJOR",
	// Duplicate value: 0: "VERSION_MINOR",
}
var APIVersion_value = map[string]int32{
	"VERSION_INVALID": 0,
	"VERSION_MAJOR":   1,
	"VERSION_MINOR":   0,
}

func (x APIVersion) String() string {
	return proto.EnumName(APIVersion_name, int32(x))
}
func (APIVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{0}
}

// *
// The enum values must NOT change. The name of the enum must match the
// request message name, uppercased, and without the req_ prefix. This naming
// standard facilitates dynamic lookup and generic frameworks.
type RequestType int32

const (
	RequestType_INVALID           RequestType = 0
	RequestType_REGISTER_CALLBACK RequestType = 1
	// core.proto message types
	RequestType_PING RequestType = 2
	// accounts.proto message types
	RequestType_ACCOUNT_PENDING RequestType = 100
	// util.proto message types
	RequestType_ADDRESS_VALID RequestType = 1000
)

var RequestType_name = map[int32]string{
	0:    "INVALID",
	1:    "REGISTER_CALLBACK",
	2:    "PING",
	100:  "ACCOUNT_PENDING",
	1000: "ADDRESS_VALID",
}
var RequestType_value = map[string]int32{
	"INVALID":           0,
	"REGISTER_CALLBACK": 1,
	"PING":              2,
	"ACCOUNT_PENDING":   100,
	"ADDRESS_VALID":     1000,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{1}
}

// *
// Request header.
// This is serialized before the actual request to tell the node what message to expect next.
// Other request meta data may be added in the future.
type Request struct {
	// * Request type
	Type                 RequestType `protobuf:"varint,1,opt,name=type,proto3,enum=nano.api.RequestType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() RequestType {
	if m != nil {
		return m.Type
	}
	return RequestType_INVALID
}

// *
// Response header.
// This is serialized before the actual response.
type Response struct {
	// *
	// For which request type is this a response? This flag allows future support for clients
	// issuing multiple concurrent requests, as well as callback messages.
	//
	// This may not be set if error_code is non-zero.
	Type RequestType `protobuf:"varint,1,opt,name=type,proto3,enum=nano.api.RequestType" json:"type,omitempty"`
	// *
	// If non-zero, an error has occurred.
	ErrorCode int32 `protobuf:"zigzag32,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// * Error message. Only set if error_code is non-zero.
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// * Error category name. Only set if error_code is non-zero.
	ErrorCategory        string   `protobuf:"bytes,4,opt,name=error_category,json=errorCategory,proto3" json:"error_category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetType() RequestType {
	if m != nil {
		return m.Type
	}
	return RequestType_INVALID
}

func (m *Response) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Response) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Response) GetErrorCategory() string {
	if m != nil {
		return m.ErrorCategory
	}
	return ""
}

// * Send ping to the node
type ReqPing struct {
	// * Ping ID. The node will respond with the same ID.
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqPing) Reset()         { *m = ReqPing{} }
func (m *ReqPing) String() string { return proto.CompactTextString(m) }
func (*ReqPing) ProtoMessage()    {}
func (*ReqPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{2}
}
func (m *ReqPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqPing.Unmarshal(m, b)
}
func (m *ReqPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqPing.Marshal(b, m, deterministic)
}
func (dst *ReqPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqPing.Merge(dst, src)
}
func (m *ReqPing) XXX_Size() int {
	return xxx_messageInfo_ReqPing.Size(m)
}
func (m *ReqPing) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqPing.DiscardUnknown(m)
}

var xxx_messageInfo_ReqPing proto.InternalMessageInfo

func (m *ReqPing) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// * Ping response
type ResPing struct {
	// * The same ID as sent in the ping request
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResPing) Reset()         { *m = ResPing{} }
func (m *ResPing) String() string { return proto.CompactTextString(m) }
func (*ResPing) ProtoMessage()    {}
func (*ResPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_e167494141e96616, []int{3}
}
func (m *ResPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResPing.Unmarshal(m, b)
}
func (m *ResPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResPing.Marshal(b, m, deterministic)
}
func (dst *ResPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResPing.Merge(dst, src)
}
func (m *ResPing) XXX_Size() int {
	return xxx_messageInfo_ResPing.Size(m)
}
func (m *ResPing) XXX_DiscardUnknown() {
	xxx_messageInfo_ResPing.DiscardUnknown(m)
}

var xxx_messageInfo_ResPing proto.InternalMessageInfo

func (m *ResPing) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "nano.api.request")
	proto.RegisterType((*Response)(nil), "nano.api.response")
	proto.RegisterType((*ReqPing)(nil), "nano.api.req_ping")
	proto.RegisterType((*ResPing)(nil), "nano.api.res_ping")
	proto.RegisterEnum("nano.api.APIVersion", APIVersion_name, APIVersion_value)
	proto.RegisterEnum("nano.api.RequestType", RequestType_name, RequestType_value)
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_core_e167494141e96616) }

var fileDescriptor_core_e167494141e96616 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x86, 0x9d, 0x54, 0xaa, 0x1e, 0x1b, 0x1b, 0xa7, 0x08, 0x41, 0x68, 0x11, 0x4b, 0xc1, 0x7a,
	0x11, 0xa1, 0xed, 0x0b, 0xa4, 0x49, 0x90, 0x54, 0x4d, 0x64, 0x62, 0xbd, 0x0d, 0xd1, 0x9c, 0x0d,
	0x81, 0xdd, 0xcc, 0x38, 0x13, 0x59, 0x7c, 0x9b, 0x7d, 0xbc, 0x7d, 0x8c, 0xc5, 0x44, 0xd9, 0x5d,
	0xd8, 0x9b, 0xbd, 0xfd, 0xfe, 0xff, 0x7c, 0x73, 0x98, 0x03, 0xb0, 0xe7, 0x12, 0x2d, 0x21, 0x79,
	0xc9, 0x69, 0xbb, 0x48, 0x0a, 0x6e, 0x25, 0x22, 0x1f, 0x7e, 0xcb, 0x38, 0xcf, 0x6e, 0x71, 0x56,
	0xf1, 0xdd, 0xf1, 0x66, 0x76, 0x2f, 0x13, 0x21, 0x50, 0xaa, 0xba, 0x39, 0xfe, 0x03, 0x2d, 0x89,
	0x87, 0x23, 0xaa, 0x92, 0xfe, 0x84, 0x66, 0x79, 0x12, 0x68, 0x92, 0x11, 0x99, 0xf4, 0x7e, 0x0d,
	0xac, 0xab, 0xc3, 0x62, 0x75, 0x61, 0x73, 0x12, 0xc8, 0xaa, 0xca, 0xf8, 0x81, 0x40, 0x5b, 0xa2,
	0x12, 0xbc, 0x50, 0xf8, 0x8e, 0x39, 0xfa, 0x15, 0x00, 0xa5, 0xe4, 0x32, 0xde, 0xf3, 0x14, 0x4d,
	0x6d, 0x44, 0x26, 0x7d, 0xd6, 0xa9, 0x88, 0xc3, 0x53, 0xa4, 0xdf, 0x41, 0xaf, 0xe3, 0x3b, 0x54,
	0x2a, 0xc9, 0xd0, 0xfc, 0x30, 0x22, 0x93, 0x0e, 0xfb, 0x54, 0xc1, 0x55, 0xcd, 0xe8, 0x0f, 0xe8,
	0x5d, 0x1c, 0x49, 0x89, 0x19, 0x97, 0x27, 0xb3, 0x59, 0xb5, 0xea, 0x51, 0xe7, 0x02, 0xc7, 0xc3,
	0xf3, 0x86, 0x87, 0x58, 0xe4, 0x45, 0x46, 0x7b, 0xa0, 0xe5, 0x69, 0xb5, 0x9f, 0xce, 0xb4, 0x3c,
	0xad, 0x33, 0xf5, 0x66, 0x36, 0x5d, 0x00, 0xd8, 0x6b, 0x7f, 0x8b, 0x52, 0xe5, 0xbc, 0xa0, 0x5f,
	0xe0, 0xf3, 0xd6, 0x63, 0x91, 0x1f, 0x06, 0xb1, 0x1f, 0x6c, 0xed, 0xa5, 0xef, 0x1a, 0x0d, 0xda,
	0x07, 0xfd, 0x0a, 0x57, 0xf6, 0xbf, 0x90, 0x19, 0xe4, 0x15, 0xf2, 0x83, 0x90, 0x19, 0x8d, 0xa1,
	0x66, 0x90, 0x69, 0x0a, 0xdd, 0x17, 0x9f, 0x40, 0xbb, 0xd0, 0x7a, 0xb6, 0x0c, 0xa0, 0xcf, 0xbc,
	0xb9, 0x1f, 0x6d, 0x3c, 0x16, 0x3b, 0xf6, 0x72, 0xf9, 0xd7, 0x76, 0x16, 0x06, 0xa1, 0x6d, 0x68,
	0xae, 0xfd, 0x60, 0x6e, 0x68, 0xe7, 0xb7, 0x6d, 0xc7, 0x09, 0xff, 0x07, 0x9b, 0x78, 0xed, 0x05,
	0xee, 0x19, 0xa6, 0x94, 0x82, 0x6e, 0xbb, 0x2e, 0xf3, 0xa2, 0x28, 0xae, 0x45, 0x8f, 0xad, 0xdd,
	0xc7, 0xea, 0x94, 0xbf, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x1b, 0x1b, 0xcb, 0x02, 0x02,
	0x00, 0x00,
}
