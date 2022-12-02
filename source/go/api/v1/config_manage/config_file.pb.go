// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file.proto

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

type ConfigFileGroup struct {
	Id                   *wrapperspb.UInt64Value   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Comment              *wrapperspb.StringValue   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue   `protobuf:"bytes,7,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue   `protobuf:"bytes,8,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	FileCount            *wrapperspb.UInt64Value   `protobuf:"bytes,9,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	UserIds              []*wrapperspb.StringValue `protobuf:"bytes,10,rep,name=user_ids,proto3" json:"user_ids,omitempty"`
	GroupIds             []*wrapperspb.StringValue `protobuf:"bytes,11,rep,name=group_ids,proto3" json:"group_ids,omitempty"`
	RemoveUserIds        []*wrapperspb.StringValue `protobuf:"bytes,13,rep,name=remove_user_ids,proto3" json:"remove_user_ids,omitempty"`
	RemoveGroupIds       []*wrapperspb.StringValue `protobuf:"bytes,14,rep,name=remove_group_ids,proto3" json:"remove_group_ids,omitempty"`
	Editable             *wrapperspb.BoolValue     `protobuf:"bytes,15,opt,name=editable,proto3" json:"editable,omitempty"`
	Owner                *wrapperspb.StringValue   `protobuf:"bytes,16,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigFileGroup) Reset()         { *m = ConfigFileGroup{} }
func (m *ConfigFileGroup) String() string { return proto.CompactTextString(m) }
func (*ConfigFileGroup) ProtoMessage()    {}
func (*ConfigFileGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{0}
}
func (m *ConfigFileGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileGroup.Unmarshal(m, b)
}
func (m *ConfigFileGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileGroup.Marshal(b, m, deterministic)
}
func (dst *ConfigFileGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileGroup.Merge(dst, src)
}
func (m *ConfigFileGroup) XXX_Size() int {
	return xxx_messageInfo_ConfigFileGroup.Size(m)
}
func (m *ConfigFileGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileGroup proto.InternalMessageInfo

func (m *ConfigFileGroup) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileGroup) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileGroup) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileGroup) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFileGroup) GetFileCount() *wrapperspb.UInt64Value {
	if m != nil {
		return m.FileCount
	}
	return nil
}

func (m *ConfigFileGroup) GetUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveUserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveGroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetEditable() *wrapperspb.BoolValue {
	if m != nil {
		return m.Editable
	}
	return nil
}

func (m *ConfigFileGroup) GetOwner() *wrapperspb.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

type ConfigFile struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Status               *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag        `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	ReleaseTime          *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	ReleaseBy            *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=release_by,json=releaseBy,proto3" json:"release_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFile) Reset()         { *m = ConfigFile{} }
func (m *ConfigFile) String() string { return proto.CompactTextString(m) }
func (*ConfigFile) ProtoMessage()    {}
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{1}
}
func (m *ConfigFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFile.Unmarshal(m, b)
}
func (m *ConfigFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFile.Marshal(b, m, deterministic)
}
func (dst *ConfigFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFile.Merge(dst, src)
}
func (m *ConfigFile) XXX_Size() int {
	return xxx_messageInfo_ConfigFile.Size(m)
}
func (m *ConfigFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFile.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFile proto.InternalMessageInfo

func (m *ConfigFile) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFile) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFile) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFile) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFile) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFile) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFile) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFile) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFile) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFile) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFile) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFile) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFile) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFile) GetReleaseTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseTime
	}
	return nil
}

func (m *ConfigFile) GetReleaseBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseBy
	}
	return nil
}

type ConfigFileTag struct {
	Key                  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTag) Reset()         { *m = ConfigFileTag{} }
