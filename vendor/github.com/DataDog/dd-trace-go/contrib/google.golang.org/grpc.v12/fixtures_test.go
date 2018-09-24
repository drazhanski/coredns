// Code generated by protoc-gen-go.
// source: fixtures.proto
// DO NOT EDIT!

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	fixtures.proto

It has these top-level messages:
	FixtureRequest
	FixtureReply
*/
package grpc

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

// The request message containing the user's name.
type FixtureRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *FixtureRequest) Reset()                    { *m = FixtureRequest{} }
func (m *FixtureRequest) String() string            { return proto.CompactTextString(m) }
func (*FixtureRequest) ProtoMessage()               {}
func (*FixtureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FixtureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type FixtureReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *FixtureReply) Reset()                    { *m = FixtureReply{} }
func (m *FixtureReply) String() string            { return proto.CompactTextString(m) }
func (*FixtureReply) ProtoMessage()               {}
func (*FixtureReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FixtureReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FixtureRequest)(nil), "grpc.FixtureRequest")
	proto.RegisterType((*FixtureReply)(nil), "grpc.FixtureReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Fixture service

type FixtureClient interface {
	Ping(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*FixtureReply, error)
}

type fixtureClient struct {
	cc *grpc.ClientConn
}

func NewFixtureClient(cc *grpc.ClientConn) FixtureClient {
	return &fixtureClient{cc}
}

func (c *fixtureClient) Ping(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*FixtureReply, error) {
	out := new(FixtureReply)
	err := grpc.Invoke(ctx, "/grpc.Fixture/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fixture service

type FixtureServer interface {
	Ping(context.Context, *FixtureRequest) (*FixtureReply, error)
}

func RegisterFixtureServer(s *grpc.Server, srv FixtureServer) {
	s.RegisterService(&fixtureServiceDesc, srv)
}

func fixturePingHandler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixtureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixtureServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fixture/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixtureServer).Ping(ctx, req.(*FixtureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var fixtureServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Fixture",
	HandlerType: (*FixtureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    fixturePingHandler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fixtures.proto",
}

func init() { proto.RegisterFile("fixtures.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xac, 0x28,
	0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x29, 0x4a, 0x4c,
	0x4e, 0x4d, 0x2f, 0x2a, 0x48, 0x56, 0x52, 0xe1, 0xe2, 0x73, 0x83, 0x48, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4a, 0x1a, 0x5c, 0x3c, 0x70, 0x55, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x65, 0x30, 0xae, 0x91, 0x3b, 0x17, 0x3b, 0x54,
	0xa5, 0x90, 0x0d, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x90, 0xa4, 0x1e, 0xdc, 0x3a, 0x3d, 0x54,
	0xbb, 0xa4, 0xc4, 0xb1, 0x49, 0x15, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0xe9, 0x70, 0x49, 0x66, 0xe6,
	0xeb, 0x81, 0x65, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0xf5, 0x4a, 0x52, 0x8b, 0x4b,
	0x40, 0x22, 0x4e, 0xbc, 0x21, 0xa9, 0xc5, 0x25, 0xee, 0x41, 0x01, 0xce, 0x01, 0x20, 0xff, 0x04,
	0x30, 0x26, 0xb1, 0x81, 0x3d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x68, 0x5d, 0x74, 0x4a,
	0xea, 0x00, 0x00, 0x00,
}
