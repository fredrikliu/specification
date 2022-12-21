// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file_response.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigSimpleResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigSimpleResponse) Reset()         { *m = ConfigSimpleResponse{} }
func (m *ConfigSimpleResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigSimpleResponse) ProtoMessage()    {}
func (*ConfigSimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_faf971ae42ada746, []int{0}
}
func (m *ConfigSimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSimpleResponse.Unmarshal(m, b)
}
func (m *ConfigSimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSimpleResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigSimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSimpleResponse.Merge(dst, src)
}
func (m *ConfigSimpleResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigSimpleResponse.Size(m)
}
func (m *ConfigSimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSimpleResponse proto.InternalMessageInfo

func (m *ConfigSimpleResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigSimpleResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ConfigResponse struct {
	Code                     *wrapperspb.UInt32Value   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                     *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	ConfigFileGroup          *ConfigFileGroup          `protobuf:"bytes,3,opt,name=configFileGroup,proto3" json:"configFileGroup,omitempty"`
	ConfigFile               *ConfigFile               `protobuf:"bytes,4,opt,name=configFile,proto3" json:"configFile,omitempty"`
	ConfigFileRelease        *ConfigFileRelease        `protobuf:"bytes,5,opt,name=configFileRelease,proto3" json:"configFileRelease,omitempty"`
	ConfigFileReleaseHistory *ConfigFileReleaseHistory `protobuf:"bytes,6,opt,name=configFileReleaseHistory,proto3" json:"configFileReleaseHistory,omitempty"`
	ConfigFileTemplate       *ConfigFileTemplate       `protobuf:"bytes,7,opt,name=configFileTemplate,proto3" json:"configFileTemplate,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_faf971ae42ada746, []int{1}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(dst, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileGroup() *ConfigFileGroup {
	if m != nil {
		return m.ConfigFileGroup
	}
	return nil
}

func (m *ConfigResponse) GetConfigFile() *ConfigFile {
	if m != nil {
		return m.ConfigFile
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileRelease() *ConfigFileRelease {
	if m != nil {
		return m.ConfigFileRelease
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileReleaseHistory() *ConfigFileReleaseHistory {
	if m != nil {
		return m.ConfigFileReleaseHistory
	}
	return nil
}

func (m *ConfigResponse) GetConfigFileTemplate() *ConfigFileTemplate {
	if m != nil {
		return m.ConfigFileTemplate
	}
	return nil
}

type ConfigBatchWriteResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Total                *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	Responses            []*ConfigResponse       `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigBatchWriteResponse) Reset()         { *m = ConfigBatchWriteResponse{} }
func (m *ConfigBatchWriteResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigBatchWriteResponse) ProtoMessage()    {}
func (*ConfigBatchWriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_faf971ae42ada746, []int{2}
}
func (m *ConfigBatchWriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigBatchWriteResponse.Unmarshal(m, b)
}
func (m *ConfigBatchWriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigBatchWriteResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigBatchWriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigBatchWriteResponse.Merge(dst, src)
}
func (m *ConfigBatchWriteResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigBatchWriteResponse.Size(m)
}
func (m *ConfigBatchWriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigBatchWriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigBatchWriteResponse proto.InternalMessageInfo

func (m *ConfigBatchWriteResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetTotal() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *ConfigBatchWriteResponse) GetResponses() []*ConfigResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

type ConfigBatchQueryResponse struct {
	Code                       *wrapperspb.UInt32Value     `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                       *wrapperspb.StringValue     `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Total                      *wrapperspb.UInt32Value     `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	ConfigFileGroups           []*ConfigFileGroup          `protobuf:"bytes,4,rep,name=configFileGroups,proto3" json:"configFileGroups,omitempty"`
	ConfigFiles                []*ConfigFile               `protobuf:"bytes,5,rep,name=configFiles,proto3" json:"configFiles,omitempty"`
	ConfigFileReleases         []*ConfigFileRelease        `protobuf:"bytes,6,rep,name=configFileReleases,proto3" json:"configFileReleases,omitempty"`
	ConfigFileReleaseHistories []*ConfigFileReleaseHistory `protobuf:"bytes,7,rep,name=configFileReleaseHistories,proto3" json:"configFileReleaseHistories,omitempty"`
	ConfigFileTemplates        []*ConfigFileTemplate       `protobuf:"bytes,8,rep,name=configFileTemplates,proto3" json:"configFileTemplates,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *ConfigBatchQueryResponse) Reset()         { *m = ConfigBatchQueryResponse{} }
func (m *ConfigBatchQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigBatchQueryResponse) ProtoMessage()    {}
func (*ConfigBatchQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_faf971ae42ada746, []int{3}
}
func (m *ConfigBatchQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigBatchQueryResponse.Unmarshal(m, b)
}
func (m *ConfigBatchQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigBatchQueryResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigBatchQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigBatchQueryResponse.Merge(dst, src)
}
func (m *ConfigBatchQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigBatchQueryResponse.Size(m)
}
func (m *ConfigBatchQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigBatchQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigBatchQueryResponse proto.InternalMessageInfo

func (m *ConfigBatchQueryResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetTotal() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileGroups() []*ConfigFileGroup {
	if m != nil {
		return m.ConfigFileGroups
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFiles() []*ConfigFile {
	if m != nil {
		return m.ConfigFiles
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileReleases() []*ConfigFileRelease {
	if m != nil {
		return m.ConfigFileReleases
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileReleaseHistories() []*ConfigFileReleaseHistory {
	if m != nil {
		return m.ConfigFileReleaseHistories
	}
	return nil
}

func (m *ConfigBatchQueryResponse) GetConfigFileTemplates() []*ConfigFileTemplate {
	if m != nil {
		return m.ConfigFileTemplates
	}
	return nil
}

type ConfigClientResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	ConfigFile           *ClientConfigFileInfo   `protobuf:"bytes,3,opt,name=configFile,proto3" json:"configFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigClientResponse) Reset()         { *m = ConfigClientResponse{} }
func (m *ConfigClientResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigClientResponse) ProtoMessage()    {}
func (*ConfigClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_response_faf971ae42ada746, []int{4}
}
func (m *ConfigClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigClientResponse.Unmarshal(m, b)
}
func (m *ConfigClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigClientResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigClientResponse.Merge(dst, src)
}
func (m *ConfigClientResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigClientResponse.Size(m)
}
func (m *ConfigClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigClientResponse proto.InternalMessageInfo

func (m *ConfigClientResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConfigClientResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ConfigClientResponse) GetConfigFile() *ClientConfigFileInfo {
	if m != nil {
		return m.ConfigFile
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigSimpleResponse)(nil), "polaris.v1.ConfigSimpleResponse")
	proto.RegisterType((*ConfigResponse)(nil), "polaris.v1.ConfigResponse")
	proto.RegisterType((*ConfigBatchWriteResponse)(nil), "polaris.v1.ConfigBatchWriteResponse")
	proto.RegisterType((*ConfigBatchQueryResponse)(nil), "polaris.v1.ConfigBatchQueryResponse")
	proto.RegisterType((*ConfigClientResponse)(nil), "polaris.v1.ConfigClientResponse")
}

func init() {
	proto.RegisterFile("config_file_response.proto", fileDescriptor_config_file_response_faf971ae42ada746)
}

var fileDescriptor_config_file_response_faf971ae42ada746 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0xd6, 0x75, 0xe0, 0x4a, 0xc0, 0x0c, 0x02, 0xab, 0xc0, 0x54, 0x55, 0x5c, 0xec,
	0xca, 0x61, 0x9d, 0x34, 0xed, 0x12, 0xb5, 0x82, 0x31, 0x21, 0xd0, 0xc8, 0xf8, 0x23, 0x71, 0x33,
	0xdc, 0xec, 0x24, 0xb5, 0x94, 0xd8, 0x96, 0xed, 0x0c, 0x8d, 0x07, 0xe1, 0x69, 0xf6, 0x48, 0x5c,
	0xf1, 0x04, 0xa8, 0x76, 0x42, 0xd2, 0xa5, 0xed, 0xb8, 0xaa, 0xb8, 0xec, 0xc9, 0xf7, 0xfd, 0x9c,
	0x2f, 0x3e, 0xe7, 0x14, 0xf5, 0x22, 0x29, 0x62, 0x9e, 0x9c, 0xc5, 0x3c, 0x85, 0x33, 0x0d, 0x46,
	0x49, 0x61, 0x80, 0x2a, 0x2d, 0xad, 0xc4, 0x48, 0xc9, 0x94, 0x69, 0x6e, 0xe8, 0xc5, 0x5e, 0x6f,
	0x27, 0x91, 0x32, 0x49, 0x21, 0x70, 0x4f, 0x26, 0x79, 0x1c, 0x7c, 0xd7, 0x4c, 0x29, 0xd0, 0xc6,
	0x6b, 0x7b, 0xdb, 0x35, 0x8e, 0x2f, 0x0d, 0x7e, 0xa0, 0x87, 0x63, 0x57, 0x3c, 0xe5, 0x99, 0x4a,
	0x21, 0x2c, 0xe0, 0xf8, 0x05, 0x6a, 0x47, 0xf2, 0x1c, 0x48, 0xab, 0xdf, 0xda, 0xed, 0x0e, 0x9f,
	0x52, 0x4f, 0xa6, 0x25, 0x99, 0x7e, 0x3a, 0x16, 0x76, 0x7f, 0xf8, 0x99, 0xa5, 0x39, 0x84, 0x4e,
	0x39, 0x73, 0x70, 0x11, 0x4b, 0x72, 0x6b, 0x89, 0xe3, 0xd4, 0x6a, 0x2e, 0x92, 0xc2, 0x31, 0x53,
	0x0e, 0x7e, 0x6f, 0xa0, 0xbb, 0xfe, 0xf0, 0x75, 0x1e, 0x8b, 0x5f, 0xa1, 0x7b, 0xfe, 0x3b, 0xbc,
	0xe6, 0x29, 0x1c, 0x69, 0x99, 0x2b, 0xb2, 0xe1, 0xcc, 0x4f, 0x68, 0xf5, 0x2d, 0xe9, 0x78, 0x5e,
	0x12, 0x5e, 0xf7, 0xe0, 0x03, 0x84, 0xaa, 0x12, 0x69, 0x3b, 0xc2, 0xa3, 0xc5, 0x84, 0xb0, 0xa6,
	0xc4, 0x6f, 0xd1, 0x76, 0xf5, 0x2b, 0x84, 0x14, 0x98, 0x01, 0xb2, 0xe9, 0xec, 0xcf, 0x96, 0xd8,
	0xbd, 0x28, 0x6c, 0xfa, 0xf0, 0x37, 0x44, 0x1a, 0xc5, 0x37, 0xdc, 0x58, 0xa9, 0x2f, 0x49, 0xc7,
	0x31, 0x9f, 0xaf, 0x64, 0x16, 0xda, 0x70, 0x29, 0x05, 0xbf, 0x47, 0xb8, 0x7a, 0xf6, 0x11, 0x32,
	0x95, 0x32, 0x0b, 0x64, 0xcb, 0xb1, 0x77, 0x16, 0xb3, 0x4b, 0x55, 0xb8, 0xc0, 0x39, 0xf8, 0xd5,
	0x42, 0xc4, 0x4b, 0x47, 0xcc, 0x46, 0xd3, 0x2f, 0x9a, 0xdb, 0xb5, 0x76, 0x1d, 0x1e, 0xa2, 0x4d,
	0x2b, 0x2d, 0x4b, 0x8b, 0x4b, 0x5f, 0x7d, 0x88, 0x97, 0xe2, 0x43, 0x74, 0xa7, 0x1c, 0x3b, 0x43,
	0xda, 0xfd, 0x8d, 0xdd, 0xee, 0xb0, 0xd7, 0xcc, 0x5e, 0xc6, 0x08, 0x2b, 0xf1, 0xe0, 0xaa, 0x3d,
	0x17, 0xf7, 0x43, 0x0e, 0xfa, 0xf2, 0xbf, 0x8f, 0x7b, 0x84, 0xee, 0x5f, 0xeb, 0xf6, 0x32, 0xf5,
	0xca, 0x11, 0x69, 0x98, 0xf0, 0x21, 0xea, 0x56, 0x35, 0x43, 0x36, 0x1d, 0x63, 0xd9, 0x90, 0xd4,
	0xa5, 0xf8, 0x5d, 0xbd, 0xed, 0x8a, 0x96, 0x34, 0xa4, 0xe3, 0x00, 0x37, 0x8c, 0xc9, 0x02, 0x23,
	0x3e, 0x2f, 0x77, 0x68, 0xa3, 0xc3, 0x39, 0x18, 0xb2, 0xe5, 0xb0, 0xff, 0x36, 0x29, 0x2b, 0x38,
	0xf8, 0x04, 0x3d, 0x68, 0x76, 0xbc, 0x21, 0xb7, 0x1d, 0xfe, 0xa6, 0x61, 0x59, 0x64, 0x1d, 0x5c,
	0xb5, 0xca, 0xfd, 0x3c, 0x4e, 0x39, 0x08, 0xbb, 0xd6, 0xd6, 0x79, 0x39, 0xb7, 0xe1, 0x7c, 0xff,
	0xf4, 0xe7, 0x52, 0xb8, 0x77, 0xaa, 0xb2, 0x1c, 0x8b, 0x58, 0xd6, 0x77, 0xdd, 0xe8, 0x67, 0x0b,
	0x1d, 0x44, 0x32, 0xa3, 0x16, 0x44, 0x04, 0xc2, 0xfe, 0xf5, 0x1a, 0x05, 0x11, 0x8f, 0x79, 0xc4,
	0x2c, 0x97, 0x82, 0x32, 0xc5, 0x67, 0x34, 0xef, 0xa3, 0x19, 0x13, 0x2c, 0x81, 0xd1, 0xe3, 0xfa,
	0x0d, 0xf8, 0xd0, 0x27, 0xb3, 0x57, 0xfd, 0x3a, 0x4e, 0xb8, 0x9d, 0xe6, 0x13, 0x1a, 0xc9, 0x2c,
	0x28, 0x78, 0x19, 0x98, 0x69, 0x30, 0xc7, 0x0c, 0x8c, 0xcc, 0x75, 0x04, 0x41, 0x22, 0x03, 0xa6,
	0x78, 0x70, 0xb1, 0x17, 0x14, 0x7f, 0x7d, 0x9e, 0x3e, 0xe9, 0xb8, 0xd8, 0xfb, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x3e, 0x47, 0x02, 0x5a, 0x07, 0x00, 0x00,
}