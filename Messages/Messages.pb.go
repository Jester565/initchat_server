// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Messages.proto

package Messages

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

type Header struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Length               int32    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Header) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type SignUpReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{1}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{2}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpireTime           uint64   `protobuf:"varint,2,opt,name=expireTime,proto3" json:"expireTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResp) Reset()         { *m = AuthResp{} }
func (m *AuthResp) String() string { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()    {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{3}
}

func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResp.Unmarshal(m, b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
}
func (m *AuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResp.Merge(m, src)
}
func (m *AuthResp) XXX_Size() int {
	return xxx_messageInfo_AuthResp.Size(m)
}
func (m *AuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResp proto.InternalMessageInfo

func (m *AuthResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResp) GetExpireTime() uint64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

type UserSearchReq struct {
	UsernamePrefix       string   `protobuf:"bytes,1,opt,name=usernamePrefix,proto3" json:"usernamePrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSearchReq) Reset()         { *m = UserSearchReq{} }
func (m *UserSearchReq) String() string { return proto.CompactTextString(m) }
func (*UserSearchReq) ProtoMessage()    {}
func (*UserSearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{4}
}

func (m *UserSearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchReq.Unmarshal(m, b)
}
func (m *UserSearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchReq.Marshal(b, m, deterministic)
}
func (m *UserSearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchReq.Merge(m, src)
}
func (m *UserSearchReq) XXX_Size() int {
	return xxx_messageInfo_UserSearchReq.Size(m)
}
func (m *UserSearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserSearchReq proto.InternalMessageInfo

func (m *UserSearchReq) GetUsernamePrefix() string {
	if m != nil {
		return m.UsernamePrefix
	}
	return ""
}

type UserSearchResp struct {
	Usernames            []string `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSearchResp) Reset()         { *m = UserSearchResp{} }
func (m *UserSearchResp) String() string { return proto.CompactTextString(m) }
func (*UserSearchResp) ProtoMessage()    {}
func (*UserSearchResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{5}
}

func (m *UserSearchResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchResp.Unmarshal(m, b)
}
func (m *UserSearchResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchResp.Marshal(b, m, deterministic)
}
func (m *UserSearchResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchResp.Merge(m, src)
}
func (m *UserSearchResp) XXX_Size() int {
	return xxx_messageInfo_UserSearchResp.Size(m)
}
func (m *UserSearchResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSearchResp proto.InternalMessageInfo

func (m *UserSearchResp) GetUsernames() []string {
	if m != nil {
		return m.Usernames
	}
	return nil
}

type TextMessageReq struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextMessageReq) Reset()         { *m = TextMessageReq{} }
func (m *TextMessageReq) String() string { return proto.CompactTextString(m) }
func (*TextMessageReq) ProtoMessage()    {}
func (*TextMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{6}
}

func (m *TextMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextMessageReq.Unmarshal(m, b)
}
func (m *TextMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextMessageReq.Marshal(b, m, deterministic)
}
func (m *TextMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextMessageReq.Merge(m, src)
}
func (m *TextMessageReq) XXX_Size() int {
	return xxx_messageInfo_TextMessageReq.Size(m)
}
func (m *TextMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TextMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_TextMessageReq proto.InternalMessageInfo

func (m *TextMessageReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TextMessage struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Time                 uint64   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextMessage) Reset()         { *m = TextMessage{} }
func (m *TextMessage) String() string { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()    {}
func (*TextMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{7}
}

func (m *TextMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextMessage.Unmarshal(m, b)
}
func (m *TextMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextMessage.Marshal(b, m, deterministic)
}
func (m *TextMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextMessage.Merge(m, src)
}
func (m *TextMessage) XXX_Size() int {
	return xxx_messageInfo_TextMessage.Size(m)
}
func (m *TextMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TextMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TextMessage proto.InternalMessageInfo

func (m *TextMessage) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TextMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TextMessage) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type FileMessageReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMessageReq) Reset()         { *m = FileMessageReq{} }
func (m *FileMessageReq) String() string { return proto.CompactTextString(m) }
func (*FileMessageReq) ProtoMessage()    {}
func (*FileMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{8}
}

func (m *FileMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMessageReq.Unmarshal(m, b)
}
func (m *FileMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMessageReq.Marshal(b, m, deterministic)
}
func (m *FileMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMessageReq.Merge(m, src)
}
func (m *FileMessageReq) XXX_Size() int {
	return xxx_messageInfo_FileMessageReq.Size(m)
}
func (m *FileMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_FileMessageReq proto.InternalMessageInfo

func (m *FileMessageReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileMessageReq) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type DownloadReq struct {
	FileID               string   `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadReq) Reset()         { *m = DownloadReq{} }
func (m *DownloadReq) String() string { return proto.CompactTextString(m) }
func (*DownloadReq) ProtoMessage()    {}
func (*DownloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{9}
}

func (m *DownloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadReq.Unmarshal(m, b)
}
func (m *DownloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadReq.Marshal(b, m, deterministic)
}
func (m *DownloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadReq.Merge(m, src)
}
func (m *DownloadReq) XXX_Size() int {
	return xxx_messageInfo_DownloadReq.Size(m)
}
func (m *DownloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadReq proto.InternalMessageInfo

func (m *DownloadReq) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

type DownloadResp struct {
	FileID               string   `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResp) Reset()         { *m = DownloadResp{} }
func (m *DownloadResp) String() string { return proto.CompactTextString(m) }
func (*DownloadResp) ProtoMessage()    {}
func (*DownloadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{10}
}

func (m *DownloadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResp.Unmarshal(m, b)
}
func (m *DownloadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResp.Marshal(b, m, deterministic)
}
func (m *DownloadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResp.Merge(m, src)
}
func (m *DownloadResp) XXX_Size() int {
	return xxx_messageInfo_DownloadResp.Size(m)
}
func (m *DownloadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResp.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResp proto.InternalMessageInfo

func (m *DownloadResp) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *DownloadResp) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type InvitesResp struct {
	Invites              []*InvitesResp_Invite `protobuf:"bytes,1,rep,name=invites,proto3" json:"invites,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InvitesResp) Reset()         { *m = InvitesResp{} }
func (m *InvitesResp) String() string { return proto.CompactTextString(m) }
func (*InvitesResp) ProtoMessage()    {}
func (*InvitesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{11}
}

func (m *InvitesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitesResp.Unmarshal(m, b)
}
func (m *InvitesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitesResp.Marshal(b, m, deterministic)
}
func (m *InvitesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitesResp.Merge(m, src)
}
func (m *InvitesResp) XXX_Size() int {
	return xxx_messageInfo_InvitesResp.Size(m)
}
func (m *InvitesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitesResp.DiscardUnknown(m)
}

var xxx_messageInfo_InvitesResp proto.InternalMessageInfo

func (m *InvitesResp) GetInvites() []*InvitesResp_Invite {
	if m != nil {
		return m.Invites
	}
	return nil
}

type InvitesResp_Invite struct {
	InviteID             string   `protobuf:"bytes,1,opt,name=inviteID,proto3" json:"inviteID,omitempty"`
	FromUsername         string   `protobuf:"bytes,2,opt,name=fromUsername,proto3" json:"fromUsername,omitempty"`
	GroupName            string   `protobuf:"bytes,3,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvitesResp_Invite) Reset()         { *m = InvitesResp_Invite{} }
func (m *InvitesResp_Invite) String() string { return proto.CompactTextString(m) }
func (*InvitesResp_Invite) ProtoMessage()    {}
func (*InvitesResp_Invite) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{11, 0}
}

func (m *InvitesResp_Invite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvitesResp_Invite.Unmarshal(m, b)
}
func (m *InvitesResp_Invite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvitesResp_Invite.Marshal(b, m, deterministic)
}
func (m *InvitesResp_Invite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvitesResp_Invite.Merge(m, src)
}
func (m *InvitesResp_Invite) XXX_Size() int {
	return xxx_messageInfo_InvitesResp_Invite.Size(m)
}
func (m *InvitesResp_Invite) XXX_DiscardUnknown() {
	xxx_messageInfo_InvitesResp_Invite.DiscardUnknown(m)
}

var xxx_messageInfo_InvitesResp_Invite proto.InternalMessageInfo

func (m *InvitesResp_Invite) GetInviteID() string {
	if m != nil {
		return m.InviteID
	}
	return ""
}

func (m *InvitesResp_Invite) GetFromUsername() string {
	if m != nil {
		return m.FromUsername
	}
	return ""
}

func (m *InvitesResp_Invite) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type InviteReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteReq) Reset()         { *m = InviteReq{} }
func (m *InviteReq) String() string { return proto.CompactTextString(m) }
func (*InviteReq) ProtoMessage()    {}
func (*InviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{12}
}

func (m *InviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteReq.Unmarshal(m, b)
}
func (m *InviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteReq.Marshal(b, m, deterministic)
}
func (m *InviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteReq.Merge(m, src)
}
func (m *InviteReq) XXX_Size() int {
	return xxx_messageInfo_InviteReq.Size(m)
}
func (m *InviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_InviteReq proto.InternalMessageInfo

func (m *InviteReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AcceptInviteReq struct {
	InviteID             string   `protobuf:"bytes,1,opt,name=inviteID,proto3" json:"inviteID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcceptInviteReq) Reset()         { *m = AcceptInviteReq{} }
func (m *AcceptInviteReq) String() string { return proto.CompactTextString(m) }
func (*AcceptInviteReq) ProtoMessage()    {}
func (*AcceptInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{13}
}

func (m *AcceptInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcceptInviteReq.Unmarshal(m, b)
}
func (m *AcceptInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcceptInviteReq.Marshal(b, m, deterministic)
}
func (m *AcceptInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptInviteReq.Merge(m, src)
}
func (m *AcceptInviteReq) XXX_Size() int {
	return xxx_messageInfo_AcceptInviteReq.Size(m)
}
func (m *AcceptInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptInviteReq proto.InternalMessageInfo

func (m *AcceptInviteReq) GetInviteID() string {
	if m != nil {
		return m.InviteID
	}
	return ""
}

type DeleteInviteReq struct {
	InviteID             string   `protobuf:"bytes,1,opt,name=inviteID,proto3" json:"inviteID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInviteReq) Reset()         { *m = DeleteInviteReq{} }
func (m *DeleteInviteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteInviteReq) ProtoMessage()    {}
func (*DeleteInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{14}
}

func (m *DeleteInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInviteReq.Unmarshal(m, b)
}
func (m *DeleteInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInviteReq.Marshal(b, m, deterministic)
}
func (m *DeleteInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInviteReq.Merge(m, src)
}
func (m *DeleteInviteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteInviteReq.Size(m)
}
func (m *DeleteInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInviteReq proto.InternalMessageInfo

func (m *DeleteInviteReq) GetInviteID() string {
	if m != nil {
		return m.InviteID
	}
	return ""
}

type CreateGroupReq struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupReq) Reset()         { *m = CreateGroupReq{} }
func (m *CreateGroupReq) String() string { return proto.CompactTextString(m) }
func (*CreateGroupReq) ProtoMessage()    {}
func (*CreateGroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{15}
}

func (m *CreateGroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupReq.Unmarshal(m, b)
}
func (m *CreateGroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupReq.Marshal(b, m, deterministic)
}
func (m *CreateGroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupReq.Merge(m, src)
}
func (m *CreateGroupReq) XXX_Size() int {
	return xxx_messageInfo_CreateGroupReq.Size(m)
}
func (m *CreateGroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupReq proto.InternalMessageInfo

func (m *CreateGroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type JoinGroupReq struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGroupReq) Reset()         { *m = JoinGroupReq{} }
func (m *JoinGroupReq) String() string { return proto.CompactTextString(m) }
func (*JoinGroupReq) ProtoMessage()    {}
func (*JoinGroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{16}
}

func (m *JoinGroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGroupReq.Unmarshal(m, b)
}
func (m *JoinGroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGroupReq.Marshal(b, m, deterministic)
}
func (m *JoinGroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGroupReq.Merge(m, src)
}
func (m *JoinGroupReq) XXX_Size() int {
	return xxx_messageInfo_JoinGroupReq.Size(m)
}
func (m *JoinGroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGroupReq proto.InternalMessageInfo

func (m *JoinGroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type GroupResp struct {
	Messages             []*TextMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GroupResp) Reset()         { *m = GroupResp{} }
func (m *GroupResp) String() string { return proto.CompactTextString(m) }
func (*GroupResp) ProtoMessage()    {}
func (*GroupResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{17}
}

func (m *GroupResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupResp.Unmarshal(m, b)
}
func (m *GroupResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupResp.Marshal(b, m, deterministic)
}
func (m *GroupResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupResp.Merge(m, src)
}
func (m *GroupResp) XXX_Size() int {
	return xxx_messageInfo_GroupResp.Size(m)
}
func (m *GroupResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupResp.DiscardUnknown(m)
}

var xxx_messageInfo_GroupResp proto.InternalMessageInfo

func (m *GroupResp) GetMessages() []*TextMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type GroupsResp struct {
	GroupNames           []string `protobuf:"bytes,1,rep,name=groupNames,proto3" json:"groupNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupsResp) Reset()         { *m = GroupsResp{} }
func (m *GroupsResp) String() string { return proto.CompactTextString(m) }
func (*GroupsResp) ProtoMessage()    {}
func (*GroupsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{18}
}

func (m *GroupsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupsResp.Unmarshal(m, b)
}
func (m *GroupsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupsResp.Marshal(b, m, deterministic)
}
func (m *GroupsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupsResp.Merge(m, src)
}
func (m *GroupsResp) XXX_Size() int {
	return xxx_messageInfo_GroupsResp.Size(m)
}
func (m *GroupsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GroupsResp proto.InternalMessageInfo

func (m *GroupsResp) GetGroupNames() []string {
	if m != nil {
		return m.GroupNames
	}
	return nil
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb86ddf19e16901, []int{19}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Header)(nil), "Header")
	proto.RegisterType((*SignUpReq)(nil), "SignUpReq")
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*AuthResp)(nil), "AuthResp")
	proto.RegisterType((*UserSearchReq)(nil), "UserSearchReq")
	proto.RegisterType((*UserSearchResp)(nil), "UserSearchResp")
	proto.RegisterType((*TextMessageReq)(nil), "TextMessageReq")
	proto.RegisterType((*TextMessage)(nil), "TextMessage")
	proto.RegisterType((*FileMessageReq)(nil), "FileMessageReq")
	proto.RegisterType((*DownloadReq)(nil), "DownloadReq")
	proto.RegisterType((*DownloadResp)(nil), "DownloadResp")
	proto.RegisterType((*InvitesResp)(nil), "InvitesResp")
	proto.RegisterType((*InvitesResp_Invite)(nil), "InvitesResp.Invite")
	proto.RegisterType((*InviteReq)(nil), "InviteReq")
	proto.RegisterType((*AcceptInviteReq)(nil), "AcceptInviteReq")
	proto.RegisterType((*DeleteInviteReq)(nil), "DeleteInviteReq")
	proto.RegisterType((*CreateGroupReq)(nil), "CreateGroupReq")
	proto.RegisterType((*JoinGroupReq)(nil), "JoinGroupReq")
	proto.RegisterType((*GroupResp)(nil), "GroupResp")
	proto.RegisterType((*GroupsResp)(nil), "GroupsResp")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("Messages.proto", fileDescriptor_9eb86ddf19e16901) }

var fileDescriptor_9eb86ddf19e16901 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xef, 0x8b, 0x13, 0x31,
	0x10, 0x65, 0x7b, 0xbd, 0x5e, 0x77, 0x5a, 0x57, 0x88, 0x72, 0x94, 0x43, 0x8e, 0x12, 0x50, 0x8b,
	0xdc, 0x15, 0x51, 0x0e, 0xbf, 0xde, 0x8f, 0xfa, 0xe3, 0x44, 0x45, 0x72, 0x57, 0xfc, 0xbc, 0x76,
	0xa7, 0x6d, 0x70, 0x9b, 0xc4, 0x24, 0xf5, 0xfa, 0x07, 0xf9, 0x87, 0x4a, 0xb2, 0xd9, 0xed, 0xf6,
	0xa0, 0xe5, 0xc0, 0x6f, 0x79, 0x2f, 0x6f, 0xde, 0x4c, 0x26, 0x93, 0x40, 0xf2, 0x15, 0x8d, 0x49,
	0x67, 0x68, 0x86, 0x4a, 0x4b, 0x2b, 0xe9, 0x6b, 0x68, 0x7d, 0xc2, 0x34, 0x43, 0x4d, 0x12, 0x68,
	0xf0, 0xac, 0x17, 0xf5, 0xa3, 0x41, 0xcc, 0x1a, 0x3c, 0x23, 0x87, 0xd0, 0xca, 0x51, 0xcc, 0xec,
	0xbc, 0xd7, 0xe8, 0x47, 0x83, 0x7d, 0x16, 0x10, 0xbd, 0x82, 0xf8, 0x86, 0xcf, 0xc4, 0x58, 0x31,
	0xfc, 0x4d, 0x8e, 0xa0, 0xbd, 0x34, 0xa8, 0x45, 0xba, 0xc0, 0x10, 0x5a, 0x61, 0xb7, 0xa7, 0x52,
	0x63, 0xee, 0xa4, 0xce, 0xbc, 0x45, 0xcc, 0x2a, 0x4c, 0x2f, 0xa1, 0xfd, 0x45, 0xce, 0xb8, 0xf8,
	0x1f, 0x8f, 0x73, 0x68, 0x5f, 0x2c, 0xed, 0x9c, 0xa1, 0x51, 0xe4, 0x29, 0xec, 0x5b, 0xf9, 0x0b,
	0x45, 0x30, 0x28, 0x00, 0x39, 0x06, 0xc0, 0x95, 0xe2, 0x1a, 0x6f, 0xf9, 0x02, 0x7d, 0x7c, 0x93,
	0xd5, 0x18, 0xfa, 0x0e, 0x1e, 0x8d, 0x0d, 0xea, 0x1b, 0x4c, 0xf5, 0x64, 0xee, 0x4a, 0x79, 0x01,
	0x49, 0x99, 0xfa, 0xbb, 0xc6, 0x29, 0x5f, 0x05, 0xbf, 0x7b, 0x2c, 0x1d, 0x42, 0x52, 0x0f, 0x34,
	0x8a, 0x3c, 0x83, 0xb8, 0xd4, 0x98, 0x5e, 0xd4, 0xdf, 0x1b, 0xc4, 0x6c, 0x4d, 0xd0, 0x57, 0x90,
	0xdc, 0xe2, 0xca, 0x86, 0xde, 0xbb, 0x4c, 0x3d, 0x38, 0x58, 0x14, 0x28, 0xa4, 0x28, 0x21, 0xfd,
	0x01, 0x9d, 0x9a, 0x76, 0x67, 0x77, 0x6a, 0x26, 0x8d, 0x0d, 0x13, 0x42, 0xa0, 0x69, 0xdd, 0x99,
	0xf7, 0xfc, 0x99, 0xfd, 0x9a, 0x9e, 0x43, 0xf2, 0x81, 0xe7, 0x58, 0x2b, 0x82, 0x40, 0xb3, 0xe6,
	0xdb, 0x2c, 0x3b, 0x3e, 0x91, 0xc2, 0xa2, 0xb0, 0xc6, 0x9b, 0x76, 0x59, 0x85, 0xe9, 0x73, 0xe8,
	0x8c, 0xe4, 0x9d, 0xc8, 0x65, 0x9a, 0xb9, 0xf0, 0x43, 0x68, 0x4d, 0x79, 0x8e, 0xd7, 0xa3, 0x60,
	0x10, 0x10, 0xbd, 0x84, 0xee, 0x5a, 0x66, 0xd4, 0x36, 0xdd, 0xce, 0x54, 0x7f, 0x23, 0xe8, 0x5c,
	0x8b, 0x3f, 0xdc, 0xa2, 0xf1, 0x1e, 0xa7, 0x70, 0xc0, 0x0b, 0xe8, 0xbb, 0xdb, 0x79, 0xf3, 0x64,
	0x58, 0xdb, 0x0e, 0x6b, 0x56, 0x6a, 0x8e, 0xa6, 0xd0, 0x2a, 0x28, 0x97, 0xa4, 0x20, 0xab, 0xf4,
	0x15, 0x26, 0x14, 0xba, 0x53, 0x2d, 0x17, 0xe3, 0xb2, 0xbf, 0x45, 0x13, 0x37, 0x38, 0x77, 0xb1,
	0x33, 0x2d, 0x97, 0xea, 0x5b, 0x1a, 0xda, 0x19, 0xb3, 0x35, 0x41, 0x5f, 0x42, 0x1c, 0x52, 0xef,
	0x1e, 0x64, 0x7a, 0x0a, 0x8f, 0x2f, 0x26, 0x13, 0x54, 0x76, 0x43, 0xbe, 0xad, 0x32, 0x27, 0x1f,
	0x61, 0x8e, 0x16, 0x1f, 0x26, 0x1f, 0x42, 0x72, 0xa5, 0x31, 0xb5, 0xf8, 0xd1, 0x55, 0xe6, 0xd4,
	0x1b, 0x65, 0x47, 0xf7, 0xcb, 0x3e, 0x81, 0xee, 0x67, 0xc9, 0xc5, 0x03, 0xd5, 0x67, 0x10, 0x07,
	0xa5, 0x51, 0x64, 0x00, 0xed, 0x30, 0x64, 0xe5, 0x4d, 0x74, 0x87, 0xf5, 0xd9, 0xae, 0x76, 0xe9,
	0x09, 0x80, 0x0f, 0x2b, 0x2e, 0xf0, 0x18, 0xa0, 0x72, 0x2c, 0x5f, 0x48, 0x8d, 0xa1, 0x67, 0xb0,
	0xff, 0x5e, 0x6b, 0xa9, 0xb7, 0xbf, 0x0c, 0x37, 0xae, 0x13, 0x99, 0x61, 0xf8, 0x8f, 0xfc, 0xfa,
	0x67, 0xcb, 0x7f, 0x63, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x82, 0x82, 0xac, 0x46, 0xd8,
	0x04, 0x00, 0x00,
}