func (m *ConfigFileTag) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTag) ProtoMessage()    {}
func (*ConfigFileTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{2}
}
func (m *ConfigFileTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTag.Unmarshal(m, b)
}
func (m *ConfigFileTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTag.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTag.Merge(dst, src)
}
func (m *ConfigFileTag) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTag.Size(m)
}
func (m *ConfigFileTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTag.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTag proto.InternalMessageInfo

func (m *ConfigFileTag) GetKey() *wrapperspb.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ConfigFileTag) GetValue() *wrapperspb.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ConfigFileRelease struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	Version              *wrapperspb.UInt64Value `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileRelease) Reset()         { *m = ConfigFileRelease{} }
func (m *ConfigFileRelease) String() string { return proto.CompactTextString(m) }
func (*ConfigFileRelease) ProtoMessage()    {}
func (*ConfigFileRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{3}
}
func (m *ConfigFileRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileRelease.Unmarshal(m, b)
}
func (m *ConfigFileRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileRelease.Marshal(b, m, deterministic)
}
func (dst *ConfigFileRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileRelease.Merge(dst, src)
}
func (m *ConfigFileRelease) XXX_Size() int {
	return xxx_messageInfo_ConfigFileRelease.Size(m)
}
func (m *ConfigFileRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileRelease proto.InternalMessageInfo

func (m *ConfigFileRelease) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileRelease) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileRelease) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileRelease) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileRelease) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileRelease) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileRelease) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileRelease) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileRelease) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileReleaseHistory struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=md5,proto3" json:"md5,omitempty"`
	Type                 *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag        `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,16,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileReleaseHistory) Reset()         { *m = ConfigFileReleaseHistory{} }
func (m *ConfigFileReleaseHistory) String() string { return proto.CompactTextString(m) }
func (*ConfigFileReleaseHistory) ProtoMessage()    {}
func (*ConfigFileReleaseHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{4}
}
func (m *ConfigFileReleaseHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileReleaseHistory.Unmarshal(m, b)
}
func (m *ConfigFileReleaseHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileReleaseHistory.Marshal(b, m, deterministic)
}
func (dst *ConfigFileReleaseHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileReleaseHistory.Merge(dst, src)
}
func (m *ConfigFileReleaseHistory) XXX_Size() int {
	return xxx_messageInfo_ConfigFileReleaseHistory.Size(m)
}
func (m *ConfigFileReleaseHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileReleaseHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileReleaseHistory proto.InternalMessageInfo

func (m *ConfigFileReleaseHistory) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetType() *wrapperspb.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileTemplate struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTemplate) Reset()         { *m = ConfigFileTemplate{} }
func (m *ConfigFileTemplate) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTemplate) ProtoMessage()    {}
func (*ConfigFileTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{5}
}
func (m *ConfigFileTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTemplate.Unmarshal(m, b)
}
func (m *ConfigFileTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTemplate.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTemplate.Merge(dst, src)
}
func (m *ConfigFileTemplate) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTemplate.Size(m)
}
func (m *ConfigFileTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTemplate proto.InternalMessageInfo

func (m *ConfigFileTemplate) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileTemplate) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileTemplate) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileTemplate) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileTemplate) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ClientConfigFileInfo struct {
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Version              *wrapperspb.UInt64Value `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientConfigFileInfo) Reset()         { *m = ClientConfigFileInfo{} }
func (m *ClientConfigFileInfo) String() string { return proto.CompactTextString(m) }
func (*ClientConfigFileInfo) ProtoMessage()    {}
func (*ClientConfigFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{6}
}
func (m *ClientConfigFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfigFileInfo.Unmarshal(m, b)
}
func (m *ClientConfigFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfigFileInfo.Marshal(b, m, deterministic)
}
func (dst *ClientConfigFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfigFileInfo.Merge(dst, src)
}
func (m *ClientConfigFileInfo) XXX_Size() int {
	return xxx_messageInfo_ClientConfigFileInfo.Size(m)
}
func (m *ClientConfigFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfigFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfigFileInfo proto.InternalMessageInfo

func (m *ClientConfigFileInfo) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ClientConfigFileInfo) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ClientConfigFileInfo) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ClientConfigFileInfo) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ClientConfigFileInfo) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ClientConfigFileInfo) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

