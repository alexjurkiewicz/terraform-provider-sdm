// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role_grants.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// RoleGrantCreateRequest specifies what kind of RoleGrants should be registered in
// the organizations fleet.
type RoleGrantCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new RoleGrant.
	RoleGrant            *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RoleGrantCreateRequest) Reset()         { *m = RoleGrantCreateRequest{} }
func (m *RoleGrantCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RoleGrantCreateRequest) ProtoMessage()    {}
func (*RoleGrantCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{0}
}

func (m *RoleGrantCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantCreateRequest.Unmarshal(m, b)
}
func (m *RoleGrantCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantCreateRequest.Marshal(b, m, deterministic)
}
func (m *RoleGrantCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantCreateRequest.Merge(m, src)
}
func (m *RoleGrantCreateRequest) XXX_Size() int {
	return xxx_messageInfo_RoleGrantCreateRequest.Size(m)
}
func (m *RoleGrantCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantCreateRequest proto.InternalMessageInfo

func (m *RoleGrantCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantCreateRequest) GetRoleGrant() *RoleGrant {
	if m != nil {
		return m.RoleGrant
	}
	return nil
}

// RoleGrantCreateResponse reports how the RoleGrants were created in the system.
type RoleGrantCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created RoleGrant.
	RoleGrant *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleGrantCreateResponse) Reset()         { *m = RoleGrantCreateResponse{} }
func (m *RoleGrantCreateResponse) String() string { return proto.CompactTextString(m) }
func (*RoleGrantCreateResponse) ProtoMessage()    {}
func (*RoleGrantCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{1}
}

