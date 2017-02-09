// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/data/schema/account/account.proto
// DO NOT EDIT!

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/data/schema/account/account.proto

It has these top-level messages:
	Account
	OrganizationMember
	Resource
	ResourceSettings
	Team
	TeamMember
	Permission
	Key
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account Schema Storage Types
type AccountType int32

const (
	AccountType_USER         AccountType = 0
	AccountType_ORGANIZATION AccountType = 1
)

var AccountType_name = map[int32]string{
	0: "USER",
	1: "ORGANIZATION",
}
var AccountType_value = map[string]int32{
	"USER":         0,
	"ORGANIZATION": 1,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}
func (AccountType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Reference Data
type ResourceType int32

const (
	ResourceType_STACK   ResourceType = 0
	ResourceType_SERVICE ResourceType = 1
)

var ResourceType_name = map[int32]string{
	0: "STACK",
	1: "SERVICE",
}
var ResourceType_value = map[string]int32{
	"STACK":   0,
	"SERVICE": 1,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GrantType int32

const (
	GrantType_READ   GrantType = 0
	GrantType_WRITE  GrantType = 1
	GrantType_DELETE GrantType = 3
	GrantType_ALL    GrantType = 4
)

var GrantType_name = map[int32]string{
	0: "READ",
	1: "WRITE",
	3: "DELETE",
	4: "ALL",
}
var GrantType_value = map[string]int32{
	"READ":   0,
	"WRITE":  1,
	"DELETE": 3,
	"ALL":    4,
}

func (x GrantType) String() string {
	return proto.EnumName(GrantType_name, int32(x))
}
func (GrantType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Account struct {
	Id           string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type         AccountType `protobuf:"varint,3,opt,name=type,enum=schema.AccountType" json:"type,omitempty"`
	Email        string      `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	PasswordHash string      `protobuf:"bytes,5,opt,name=password_hash,json=passwordHash" json:"password_hash,omitempty"`
	IsVerified   bool        `protobuf:"varint,6,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	// TODO Billing Attributes;
	CreateDt *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_USER
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *Account) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *Account) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type OrganizationMember struct {
	Id            string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId  string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	UserAccountId string                     `protobuf:"bytes,3,opt,name=user_account_id,json=userAccountId" json:"user_account_id,omitempty"`
	CreateDt      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *OrganizationMember) Reset()                    { *m = OrganizationMember{} }
func (m *OrganizationMember) String() string            { return proto.CompactTextString(m) }
func (*OrganizationMember) ProtoMessage()               {}
func (*OrganizationMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrganizationMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrganizationMember) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *OrganizationMember) GetUserAccountId() string {
	if m != nil {
		return m.UserAccountId
	}
	return ""
}

func (m *OrganizationMember) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Resource struct {
	Id           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	TeamId       string                     `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	Name         string                     `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Type         ResourceType               `protobuf:"varint,5,opt,name=type,enum=schema.ResourceType" json:"type,omitempty"`
	CreateDt     *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *Resource) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_STACK
}

func (m *Resource) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type ResourceSettings struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ResourceId string                     `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	Key        string                     `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value      string                     `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	CreateDt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *ResourceSettings) Reset()                    { *m = ResourceSettings{} }
func (m *ResourceSettings) String() string            { return proto.CompactTextString(m) }
func (*ResourceSettings) ProtoMessage()               {}
func (*ResourceSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResourceSettings) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceSettings) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ResourceSettings) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResourceSettings) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ResourceSettings) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Team struct {
	Id           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	Name         string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Desc         string                     `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	CreateDt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Team) Reset()                    { *m = Team{} }
func (m *Team) String() string            { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()               {}
func (*Team) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Team) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type TeamMember struct {
	Id            string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserAccountId string                     `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId" json:"user_account_id,omitempty"`
	TeamId        string                     `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	CreateDt      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *TeamMember) Reset()                    { *m = TeamMember{} }
func (m *TeamMember) String() string            { return proto.CompactTextString(m) }
func (*TeamMember) ProtoMessage()               {}
func (*TeamMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TeamMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TeamMember) GetUserAccountId() string {
	if m != nil {
		return m.UserAccountId
	}
	return ""
}

func (m *TeamMember) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *TeamMember) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Permission struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TeamId     string                     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	ResourceId string                     `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	GrantType  GrantType                  `protobuf:"varint,4,opt,name=grant_type,json=grantType,enum=schema.GrantType" json:"grant_type,omitempty"`
	CreateDt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (m *Permission) String() string            { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Permission) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Permission) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Permission) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Permission) GetGrantType() GrantType {
	if m != nil {
		return m.GrantType
	}
	return GrantType_READ
}

func (m *Permission) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

