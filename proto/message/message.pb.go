// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessSrvErrors int32

const (
	MessSrvErrors_None          MessSrvErrors = 0
	MessSrvErrors_InvalidUserId MessSrvErrors = 1001
)

var MessSrvErrors_name = map[int32]string{
	0:    "None",
	1001: "InvalidUserId",
}
var MessSrvErrors_value = map[string]int32{
	"None":          0,
	"InvalidUserId": 1001,
}

func (x MessSrvErrors) String() string {
	return proto.EnumName(MessSrvErrors_name, int32(x))
}
func (MessSrvErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{0}
}

type SendRequest struct {
	Userinfo             *UserInfo `protobuf:"bytes,1,opt,name=userinfo" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{0}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (dst *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(dst, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetUserinfo() *UserInfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

type SendReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReply) Reset()         { *m = SendReply{} }
func (m *SendReply) String() string { return proto.CompactTextString(m) }
func (*SendReply) ProtoMessage()    {}
func (*SendReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{1}
}
func (m *SendReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendReply.Unmarshal(m, b)
}
func (m *SendReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendReply.Marshal(b, m, deterministic)
}
func (dst *SendReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReply.Merge(dst, src)
}
func (m *SendReply) XXX_Size() int {
	return xxx_messageInfo_SendReply.Size(m)
}
func (m *SendReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendReply proto.InternalMessageInfo

func (m *SendReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ReceiveReply struct {
	Userinfo             []*UserInfo `protobuf:"bytes,1,rep,name=userinfo" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceiveReply) Reset()         { *m = ReceiveReply{} }
func (m *ReceiveReply) String() string { return proto.CompactTextString(m) }
func (*ReceiveReply) ProtoMessage()    {}
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{3}
}
func (m *ReceiveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveReply.Unmarshal(m, b)
}
func (m *ReceiveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveReply.Marshal(b, m, deterministic)
}
func (dst *ReceiveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveReply.Merge(dst, src)
}
func (m *ReceiveReply) XXX_Size() int {
	return xxx_messageInfo_ReceiveReply.Size(m)
}
func (m *ReceiveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveReply proto.InternalMessageInfo

func (m *ReceiveReply) GetUserinfo() []*UserInfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

type UserInfo struct {
	Userid               int32    `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Intro                string   `protobuf:"bytes,3,opt,name=intro" json:"intro,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_6ea518f37022e2ae, []int{4}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfo) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "message.SendRequest")
	proto.RegisterType((*SendReply)(nil), "message.SendReply")
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*ReceiveReply)(nil), "message.ReceiveReply")
	proto.RegisterType((*UserInfo)(nil), "message.UserInfo")
	proto.RegisterEnum("message.MessSrvErrors", MessSrvErrors_name, MessSrvErrors_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessSrv service

type MessSrvClient interface {
	// 发送
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error)
	// 接收
	Receive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReceiveReply, error)
}

type messSrvClient struct {
	cc *grpc.ClientConn
}

func NewMessSrvClient(cc *grpc.ClientConn) MessSrvClient {
	return &messSrvClient{cc}
}

func (c *messSrvClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := grpc.Invoke(ctx, "/message.MessSrv/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messSrvClient) Receive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReceiveReply, error) {
	out := new(ReceiveReply)
	err := grpc.Invoke(ctx, "/message.MessSrv/Receive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessSrv service

type MessSrvServer interface {
	// 发送
	Send(context.Context, *SendRequest) (*SendReply, error)
	// 接收
	Receive(context.Context, *Empty) (*ReceiveReply, error)
}

func RegisterMessSrvServer(s *grpc.Server, srv MessSrvServer) {
	s.RegisterService(&_MessSrv_serviceDesc, srv)
}

func _MessSrv_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessSrvServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessSrv/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessSrvServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessSrv_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessSrvServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessSrv/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessSrvServer).Receive(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessSrv",
	HandlerType: (*MessSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MessSrv_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _MessSrv_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_6ea518f37022e2ae) }

var fileDescriptor_message_6ea518f37022e2ae = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0xe1, 0x6f, 0x29, 0xf4, 0xf6, 0xc7, 0xd4, 0x9b, 0x6a, 0x08, 0xab, 0x66, 0xdc, 0x34,
	0x26, 0xed, 0x02, 0xb7, 0xba, 0xec, 0xa2, 0x0b, 0x5d, 0x4c, 0xf5, 0x01, 0x50, 0x6e, 0x0d, 0x09,
	0xcc, 0xd0, 0x19, 0x20, 0xe1, 0x91, 0x7d, 0x0b, 0xd3, 0x81, 0x4e, 0xaa, 0x1b, 0x97, 0xdf, 0xe1,
	0x1c, 0xce, 0xcc, 0x19, 0x08, 0x4b, 0xd2, 0x3a, 0xfd, 0xa4, 0x4d, 0xa5, 0x64, 0x2d, 0xd1, 0x1f,
	0x90, 0x3d, 0xc2, 0x6c, 0x4f, 0x22, 0xe3, 0x74, 0x6c, 0x48, 0xd7, 0xb8, 0x86, 0xa0, 0xd1, 0xa4,
	0x72, 0x71, 0x90, 0x91, 0xbb, 0x74, 0x57, 0xb3, 0xe4, 0x7a, 0x73, 0x4e, 0xbe, 0x69, 0x52, 0x3b,
	0x71, 0x90, 0xdc, 0x5a, 0xd8, 0x1d, 0x4c, 0xfb, 0x74, 0x55, 0x74, 0x78, 0x0b, 0x13, 0x45, 0xba,
	0x29, 0x6a, 0x93, 0x0c, 0xf8, 0x40, 0xcc, 0x07, 0x6f, 0x5b, 0x56, 0x75, 0xc7, 0x9e, 0xe0, 0x3f,
	0xa7, 0x0f, 0xca, 0x5b, 0xea, 0x03, 0x3f, 0xcb, 0x46, 0x7f, 0x95, 0xbd, 0x42, 0x70, 0x56, 0x4f,
	0x5d, 0x46, 0xcf, 0x4c, 0x97, 0xc7, 0x07, 0xc2, 0xb8, 0xff, 0xa5, 0x48, 0x4b, 0x8a, 0xfe, 0x2d,
	0xdd, 0xd5, 0x94, 0x5b, 0xc6, 0x05, 0x78, 0xb9, 0xa8, 0x95, 0x8c, 0x46, 0xe6, 0x43, 0x0f, 0xf7,
	0x6b, 0x08, 0x9f, 0x49, 0xeb, 0xbd, 0x6a, 0xb7, 0x4a, 0x49, 0xa5, 0x31, 0x80, 0xf1, 0x8b, 0x14,
	0x34, 0x77, 0x10, 0x21, 0xdc, 0x89, 0x36, 0x2d, 0xf2, 0xcc, 0xf4, 0x66, 0xf3, 0x2f, 0x3f, 0x39,
	0x82, 0x3f, 0xd8, 0x31, 0x81, 0xf1, 0xe9, 0xf2, 0xb8, 0xb0, 0x87, 0xbe, 0x58, 0x32, 0xc6, 0x5f,
	0x6a, 0x55, 0x74, 0xcc, 0xc1, 0x04, 0xfc, 0x61, 0x02, 0xbc, 0xb2, 0x06, 0xb3, 0x4e, 0x7c, 0x63,
	0xf9, 0x72, 0x24, 0xe6, 0xbc, 0x4f, 0xcc, 0x93, 0x3d, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf0,
	0x07, 0x80, 0xb7, 0xc3, 0x01, 0x00, 0x00,
}