type ClientWatchConfigFileRequest struct {
	ClientIp             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ServiceName          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	WatchFiles           []*ClientConfigFileInfo `protobuf:"bytes,3,rep,name=watch_files,json=watchFiles,proto3" json:"watch_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientWatchConfigFileRequest) Reset()         { *m = ClientWatchConfigFileRequest{} }
func (m *ClientWatchConfigFileRequest) String() string { return proto.CompactTextString(m) }
func (*ClientWatchConfigFileRequest) ProtoMessage()    {}
func (*ClientWatchConfigFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_6fc466f54e79f210, []int{7}
}
func (m *ClientWatchConfigFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Unmarshal(m, b)
}
func (m *ClientWatchConfigFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Marshal(b, m, deterministic)
}
func (dst *ClientWatchConfigFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientWatchConfigFileRequest.Merge(dst, src)
}
func (m *ClientWatchConfigFileRequest) XXX_Size() int {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Size(m)
}
func (m *ClientWatchConfigFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientWatchConfigFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientWatchConfigFileRequest proto.InternalMessageInfo

func (m *ClientWatchConfigFileRequest) GetClientIp() *wrapperspb.StringValue {
	if m != nil {
		return m.ClientIp
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetServiceName() *wrapperspb.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetWatchFiles() []*ClientConfigFileInfo {
	if m != nil {
		return m.WatchFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigFileGroup)(nil), "v1.ConfigFileGroup")
	proto.RegisterType((*ConfigFile)(nil), "v1.ConfigFile")
	proto.RegisterType((*ConfigFileTag)(nil), "v1.ConfigFileTag")
	proto.RegisterType((*ConfigFileRelease)(nil), "v1.ConfigFileRelease")
	proto.RegisterType((*ConfigFileReleaseHistory)(nil), "v1.ConfigFileReleaseHistory")
	proto.RegisterType((*ConfigFileTemplate)(nil), "v1.ConfigFileTemplate")
	proto.RegisterType((*ClientConfigFileInfo)(nil), "v1.ClientConfigFileInfo")
	proto.RegisterType((*ClientWatchConfigFileRequest)(nil), "v1.ClientWatchConfigFileRequest")
}

func init() { proto.RegisterFile("config_file.proto", fileDescriptor_config_file_6fc466f54e79f210) }

var fileDescriptor_config_file_6fc466f54e79f210 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x5d, 0x8b, 0x23, 0x45,
	0x14, 0x25, 0x5f, 0x9d, 0xe4, 0x66, 0x3e, 0x0b, 0x1f, 0x9a, 0x65, 0x91, 0x25, 0x20, 0xf8, 0x20,
	0xdd, 0x3b, 0xe3, 0x38, 0x38, 0x8a, 0x08, 0x19, 0x58, 0x77, 0x5e, 0x44, 0xe2, 0xaa, 0xe0, 0x4b,
	0xa8, 0x74, 0x6e, 0x7a, 0x0a, 0xbb, 0xbb, 0xda, 0xaa, 0xea, 0x0c, 0xfd, 0x13, 0x7c, 0xd9, 0xff,
	0xe6, 0x83, 0x3f, 0xc0, 0x57, 0x7f, 0x84, 0x48, 0x55, 0x75, 0xec, 0xcc, 0x8e, 0xcb, 0x54, 0xba,
	0x11, 0x44, 0x7d, 0x1a, 0x92, 0x9c, 0x73, 0x6f, 0xe5, 0xf6, 0x39, 0xe7, 0x66, 0x0a, 0x4e, 0x23,
	0x9e, 0xad, 0x59, 0xbc, 0x58, 0xb3, 0x04, 0x83, 0x5c, 0x70, 0xc5, 0x49, 0x77, 0x73, 0xf6, 0xe4,
	0xdd, 0x98, 0xf3, 0x38, 0xc1, 0xd0, 0xbc, 0xb3, 0x2c, 0xd6, 0xe1, 0x9d, 0xa0, 0x79, 0x8e, 0x42,
	0x5a, 0xcc, 0xf4, 0xa7, 0x21, 0x1c, 0x5f, 0x1b, 0xe6, 0x0b, 0x96, 0xe0, 0x17, 0x82, 0x17, 0x39,
	0xf9, 0x00, 0xba, 0x6c, 0xe5, 0x77, 0x9e, 0x75, 0xde, 0x9f, 0x9c, 0x3f, 0x0d, 0x6c, 0x81, 0x60,
	0x5b, 0x20, 0xf8, 0xe6, 0x26, 0x53, 0x97, 0x17, 0xdf, 0xd2, 0xa4, 0xc0, 0x79, 0x97, 0xad, 0xc8,
	0x73, 0xe8, 0x67, 0x34, 0x45, 0xbf, 0xfb, 0x16, 0xfc, 0xd7, 0x4a, 0xb0, 0x2c, 0xb6, 0x78, 0x83,
	0x24, 0x9f, 0xc0, 0x58, 0xff, 0x95, 0x39, 0x8d, 0xd0, 0xef, 0x39, 0xd0, 0x6a, 0x38, 0xb9, 0x84,
	0x61, 0xc4, 0xd3, 0x14, 0x33, 0xe5, 0xf7, 0x1d, 0x98, 0x5b, 0x30, 0xf9, 0x0c, 0x26, 0x91, 0x40,
	0xaa, 0x70, 0xa1, 0x58, 0x8a, 0xfe, 0xc0, 0x81, 0x0b, 0x96, 0xf0, 0x8a, 0xa5, 0x48, 0xae, 0x60,
	0x5c, 0xd1, 0x97, 0xa5, 0xef, 0x39, 0x90, 0x47, 0x16, 0x3e, 0x2b, 0x75, 0xe7, 0x94, 0xaf, 0xd8,
	0xba, 0xb4, 0x9d, 0x87, 0x2e, 0x9d, 0x2d, 0x61, 0xdb, 0xb9, 0xa2, 0x2f, 0x4b, 0x7f, 0xe4, 0xd2,
	0xd9, 0xc2, 0x67, 0xa5, 0x9e, 0xb3, 0x56, 0xc3, 0x35, 0x2f, 0x32, 0xe5, 0x8f, 0x1d, 0x1e, 0x67,
	0x0d, 0x27, 0x1f, 0xc3, 0xa8, 0x90, 0x28, 0x16, 0x6c, 0x25, 0x7d, 0x78, 0xd6, 0x7b, 0xbc, 0xeb,
	0x16, 0xad, 0xbb, 0xc6, 0x5a, 0x46, 0x86, 0x3a, 0x71, 0xa0, 0xd6, 0x70, 0xf2, 0x02, 0x8e, 0x05,
	0xa6, 0x7c, 0x83, 0x8b, 0x3f, 0x9b, 0x1f, 0x3a, 0x54, 0x78, 0x93, 0x44, 0x5e, 0xc2, 0x49, 0xf5,
	0x56, 0x7d, 0x94, 0x23, 0x87, 0x42, 0x0f, 0x58, 0xe4, 0x12, 0x46, 0xb8, 0x62, 0x8a, 0x2e, 0x13,
	0xf4, 0x8f, 0xcd, 0x08, 0x9f, 0x3c, 0xa8, 0x30, 0xe3, 0x3c, 0xa9, 0xa6, 0xb0, 0xc5, 0x92, 0x73,
	0x18, 0xf0, 0xbb, 0x0c, 0x85, 0x7f, 0xe2, 0xf0, 0xc8, 0x2c, 0x74, 0xfa, 0xb3, 0x07, 0x50, 0x7b,
	0xf1, 0x1f, 0x6d, 0xc3, 0x73, 0x18, 0x98, 0x19, 0x39, 0x99, 0xd0, 0x42, 0xad, 0x75, 0x33, 0xa5,
	0xad, 0x3b, 0x70, 0xb3, 0xae, 0x01, 0x93, 0x0b, 0xf0, 0xd6, 0x5c, 0xa4, 0x54, 0x39, 0x19, 0xaf,
	0xc2, 0xee, 0x06, 0xc5, 0x70, 0x9f, 0xa0, 0xb8, 0x00, 0x4f, 0x2a, 0xaa, 0x0a, 0xe9, 0x64, 0xb6,
	0x0a, 0x4b, 0xde, 0x83, 0xbe, 0xa2, 0xb1, 0xf4, 0xc7, 0x46, 0x64, 0xa7, 0xc1, 0xe6, 0x2c, 0xa8,
	0x9f, 0xe4, 0x2b, 0x1a, 0xcf, 0xcd, 0xc7, 0x6f, 0xa6, 0x10, 0xb4, 0x49, 0xa1, 0x49, 0x9b, 0x14,
	0x3a, 0x68, 0x93, 0x42, 0x87, 0x7b, 0xa5, 0xd0, 0xe7, 0x70, 0x20, 0x30, 0x41, 0x2a, 0xab, 0x2f,
	0x7d, 0xe4, 0xc0, 0x9e, 0x54, 0x0c, 0xd3, 0xfb, 0x53, 0x80, 0x6d, 0x81, 0x65, 0x59, 0x99, 0xf0,
	0x11, 0xa1, 0x56, 0xf8, 0x59, 0x39, 0x95, 0x70, 0x78, 0xef, 0x41, 0x90, 0x00, 0x7a, 0x3f, 0x60,
	0xf9, 0x56, 0x5b, 0xed, 0x96, 0xd1, 0x40, 0xad, 0xf4, 0x8d, 0x7e, 0xe5, 0x64, 0x2c, 0x0b, 0x9d,
	0xfe, 0x36, 0x80, 0xd3, 0xba, 0xeb, 0xdc, 0x1e, 0xe6, 0x5f, 0xe7, 0xe7, 0x2b, 0xbb, 0x5e, 0x16,
	0xe6, 0x98, 0x2e, 0x8e, 0x1e, 0x69, 0xf8, 0x97, 0xfa, 0xa8, 0x3b, 0x51, 0xe0, 0xed, 0x13, 0x05,
	0x4d, 0x4d, 0x1d, 0x40, 0x2f, 0x5d, 0x7d, 0xe4, 0xe4, 0x68, 0x0d, 0xd4, 0x7d, 0x36, 0x28, 0x24,
	0xe3, 0x99, 0xd3, 0xde, 0xdc, 0x82, 0xff, 0x8b, 0xfe, 0x9e, 0xfe, 0xee, 0x81, 0xff, 0x40, 0xec,
	0x2f, 0x99, 0x54, 0x5c, 0x94, 0xff, 0x6b, 0xbe, 0xbd, 0xe6, 0xeb, 0xf5, 0x37, 0x6c, 0xb6, 0xfe,
	0x46, 0x0d, 0x9c, 0x32, 0x76, 0x75, 0xca, 0x73, 0xe8, 0xab, 0x32, 0x77, 0x93, 0xba, 0x41, 0xee,
	0x2c, 0xd8, 0x49, 0x83, 0x05, 0x7b, 0xb0, 0xd7, 0x82, 0x3d, 0x6c, 0x63, 0xc0, 0xa3, 0x36, 0x06,
	0x3c, 0x6e, 0x63, 0xc0, 0x93, 0xbd, 0x0c, 0xf8, 0xba, 0x0f, 0x64, 0x67, 0x16, 0x98, 0xe6, 0x09,
	0x55, 0x7f, 0xff, 0xba, 0xd9, 0xd1, 0x73, 0xaf, 0x99, 0x9e, 0xfb, 0xcd, 0xf4, 0x3c, 0x68, 0xf1,
	0x7f, 0x9f, 0xd7, 0x46, 0x10, 0xc3, 0x36, 0x82, 0x18, 0xb5, 0x11, 0xc4, 0x78, 0x2f, 0x41, 0xfc,
	0xda, 0x85, 0x77, 0xae, 0x13, 0x86, 0x99, 0xaa, 0x65, 0x71, 0x93, 0xad, 0xf9, 0xfd, 0xb4, 0xec,
	0x34, 0x4c, 0xcb, 0x6e, 0xc3, 0xb4, 0xec, 0x35, 0x4d, 0xcb, 0xfe, 0x9e, 0xbf, 0x10, 0xb6, 0x9b,
	0x7b, 0xb0, 0xcf, 0xe6, 0xae, 0x72, 0xcf, 0x73, 0xcc, 0xbd, 0xe9, 0x2f, 0x1d, 0x78, 0x6a, 0x67,
	0xfc, 0x1d, 0x55, 0xd1, 0xed, 0xee, 0x02, 0xfc, 0xb1, 0x40, 0xa9, 0x8c, 0x72, 0xcc, 0xe7, 0x0b,
	0x96, 0x3b, 0xcd, 0x7a, 0x64, 0xe1, 0x37, 0xb9, 0xfe, 0xc5, 0x2c, 0x51, 0x6c, 0x58, 0x54, 0x4d,
	0xce, 0x65, 0xe2, 0x93, 0x8a, 0x61, 0x86, 0x77, 0x05, 0x93, 0x3b, 0x7d, 0x2a, 0x73, 0x19, 0x24,
	0xfd, 0x9e, 0xc9, 0x4c, 0xdf, 0x64, 0xe6, 0x5f, 0xc8, 0x62, 0x0e, 0x06, 0xac, 0x5f, 0xca, 0xd9,
	0xeb, 0x0e, 0x5c, 0x46, 0x3c, 0x0d, 0x14, 0x66, 0x11, 0x66, 0x2a, 0xc8, 0x79, 0x42, 0x05, 0x93,
	0x81, 0xcc, 0x31, 0x62, 0x6b, 0x16, 0x51, 0xc5, 0x78, 0x16, 0xd0, 0x9c, 0xe9, 0x6a, 0xf6, 0xca,
	0x29, 0x48, 0x69, 0x46, 0x63, 0x9c, 0xed, 0xdc, 0x23, 0x7d, 0xa5, 0x8f, 0xf8, 0xfd, 0x75, 0xcc,
	0xd4, 0x6d, 0xb1, 0x0c, 0x22, 0x9e, 0x86, 0x55, 0x9d, 0x14, 0xe5, 0x6d, 0x78, 0xaf, 0x56, 0x28,
	0x79, 0x21, 0x22, 0x0c, 0x63, 0x1e, 0xd2, 0x9c, 0x85, 0x9b, 0xb3, 0xb0, 0xba, 0xc8, 0xb2, 0x55,
	0x97, 0x9e, 0xf9, 0xba, 0x1f, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xb9, 0x3d, 0xf1, 0xe0,
	0x12, 0x00, 0x00,
}
