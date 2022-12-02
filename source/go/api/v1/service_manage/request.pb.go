// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

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

type DiscoverRequest_DiscoverRequestType int32

const (
	DiscoverRequest_UNKNOWN         DiscoverRequest_DiscoverRequestType = 0
	DiscoverRequest_INSTANCE        DiscoverRequest_DiscoverRequestType = 1
	DiscoverRequest_CLUSTER         DiscoverRequest_DiscoverRequestType = 2
	DiscoverRequest_ROUTING         DiscoverRequest_DiscoverRequestType = 3
	DiscoverRequest_RATE_LIMIT      DiscoverRequest_DiscoverRequestType = 4
	DiscoverRequest_CIRCUIT_BREAKER DiscoverRequest_DiscoverRequestType = 5
	DiscoverRequest_SERVICES        DiscoverRequest_DiscoverRequestType = 6
	DiscoverRequest_NAMESPACES      DiscoverRequest_DiscoverRequestType = 12
)

var DiscoverRequest_DiscoverRequestType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "INSTANCE",
	2:  "CLUSTER",
	3:  "ROUTING",
	4:  "RATE_LIMIT",
	5:  "CIRCUIT_BREAKER",
	6:  "SERVICES",
	12: "NAMESPACES",
}
var DiscoverRequest_DiscoverRequestType_value = map[string]int32{
	"UNKNOWN":         0,
	"INSTANCE":        1,
	"CLUSTER":         2,
	"ROUTING":         3,
	"RATE_LIMIT":      4,
	"CIRCUIT_BREAKER": 5,
	"SERVICES":        6,
	"NAMESPACES":      12,
}

func (x DiscoverRequest_DiscoverRequestType) String() string {
	return proto.EnumName(DiscoverRequest_DiscoverRequestType_name, int32(x))
}
func (DiscoverRequest_DiscoverRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_request_c646cb0384284a57, []int{0, 0}
}

type DiscoverRequest struct {
	Type                 DiscoverRequest_DiscoverRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.DiscoverRequest_DiscoverRequestType" json:"type,omitempty"`
	Service              *Service                            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *DiscoverRequest) Reset()         { *m = DiscoverRequest{} }
func (m *DiscoverRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoverRequest) ProtoMessage()    {}
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_c646cb0384284a57, []int{0}
}
func (m *DiscoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRequest.Unmarshal(m, b)
}
func (m *DiscoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRequest.Merge(dst, src)
}
func (m *DiscoverRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoverRequest.Size(m)
}
func (m *DiscoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRequest proto.InternalMessageInfo

func (m *DiscoverRequest) GetType() DiscoverRequest_DiscoverRequestType {
	if m != nil {
		return m.Type
	}
	return DiscoverRequest_UNKNOWN
}

func (m *DiscoverRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoverRequest)(nil), "v1.DiscoverRequest")
	proto.RegisterEnum("v1.DiscoverRequest_DiscoverRequestType", DiscoverRequest_DiscoverRequestType_name, DiscoverRequest_DiscoverRequestType_value)
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_request_c646cb0384284a57) }

var fileDescriptor_request_c646cb0384284a57 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6f, 0xb2, 0x40,
	0x14, 0x85, 0x5f, 0x10, 0x3f, 0x32, 0xf8, 0x31, 0xc1, 0x8d, 0x79, 0x57, 0xc6, 0xa4, 0xa9, 0xab,
	0x21, 0xda, 0x45, 0x17, 0x5d, 0x21, 0x4e, 0x1a, 0xaa, 0xa2, 0x19, 0xa0, 0x4d, 0xba, 0x31, 0x38,
	0x99, 0xea, 0x24, 0x85, 0xa1, 0x80, 0x24, 0xfe, 0x88, 0x2e, 0x9a, 0xfe, 0xe1, 0x66, 0x04, 0x17,
	0x35, 0x5d, 0x9e, 0x7b, 0xef, 0x79, 0x26, 0xe7, 0x0c, 0xe8, 0xa4, 0xec, 0xe3, 0xc8, 0xb2, 0x1c,
	0x25, 0xa9, 0xc8, 0x85, 0xa1, 0x16, 0x93, 0xff, 0x9d, 0x8c, 0xa5, 0x05, 0xa7, 0xac, 0x1c, 0x8d,
	0xbe, 0x54, 0xd0, 0x9b, 0xf3, 0x8c, 0x8a, 0x82, 0xa5, 0xa4, 0x3c, 0x36, 0x1e, 0x80, 0x96, 0x9f,
	0x12, 0x36, 0x50, 0x86, 0xca, 0xb8, 0x3b, 0xbd, 0x45, 0xc5, 0x04, 0x5d, 0x9d, 0x5c, 0x6b, 0xff,
	0x94, 0x30, 0x72, 0x36, 0x19, 0x37, 0xa0, 0x59, 0xbd, 0x30, 0x50, 0x87, 0xca, 0x58, 0x9f, 0xea,
	0xd2, 0xef, 0x95, 0x23, 0x72, 0xd9, 0x8d, 0xbe, 0x15, 0xd0, 0xff, 0x03, 0x62, 0xe8, 0xa0, 0x19,
	0xb8, 0x0b, 0x77, 0xfd, 0xe2, 0xc2, 0x7f, 0x46, 0x1b, 0xb4, 0x1c, 0xd7, 0xf3, 0x2d, 0xd7, 0xc6,
	0x50, 0x91, 0x2b, 0x7b, 0x19, 0x78, 0x3e, 0x26, 0x50, 0x95, 0x82, 0xac, 0x03, 0xdf, 0x71, 0x1f,
	0x61, 0xcd, 0xe8, 0x02, 0x40, 0x2c, 0x1f, 0x6f, 0x97, 0xce, 0xca, 0xf1, 0xa1, 0x66, 0xf4, 0x41,
	0xcf, 0x76, 0x88, 0x1d, 0x38, 0xfe, 0x76, 0x46, 0xb0, 0xb5, 0xc0, 0x04, 0xd6, 0x25, 0xcc, 0xc3,
	0xe4, 0xd9, 0xb1, 0xb1, 0x07, 0x1b, 0xd2, 0xe2, 0x5a, 0x2b, 0xec, 0x6d, 0x2c, 0xa9, 0xdb, 0x23,
	0xad, 0xd5, 0x84, 0xfa, 0x93, 0xd6, 0xaa, 0xc1, 0xfa, 0xec, 0x53, 0x01, 0xf7, 0x54, 0x44, 0x28,
	0x67, 0x31, 0x65, 0x71, 0x8e, 0x12, 0xf1, 0x1e, 0xa6, 0x3c, 0x43, 0x59, 0xc2, 0x28, 0x7f, 0xe3,
	0x34, 0xcc, 0xb9, 0x88, 0x51, 0x98, 0x70, 0x99, 0xec, 0x52, 0x67, 0x14, 0xc6, 0xe1, 0x9e, 0xcd,
	0xda, 0x55, 0x98, 0x8d, 0x6c, 0xf7, 0x75, 0xbe, 0xe7, 0xf9, 0xe1, 0xb8, 0x43, 0x54, 0x44, 0x66,
	0x45, 0x89, 0x58, 0x76, 0x30, 0x7f, 0x91, 0xcc, 0x4c, 0x1c, 0x53, 0xca, 0xcc, 0xbd, 0x30, 0xc3,
	0x84, 0x9b, 0xc5, 0xc4, 0xac, 0x98, 0xdb, 0x92, 0xb9, 0x6b, 0x9c, 0xbf, 0xea, 0xee, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0x20, 0x24, 0xf5, 0xce, 0x01, 0x00, 0x00,
}
