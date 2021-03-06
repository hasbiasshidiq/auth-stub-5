// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth.proto

/*
Package proto_auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth.proto

It has these top-level messages:
	CreateTokenRequest
	IntrospectTokenRequest
	CreateTokenResponse
	IntrospectTokenResponse
*/
package auth_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthStatusCode int32

const (
	AuthStatusCode_SUCCESS               AuthStatusCode = 0
	AuthStatusCode_INTERNAL_SERVER_ERROR AuthStatusCode = 1
	AuthStatusCode_INVALID_TOKEN         AuthStatusCode = 2
	AuthStatusCode_EXPIRATE_TOKEN        AuthStatusCode = 3
)

var AuthStatusCode_name = map[int32]string{
	0: "SUCCESS",
	1: "INTERNAL_SERVER_ERROR",
	2: "INVALID_TOKEN",
	3: "EXPIRATE_TOKEN",
}
var AuthStatusCode_value = map[string]int32{
	"SUCCESS":               0,
	"INTERNAL_SERVER_ERROR": 1,
	"INVALID_TOKEN":         2,
	"EXPIRATE_TOKEN":        3,
}

func (x AuthStatusCode) String() string {
	return proto.EnumName(AuthStatusCode_name, int32(x))
}
func (AuthStatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateTokenRequest struct {
	Issuer       string `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	Iat          string `protobuf:"bytes,2,opt,name=iat" json:"iat,omitempty"`
	Exp          string `protobuf:"bytes,3,opt,name=exp" json:"exp,omitempty"`
	ReferralLink string `protobuf:"bytes,4,opt,name=referral_link,json=referralLink" json:"referral_link,omitempty"`
	Role         string `protobuf:"bytes,5,opt,name=role" json:"role,omitempty"`
}

func (m *CreateTokenRequest) Reset()                    { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()               {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTokenRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *CreateTokenRequest) GetIat() string {
	if m != nil {
		return m.Iat
	}
	return ""
}

func (m *CreateTokenRequest) GetExp() string {
	if m != nil {
		return m.Exp
	}
	return ""
}

func (m *CreateTokenRequest) GetReferralLink() string {
	if m != nil {
		return m.ReferralLink
	}
	return ""
}

func (m *CreateTokenRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type IntrospectTokenRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *IntrospectTokenRequest) Reset()                    { *m = IntrospectTokenRequest{} }
func (m *IntrospectTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*IntrospectTokenRequest) ProtoMessage()               {}
func (*IntrospectTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IntrospectTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type CreateTokenResponse struct {
	StatusCode  AuthStatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,enum=auth.AuthStatusCode" json:"status_code,omitempty"`
	AccessToken string         `protobuf:"bytes,2,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *CreateTokenResponse) Reset()                    { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()               {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateTokenResponse) GetStatusCode() AuthStatusCode {
	if m != nil {
		return m.StatusCode
	}
	return AuthStatusCode_SUCCESS
}

func (m *CreateTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type IntrospectTokenResponse struct {
	StatusCode   AuthStatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,enum=auth.AuthStatusCode" json:"status_code,omitempty"`
	ReferralLink string         `protobuf:"bytes,2,opt,name=referral_link,json=referralLink" json:"referral_link,omitempty"`
	Role         string         `protobuf:"bytes,3,opt,name=role" json:"role,omitempty"`
}

func (m *IntrospectTokenResponse) Reset()                    { *m = IntrospectTokenResponse{} }
func (m *IntrospectTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*IntrospectTokenResponse) ProtoMessage()               {}
func (*IntrospectTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IntrospectTokenResponse) GetStatusCode() AuthStatusCode {
	if m != nil {
		return m.StatusCode
	}
	return AuthStatusCode_SUCCESS
}

func (m *IntrospectTokenResponse) GetReferralLink() string {
	if m != nil {
		return m.ReferralLink
	}
	return ""
}

func (m *IntrospectTokenResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "auth.CreateTokenRequest")
	proto.RegisterType((*IntrospectTokenRequest)(nil), "auth.IntrospectTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "auth.CreateTokenResponse")
	proto.RegisterType((*IntrospectTokenResponse)(nil), "auth.IntrospectTokenResponse")
	proto.RegisterEnum("auth.AuthStatusCode", AuthStatusCode_name, AuthStatusCode_value)
}

func init() { proto.RegisterFile("proto/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xdf, 0x8e, 0x93, 0x40,
	0x14, 0xc6, 0x05, 0xea, 0x1a, 0x0f, 0xdd, 0x8a, 0x47, 0x5d, 0xd9, 0x8d, 0x26, 0x5a, 0x6f, 0x8c,
	0x17, 0x6b, 0xb2, 0xc6, 0xab, 0xbd, 0xa2, 0x38, 0x17, 0xc4, 0x86, 0x9a, 0x01, 0x1b, 0xe3, 0x0d,
	0x41, 0x7a, 0x4c, 0x49, 0x1b, 0x06, 0x67, 0x86, 0xa4, 0xaf, 0xd0, 0x57, 0xf0, 0x69, 0x0d, 0x03,
	0xfe, 0xa9, 0x54, 0x6f, 0xf6, 0xee, 0xe3, 0xfb, 0x0e, 0x73, 0xce, 0xfc, 0xe6, 0x80, 0x57, 0x4b,
	0xa1, 0xc5, 0xeb, 0xbc, 0xd1, 0xeb, 0x4b, 0x23, 0x71, 0xd4, 0xea, 0xe9, 0xde, 0x02, 0x0c, 0x25,
	0xe5, 0x9a, 0x52, 0xb1, 0xa1, 0x8a, 0xd3, 0xb7, 0x86, 0x94, 0xc6, 0x33, 0x38, 0x29, 0x95, 0x6a,
	0x48, 0xfa, 0xd6, 0x33, 0xeb, 0xe5, 0x5d, 0xde, 0x7f, 0xa1, 0x07, 0x4e, 0x99, 0x6b, 0xdf, 0x36,
	0x66, 0x2b, 0x5b, 0x87, 0x76, 0xb5, 0xef, 0x74, 0x0e, 0xed, 0x6a, 0x7c, 0x01, 0xa7, 0x92, 0xbe,
	0x92, 0x94, 0xf9, 0x36, 0xdb, 0x96, 0xd5, 0xc6, 0x1f, 0x99, 0x6c, 0xfc, 0xd3, 0x9c, 0x97, 0xd5,
	0x06, 0x11, 0x46, 0x52, 0x6c, 0xc9, 0xbf, 0x6d, 0x32, 0xa3, 0xa7, 0xd7, 0x70, 0x16, 0x55, 0x5a,
	0x0a, 0x55, 0x53, 0xa1, 0x0f, 0xc6, 0x79, 0x0e, 0xe3, 0xbc, 0x28, 0x48, 0xa9, 0x4c, 0xb7, 0x76,
	0x3f, 0x94, 0xdb, 0x79, 0xa6, 0x72, 0x2a, 0xe0, 0xc1, 0xc1, 0x3d, 0x54, 0x2d, 0x2a, 0x45, 0xf8,
	0x16, 0x5c, 0xa5, 0x73, 0xdd, 0xa8, 0xac, 0x10, 0x2b, 0x32, 0x3f, 0x4e, 0xae, 0x1e, 0x5e, 0x1a,
	0x0e, 0x41, 0xa3, 0xd7, 0x89, 0x09, 0x43, 0xb1, 0x22, 0x0e, 0xea, 0x97, 0x1e, 0x34, 0xb4, 0x87,
	0x0d, 0xf7, 0x16, 0x3c, 0x1e, 0x8c, 0x7b, 0xb3, 0xae, 0x03, 0x72, 0xf6, 0x7f, 0xc8, 0x39, 0xbf,
	0xc9, 0xbd, 0xca, 0x60, 0x72, 0x78, 0x2c, 0xba, 0x70, 0x27, 0xf9, 0x18, 0x86, 0x2c, 0x49, 0xbc,
	0x5b, 0x78, 0x0e, 0x8f, 0xa2, 0x38, 0x65, 0x3c, 0x0e, 0xe6, 0x59, 0xc2, 0xf8, 0x92, 0xf1, 0x8c,
	0x71, 0xbe, 0xe0, 0x9e, 0x85, 0xf7, 0xe1, 0x34, 0x8a, 0x97, 0xc1, 0x3c, 0x7a, 0x97, 0xa5, 0x8b,
	0xf7, 0x2c, 0xf6, 0x6c, 0x44, 0x98, 0xb0, 0x4f, 0x1f, 0x22, 0x1e, 0xa4, 0xac, 0xf7, 0x9c, 0xab,
	0xef, 0x16, 0x98, 0x7d, 0xc1, 0x19, 0xb8, 0x7f, 0x60, 0x46, 0xbf, 0xbb, 0xd3, 0x70, 0x83, 0x2e,
	0xce, 0x8f, 0x24, 0x3d, 0x9d, 0x18, 0xee, 0xfd, 0x05, 0x0e, 0x9f, 0x74, 0xd5, 0xc7, 0x9f, 0xff,
	0xe2, 0xe9, 0x3f, 0xd2, 0xee, 0xbc, 0xd9, 0xf8, 0x33, 0x98, 0x95, 0xbe, 0x6e, 0xab, 0xbe, 0x9c,
	0x18, 0xfd, 0xe6, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xff, 0x69, 0xa6, 0xf2, 0x02, 0x00,
	0x00,
}
