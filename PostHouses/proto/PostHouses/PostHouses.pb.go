// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/PostHouses/PostHouses.proto

package go_micro_srv_PostHouses

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	return fileDescriptor_fc9a2ca88e02e7ac, []int{0}
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
	HouseInfo            []byte   `protobuf:"bytes,2,opt,name=HouseInfo,proto3" json:"HouseInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc9a2ca88e02e7ac, []int{1}
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

func (m *Request) GetHouseInfo() []byte {
	if m != nil {
		return m.HouseInfo
	}
	return nil
}

type Response struct {
	Error                string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	HousID               string   `protobuf:"bytes,3,opt,name=HousID,proto3" json:"HousID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc9a2ca88e02e7ac, []int{2}
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

func (m *Response) GetHousID() string {
	if m != nil {
		return m.HousID
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostHouses.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostHouses.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostHouses.Response")
}

func init() { proto.RegisterFile("proto/PostHouses/PostHouses.proto", fileDescriptor_fc9a2ca88e02e7ac) }

var fileDescriptor_fc9a2ca88e02e7ac = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x8d, 0xc5, 0xd6, 0x0c, 0x22, 0xb2, 0x88, 0x06, 0xf5, 0x90, 0xee, 0xa9, 0xa7, 0x15,
	0xf4, 0x23, 0xd8, 0x80, 0x39, 0x14, 0xca, 0x7a, 0xf0, 0x1c, 0x65, 0x0c, 0x85, 0x98, 0x89, 0x33,
	0x89, 0xe0, 0xb7, 0x97, 0xfd, 0x03, 0x9b, 0x8b, 0xde, 0xe6, 0xfd, 0xde, 0xec, 0x63, 0xdf, 0xc0,
	0x7a, 0x60, 0x1a, 0xe9, 0x7e, 0x4f, 0x32, 0x3e, 0xd3, 0x24, 0x28, 0xb3, 0xd1, 0x78, 0x4f, 0x5d,
	0xb7, 0x64, 0x3e, 0x0f, 0xef, 0x4c, 0x46, 0xf8, 0xdb, 0x24, 0x5b, 0xdf, 0xc2, 0x6a, 0x87, 0x22,
	0x4d, 0x8b, 0xea, 0x02, 0x16, 0xd2, 0xfc, 0x14, 0x59, 0x99, 0x6d, 0x72, 0xeb, 0x46, 0x5d, 0xc1,
	0xca, 0xe2, 0xd7, 0x84, 0x32, 0xaa, 0x3b, 0xc8, 0x5f, 0x50, 0xe4, 0x40, 0x7d, 0xbd, 0x8d, 0x2b,
	0x09, 0x38, 0xd7, 0xe7, 0xd5, 0xfd, 0x07, 0x15, 0xc7, 0x65, 0xb6, 0x39, 0xb3, 0x09, 0xe8, 0x3d,
	0x9c, 0x5a, 0x94, 0x81, 0x7a, 0x41, 0x75, 0x09, 0x27, 0x15, 0x33, 0x71, 0xcc, 0x08, 0x42, 0x5d,
	0xc1, 0xb2, 0x62, 0xde, 0x49, 0xeb, 0x1f, 0xe7, 0x36, 0x2a, 0xc7, 0x5d, 0x4c, 0xbd, 0x2d, 0x16,
	0x81, 0x07, 0xf5, 0x80, 0x00, 0xa9, 0x83, 0x7a, 0x85, 0xf3, 0xa7, 0xa6, 0xeb, 0x66, 0xa4, 0x34,
	0x7f, 0xf4, 0x35, 0xb1, 0xcf, 0xcd, 0xfa, 0x9f, 0x8d, 0xf0, 0x55, 0x7d, 0xf4, 0xb6, 0xf4, 0xc7,
	0x7b, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xe4, 0x2f, 0x9e, 0x61, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostHousesClient is the client API for PostHouses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostHousesClient interface {
	CallPostHouses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type postHousesClient struct {
	cc *grpc.ClientConn
}

func NewPostHousesClient(cc *grpc.ClientConn) PostHousesClient {
	return &postHousesClient{cc}
}

func (c *postHousesClient) CallPostHouses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.PostHouses.PostHouses/CallPostHouses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostHousesServer is the server API for PostHouses service.
type PostHousesServer interface {
	CallPostHouses(context.Context, *Request) (*Response, error)
}

// UnimplementedPostHousesServer can be embedded to have forward compatible implementations.
type UnimplementedPostHousesServer struct {
}

func (*UnimplementedPostHousesServer) CallPostHouses(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallPostHouses not implemented")
}

func RegisterPostHousesServer(s *grpc.Server, srv PostHousesServer) {
	s.RegisterService(&_PostHouses_serviceDesc, srv)
}

func _PostHouses_CallPostHouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostHousesServer).CallPostHouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.PostHouses.PostHouses/CallPostHouses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostHousesServer).CallPostHouses(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostHouses_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.PostHouses.PostHouses",
	HandlerType: (*PostHousesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallPostHouses",
			Handler:    _PostHouses_CallPostHouses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/PostHouses/PostHouses.proto",
}