func (m *RoleGrantCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantCreateResponse.Unmarshal(m, b)
}
func (m *RoleGrantCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantCreateResponse.Marshal(b, m, deterministic)
}
func (m *RoleGrantCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantCreateResponse.Merge(m, src)
}
func (m *RoleGrantCreateResponse) XXX_Size() int {
	return xxx_messageInfo_RoleGrantCreateResponse.Size(m)
}
func (m *RoleGrantCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantCreateResponse proto.InternalMessageInfo

func (m *RoleGrantCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantCreateResponse) GetRoleGrant() *RoleGrant {
	if m != nil {
		return m.RoleGrant
	}
	return nil
}

func (m *RoleGrantCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleGrantGetRequest specifies which RoleGrant to retrieve.
type RoleGrantGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleGrant to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleGrantGetRequest) Reset()         { *m = RoleGrantGetRequest{} }
func (m *RoleGrantGetRequest) String() string { return proto.CompactTextString(m) }
func (*RoleGrantGetRequest) ProtoMessage()    {}
func (*RoleGrantGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{2}
}

func (m *RoleGrantGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantGetRequest.Unmarshal(m, b)
}
func (m *RoleGrantGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantGetRequest.Marshal(b, m, deterministic)
}
func (m *RoleGrantGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantGetRequest.Merge(m, src)
}
func (m *RoleGrantGetRequest) XXX_Size() int {
	return xxx_messageInfo_RoleGrantGetRequest.Size(m)
}
func (m *RoleGrantGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantGetRequest proto.InternalMessageInfo

func (m *RoleGrantGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleGrantGetResponse returns a requested RoleGrant.
type RoleGrantGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested RoleGrant.
	RoleGrant *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleGrantGetResponse) Reset()         { *m = RoleGrantGetResponse{} }
func (m *RoleGrantGetResponse) String() string { return proto.CompactTextString(m) }
func (*RoleGrantGetResponse) ProtoMessage()    {}
func (*RoleGrantGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{3}
}

func (m *RoleGrantGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantGetResponse.Unmarshal(m, b)
}
func (m *RoleGrantGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantGetResponse.Marshal(b, m, deterministic)
}
func (m *RoleGrantGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantGetResponse.Merge(m, src)
}
func (m *RoleGrantGetResponse) XXX_Size() int {
	return xxx_messageInfo_RoleGrantGetResponse.Size(m)
}
func (m *RoleGrantGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantGetResponse proto.InternalMessageInfo

func (m *RoleGrantGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantGetResponse) GetRoleGrant() *RoleGrant {
	if m != nil {
		return m.RoleGrant
	}
	return nil
}

func (m *RoleGrantGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleGrantDeleteRequest identifies a RoleGrant by ID to delete.
type RoleGrantDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleGrant to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleGrantDeleteRequest) Reset()         { *m = RoleGrantDeleteRequest{} }
func (m *RoleGrantDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RoleGrantDeleteRequest) ProtoMessage()    {}
func (*RoleGrantDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{4}
}

func (m *RoleGrantDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantDeleteRequest.Unmarshal(m, b)
}
func (m *RoleGrantDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantDeleteRequest.Marshal(b, m, deterministic)
}
func (m *RoleGrantDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantDeleteRequest.Merge(m, src)
}
func (m *RoleGrantDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_RoleGrantDeleteRequest.Size(m)
}
func (m *RoleGrantDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantDeleteRequest proto.InternalMessageInfo

func (m *RoleGrantDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleGrantDeleteResponse returns information about a RoleGrant that was deleted.
type RoleGrantDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleGrantDeleteResponse) Reset()         { *m = RoleGrantDeleteResponse{} }
func (m *RoleGrantDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*RoleGrantDeleteResponse) ProtoMessage()    {}
func (*RoleGrantDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{5}
}

func (m *RoleGrantDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantDeleteResponse.Unmarshal(m, b)
}
func (m *RoleGrantDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantDeleteResponse.Marshal(b, m, deterministic)
}
func (m *RoleGrantDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantDeleteResponse.Merge(m, src)
}
func (m *RoleGrantDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_RoleGrantDeleteResponse.Size(m)
}
func (m *RoleGrantDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantDeleteResponse proto.InternalMessageInfo

func (m *RoleGrantDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleGrantListRequest specifies criteria for retrieving a list of RoleGrants.
type RoleGrantListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleGrantListRequest) Reset()         { *m = RoleGrantListRequest{} }
func (m *RoleGrantListRequest) String() string { return proto.CompactTextString(m) }
func (*RoleGrantListRequest) ProtoMessage()    {}
func (*RoleGrantListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{6}
}

func (m *RoleGrantListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantListRequest.Unmarshal(m, b)
}
func (m *RoleGrantListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantListRequest.Marshal(b, m, deterministic)
}
func (m *RoleGrantListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantListRequest.Merge(m, src)
}
func (m *RoleGrantListRequest) XXX_Size() int {
	return xxx_messageInfo_RoleGrantListRequest.Size(m)
}
func (m *RoleGrantListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantListRequest proto.InternalMessageInfo

func (m *RoleGrantListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// RoleGrantListResponse returns a list of RoleGrants that meet the criteria of a
// RoleGrantListRequest.
type RoleGrantListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	RoleGrants []*RoleGrant `protobuf:"bytes,2,rep,name=role_grants,json=roleGrants,proto3" json:"role_grants,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleGrantListResponse) Reset()         { *m = RoleGrantListResponse{} }
func (m *RoleGrantListResponse) String() string { return proto.CompactTextString(m) }
func (*RoleGrantListResponse) ProtoMessage()    {}
func (*RoleGrantListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{7}
}

func (m *RoleGrantListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrantListResponse.Unmarshal(m, b)
}
func (m *RoleGrantListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrantListResponse.Marshal(b, m, deterministic)
}
func (m *RoleGrantListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrantListResponse.Merge(m, src)
}
func (m *RoleGrantListResponse) XXX_Size() int {
	return xxx_messageInfo_RoleGrantListResponse.Size(m)
}
func (m *RoleGrantListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrantListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrantListResponse proto.InternalMessageInfo

func (m *RoleGrantListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGrantListResponse) GetRoleGrants() []*RoleGrant {
	if m != nil {
		return m.RoleGrants
	}
	return nil
}

func (m *RoleGrantListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// A RoleGrant connects a resource to a role, granting members of the role
// access to that resource.
type RoleGrant struct {
	// Unique identifier of the RoleGrant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the resource of this RoleGrant.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The id of the attached role of this RoleGrant.
	RoleId               string   `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleGrant) Reset()         { *m = RoleGrant{} }
func (m *RoleGrant) String() string { return proto.CompactTextString(m) }
func (*RoleGrant) ProtoMessage()    {}
func (*RoleGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_530cb39924523b64, []int{8}
}

func (m *RoleGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGrant.Unmarshal(m, b)
}
func (m *RoleGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGrant.Marshal(b, m, deterministic)
}
func (m *RoleGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGrant.Merge(m, src)
}
func (m *RoleGrant) XXX_Size() int {
	return xxx_messageInfo_RoleGrant.Size(m)
}
func (m *RoleGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGrant.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGrant proto.InternalMessageInfo

func (m *RoleGrant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoleGrant) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *RoleGrant) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*RoleGrantCreateRequest)(nil), "v1.RoleGrantCreateRequest")
	proto.RegisterType((*RoleGrantCreateResponse)(nil), "v1.RoleGrantCreateResponse")
	proto.RegisterType((*RoleGrantGetRequest)(nil), "v1.RoleGrantGetRequest")
	proto.RegisterType((*RoleGrantGetResponse)(nil), "v1.RoleGrantGetResponse")
	proto.RegisterType((*RoleGrantDeleteRequest)(nil), "v1.RoleGrantDeleteRequest")
	proto.RegisterType((*RoleGrantDeleteResponse)(nil), "v1.RoleGrantDeleteResponse")
	proto.RegisterType((*RoleGrantListRequest)(nil), "v1.RoleGrantListRequest")
	proto.RegisterType((*RoleGrantListResponse)(nil), "v1.RoleGrantListResponse")
	proto.RegisterType((*RoleGrant)(nil), "v1.RoleGrant")
}

func init() { proto.RegisterFile("role_grants.proto", fileDescriptor_530cb39924523b64) }

var fileDescriptor_530cb39924523b64 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xae, 0x83, 0x57, 0x9e, 0xb4, 0x11, 0x4c, 0x48, 0xe2, 0x6e, 0x0b, 0x5d, 0x4d, 0x10,
	0x42, 0x26, 0xf6, 0xca, 0xc6, 0x12, 0xc8, 0x2a, 0x82, 0x84, 0x48, 0x51, 0xa4, 0x20, 0x45, 0xae,
	0x38, 0x16, 0x6b, 0xe2, 0x9d, 0x6c, 0x47, 0x78, 0x77, 0x96, 0x99, 0x71, 0x5c, 0x29, 0xaa, 0x04,
	0x5c, 0xb8, 0x97, 0xff, 0xa0, 0x5c, 0x10, 0x27, 0x24, 0x5f, 0x38, 0x21, 0x54, 0x6e, 0xbd, 0x82,
	0xc4, 0x89, 0x1b, 0x67, 0x0e, 0xf8, 0x42, 0x39, 0xa1, 0x19, 0xef, 0xaf, 0x71, 0x9c, 0x22, 0x94,
	0x13, 0x27, 0x7b, 0xe7, 0xbd, 0xf7, 0x7d, 0xef, 0x7d, 0xf3, 0xed, 0xcc, 0x82, 0x97, 0x38, 0x1b,
	0x91, 0x41, 0xc8, 0x71, 0x2c, 0x45, 0x2b, 0xe1, 0x4c, 0x32, 0x68, 0x9f, 0xb5, 0xdd, 0x5b, 0x21,
	0x63, 0xe1, 0x88, 0xf8, 0x38, 0xa1, 0x3e, 0x8e, 0x63, 0x26, 0xb1, 0xa4, 0x2c, 0x4e, 0x33, 0xdc,
	0x1d, 0xfd, 0x33, 0x6c, 0x86, 0x24, 0x6e, 0x8a, 0x09, 0x0e, 0x43, 0xc2, 0x7d, 0x96, 0xe8, 0x8c,
	0x25, 0xd9, 0xb7, 0x53, 0x2c, 0xfd, 0x74, 0x32, 0x3e, 0xf5, 0x25, 0x8d, 0x88, 0x90, 0x38, 0x4a,
	0xd2, 0x84, 0xeb, 0x69, 0x6d, 0xfa, 0x08, 0x44, 0x42, 0x86, 0xf3, 0xff, 0xe8, 0x73, 0x0b, 0x6c,
	0xf6, 0xd9, 0x88, 0x1c, 0xa8, 0x06, 0x3f, 0xe0, 0x04, 0x4b, 0xd2, 0x27, 0x9f, 0x8e, 0x89, 0x90,
	0xb0, 0x09, 0x56, 0x22, 0x22, 0x71, 0xdd, 0xf2, 0xac, 0x37, 0x56, 0x3b, 0x37, 0x5a, 0x67, 0xed,
	0x96, 0x91, 0xf0, 0x21, 0x91, 0x38, 0xc0, 0x12, 0xf7, 0x75, 0x1a, 0x7c, 0x07, 0x80, 0x62, 0xd4,
	0xba, 0xad, 0x8b, 0xae, 0xab, 0xa2, 0x1c, 0x7e, 0x0f, 0xfc, 0xf9, 0xd7, 0xd4, 0x79, 0xe1, 0xbb,
	0xd9, 0xd4, 0xb1, 0xfa, 0x35, 0x9e, 0x2d, 0xa3, 0xdf, 0x2c, 0xb0, 0x75, 0xa1, 0x07, 0x91, 0xb0,
	0x58, 0x10, 0xd8, 0x33, 0x9a, 0x70, 0xcb, 0x4d, 0xcc, 0x33, 0xb2, 0x2e, 0x0c, 0xf0, 0x2b, 0x76,
	0x04, 0xdf, 0x07, 0x80, 0x63, 0x49, 0x06, 0x23, 0x1a, 0x51, 0x59, 0xaf, 0xe8, 0xca, 0x0d, 0x5d,
	0x89, 0x25, 0x39, 0x52, 0x8b, 0x4b, 0x69, 0x6b, 0x3c, 0x0b, 0xf7, 0xc0, 0xdf, 0x6a, 0xf9, 0x1b,
	0xb5, 0x8c, 0xee, 0x81, 0xf5, 0x9c, 0xf1, 0x80, 0xc8, 0x4c, 0xdf, 0x86, 0x31, 0xda, 0xa6, 0x82,
	0x2f, 0xa2, 0x0b, 0xe2, 0xba, 0xc0, 0xa6, 0x81, 0x1e, 0xa1, 0x66, 0x30, 0xda, 0x34, 0x40, 0xbf,
	0x5a, 0xe0, 0x65, 0x13, 0x3f, 0xd5, 0xee, 0x6d, 0x83, 0x60, 0x2b, 0x27, 0xf8, 0xdf, 0x08, 0x37,
	0x2c, 0x79, 0x73, 0x9f, 0x8c, 0xc8, 0x73, 0xbd, 0x69, 0x24, 0xfc, 0x07, 0xf9, 0xbe, 0x2e, 0xbb,
	0x2f, 0x03, 0xb9, 0xdc, 0x7d, 0x66, 0xc6, 0x73, 0x44, 0x34, 0xa5, 0xb0, 0xaf, 0x28, 0x45, 0x58,
	0xda, 0xe3, 0x23, 0x2a, 0x72, 0x13, 0xbd, 0x79, 0x71, 0x8f, 0x4b, 0xe1, 0x05, 0x19, 0x10, 0xa8,
	0x9e, 0xd2, 0x91, 0x24, 0x7c, 0x89, 0x14, 0x69, 0x04, 0xfd, 0x64, 0x81, 0x8d, 0x05, 0xa6, 0x54,
	0x8c, 0x1d, 0x83, 0xaa, 0x5e, 0x50, 0x99, 0x52, 0xa4, 0x5c, 0x3d, 0xb0, 0x5a, 0x3a, 0xf9, 0xea,
	0xb6, 0x57, 0xb9, 0xcc, 0x44, 0xdf, 0x6b, 0x7e, 0x90, 0x9b, 0x48, 0x5c, 0xdd, 0x45, 0xe8, 0x97,
	0x0a, 0xa8, 0xe5, 0x3c, 0xb0, 0xab, 0xb7, 0xdf, 0xd2, 0x33, 0xbf, 0xa6, 0x0a, 0x6e, 0x3f, 0x9e,
	0x4d, 0x1d, 0xfb, 0x70, 0x5f, 0xd7, 0x3d, 0x9d, 0x4d, 0x9d, 0x35, 0x95, 0x7c, 0x4c, 0x78, 0x44,
	0x85, 0xa0, 0x2c, 0x56, 0xc6, 0x80, 0xc7, 0x60, 0x95, 0x13, 0xc1, 0xc6, 0x7c, 0x48, 0x06, 0xb9,
	0x7b, 0x7c, 0x55, 0xde, 0x50, 0xe5, 0xd7, 0xf6, 0xb1, 0xc4, 0xf3, 0x70, 0x0a, 0xf4, 0x63, 0x86,
	0x06, 0x8a, 0x50, 0x1f, 0x64, 0x18, 0x87, 0x01, 0xbc, 0x03, 0x1c, 0xad, 0x09, 0x0d, 0xf4, 0x50,
	0xb5, 0xbd, 0x6d, 0x85, 0xf6, 0xaa, 0x42, 0xab, 0xaa, 0x06, 0x16, 0x71, 0x56, 0xd4, 0x62, 0xbf,
	0xaa, 0x6a, 0x0e, 0x83, 0xde, 0x67, 0xf6, 0xa3, 0xdd, 0x2f, 0xad, 0xc6, 0x3e, 0xb8, 0xb6, 0xeb,
	0xe5, 0xb3, 0xb5, 0x60, 0xf7, 0xbe, 0x94, 0x89, 0xe8, 0xf9, 0xfe, 0x64, 0x32, 0x69, 0x09, 0xc9,
	0x59, 0x1c, 0x06, 0x51, 0x6b, 0xc8, 0x22, 0x3f, 0x60, 0x43, 0xa1, 0x2f, 0x18, 0x12, 0x4b, 0x2a,
	0x29, 0x11, 0xdb, 0x79, 0x59, 0xe7, 0x3d, 0xf8, 0xee, 0xb9, 0x87, 0x68, 0x80, 0x7a, 0x1e, 0x1a,
	0x87, 0xcd, 0x4e, 0xb7, 0x8b, 0x76, 0x3c, 0x54, 0x9a, 0x57, 0x05, 0xb8, 0x68, 0xb6, 0xdb, 0x6d,
	0x1d, 0x98, 0xb7, 0x8d, 0x7a, 0x88, 0x37, 0xbb, 0xdd, 0x2e, 0x7a, 0xa8, 0xac, 0x99, 0x3c, 0xbe,
	0xa8, 0x9c, 0x76, 0xeb, 0x93, 0xd9, 0xd4, 0xb9, 0xab, 0x82, 0xdb, 0xf2, 0x74, 0x40, 0x1e, 0xe0,
	0x28, 0x19, 0x11, 0xe1, 0x17, 0x76, 0x18, 0x64, 0x3c, 0x2d, 0xf9, 0x40, 0x7e, 0x3b, 0x9b, 0x3a,
	0xaf, 0x5f, 0x92, 0xa7, 0xb6, 0x78, 0x50, 0xa4, 0x76, 0xfe, 0xa8, 0x00, 0xd0, 0x2f, 0x7c, 0xf2,
	0x83, 0x05, 0xaa, 0xf3, 0xdb, 0x00, 0xba, 0x86, 0xb3, 0x8c, 0x7b, 0xca, 0xbd, 0xb9, 0x34, 0x36,
	0x77, 0x2d, 0x3a, 0x7f, 0xb4, 0x3b, 0x40, 0xf7, 0xc0, 0x2b, 0x47, 0x04, 0xf3, 0xd8, 0xbb, 0xcf,
	0x26, 0x9e, 0x64, 0x5e, 0x84, 0x3f, 0x21, 0x1e, 0x2e, 0x54, 0x86, 0x77, 0xfe, 0x5d, 0x64, 0x41,
	0xf8, 0x19, 0x1d, 0x12, 0xe1, 0x17, 0x0d, 0x6e, 0xcf, 0xb9, 0xbe, 0xf8, 0xf9, 0xf7, 0xaf, 0xec,
	0x75, 0xb4, 0xe6, 0x9f, 0xb5, 0xfd, 0xc2, 0xe5, 0x3d, 0xab, 0x01, 0x3f, 0x02, 0x95, 0x03, 0x22,
	0xe1, 0x96, 0xd1, 0x60, 0x71, 0x09, 0xb8, 0xf5, 0x8b, 0x81, 0xb4, 0xed, 0x9b, 0x1a, 0x74, 0x03,
	0xae, 0x9b, 0xa0, 0xfe, 0x39, 0x0d, 0x1e, 0xc2, 0x8f, 0x41, 0x75, 0x7e, 0x4c, 0x2d, 0xc8, 0x62,
	0x1c, 0x91, 0x0b, 0xb2, 0x98, 0xe7, 0x5a, 0x86, 0xdf, 0x58, 0x8a, 0x7f, 0x17, 0xac, 0xa8, 0x37,
	0x1f, 0x9a, 0xed, 0x95, 0xce, 0x1d, 0xf7, 0xc6, 0x92, 0x48, 0x8a, 0xbc, 0xa9, 0x91, 0x5f, 0x84,
	0x0b, 0x72, 0xb8, 0xeb, 0x4f, 0x9f, 0x4d, 0x9d, 0xb5, 0x27, 0xcf, 0xa6, 0x4e, 0xf1, 0xe6, 0xee,
	0x75, 0xc0, 0xad, 0x21, 0x8b, 0x0a, 0xc5, 0x71, 0x42, 0x15, 0x72, 0x32, 0x1a, 0x47, 0x27, 0x34,
	0x0e, 0xf7, 0x60, 0x21, 0xf6, 0x71, 0xba, 0x76, 0x52, 0xd5, 0x5f, 0x36, 0x6f, 0xfd, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x59, 0xc4, 0x36, 0xdc, 0x7a, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleGrantsClient is the client API for RoleGrants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleGrantsClient interface {
	// Create registers a new RoleGrant.
	Create(ctx context.Context, in *RoleGrantCreateRequest, opts ...grpc.CallOption) (*RoleGrantCreateResponse, error)
	// Get reads one RoleGrant by ID.
	Get(ctx context.Context, in *RoleGrantGetRequest, opts ...grpc.CallOption) (*RoleGrantGetResponse, error)
	// Delete removes a RoleGrant by ID.
	Delete(ctx context.Context, in *RoleGrantDeleteRequest, opts ...grpc.CallOption) (*RoleGrantDeleteResponse, error)
	// List gets a list of RoleGrants matching a given set of criteria.
	List(ctx context.Context, in *RoleGrantListRequest, opts ...grpc.CallOption) (*RoleGrantListResponse, error)
}

type roleGrantsClient struct {
	cc *grpc.ClientConn
}

func NewRoleGrantsClient(cc *grpc.ClientConn) RoleGrantsClient {
	return &roleGrantsClient{cc}
}

func (c *roleGrantsClient) Create(ctx context.Context, in *RoleGrantCreateRequest, opts ...grpc.CallOption) (*RoleGrantCreateResponse, error) {
	out := new(RoleGrantCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) Get(ctx context.Context, in *RoleGrantGetRequest, opts ...grpc.CallOption) (*RoleGrantGetResponse, error) {
	out := new(RoleGrantGetResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) Delete(ctx context.Context, in *RoleGrantDeleteRequest, opts ...grpc.CallOption) (*RoleGrantDeleteResponse, error) {
	out := new(RoleGrantDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) List(ctx context.Context, in *RoleGrantListRequest, opts ...grpc.CallOption) (*RoleGrantListResponse, error) {
	out := new(RoleGrantListResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleGrantsServer is the server API for RoleGrants service.
type RoleGrantsServer interface {
	// Create registers a new RoleGrant.
	Create(context.Context, *RoleGrantCreateRequest) (*RoleGrantCreateResponse, error)
	// Get reads one RoleGrant by ID.
	Get(context.Context, *RoleGrantGetRequest) (*RoleGrantGetResponse, error)
	// Delete removes a RoleGrant by ID.
	Delete(context.Context, *RoleGrantDeleteRequest) (*RoleGrantDeleteResponse, error)
	// List gets a list of RoleGrants matching a given set of criteria.
	List(context.Context, *RoleGrantListRequest) (*RoleGrantListResponse, error)
}

// UnimplementedRoleGrantsServer can be embedded to have forward compatible implementations.
type UnimplementedRoleGrantsServer struct {
}

func (*UnimplementedRoleGrantsServer) Create(ctx context.Context, req *RoleGrantCreateRequest) (*RoleGrantCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRoleGrantsServer) Get(ctx context.Context, req *RoleGrantGetRequest) (*RoleGrantGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRoleGrantsServer) Delete(ctx context.Context, req *RoleGrantDeleteRequest) (*RoleGrantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRoleGrantsServer) List(ctx context.Context, req *RoleGrantListRequest) (*RoleGrantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRoleGrantsServer(s *grpc.Server, srv RoleGrantsServer) {
	s.RegisterService(&_RoleGrants_serviceDesc, srv)
}

func _RoleGrants_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Create(ctx, req.(*RoleGrantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Get(ctx, req.(*RoleGrantGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Delete(ctx, req.(*RoleGrantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).List(ctx, req.(*RoleGrantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleGrants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RoleGrants",
	HandlerType: (*RoleGrantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoleGrants_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoleGrants_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleGrants_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RoleGrants_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_grants.proto",
}
