// Code generated by protoc-gen-go. DO NOT EDIT.
// source: knowing.proto

package knowing

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

type FindRequest struct {
	Tokens               []string `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
	ArticleId            string   `protobuf:"bytes,2,opt,name=articleId" json:"articleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowing_76304af7d38551f8, []int{0}
}
func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (dst *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(dst, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *FindRequest) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

type HasOptionResponse struct {
	Items                []*HasOption `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HasOptionResponse) Reset()         { *m = HasOptionResponse{} }
func (m *HasOptionResponse) String() string { return proto.CompactTextString(m) }
func (*HasOptionResponse) ProtoMessage()    {}
func (*HasOptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowing_76304af7d38551f8, []int{1}
}
func (m *HasOptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasOptionResponse.Unmarshal(m, b)
}
func (m *HasOptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasOptionResponse.Marshal(b, m, deterministic)
}
func (dst *HasOptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasOptionResponse.Merge(dst, src)
}
func (m *HasOptionResponse) XXX_Size() int {
	return xxx_messageInfo_HasOptionResponse.Size(m)
}
func (m *HasOptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasOptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasOptionResponse proto.InternalMessageInfo

func (m *HasOptionResponse) GetItems() []*HasOption {
	if m != nil {
		return m.Items
	}
	return nil
}

type HasOption struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Listen               int64    `protobuf:"varint,2,opt,name=listen" json:"listen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasOption) Reset()         { *m = HasOption{} }
func (m *HasOption) String() string { return proto.CompactTextString(m) }
func (*HasOption) ProtoMessage()    {}
func (*HasOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_knowing_76304af7d38551f8, []int{2}
}
func (m *HasOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasOption.Unmarshal(m, b)
}
func (m *HasOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasOption.Marshal(b, m, deterministic)
}
func (dst *HasOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasOption.Merge(dst, src)
}
func (m *HasOption) XXX_Size() int {
	return xxx_messageInfo_HasOption.Size(m)
}
func (m *HasOption) XXX_DiscardUnknown() {
	xxx_messageInfo_HasOption.DiscardUnknown(m)
}

var xxx_messageInfo_HasOption proto.InternalMessageInfo

func (m *HasOption) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *HasOption) GetListen() int64 {
	if m != nil {
		return m.Listen
	}
	return 0
}

func init() {
	proto.RegisterType((*FindRequest)(nil), "knowing.FindRequest")
	proto.RegisterType((*HasOptionResponse)(nil), "knowing.HasOptionResponse")
	proto.RegisterType((*HasOption)(nil), "knowing.HasOption")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RegionHandler service

type RegionHandlerClient interface {
	GetListenAudio(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*HasOptionResponse, error)
}

type regionHandlerClient struct {
	cc *grpc.ClientConn
}

func NewRegionHandlerClient(cc *grpc.ClientConn) RegionHandlerClient {
	return &regionHandlerClient{cc}
}

func (c *regionHandlerClient) GetListenAudio(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*HasOptionResponse, error) {
	out := new(HasOptionResponse)
	err := grpc.Invoke(ctx, "/knowing.RegionHandler/GetListenAudio", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegionHandler service

type RegionHandlerServer interface {
	GetListenAudio(context.Context, *FindRequest) (*HasOptionResponse, error)
}

func RegisterRegionHandlerServer(s *grpc.Server, srv RegionHandlerServer) {
	s.RegisterService(&_RegionHandler_serviceDesc, srv)
}

func _RegionHandler_GetListenAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionHandlerServer).GetListenAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knowing.RegionHandler/GetListenAudio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionHandlerServer).GetListenAudio(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegionHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "knowing.RegionHandler",
	HandlerType: (*RegionHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListenAudio",
			Handler:    _RegionHandler_GetListenAudio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "knowing.proto",
}

func init() { proto.RegisterFile("knowing.proto", fileDescriptor_knowing_76304af7d38551f8) }

var fileDescriptor_knowing_76304af7d38551f8 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0x89, 0xa5, 0x95, 0x9d, 0x52, 0xc1, 0xa1, 0x48, 0x28, 0x1e, 0x4a, 0x4e, 0x39, 0xf5,
	0x50, 0x4f, 0x1e, 0x3c, 0xa8, 0xa0, 0x2d, 0x08, 0xc2, 0xfa, 0x05, 0xd1, 0x1d, 0xca, 0xd2, 0x38,
	0x13, 0xb3, 0x13, 0xfc, 0x7d, 0x71, 0xb3, 0x26, 0x1e, 0x7a, 0x7c, 0xb3, 0xc3, 0xbe, 0xb7, 0x0b,
	0x8b, 0x23, 0xcb, 0xb7, 0xe7, 0xc3, 0xa6, 0x69, 0x45, 0x05, 0xcf, 0x13, 0x16, 0x8f, 0x30, 0x7f,
	0xf2, 0xec, 0x2c, 0x7d, 0x75, 0x14, 0x14, 0xaf, 0x60, 0xa6, 0x72, 0x24, 0x0e, 0x79, 0xb6, 0x9e,
	0x94, 0xc6, 0x26, 0xc2, 0x6b, 0x30, 0x55, 0xab, 0xfe, 0xa3, 0xa6, 0xbd, 0xcb, 0xcf, 0xd6, 0x59,
	0x69, 0xec, 0x38, 0x28, 0xee, 0xe0, 0x72, 0x57, 0x85, 0xd7, 0x46, 0xbd, 0xb0, 0xa5, 0xd0, 0x08,
	0x07, 0xc2, 0x12, 0xa6, 0x7b, 0xa5, 0xcf, 0xfe, 0xa6, 0xf9, 0x16, 0x37, 0x7f, 0x05, 0xe3, 0x6a,
	0xbf, 0x50, 0xdc, 0x82, 0x19, 0x66, 0xb8, 0x84, 0x69, 0x74, 0xe6, 0x59, 0xb4, 0xf4, 0xf0, 0xdb,
	0x55, 0xfb, 0xa0, 0xc4, 0x51, 0x3e, 0xb1, 0x89, 0xb6, 0x6f, 0xb0, 0xb0, 0x74, 0xf0, 0xc2, 0xbb,
	0x8a, 0x5d, 0x4d, 0x2d, 0x3e, 0xc0, 0xc5, 0x33, 0xe9, 0x4b, 0x3c, 0xbd, 0xef, 0x9c, 0x17, 0x5c,
	0x0e, 0xe2, 0x7f, 0x0f, 0x5d, 0xad, 0x4e, 0xe4, 0xa4, 0xf2, 0xf7, 0x59, 0xfc, 0xa3, 0x9b, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x5b, 0x23, 0xb4, 0x34, 0x01, 0x00, 0x00,
}
