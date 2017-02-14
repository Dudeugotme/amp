// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/account/account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/account/account.proto

It has these top-level messages:
	AccountEntry
	SignUpRequest
	SignUpReply
	VerificationRequest
	LogInRequest
	LogInReply
	PasswordResetRequest
	PasswordResetReply
	PasswordSetRequest
	PasswordChangeRequest
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import schema "github.com/appcelerator/amp/data/account/schema"

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

type AccountEntry struct {
	UserName      string             `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Email         string             `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	EmailVerified bool               `protobuf:"varint,3,opt,name=email_verified,json=emailVerified" json:"email_verified,omitempty"`
	AccountType   schema.AccountType `protobuf:"varint,4,opt,name=account_type,json=accountType,enum=schema.AccountType" json:"account_type,omitempty"`
}

func (m *AccountEntry) Reset()                    { *m = AccountEntry{} }
func (m *AccountEntry) String() string            { return proto.CompactTextString(m) }
func (*AccountEntry) ProtoMessage()               {}
func (*AccountEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountEntry) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AccountEntry) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountEntry) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *AccountEntry) GetAccountType() schema.AccountType {
	if m != nil {
		return m.AccountType
	}
	return schema.AccountType_USER
}

type SignUpRequest struct {
	UserName    string             `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password    string             `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email       string             `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	AccountType schema.AccountType `protobuf:"varint,4,opt,name=account_type,json=accountType,enum=schema.AccountType" json:"account_type,omitempty"`
}

func (m *SignUpRequest) Reset()                    { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string            { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()               {}
func (*SignUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignUpRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetAccountType() schema.AccountType {
	if m != nil {
		return m.AccountType
	}
	return schema.AccountType_USER
}

type SignUpReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *SignUpReply) Reset()                    { *m = SignUpReply{} }
func (m *SignUpReply) String() string            { return proto.CompactTextString(m) }
func (*SignUpReply) ProtoMessage()               {}
func (*SignUpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SignUpReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerificationRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *VerificationRequest) Reset()                    { *m = VerificationRequest{} }
func (m *VerificationRequest) String() string            { return proto.CompactTextString(m) }
func (*VerificationRequest) ProtoMessage()               {}
func (*VerificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VerificationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogInRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LogInRequest) Reset()                    { *m = LogInRequest{} }
func (m *LogInRequest) String() string            { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()               {}
func (*LogInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogInRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogInReply) Reset()                    { *m = LogInReply{} }
func (m *LogInReply) String() string            { return proto.CompactTextString(m) }
func (*LogInReply) ProtoMessage()               {}
func (*LogInReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogInReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordResetRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *PasswordResetRequest) Reset()                    { *m = PasswordResetRequest{} }
func (m *PasswordResetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordResetRequest) ProtoMessage()               {}
func (*PasswordResetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PasswordResetRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type PasswordResetReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *PasswordResetReply) Reset()                    { *m = PasswordResetReply{} }
func (m *PasswordResetReply) String() string            { return proto.CompactTextString(m) }
func (*PasswordResetReply) ProtoMessage()               {}
func (*PasswordResetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PasswordResetReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordSetRequest struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *PasswordSetRequest) Reset()                    { *m = PasswordSetRequest{} }
func (m *PasswordSetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordSetRequest) ProtoMessage()               {}
func (*PasswordSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PasswordSetRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PasswordSetRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PasswordChangeRequest struct {
	UserName         string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	ExistingPassword string `protobuf:"bytes,2,opt,name=existingPassword" json:"existingPassword,omitempty"`
	NewPassword      string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
}

func (m *PasswordChangeRequest) Reset()                    { *m = PasswordChangeRequest{} }
func (m *PasswordChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordChangeRequest) ProtoMessage()               {}
func (*PasswordChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PasswordChangeRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PasswordChangeRequest) GetExistingPassword() string {
	if m != nil {
		return m.ExistingPassword
	}
	return ""
}

func (m *PasswordChangeRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountEntry)(nil), "account.AccountEntry")
	proto.RegisterType((*SignUpRequest)(nil), "account.SignUpRequest")
	proto.RegisterType((*SignUpReply)(nil), "account.SignUpReply")
	proto.RegisterType((*VerificationRequest)(nil), "account.VerificationRequest")
	proto.RegisterType((*LogInRequest)(nil), "account.LogInRequest")
	proto.RegisterType((*LogInReply)(nil), "account.LogInReply")
	proto.RegisterType((*PasswordResetRequest)(nil), "account.PasswordResetRequest")
	proto.RegisterType((*PasswordResetReply)(nil), "account.PasswordResetReply")
	proto.RegisterType((*PasswordSetRequest)(nil), "account.PasswordSetRequest")
	proto.RegisterType((*PasswordChangeRequest)(nil), "account.PasswordChangeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Account service

type AccountClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error)
	Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInReply, error)
	PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetReply, error)
	PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpReply, error) {
	out := new(SignUpReply)
	err := grpc.Invoke(ctx, "/account.Account/SignUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInReply, error) {
	out := new(LogInReply)
	err := grpc.Invoke(ctx, "/account.Account/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetReply, error) {
	out := new(PasswordResetReply)
	err := grpc.Invoke(ctx, "/account.Account/PasswordReset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpReply, error)
	Verify(context.Context, *VerificationRequest) (*google_protobuf1.Empty, error)
	Login(context.Context, *LogInRequest) (*LogInReply, error)
	PasswordReset(context.Context, *PasswordResetRequest) (*PasswordResetReply, error)
	PasswordSet(context.Context, *PasswordSetRequest) (*google_protobuf1.Empty, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*google_protobuf1.Empty, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Verify(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordReset(ctx, req.(*PasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordSet(ctx, req.(*PasswordSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordChange(ctx, req.(*PasswordChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Account_SignUp_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Account_Verify_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "PasswordReset",
			Handler:    _Account_PasswordReset_Handler,
		},
		{
			MethodName: "PasswordSet",
			Handler:    _Account_PasswordSet_Handler,
		},
		{
			MethodName: "PasswordChange",
			Handler:    _Account_PasswordChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/account/account.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x18, 0x25, 0xfd, 0xef, 0xb7, 0x6d, 0xd1, 0xe9, 0xdf, 0x92, 0x6e, 0x65, 0x19, 0xa9, 0x2c, 0x5b,
	0x48, 0x6c, 0x0b, 0x82, 0x05, 0x2f, 0x44, 0xaa, 0x08, 0x45, 0xea, 0xd6, 0x7a, 0x27, 0xcb, 0x6c,
	0x76, 0x36, 0x1b, 0x4c, 0x66, 0xc6, 0x64, 0xd2, 0x1a, 0x4a, 0x6f, 0xfa, 0x0a, 0x82, 0xd7, 0xbe,
	0x93, 0xaf, 0xe0, 0x33, 0x78, 0x2d, 0x99, 0xfc, 0x77, 0x7f, 0x2c, 0x7a, 0x95, 0x7c, 0x67, 0xbe,
	0x39, 0xe7, 0xcc, 0xcc, 0x99, 0x81, 0xe7, 0xb6, 0x23, 0x87, 0x61, 0xcf, 0xb0, 0xb8, 0x67, 0x12,
	0x21, 0x2c, 0xea, 0x52, 0x9f, 0x48, 0xee, 0x9b, 0xc4, 0x13, 0x26, 0x11, 0x8e, 0xe9, 0x0b, 0xcb,
	0x24, 0x96, 0xc5, 0x43, 0x26, 0xb3, 0xaf, 0x21, 0x7c, 0x2e, 0x39, 0x5a, 0x4c, 0x4b, 0xbd, 0x61,
	0x73, 0x6e, 0xbb, 0x54, 0xb5, 0x13, 0xc6, 0xb8, 0x24, 0xd2, 0xe1, 0x2c, 0x48, 0xda, 0xf4, 0x9d,
	0x74, 0x54, 0x55, 0xbd, 0x70, 0x60, 0x52, 0x4f, 0xc8, 0x28, 0x1d, 0x7c, 0x31, 0x4d, 0xbe, 0x4f,
	0x24, 0xc9, 0xb5, 0x03, 0x6b, 0x48, 0x3d, 0x52, 0xb5, 0x80, 0x7f, 0x68, 0xb0, 0xf2, 0x32, 0x41,
	0x4e, 0x98, 0xf4, 0x23, 0xb4, 0x03, 0xcb, 0x61, 0x40, 0xfd, 0x2e, 0x23, 0x1e, 0xad, 0x6b, 0x4d,
	0xad, 0xb5, 0xdc, 0x59, 0x8a, 0x81, 0x77, 0xc4, 0xa3, 0x68, 0x03, 0xe6, 0xa9, 0x47, 0x1c, 0xb7,
	0x3e, 0xa3, 0x06, 0x92, 0x02, 0xed, 0xc1, 0x9a, 0xfa, 0xe9, 0x5e, 0x52, 0xdf, 0x19, 0x38, 0xb4,
	0x5f, 0x9f, 0x6d, 0x6a, 0xad, 0xa5, 0xce, 0xaa, 0x42, 0x3f, 0xa6, 0x20, 0x7a, 0x06, 0x2b, 0xa9,
	0x76, 0x57, 0x46, 0x82, 0xd6, 0xe7, 0x9a, 0x5a, 0x6b, 0xed, 0x70, 0xdd, 0x48, 0x7c, 0x19, 0xa9,
	0x8b, 0x0f, 0x91, 0xa0, 0x9d, 0x1a, 0x29, 0x0a, 0xfc, 0x5d, 0x83, 0xd5, 0x73, 0xc7, 0x66, 0x17,
	0xa2, 0x43, 0xbf, 0x84, 0x34, 0x90, 0xd3, 0x3d, 0xea, 0xb0, 0x24, 0x48, 0x10, 0x5c, 0x71, 0xbf,
	0x9f, 0xda, 0xcc, 0xeb, 0xc2, 0xff, 0x6c, 0xd9, 0xff, 0xbf, 0x1a, 0x7b, 0x0c, 0xb5, 0xcc, 0x97,
	0x70, 0xa3, 0x98, 0x5c, 0xf2, 0xcf, 0x94, 0xa5, 0x8e, 0x92, 0x02, 0xef, 0xc3, 0x7a, 0xb2, 0x03,
	0x96, 0x3a, 0xd3, 0x6c, 0x09, 0xe3, 0x9b, 0xdf, 0xc0, 0xca, 0x29, 0xb7, 0xdf, 0xb2, 0xff, 0x5d,
	0x28, 0xc6, 0x00, 0x29, 0xd1, 0x64, 0x67, 0x47, 0xb0, 0x71, 0x96, 0xf6, 0x77, 0x68, 0x40, 0xe5,
	0x7d, 0x44, 0x71, 0x1b, 0xd0, 0x9d, 0x49, 0x93, 0x05, 0x5e, 0x17, 0xbd, 0xe7, 0x05, 0xfd, 0xd8,
	0xde, 0xa9, 0x8b, 0xb9, 0xd5, 0x60, 0x33, 0x23, 0x7a, 0x35, 0x24, 0xcc, 0xa6, 0xf7, 0xda, 0x9f,
	0x36, 0x3c, 0xa0, 0x5f, 0x9d, 0x40, 0x3a, 0xcc, 0x3e, 0xab, 0x52, 0x8f, 0xe0, 0xa8, 0x09, 0x35,
	0x46, 0xaf, 0xf2, 0xb6, 0x24, 0x1e, 0x65, 0xe8, 0xf0, 0xf7, 0x1c, 0x2c, 0xa6, 0x49, 0x40, 0x17,
	0xb0, 0x90, 0x1c, 0x3c, 0xda, 0x32, 0xb2, 0xeb, 0x54, 0x49, 0xa8, 0xbe, 0x31, 0x82, 0x0b, 0x37,
	0xc2, 0xbb, 0xb7, 0x3f, 0x7f, 0x7d, 0x9b, 0xd9, 0xc6, 0xc8, 0xbc, 0x3c, 0x28, 0xae, 0xa5, 0x63,
	0xb3, 0x50, 0x1c, 0x6b, 0x6d, 0xf4, 0x09, 0x16, 0x54, 0x54, 0x22, 0xd4, 0xc8, 0xa7, 0x8f, 0xc9,
	0x8e, 0xbe, 0x65, 0x24, 0x0f, 0x82, 0x91, 0x3d, 0x08, 0xc6, 0x49, 0xfc, 0x20, 0x8c, 0xa7, 0x57,
	0x77, 0x32, 0x8a, 0xe9, 0xdf, 0xc3, 0xfc, 0x29, 0xb7, 0x1d, 0x86, 0x36, 0x73, 0xf6, 0x72, 0xd8,
	0xf4, 0xf5, 0xbb, 0x70, 0x6c, 0xb9, 0xa1, 0x38, 0xb7, 0xf0, 0xc3, 0x32, 0xa7, 0x1b, 0xd3, 0xc4,
	0x94, 0x37, 0xb0, 0x5a, 0x49, 0x03, 0xda, 0xcd, 0x39, 0xc6, 0x45, 0x4b, 0xdf, 0x99, 0x34, 0x1c,
	0x4b, 0x3d, 0x55, 0x52, 0x6d, 0xbc, 0x57, 0x96, 0xba, 0xce, 0xcf, 0xf7, 0xc6, 0x14, 0xe5, 0x39,
	0xb1, 0xfc, 0x00, 0x6a, 0xa5, 0x80, 0xa1, 0x51, 0xf6, 0x22, 0x76, 0x13, 0x37, 0x0d, 0x2b, 0xd5,
	0x06, 0xde, 0x2e, 0xab, 0x8a, 0x62, 0x7e, 0xac, 0x73, 0x0d, 0x6b, 0xd5, 0xfc, 0xa1, 0x47, 0x23,
	0x52, 0x95, 0x60, 0x4e, 0x54, 0x3b, 0x50, 0x6a, 0xfb, 0xf8, 0xc9, 0xdf, 0xd6, 0x98, 0xd0, 0x1d,
	0x6b, 0xed, 0xde, 0x82, 0xa2, 0x38, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x66, 0xe7, 0x32,
	0x68, 0x06, 0x00, 0x00,
}
