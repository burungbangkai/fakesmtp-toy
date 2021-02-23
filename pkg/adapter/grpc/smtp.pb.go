// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smtp.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RawEmailReceivedEvent struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InboxName            string   `protobuf:"bytes,2,opt,name=inbox_name,json=inboxName,proto3" json:"inbox_name,omitempty"`
	Raw                  []byte   `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawEmailReceivedEvent) Reset()         { *m = RawEmailReceivedEvent{} }
func (m *RawEmailReceivedEvent) String() string { return proto.CompactTextString(m) }
func (*RawEmailReceivedEvent) ProtoMessage()    {}
func (*RawEmailReceivedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{0}
}

func (m *RawEmailReceivedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawEmailReceivedEvent.Unmarshal(m, b)
}
func (m *RawEmailReceivedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawEmailReceivedEvent.Marshal(b, m, deterministic)
}
func (m *RawEmailReceivedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawEmailReceivedEvent.Merge(m, src)
}
func (m *RawEmailReceivedEvent) XXX_Size() int {
	return xxx_messageInfo_RawEmailReceivedEvent.Size(m)
}
func (m *RawEmailReceivedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RawEmailReceivedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RawEmailReceivedEvent proto.InternalMessageInfo

func (m *RawEmailReceivedEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RawEmailReceivedEvent) GetInboxName() string {
	if m != nil {
		return m.InboxName
	}
	return ""
}

func (m *RawEmailReceivedEvent) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

type GetUserInboxsReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInboxsReq) Reset()         { *m = GetUserInboxsReq{} }
func (m *GetUserInboxsReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInboxsReq) ProtoMessage()    {}
func (*GetUserInboxsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{1}
}

func (m *GetUserInboxsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInboxsReq.Unmarshal(m, b)
}
func (m *GetUserInboxsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInboxsReq.Marshal(b, m, deterministic)
}
func (m *GetUserInboxsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInboxsReq.Merge(m, src)
}
func (m *GetUserInboxsReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInboxsReq.Size(m)
}
func (m *GetUserInboxsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInboxsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInboxsReq proto.InternalMessageInfo

func (m *GetUserInboxsReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UserInbox struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInbox) Reset()         { *m = UserInbox{} }
func (m *UserInbox) String() string { return proto.CompactTextString(m) }
func (*UserInbox) ProtoMessage()    {}
func (*UserInbox) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{2}
}

func (m *UserInbox) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInbox.Unmarshal(m, b)
}
func (m *UserInbox) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInbox.Marshal(b, m, deterministic)
}
func (m *UserInbox) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInbox.Merge(m, src)
}
func (m *UserInbox) XXX_Size() int {
	return xxx_messageInfo_UserInbox.Size(m)
}
func (m *UserInbox) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInbox.DiscardUnknown(m)
}

var xxx_messageInfo_UserInbox proto.InternalMessageInfo

func (m *UserInbox) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInbox) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInbox) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type GetUserInboxsResp struct {
	UserId               string       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Inboxs               []*UserInbox `protobuf:"bytes,2,rep,name=inboxs,proto3" json:"inboxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserInboxsResp) Reset()         { *m = GetUserInboxsResp{} }
func (m *GetUserInboxsResp) String() string { return proto.CompactTextString(m) }
func (*GetUserInboxsResp) ProtoMessage()    {}
func (*GetUserInboxsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{3}
}

