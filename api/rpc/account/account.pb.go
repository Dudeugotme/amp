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
	// Specific to organization accounts
	OwnerId string `protobuf:"bytes,5,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
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

func (m *AccountEntry) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
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

// LogIn function
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

// PasswordChange function
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
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0x4e, 0xf9, 0x5d, 0xce, 0x02, 0xd1, 0xe1, 0x6f, 0x2d, 0x8b, 0xd9, 0x8c, 0xc1, 0x90, 0x25,
	0x69, 0x05, 0x12, 0x13, 0x49, 0xbc, 0x30, 0x06, 0x0d, 0x09, 0x31, 0xb8, 0x88, 0x77, 0x66, 0x33,
	0x74, 0x87, 0x32, 0x71, 0x3b, 0x33, 0xb6, 0x53, 0xb0, 0x21, 0xdc, 0xf0, 0x0a, 0x26, 0xbe, 0x8c,
	0x8f, 0xe1, 0x2b, 0xf8, 0x0c, 0x5e, 0x9b, 0x4e, 0xff, 0xd9, 0x1f, 0x89, 0x5e, 0xb5, 0xe7, 0xcc,
	0x39, 0xdf, 0xf7, 0x9d, 0x99, 0x6f, 0x06, 0x5e, 0xb8, 0x4c, 0x5d, 0x84, 0x67, 0x96, 0x23, 0x3c,
	0x9b, 0x48, 0xe9, 0xd0, 0x3e, 0xf5, 0x89, 0x12, 0xbe, 0x4d, 0x3c, 0x69, 0x13, 0xc9, 0x6c, 0x5f,
	0x3a, 0x36, 0x71, 0x1c, 0x11, 0x72, 0x95, 0x7d, 0x2d, 0xe9, 0x0b, 0x25, 0xd0, 0x6c, 0x1a, 0x9a,
	0x4d, 0x57, 0x08, 0xb7, 0x4f, 0x75, 0x39, 0xe1, 0x5c, 0x28, 0xa2, 0x98, 0xe0, 0x41, 0x52, 0x66,
	0xae, 0xa7, 0xab, 0x3a, 0x3a, 0x0b, 0xcf, 0x6d, 0xea, 0x49, 0x15, 0xa5, 0x8b, 0x2f, 0xc7, 0xd1,
	0xf7, 0x88, 0x22, 0x39, 0x77, 0xe0, 0x5c, 0x50, 0x8f, 0x54, 0x25, 0xe0, 0x1f, 0x06, 0xcc, 0xbf,
	0x4a, 0x32, 0x07, 0x5c, 0xf9, 0x11, 0x5a, 0x87, 0xb9, 0x30, 0xa0, 0x7e, 0x97, 0x13, 0x8f, 0x36,
	0x8c, 0x96, 0xb1, 0x35, 0xd7, 0xa9, 0xc5, 0x89, 0x77, 0xc4, 0xa3, 0x68, 0x19, 0xa6, 0xa9, 0x47,
	0x58, 0xbf, 0x31, 0xa1, 0x17, 0x92, 0x00, 0x6d, 0xc2, 0xa2, 0xfe, 0xe9, 0x5e, 0x52, 0x9f, 0x9d,
	0x33, 0xda, 0x6b, 0x4c, 0xb6, 0x8c, 0xad, 0x5a, 0x67, 0x41, 0x67, 0x3f, 0xa6, 0x49, 0xf4, 0x1c,
	0xe6, 0x53, 0xee, 0xae, 0x8a, 0x24, 0x6d, 0x4c, 0xb5, 0x8c, 0xad, 0xc5, 0xdd, 0x25, 0x2b, 0xd1,
	0x65, 0xa5, 0x2a, 0x3e, 0x44, 0x92, 0x76, 0xea, 0xa4, 0x08, 0xd0, 0x23, 0xa8, 0x89, 0x2b, 0x4e,
	0xfd, 0x2e, 0xeb, 0x35, 0xa6, 0x35, 0xef, 0xac, 0x8e, 0x0f, 0x7b, 0xf8, 0xbb, 0x01, 0x0b, 0x27,
	0xcc, 0xe5, 0xa7, 0xb2, 0x43, 0xbf, 0x84, 0x34, 0x50, 0xe3, 0xe5, 0x9b, 0x50, 0x93, 0x24, 0x08,
	0xae, 0x84, 0xdf, 0x4b, 0x27, 0xc8, 0xe3, 0x62, 0xb4, 0xc9, 0xf2, 0x68, 0xff, 0xa8, 0x19, 0x3f,
	0x81, 0x7a, 0xa6, 0x4b, 0xf6, 0xa3, 0x18, 0x5c, 0x89, 0xcf, 0x94, 0xa7, 0x8a, 0x92, 0x00, 0x6f,
	0xc3, 0x52, 0xb2, 0x39, 0x8e, 0x3e, 0xee, 0x6c, 0x84, 0xe1, 0xc5, 0x6f, 0x61, 0xfe, 0x48, 0xb8,
	0x87, 0xfc, 0x7f, 0x07, 0xc5, 0x18, 0x20, 0x05, 0x1a, 0xad, 0x6c, 0x0f, 0x96, 0x8f, 0xd3, 0xfa,
	0x0e, 0x0d, 0xa8, 0xba, 0x0f, 0x29, 0x6e, 0x03, 0xba, 0xd3, 0x34, 0x9a, 0xe0, 0x4d, 0x51, 0x7b,
	0x52, 0xc0, 0x0f, 0xad, 0x1d, 0x3b, 0xcc, 0xad, 0x01, 0x2b, 0x19, 0xd0, 0xeb, 0x0b, 0xc2, 0x5d,
	0x7a, 0xaf, 0xfd, 0x69, 0xc3, 0x03, 0xfa, 0x95, 0x05, 0x8a, 0x71, 0xf7, 0xb8, 0x0a, 0x3d, 0x90,
	0x47, 0x2d, 0xa8, 0x73, 0x7a, 0x95, 0x97, 0x25, 0xf6, 0x28, 0xa7, 0x76, 0x7f, 0x4f, 0xc1, 0x6c,
	0xea, 0x04, 0x74, 0x0a, 0x33, 0xc9, 0xc1, 0xa3, 0x55, 0x2b, 0xbb, 0x69, 0x15, 0x87, 0x9a, 0xcb,
	0x03, 0x79, 0xd9, 0x8f, 0xf0, 0xc6, 0xed, 0xcf, 0x5f, 0xdf, 0x26, 0xd6, 0x30, 0xb2, 0x2f, 0x77,
	0x8a, 0x1b, 0xcb, 0x5c, 0x1e, 0xca, 0x7d, 0xa3, 0x8d, 0x3e, 0xc1, 0x8c, 0xb6, 0x4a, 0x84, 0x9a,
	0x79, 0xfb, 0x10, 0xef, 0x98, 0xab, 0x56, 0xf2, 0x56, 0x58, 0xd9, 0x5b, 0x61, 0x1d, 0xc4, 0x6f,
	0xc5, 0x70, 0x78, 0x7d, 0x5d, 0xa3, 0x18, 0xfe, 0x3d, 0x4c, 0x1f, 0x09, 0x97, 0x71, 0xb4, 0x92,
	0xa3, 0x97, 0xcd, 0x66, 0x2e, 0xdd, 0x4d, 0xc7, 0x92, 0x9b, 0x1a, 0x73, 0x15, 0x3f, 0x2c, 0x63,
	0xf6, 0x63, 0x98, 0x18, 0xf2, 0x06, 0x16, 0x2a, 0x6e, 0x40, 0x1b, 0x39, 0xc6, 0x30, 0x6b, 0x99,
	0xeb, 0xa3, 0x96, 0x63, 0xaa, 0x67, 0x9a, 0xaa, 0x8d, 0x37, 0xcb, 0x54, 0xd7, 0xf9, 0xf9, 0xde,
	0xd8, 0xb2, 0xdc, 0x13, 0xd3, 0x9f, 0x43, 0xbd, 0x64, 0x30, 0x34, 0x88, 0x5e, 0xd8, 0x6e, 0xe4,
	0xa6, 0x61, 0xcd, 0xda, 0xc4, 0x6b, 0x65, 0x56, 0x59, 0xf4, 0xc7, 0x3c, 0xd7, 0xb0, 0x58, 0xf5,
	0x1f, 0x7a, 0x3c, 0x40, 0x55, 0x31, 0xe6, 0x48, 0xb6, 0x1d, 0xcd, 0xb6, 0x8d, 0x9f, 0xfe, 0x6d,
	0xc6, 0x04, 0x6e, 0xdf, 0x68, 0x9f, 0xcd, 0x68, 0x88, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0x4f, 0xca, 0x18, 0x83, 0x06, 0x00, 0x00,
}
