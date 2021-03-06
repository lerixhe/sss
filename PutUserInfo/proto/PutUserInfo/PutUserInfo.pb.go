// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/PutUserInfo/PutUserInfo.proto

package go_micro_srv_PutUserInfo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6116466ada533dc8, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	NewName              string   `protobuf:"bytes,2,opt,name=NewName,proto3" json:"NewName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6116466ada533dc8, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *Request) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

type Response struct {
	Error                string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	NewName              string   `protobuf:"bytes,3,opt,name=NewName,proto3" json:"NewName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6116466ada533dc8, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *Response) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PutUserInfo.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutUserInfo.Response")
}

func init() {
	proto.RegisterFile("proto/PutUserInfo/PutUserInfo.proto", fileDescriptor_6116466ada533dc8)
}

var fileDescriptor_6116466ada533dc8 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbf, 0x4e, 0x87, 0x30,
	0x10, 0x16, 0x89, 0x20, 0xe7, 0xa0, 0x69, 0x8c, 0x69, 0xd4, 0x41, 0xeb, 0xe2, 0x54, 0x13, 0x7d,
	0x02, 0xa3, 0x0c, 0x0c, 0x10, 0x53, 0xe3, 0xe2, 0x56, 0xcd, 0x49, 0x48, 0x80, 0xe2, 0x1d, 0x68,
	0x7c, 0x7b, 0x23, 0xd4, 0x58, 0x06, 0x7f, 0xdb, 0x7d, 0x7f, 0xd3, 0xaf, 0x70, 0x31, 0x90, 0x1b,
	0xdd, 0xd5, 0xc3, 0x34, 0x3e, 0x31, 0x52, 0xd1, 0xbf, 0xad, 0x6e, 0x3d, 0xab, 0x42, 0xd6, 0x4e,
	0x77, 0xcd, 0x2b, 0x39, 0xcd, 0xf4, 0xa1, 0x03, 0x5d, 0x9d, 0x40, 0x5a, 0x22, 0xb3, 0xad, 0x51,
	0x1c, 0x40, 0xcc, 0xf6, 0x4b, 0x46, 0x67, 0xd1, 0x65, 0x66, 0x7e, 0x4e, 0x75, 0x0b, 0xa9, 0xc1,
	0xf7, 0x09, 0x79, 0x14, 0xa7, 0x90, 0x3d, 0x22, 0x73, 0xe3, 0xfa, 0xe2, 0xde, 0x5b, 0xfe, 0x08,
	0x21, 0x21, 0xad, 0xf0, 0xb3, 0xb2, 0x1d, 0xca, 0xed, 0x59, 0xfb, 0x85, 0xca, 0xc0, 0xae, 0x41,
	0x1e, 0x5c, 0xcf, 0x28, 0x0e, 0x61, 0x27, 0x27, 0x72, 0xe4, 0xf3, 0x0b, 0x10, 0x47, 0x90, 0xe4,
	0x44, 0x25, 0xd7, 0x3e, 0xea, 0x51, 0xd8, 0x19, 0xaf, 0x3a, 0xaf, 0x1b, 0xd8, 0x0b, 0x26, 0x88,
	0x67, 0xd8, 0xbf, 0xb3, 0x6d, 0x1b, 0x52, 0xe7, 0xfa, 0xbf, 0xc1, 0xda, 0x0f, 0x3a, 0x56, 0x9b,
	0x2c, 0xcb, 0x83, 0xd5, 0xd6, 0x4b, 0x32, 0xff, 0xdf, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x67, 0x60, 0xc8, 0xe9, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PutUserInfoClient is the client API for PutUserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PutUserInfoClient interface {
	CallPutUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type putUserInfoClient struct {
	cc *grpc.ClientConn
}

func NewPutUserInfoClient(cc *grpc.ClientConn) PutUserInfoClient {
	return &putUserInfoClient{cc}
}

func (c *putUserInfoClient) CallPutUserInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.PutUserInfo.PutUserInfo/CallPutUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutUserInfoServer is the server API for PutUserInfo service.
type PutUserInfoServer interface {
	CallPutUserInfo(context.Context, *Request) (*Response, error)
}

// UnimplementedPutUserInfoServer can be embedded to have forward compatible implementations.
type UnimplementedPutUserInfoServer struct {
}

func (*UnimplementedPutUserInfoServer) CallPutUserInfo(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallPutUserInfo not implemented")
}

func RegisterPutUserInfoServer(s *grpc.Server, srv PutUserInfoServer) {
	s.RegisterService(&_PutUserInfo_serviceDesc, srv)
}

func _PutUserInfo_CallPutUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutUserInfoServer).CallPutUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.PutUserInfo.PutUserInfo/CallPutUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutUserInfoServer).CallPutUserInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _PutUserInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.PutUserInfo.PutUserInfo",
	HandlerType: (*PutUserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallPutUserInfo",
			Handler:    _PutUserInfo_CallPutUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/PutUserInfo/PutUserInfo.proto",
}