func (m *GetUserInboxsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInboxsResp.Unmarshal(m, b)
}
func (m *GetUserInboxsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInboxsResp.Marshal(b, m, deterministic)
}
func (m *GetUserInboxsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInboxsResp.Merge(m, src)
}
func (m *GetUserInboxsResp) XXX_Size() int {
	return xxx_messageInfo_GetUserInboxsResp.Size(m)
}
func (m *GetUserInboxsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInboxsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInboxsResp proto.InternalMessageInfo

func (m *GetUserInboxsResp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetUserInboxsResp) GetInboxs() []*UserInbox {
	if m != nil {
		return m.Inboxs
	}
	return nil
}

type DeleteUserInboxReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InboxName            string   `protobuf:"bytes,2,opt,name=inbox_name,json=inboxName,proto3" json:"inbox_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserInboxReq) Reset()         { *m = DeleteUserInboxReq{} }
func (m *DeleteUserInboxReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserInboxReq) ProtoMessage()    {}
func (*DeleteUserInboxReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{4}
}

func (m *DeleteUserInboxReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserInboxReq.Unmarshal(m, b)
}
func (m *DeleteUserInboxReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserInboxReq.Marshal(b, m, deterministic)
}
func (m *DeleteUserInboxReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserInboxReq.Merge(m, src)
}
func (m *DeleteUserInboxReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserInboxReq.Size(m)
}
func (m *DeleteUserInboxReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserInboxReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserInboxReq proto.InternalMessageInfo

func (m *DeleteUserInboxReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteUserInboxReq) GetInboxName() string {
	if m != nil {
		return m.InboxName
	}
	return ""
}

type AddUserInboxReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserInboxReq) Reset()         { *m = AddUserInboxReq{} }
func (m *AddUserInboxReq) String() string { return proto.CompactTextString(m) }
func (*AddUserInboxReq) ProtoMessage()    {}
func (*AddUserInboxReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{5}
}

func (m *AddUserInboxReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserInboxReq.Unmarshal(m, b)
}
func (m *AddUserInboxReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserInboxReq.Marshal(b, m, deterministic)
}
func (m *AddUserInboxReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserInboxReq.Merge(m, src)
}
func (m *AddUserInboxReq) XXX_Size() int {
	return xxx_messageInfo_AddUserInboxReq.Size(m)
}
func (m *AddUserInboxReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserInboxReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserInboxReq proto.InternalMessageInfo

func (m *AddUserInboxReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type AddUserInboxResp struct {
	UserId               string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Inbox                *UserInbox `protobuf:"bytes,2,opt,name=inbox,proto3" json:"inbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddUserInboxResp) Reset()         { *m = AddUserInboxResp{} }
func (m *AddUserInboxResp) String() string { return proto.CompactTextString(m) }
func (*AddUserInboxResp) ProtoMessage()    {}
func (*AddUserInboxResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{6}
}

func (m *AddUserInboxResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserInboxResp.Unmarshal(m, b)
}
func (m *AddUserInboxResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserInboxResp.Marshal(b, m, deterministic)
}
func (m *AddUserInboxResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserInboxResp.Merge(m, src)
}
func (m *AddUserInboxResp) XXX_Size() int {
	return xxx_messageInfo_AddUserInboxResp.Size(m)
}
func (m *AddUserInboxResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserInboxResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserInboxResp proto.InternalMessageInfo

func (m *AddUserInboxResp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddUserInboxResp) GetInbox() *UserInbox {
	if m != nil {
		return m.Inbox
	}
	return nil
}

type UpdateUserInboxReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InboxName            string   `protobuf:"bytes,2,opt,name=inbox_name,json=inboxName,proto3" json:"inbox_name,omitempty"`
	NewInboxName         string   `protobuf:"bytes,3,opt,name=new_inbox_name,json=newInboxName,proto3" json:"new_inbox_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInboxReq) Reset()         { *m = UpdateUserInboxReq{} }
func (m *UpdateUserInboxReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInboxReq) ProtoMessage()    {}
func (*UpdateUserInboxReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{7}
}

func (m *UpdateUserInboxReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInboxReq.Unmarshal(m, b)
}
func (m *UpdateUserInboxReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInboxReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserInboxReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInboxReq.Merge(m, src)
}
func (m *UpdateUserInboxReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInboxReq.Size(m)
}
func (m *UpdateUserInboxReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInboxReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInboxReq proto.InternalMessageInfo

func (m *UpdateUserInboxReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateUserInboxReq) GetInboxName() string {
	if m != nil {
		return m.InboxName
	}
	return ""
}

func (m *UpdateUserInboxReq) GetNewInboxName() string {
	if m != nil {
		return m.NewInboxName
	}
	return ""
}

type UpdateUserInboxResp struct {
	UserId               string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Inbox                *UserInbox `protobuf:"bytes,2,opt,name=inbox,proto3" json:"inbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateUserInboxResp) Reset()         { *m = UpdateUserInboxResp{} }
func (m *UpdateUserInboxResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInboxResp) ProtoMessage()    {}
func (*UpdateUserInboxResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e017a3db57e14, []int{8}
}

func (m *UpdateUserInboxResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInboxResp.Unmarshal(m, b)
}
func (m *UpdateUserInboxResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInboxResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserInboxResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInboxResp.Merge(m, src)
}
func (m *UpdateUserInboxResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInboxResp.Size(m)
}
func (m *UpdateUserInboxResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInboxResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInboxResp proto.InternalMessageInfo

func (m *UpdateUserInboxResp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateUserInboxResp) GetInbox() *UserInbox {
	if m != nil {
		return m.Inbox
	}
	return nil
}

func init() {
	proto.RegisterType((*RawEmailReceivedEvent)(nil), "stmp.RawEmailReceivedEvent")
	proto.RegisterType((*GetUserInboxsReq)(nil), "stmp.GetUserInboxsReq")
	proto.RegisterType((*UserInbox)(nil), "stmp.UserInbox")
	proto.RegisterType((*GetUserInboxsResp)(nil), "stmp.GetUserInboxsResp")
	proto.RegisterType((*DeleteUserInboxReq)(nil), "stmp.DeleteUserInboxReq")
	proto.RegisterType((*AddUserInboxReq)(nil), "stmp.AddUserInboxReq")
	proto.RegisterType((*AddUserInboxResp)(nil), "stmp.AddUserInboxResp")
	proto.RegisterType((*UpdateUserInboxReq)(nil), "stmp.UpdateUserInboxReq")
	proto.RegisterType((*UpdateUserInboxResp)(nil), "stmp.UpdateUserInboxResp")
}

func init() { proto.RegisterFile("smtp.proto", fileDescriptor_b40e017a3db57e14) }

var fileDescriptor_b40e017a3db57e14 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0x26, 0x49, 0x8d, 0x66, 0x1a, 0x4d, 0x1c, 0x69, 0x7b, 0x5e, 0x11, 0xc2, 0xa2, 0x18, 0x14,
	0x2e, 0x50, 0x9f, 0x05, 0x7f, 0x45, 0x29, 0x88, 0x0f, 0x57, 0x0e, 0xc4, 0x97, 0xb0, 0xc9, 0x8e,
	0xc7, 0x61, 0xee, 0x6e, 0xdd, 0xdd, 0xe6, 0xf4, 0x6f, 0xf1, 0x9f, 0x95, 0xdd, 0x4d, 0xcf, 0xf6,
	0x92, 0x54, 0xc1, 0xbe, 0xed, 0xce, 0xf7, 0xcd, 0xcc, 0x7e, 0xf3, 0xed, 0x00, 0xe8, 0xdc, 0xc8,
	0x48, 0xaa, 0xd2, 0x94, 0xb8, 0xa7, 0x4d, 0x2e, 0xc3, 0xe3, 0xb4, 0x2c, 0xd3, 0x25, 0x4d, 0x5c,
	0x6c, 0x7e, 0xfe, 0x75, 0x42, 0xb9, 0x34, 0x3f, 0x3d, 0x85, 0x71, 0x38, 0x88, 0x79, 0x35, 0xcd,
	0x79, 0xb6, 0x8c, 0x69, 0x41, 0xd9, 0x8a, 0xc4, 0x74, 0x45, 0x85, 0xc1, 0x23, 0xb8, 0x7d, 0xae,
	0x49, 0xcd, 0x32, 0x11, 0xb4, 0x46, 0xad, 0x71, 0x2f, 0xee, 0xda, 0xeb, 0xa9, 0xc0, 0x47, 0x00,
	0x59, 0x31, 0x2f, 0x7f, 0xcc, 0x0a, 0x9e, 0x53, 0xd0, 0x76, 0x58, 0xcf, 0x45, 0x3e, 0xf1, 0x9c,
	0x70, 0x08, 0x1d, 0xc5, 0xab, 0xa0, 0x33, 0x6a, 0x8d, 0xfb, 0xb1, 0x3d, 0xb2, 0xe7, 0x30, 0xfc,
	0x40, 0x26, 0xb1, 0xd9, 0x96, 0xa5, 0x63, 0xfa, 0xbe, 0xb3, 0x3a, 0xfb, 0x0c, 0xbd, 0x9a, 0x89,
	0x08, 0x7b, 0xae, 0x89, 0xa7, 0xb8, 0x33, 0x86, 0x70, 0x47, 0x72, 0xad, 0xab, 0x52, 0x89, 0x75,
	0xf3, 0xfa, 0x8e, 0xc7, 0xd0, 0x73, 0x55, 0x5d, 0x52, 0xc7, 0x83, 0x36, 0x60, 0x1f, 0xc6, 0x12,
	0xb8, 0xdf, 0x78, 0x86, 0x96, 0xbb, 0x55, 0x3e, 0x85, 0xae, 0xd3, 0xa4, 0x83, 0xf6, 0xa8, 0x33,
	0xde, 0x3f, 0x19, 0x44, 0x76, 0x96, 0x51, 0x9d, 0x1e, 0xaf, 0x61, 0xf6, 0x11, 0xf0, 0x1d, 0x2d,
	0xc9, 0xd0, 0x1f, 0xe8, 0x1a, 0x7d, 0x7f, 0x99, 0x1e, 0x7b, 0x06, 0x83, 0xd7, 0x42, 0xfc, 0x53,
	0x29, 0x16, 0xc3, 0xf0, 0x2a, 0xf7, 0x3a, 0x3d, 0x4f, 0xe0, 0x96, 0xeb, 0xe2, 0x5a, 0x6e, 0x91,
	0xe3, 0x51, 0xa6, 0x00, 0x13, 0x29, 0xf8, 0xcd, 0xa8, 0xc1, 0xc7, 0x70, 0xaf, 0xa0, 0x6a, 0x76,
	0x89, 0xe2, 0x4d, 0xe9, 0x17, 0x54, 0x9d, 0xd6, 0x9a, 0x13, 0x78, 0xb0, 0xd1, 0xf3, 0xff, 0xa5,
	0x9c, 0xfc, 0x6a, 0xc3, 0xfe, 0x59, 0x6e, 0xe4, 0x19, 0xa9, 0x55, 0xb6, 0x20, 0x7c, 0x0b, 0x83,
	0x86, 0x51, 0x18, 0xf8, 0xd4, 0x4d, 0xff, 0xc2, 0xc3, 0xc8, 0x2f, 0x4d, 0x74, 0xb1, 0x34, 0xd1,
	0xd4, 0x2e, 0x0d, 0xbe, 0x84, 0xfe, 0xe5, 0x99, 0xe3, 0x81, 0xaf, 0xd0, 0xf0, 0x2c, 0x3c, 0xdc,
	0x16, 0xd6, 0x12, 0x5f, 0xc1, 0xdd, 0x2b, 0x7f, 0x10, 0xd7, 0xc4, 0xe6, 0x7e, 0x84, 0x47, 0x5b,
	0xe3, 0x5a, 0xe2, 0x7b, 0x18, 0x34, 0x86, 0x75, 0xa1, 0x62, 0xd3, 0xb7, 0xf0, 0xe1, 0x0e, 0x44,
	0xcb, 0x37, 0xf8, 0x65, 0x28, 0xbf, 0xa5, 0x13, 0x2e, 0xb8, 0x34, 0xa4, 0x26, 0xa9, 0x92, 0x8b,
	0x79, 0xd7, 0x89, 0x7d, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x86, 0x2a, 0xa8, 0xd1, 0x43, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmtpServiceClient is the client API for SmtpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmtpServiceClient interface {
	DeleteUserInbox(ctx context.Context, in *DeleteUserInboxReq, opts ...grpc.CallOption) (*empty.Empty, error)
	AddUserInbox(ctx context.Context, in *AddUserInboxReq, opts ...grpc.CallOption) (*AddUserInboxResp, error)
	GetUserInboxs(ctx context.Context, in *GetUserInboxsReq, opts ...grpc.CallOption) (*GetUserInboxsResp, error)
	UpdateUserInbox(ctx context.Context, in *UpdateUserInboxReq, opts ...grpc.CallOption) (*UpdateUserInboxResp, error)
}

type smtpServiceClient struct {
	cc *grpc.ClientConn
}

func NewSmtpServiceClient(cc *grpc.ClientConn) SmtpServiceClient {
	return &smtpServiceClient{cc}
}

func (c *smtpServiceClient) DeleteUserInbox(ctx context.Context, in *DeleteUserInboxReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/stmp.SmtpService/DeleteUserInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smtpServiceClient) AddUserInbox(ctx context.Context, in *AddUserInboxReq, opts ...grpc.CallOption) (*AddUserInboxResp, error) {
	out := new(AddUserInboxResp)
	err := c.cc.Invoke(ctx, "/stmp.SmtpService/AddUserInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smtpServiceClient) GetUserInboxs(ctx context.Context, in *GetUserInboxsReq, opts ...grpc.CallOption) (*GetUserInboxsResp, error) {
	out := new(GetUserInboxsResp)
	err := c.cc.Invoke(ctx, "/stmp.SmtpService/GetUserInboxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smtpServiceClient) UpdateUserInbox(ctx context.Context, in *UpdateUserInboxReq, opts ...grpc.CallOption) (*UpdateUserInboxResp, error) {
	out := new(UpdateUserInboxResp)
	err := c.cc.Invoke(ctx, "/stmp.SmtpService/UpdateUserInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmtpServiceServer is the server API for SmtpService service.
type SmtpServiceServer interface {
	DeleteUserInbox(context.Context, *DeleteUserInboxReq) (*empty.Empty, error)
	AddUserInbox(context.Context, *AddUserInboxReq) (*AddUserInboxResp, error)
	GetUserInboxs(context.Context, *GetUserInboxsReq) (*GetUserInboxsResp, error)
	UpdateUserInbox(context.Context, *UpdateUserInboxReq) (*UpdateUserInboxResp, error)
}

// UnimplementedSmtpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmtpServiceServer struct {
}

func (*UnimplementedSmtpServiceServer) DeleteUserInbox(ctx context.Context, req *DeleteUserInboxReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserInbox not implemented")
}
func (*UnimplementedSmtpServiceServer) AddUserInbox(ctx context.Context, req *AddUserInboxReq) (*AddUserInboxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserInbox not implemented")
}
func (*UnimplementedSmtpServiceServer) GetUserInboxs(ctx context.Context, req *GetUserInboxsReq) (*GetUserInboxsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInboxs not implemented")
}
func (*UnimplementedSmtpServiceServer) UpdateUserInbox(ctx context.Context, req *UpdateUserInboxReq) (*UpdateUserInboxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInbox not implemented")
}

func RegisterSmtpServiceServer(s *grpc.Server, srv SmtpServiceServer) {
	s.RegisterService(&_SmtpService_serviceDesc, srv)
}

func _SmtpService_DeleteUserInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserInboxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).DeleteUserInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stmp.SmtpService/DeleteUserInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).DeleteUserInbox(ctx, req.(*DeleteUserInboxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmtpService_AddUserInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserInboxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).AddUserInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stmp.SmtpService/AddUserInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).AddUserInbox(ctx, req.(*AddUserInboxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmtpService_GetUserInboxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInboxsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).GetUserInboxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stmp.SmtpService/GetUserInboxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).GetUserInboxs(ctx, req.(*GetUserInboxsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmtpService_UpdateUserInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInboxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).UpdateUserInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stmp.SmtpService/UpdateUserInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).UpdateUserInbox(ctx, req.(*UpdateUserInboxReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmtpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stmp.SmtpService",
	HandlerType: (*SmtpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserInbox",
			Handler:    _SmtpService_DeleteUserInbox_Handler,
		},
		{
			MethodName: "AddUserInbox",
			Handler:    _SmtpService_AddUserInbox_Handler,
		},
		{
			MethodName: "GetUserInboxs",
			Handler:    _SmtpService_GetUserInboxs_Handler,
		},
		{
			MethodName: "UpdateUserInbox",
			Handler:    _SmtpService_UpdateUserInbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smtp.proto",
}