// Protobuf Type to store the foreign key relationships
type Key struct {
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Key) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "schema.Account")
	proto.RegisterType((*OrganizationMember)(nil), "schema.OrganizationMember")
	proto.RegisterType((*Resource)(nil), "schema.Resource")
	proto.RegisterType((*ResourceSettings)(nil), "schema.ResourceSettings")
	proto.RegisterType((*Team)(nil), "schema.Team")
	proto.RegisterType((*TeamMember)(nil), "schema.TeamMember")
	proto.RegisterType((*Permission)(nil), "schema.Permission")
	proto.RegisterType((*Key)(nil), "schema.Key")
	proto.RegisterEnum("schema.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("schema.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("schema.GrantType", GrantType_name, GrantType_value)
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/data/schema/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x18, 0xed, 0xc4, 0xce, 0xdf, 0x97, 0xb4, 0xd7, 0x77, 0x6e, 0xaf, 0x88, 0x2a, 0xa4, 0x46, 0x05,
	0x95, 0xd0, 0x85, 0x8d, 0x8a, 0x10, 0x2b, 0x16, 0x51, 0x6b, 0x15, 0xab, 0xa5, 0x45, 0x8e, 0x29,
	0x12, 0x1b, 0x6b, 0x62, 0x4f, 0x9d, 0x51, 0xe3, 0x1f, 0xcd, 0x4c, 0x8a, 0xc2, 0x83, 0xb0, 0x63,
	0xc7, 0x8e, 0xa7, 0xe0, 0x0d, 0x78, 0x0f, 0x5e, 0x02, 0xf9, 0xaf, 0x71, 0x48, 0xbb, 0x68, 0x56,
	0xf9, 0xe6, 0x9b, 0x33, 0x33, 0xe7, 0x7c, 0xe7, 0xc4, 0xf0, 0x26, 0x60, 0x72, 0x32, 0x1b, 0xeb,
	0x5e, 0x1c, 0x1a, 0x24, 0x49, 0x3c, 0x3a, 0xa5, 0x9c, 0xc8, 0x98, 0x1b, 0x24, 0x4c, 0x0c, 0x9f,
	0x48, 0x62, 0x08, 0x6f, 0x42, 0x43, 0x62, 0x10, 0xcf, 0x8b, 0x67, 0x91, 0x2c, 0x7f, 0xf5, 0x84,
	0xc7, 0x32, 0xc6, 0x8d, 0x7c, 0x77, 0x67, 0x37, 0x88, 0xe3, 0x60, 0x4a, 0x8d, 0xac, 0x3b, 0x9e,
	0x5d, 0x19, 0x92, 0x85, 0x54, 0x48, 0x12, 0x26, 0x39, 0x70, 0xef, 0x37, 0x82, 0xe6, 0x30, 0x3f,
	0x8a, 0xb7, 0xa0, 0xc6, 0xfc, 0x1e, 0xea, 0xa3, 0x41, 0xdb, 0xae, 0x31, 0x1f, 0x63, 0x50, 0x23,
	0x12, 0xd2, 0x5e, 0x2d, 0xeb, 0x64, 0x35, 0x7e, 0x06, 0xaa, 0x9c, 0x27, 0xb4, 0xa7, 0xf4, 0xd1,
	0x60, 0xeb, 0xf0, 0x3f, 0x3d, 0x7f, 0x47, 0x2f, 0xae, 0x70, 0xe6, 0x09, 0xb5, 0x33, 0x00, 0xde,
	0x86, 0x3a, 0x0d, 0x09, 0x9b, 0xf6, 0xd4, 0xec, 0x74, 0xbe, 0xc0, 0x4f, 0x60, 0x33, 0x21, 0x42,
	0x7c, 0x8e, 0xb9, 0xef, 0x4e, 0x88, 0x98, 0xf4, 0xea, 0xd9, 0x6e, 0xb7, 0x6c, 0xbe, 0x25, 0x62,
	0x82, 0x77, 0xa1, 0xc3, 0x84, 0x7b, 0x43, 0x39, 0xbb, 0x62, 0xd4, 0xef, 0x35, 0xfa, 0x68, 0xd0,
	0xb2, 0x81, 0x89, 0xcb, 0xa2, 0x83, 0x5f, 0x43, 0xdb, 0xe3, 0x94, 0x48, 0xea, 0xfa, 0xb2, 0xd7,
	0xec, 0xa3, 0x41, 0xe7, 0x70, 0x47, 0xcf, 0x95, 0xea, 0xa5, 0x52, 0xdd, 0x29, 0x95, 0xda, 0xad,
	0x1c, 0x7c, 0x2c, 0xf7, 0x7e, 0x20, 0xc0, 0x17, 0x3c, 0x20, 0x11, 0xfb, 0x42, 0x24, 0x8b, 0xa3,
	0x77, 0x34, 0x1c, 0x53, 0xbe, 0x22, 0xfc, 0x29, 0x6c, 0xc5, 0x3c, 0x70, 0x8b, 0x91, 0xba, 0xcc,
	0x2f, 0x46, 0xd0, 0x8d, 0x79, 0x50, 0x28, 0xb5, 0x7c, 0xbc, 0x0f, 0xff, 0xcc, 0x04, 0xe5, 0x55,
	0x98, 0x92, 0xc1, 0x36, 0xd3, 0xf6, 0x02, 0xb7, 0xc4, 0x56, 0x7d, 0x00, 0xdb, 0x5f, 0x08, 0x5a,
	0x36, 0x15, 0xf1, 0x8c, 0x7b, 0x74, 0x4d, 0x8e, 0x8f, 0xa0, 0x29, 0x29, 0x09, 0x17, 0xdc, 0x1a,
	0xe9, 0xd2, 0x5a, 0x78, 0xab, 0x56, 0xbc, 0x1d, 0x14, 0xde, 0xd6, 0x33, 0x6f, 0xb7, 0x4b, 0x6f,
	0x4b, 0x0a, 0x15, 0x73, 0x97, 0x24, 0x35, 0x1e, 0x20, 0xe9, 0x3b, 0x02, 0xad, 0xbc, 0x6f, 0x44,
	0xa5, 0x64, 0x51, 0x20, 0x56, 0xa4, 0xed, 0x42, 0x87, 0x17, 0x98, 0x85, 0x2e, 0x28, 0x5b, 0x96,
	0x8f, 0x35, 0x50, 0xae, 0xe9, 0xbc, 0x50, 0x94, 0x96, 0x69, 0xda, 0x6e, 0xc8, 0x74, 0x56, 0xea,
	0xc9, 0x17, 0xcb, 0x34, 0xeb, 0x0f, 0xa0, 0xf9, 0x0d, 0x81, 0xea, 0x50, 0x12, 0xae, 0x39, 0xf5,
	0x72, 0xb8, 0x4a, 0x65, 0xb8, 0x18, 0x54, 0x9f, 0x0a, 0xaf, 0x1c, 0x78, 0x5a, 0xaf, 0xcf, 0xef,
	0x2b, 0x02, 0x48, 0xf9, 0xdd, 0x93, 0xdf, 0x3b, 0x92, 0x59, 0xbb, 0x2b, 0x99, 0xf7, 0xa6, 0x63,
	0xed, 0xc8, 0xfe, 0x44, 0x00, 0xef, 0x29, 0x0f, 0x99, 0x10, 0x2c, 0x8e, 0x56, 0x88, 0x55, 0x1e,
	0xac, 0x2d, 0x3d, 0xf8, 0x97, 0xe5, 0xca, 0x8a, 0xe5, 0x2f, 0x00, 0x02, 0x4e, 0x22, 0xe9, 0x66,
	0x09, 0x55, 0xb3, 0x84, 0xfe, 0x5b, 0x26, 0xf4, 0x24, 0xdd, 0xc9, 0xe2, 0xd9, 0x0e, 0xca, 0x72,
	0xfd, 0xe1, 0x3e, 0x06, 0xe5, 0x94, 0xce, 0xf1, 0xff, 0xd0, 0xb8, 0xa6, 0x73, 0xf7, 0x96, 0x7f,
	0xfd, 0x9a, 0xce, 0x2d, 0xff, 0xe0, 0x39, 0x74, 0x2a, 0x1f, 0x3b, 0xdc, 0x02, 0xf5, 0xc3, 0xc8,
	0xb4, 0xb5, 0x0d, 0xac, 0x41, 0xf7, 0xc2, 0x3e, 0x19, 0x9e, 0x5b, 0x9f, 0x86, 0x8e, 0x75, 0x71,
	0xae, 0xa1, 0x83, 0x7d, 0xe8, 0x56, 0xff, 0x3b, 0xb8, 0x0d, 0xf5, 0x91, 0x33, 0x3c, 0x3a, 0xd5,
	0x36, 0x70, 0x07, 0x9a, 0x23, 0xd3, 0xbe, 0xb4, 0x8e, 0x4c, 0x0d, 0x1d, 0xbc, 0x82, 0xf6, 0xad,
	0x82, 0xf4, 0x42, 0xdb, 0x1c, 0x1e, 0x6b, 0x1b, 0x29, 0xfc, 0xa3, 0x6d, 0x39, 0xa6, 0x86, 0x30,
	0x40, 0xe3, 0xd8, 0x3c, 0x33, 0x1d, 0x53, 0x53, 0x70, 0x13, 0x94, 0xe1, 0xd9, 0x99, 0xa6, 0x8e,
	0x1b, 0x99, 0x8a, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x18, 0x3f, 0xcc, 0x2b, 0x06,
	0x00, 0x00,
}
