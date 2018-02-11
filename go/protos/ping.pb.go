// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/ping.proto

/*
Package ping is a generated protocol buffer package.

It is generated from these files:
	protos/ping.proto

It has these top-level messages:
	HelloReqest
	HelloResponse
*/
package ping

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

type HelloReqest struct {
	ToMessage string `protobuf:"bytes,1,opt,name=toMessage" json:"toMessage,omitempty"`
}

func (m *HelloReqest) Reset()                    { *m = HelloReqest{} }
func (m *HelloReqest) String() string            { return proto.CompactTextString(m) }
func (*HelloReqest) ProtoMessage()               {}
func (*HelloReqest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloReqest) GetToMessage() string {
	if m != nil {
		return m.ToMessage
	}
	return ""
}

type HelloResponse struct {
	ResMessage string `protobuf:"bytes,1,opt,name=resMessage" json:"resMessage,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetResMessage() string {
	if m != nil {
		return m.ResMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReqest)(nil), "ping.HelloReqest")
	proto.RegisterType((*HelloResponse)(nil), "ping.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ping service

type PingClient interface {
	Hello(ctx context.Context, in *HelloReqest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) Hello(ctx context.Context, in *HelloReqest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/ping.Ping/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingServer interface {
	Hello(context.Context, *HelloReqest) (*HelloResponse, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ping.Ping/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Hello(ctx, req.(*HelloReqest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ping.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Ping_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/ping.proto",
}

func init() { proto.RegisterFile("protos/ping.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xc8, 0xcc, 0x4b, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0x40, 0x6c, 0x25, 0x6d,
	0x2e, 0x6e, 0x8f, 0xd4, 0x9c, 0x9c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e,
	0xce, 0x92, 0x7c, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x84, 0x80, 0x92, 0x3e, 0x17, 0x2f, 0x54, 0x71, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90,
	0x1c, 0x17, 0x57, 0x51, 0x6a, 0x31, 0xaa, 0x7a, 0x24, 0x11, 0x23, 0x4b, 0x2e, 0x96, 0x80, 0xcc,
	0xbc, 0x74, 0x21, 0x43, 0x2e, 0x56, 0xb0, 0x46, 0x21, 0x41, 0x3d, 0xb0, 0x0b, 0x90, 0xac, 0x94,
	0x12, 0x46, 0x11, 0x82, 0x18, 0xac, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xa5, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xcf, 0xad, 0x53, 0xba, 0x00, 0x00, 0x00,
}