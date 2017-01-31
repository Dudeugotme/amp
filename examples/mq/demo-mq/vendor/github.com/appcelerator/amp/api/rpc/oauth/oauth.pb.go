// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/oauth/oauth.proto
// DO NOT EDIT!

/*
Package oauth is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/oauth/oauth.proto

It has these top-level messages:
	AuthRequest
	AuthReply
	Token
*/
package oauth

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

type AuthRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Otp      string `protobuf:"bytes,3,opt,name=otp" json:"otp,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetOtp() string {
	if m != nil {
		return m.Otp
	}
	return ""
}

type AuthReply struct {
	SessionKey string `protobuf:"bytes,1,opt,name=sessionKey" json:"sessionKey,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *AuthReply) Reset()                    { *m = AuthReply{} }
func (m *AuthReply) String() string            { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()               {}
func (*AuthReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthReply) GetSessionKey() string {
	if m != nil {
		return m.SessionKey
	}
	return ""
}

func (m *AuthReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Token struct {
	Id             int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Token          string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	TokenLastEight string `protobuf:"bytes,3,opt,name=tokenLastEight" json:"tokenLastEight,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Token) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetTokenLastEight() string {
	if m != nil {
		return m.TokenLastEight
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "oauth.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "oauth.AuthReply")
	proto.RegisterType((*Token)(nil), "oauth.Token")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Github service

type GithubClient interface {
	Create(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type githubClient struct {
	cc *grpc.ClientConn
}

func NewGithubClient(cc *grpc.ClientConn) GithubClient {
	return &githubClient{cc}
}

func (c *githubClient) Create(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/oauth.Github/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Github service

type GithubServer interface {
	Create(context.Context, *AuthRequest) (*AuthReply, error)
}

func RegisterGithubServer(s *grpc.Server, srv GithubServer) {
	s.RegisterService(&_Github_serviceDesc, srv)
}

func _Github_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.Github/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServer).Create(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Github_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.Github",
	HandlerType: (*GithubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Github_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/oauth/oauth.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/oauth/oauth.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x50, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0xf5, 0xee, 0x4c, 0xf0, 0x46, 0x38, 0x8e, 0xc1, 0x22, 0x5c, 0x21, 0x92, 0x42, 0xac, 0x12,
	0x51, 0x6c, 0x6c, 0x44, 0x44, 0x2c, 0xb4, 0x0a, 0x8a, 0xf5, 0x5e, 0x32, 0x5c, 0x16, 0x93, 0xec,
	0xb8, 0x3b, 0x41, 0xf2, 0xef, 0x25, 0xbb, 0x51, 0x4e, 0x9b, 0xe5, 0x7d, 0xec, 0xbc, 0x19, 0x1e,
	0xdc, 0xec, 0xb4, 0xd4, 0xfd, 0x36, 0x2b, 0x4d, 0x9b, 0x2b, 0xe6, 0x92, 0x1a, 0xb2, 0x4a, 0x8c,
	0xcd, 0x55, 0xcb, 0xb9, 0x62, 0x9d, 0x5b, 0x2e, 0x73, 0xa3, 0x7a, 0xa9, 0xc3, 0x9b, 0xb1, 0x35,
	0x62, 0x30, 0xf2, 0x24, 0x7d, 0x87, 0xe3, 0xfb, 0x5e, 0xea, 0x82, 0x3e, 0x7b, 0x72, 0x82, 0x1b,
	0x38, 0xea, 0x1d, 0xd9, 0x4e, 0xb5, 0x94, 0xcc, 0xce, 0x66, 0x17, 0xcb, 0xe2, 0x97, 0x8f, 0x1e,
	0x2b, 0xe7, 0xbe, 0x8c, 0xad, 0x92, 0x79, 0xf0, 0x7e, 0x38, 0xae, 0x61, 0x61, 0x84, 0x93, 0x85,
	0x97, 0x47, 0x98, 0xde, 0xc1, 0x32, 0x04, 0x73, 0x33, 0xe0, 0x29, 0x80, 0x23, 0xe7, 0xb4, 0xe9,
	0x9e, 0x69, 0x98, 0x82, 0xf7, 0x14, 0x44, 0x38, 0xf4, 0x2b, 0x43, 0xac, 0xc7, 0xe9, 0x1b, 0x44,
	0xaf, 0xe6, 0x83, 0x3a, 0x5c, 0xc1, 0x5c, 0x57, 0x7e, 0x28, 0x2a, 0xe6, 0xba, 0xc2, 0x13, 0x88,
	0x64, 0x34, 0xa6, 0xdf, 0x81, 0xe0, 0x39, 0xac, 0x3c, 0x78, 0x51, 0x4e, 0x1e, 0xf5, 0xae, 0x96,
	0xe9, 0x98, 0x7f, 0xea, 0xd5, 0x2d, 0xc4, 0x4f, 0xbe, 0x30, 0xbc, 0x84, 0xf8, 0xc1, 0x92, 0x12,
	0x42, 0xcc, 0x42, 0x33, 0x7b, 0x4d, 0x6c, 0xd6, 0x7f, 0x34, 0x6e, 0x86, 0xf4, 0x60, 0x1b, 0xfb,
	0xea, 0xae, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0xd9, 0x1a, 0x77, 0x73, 0x01, 0x00, 0x00,
}
