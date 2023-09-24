// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NetworkInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Plugins              []string `protobuf:"bytes,3,rep,name=plugins,proto3" json:"plugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfo) Reset()         { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()    {}
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}
func (m *NetworkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfo.Unmarshal(m, b)
}
func (m *NetworkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfo.Marshal(b, m, deterministic)
}
func (m *NetworkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfo.Merge(m, src)
}
func (m *NetworkInfo) XXX_Size() int {
	return xxx_messageInfo_NetworkInfo.Size(m)
}
func (m *NetworkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfo proto.InternalMessageInfo

func (m *NetworkInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NetworkInfo) GetPlugins() []string {
	if m != nil {
		return m.Plugins
	}
	return nil
}

type NetworkCreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Driver               string   `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Internal             bool     `protobuf:"varint,4,opt,name=internal,proto3" json:"internal,omitempty"`
	Subnet               string   `protobuf:"bytes,5,opt,name=subnet,proto3" json:"subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkCreateRequest) Reset()         { *m = NetworkCreateRequest{} }
func (m *NetworkCreateRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkCreateRequest) ProtoMessage()    {}
func (*NetworkCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}
func (m *NetworkCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkCreateRequest.Unmarshal(m, b)
}
func (m *NetworkCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkCreateRequest.Marshal(b, m, deterministic)
}
func (m *NetworkCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkCreateRequest.Merge(m, src)
}
func (m *NetworkCreateRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkCreateRequest.Size(m)
}
func (m *NetworkCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkCreateRequest proto.InternalMessageInfo

func (m *NetworkCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkCreateRequest) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *NetworkCreateRequest) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NetworkCreateRequest) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

func (m *NetworkCreateRequest) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

type NetworkCreateResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cc                   uint32   `protobuf:"varint,2,opt,name=cc,proto3" json:"cc,omitempty"`
	Errmsg               string   `protobuf:"bytes,3,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkCreateResponse) Reset()         { *m = NetworkCreateResponse{} }
func (m *NetworkCreateResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkCreateResponse) ProtoMessage()    {}
func (*NetworkCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}
func (m *NetworkCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkCreateResponse.Unmarshal(m, b)
}
func (m *NetworkCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkCreateResponse.Marshal(b, m, deterministic)
}
func (m *NetworkCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkCreateResponse.Merge(m, src)
}
func (m *NetworkCreateResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkCreateResponse.Size(m)
}
func (m *NetworkCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkCreateResponse proto.InternalMessageInfo

func (m *NetworkCreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkCreateResponse) GetCc() uint32 {
	if m != nil {
		return m.Cc
	}
	return 0
}

func (m *NetworkCreateResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type NetworkInspectRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInspectRequest) Reset()         { *m = NetworkInspectRequest{} }
func (m *NetworkInspectRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkInspectRequest) ProtoMessage()    {}
func (*NetworkInspectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}
func (m *NetworkInspectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInspectRequest.Unmarshal(m, b)
}
func (m *NetworkInspectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInspectRequest.Marshal(b, m, deterministic)
}
func (m *NetworkInspectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInspectRequest.Merge(m, src)
}
func (m *NetworkInspectRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkInspectRequest.Size(m)
}
func (m *NetworkInspectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInspectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInspectRequest proto.InternalMessageInfo

func (m *NetworkInspectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetworkInspectResponse struct {
	NetworkJSON          string   `protobuf:"bytes,1,opt,name=NetworkJSON,proto3" json:"NetworkJSON,omitempty"`
	Cc                   uint32   `protobuf:"varint,2,opt,name=cc,proto3" json:"cc,omitempty"`
	Errmsg               string   `protobuf:"bytes,3,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInspectResponse) Reset()         { *m = NetworkInspectResponse{} }
func (m *NetworkInspectResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInspectResponse) ProtoMessage()    {}
func (*NetworkInspectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}
func (m *NetworkInspectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInspectResponse.Unmarshal(m, b)
}
func (m *NetworkInspectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInspectResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInspectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInspectResponse.Merge(m, src)
}
func (m *NetworkInspectResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInspectResponse.Size(m)
}
func (m *NetworkInspectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInspectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInspectResponse proto.InternalMessageInfo

func (m *NetworkInspectResponse) GetNetworkJSON() string {
	if m != nil {
		return m.NetworkJSON
	}
	return ""
}

func (m *NetworkInspectResponse) GetCc() uint32 {
	if m != nil {
		return m.Cc
	}
	return 0
}

func (m *NetworkInspectResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type NetworkListRequest struct {
	Filters              map[string]string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NetworkListRequest) Reset()         { *m = NetworkListRequest{} }
func (m *NetworkListRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkListRequest) ProtoMessage()    {}
func (*NetworkListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}
func (m *NetworkListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkListRequest.Unmarshal(m, b)
}
func (m *NetworkListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkListRequest.Marshal(b, m, deterministic)
}
func (m *NetworkListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkListRequest.Merge(m, src)
}
func (m *NetworkListRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkListRequest.Size(m)
}
func (m *NetworkListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkListRequest proto.InternalMessageInfo

func (m *NetworkListRequest) GetFilters() map[string]string {
	if m != nil {
		return m.Filters
	}
	return nil
}

type NetworkListResponse struct {
	Networks             []*NetworkInfo `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	Cc                   uint32         `protobuf:"varint,2,opt,name=cc,proto3" json:"cc,omitempty"`
	Errmsg               string         `protobuf:"bytes,3,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NetworkListResponse) Reset()         { *m = NetworkListResponse{} }
func (m *NetworkListResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkListResponse) ProtoMessage()    {}
func (*NetworkListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}
func (m *NetworkListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkListResponse.Unmarshal(m, b)
}
func (m *NetworkListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkListResponse.Marshal(b, m, deterministic)
}
func (m *NetworkListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkListResponse.Merge(m, src)
}
func (m *NetworkListResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkListResponse.Size(m)
}
func (m *NetworkListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkListResponse proto.InternalMessageInfo

func (m *NetworkListResponse) GetNetworks() []*NetworkInfo {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *NetworkListResponse) GetCc() uint32 {
	if m != nil {
		return m.Cc
	}
	return 0
}

func (m *NetworkListResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type NetworkRemoveRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkRemoveRequest) Reset()         { *m = NetworkRemoveRequest{} }
func (m *NetworkRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkRemoveRequest) ProtoMessage()    {}
func (*NetworkRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}
func (m *NetworkRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRemoveRequest.Unmarshal(m, b)
}
func (m *NetworkRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRemoveRequest.Marshal(b, m, deterministic)
}
func (m *NetworkRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRemoveRequest.Merge(m, src)
}
func (m *NetworkRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkRemoveRequest.Size(m)
}
func (m *NetworkRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRemoveRequest proto.InternalMessageInfo

func (m *NetworkRemoveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetworkRemoveResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cc                   uint32   `protobuf:"varint,2,opt,name=cc,proto3" json:"cc,omitempty"`
	Errmsg               string   `protobuf:"bytes,3,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkRemoveResponse) Reset()         { *m = NetworkRemoveResponse{} }
func (m *NetworkRemoveResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkRemoveResponse) ProtoMessage()    {}
func (*NetworkRemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{8}
}
func (m *NetworkRemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRemoveResponse.Unmarshal(m, b)
}
func (m *NetworkRemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRemoveResponse.Marshal(b, m, deterministic)
}
func (m *NetworkRemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRemoveResponse.Merge(m, src)
}
func (m *NetworkRemoveResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkRemoveResponse.Size(m)
}
func (m *NetworkRemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRemoveResponse proto.InternalMessageInfo

func (m *NetworkRemoveResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkRemoveResponse) GetCc() uint32 {
	if m != nil {
		return m.Cc
	}
	return 0
}

func (m *NetworkRemoveResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkInfo)(nil), "network.NetworkInfo")
	proto.RegisterType((*NetworkCreateRequest)(nil), "network.NetworkCreateRequest")
	proto.RegisterType((*NetworkCreateResponse)(nil), "network.NetworkCreateResponse")
	proto.RegisterType((*NetworkInspectRequest)(nil), "network.NetworkInspectRequest")
	proto.RegisterType((*NetworkInspectResponse)(nil), "network.NetworkInspectResponse")
	proto.RegisterType((*NetworkListRequest)(nil), "network.NetworkListRequest")
	proto.RegisterMapType((map[string]string)(nil), "network.NetworkListRequest.FiltersEntry")
	proto.RegisterType((*NetworkListResponse)(nil), "network.NetworkListResponse")
	proto.RegisterType((*NetworkRemoveRequest)(nil), "network.NetworkRemoveRequest")
	proto.RegisterType((*NetworkRemoveResponse)(nil), "network.NetworkRemoveResponse")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0xae, 0x2d, 0x77, 0x6c, 0x42, 0xa6, 0x4c, 0x56, 0x80, 0x11, 0xe5, 0xa9, 0x02,
	0xa9, 0x42, 0xe3, 0x05, 0xed, 0x8d, 0x21, 0x18, 0x9b, 0xd0, 0x90, 0x52, 0xf1, 0x01, 0x69, 0xb8,
	0xab, 0xa2, 0xb5, 0x76, 0xb1, 0xdd, 0x4c, 0xfd, 0x0b, 0x1e, 0xf8, 0x11, 0xfe, 0x10, 0xb9, 0xbe,
	0x89, 0x9a, 0xa4, 0xab, 0x2a, 0xed, 0xcd, 0xc7, 0xe7, 0xfa, 0xf8, 0xdc, 0xeb, 0x93, 0xc0, 0x91,
	0x40, 0x73, 0x2f, 0xd5, 0xdd, 0x68, 0xa1, 0xa4, 0x91, 0xac, 0x47, 0x30, 0xfe, 0x09, 0x87, 0x37,
	0x6e, 0x79, 0x25, 0x6e, 0x25, 0x63, 0xd0, 0x11, 0xe9, 0x1c, 0xb9, 0x17, 0x79, 0xc3, 0x27, 0xc9,
	0x7a, 0xcd, 0x38, 0xf4, 0x0a, 0x54, 0x3a, 0x97, 0x82, 0xfb, 0xeb, 0xed, 0x12, 0x5a, 0x66, 0x31,
	0x5b, 0x4e, 0x73, 0xa1, 0x79, 0x10, 0x05, 0x96, 0x21, 0x18, 0xff, 0xf1, 0x60, 0x40, 0xba, 0x9f,
	0x15, 0xa6, 0x06, 0x13, 0xfc, 0xbd, 0x44, 0x6d, 0xb6, 0x5e, 0x70, 0x02, 0xdd, 0x5f, 0x2a, 0x2f,
	0x50, 0x91, 0x3e, 0x21, 0x2b, 0x3f, 0x4d, 0x0d, 0xde, 0xa7, 0x2b, 0x1e, 0xb8, 0x8b, 0x09, 0xb2,
	0x10, 0xfa, 0xb9, 0x30, 0xa8, 0x44, 0x3a, 0xe3, 0x9d, 0xc8, 0x1b, 0xf6, 0x93, 0x0a, 0x5b, 0x35,
	0xbd, 0x9c, 0x08, 0x34, 0xfc, 0xc0, 0xa9, 0x39, 0x14, 0x8f, 0xe1, 0x45, 0xc3, 0x91, 0x5e, 0x48,
	0xa1, 0x71, 0xab, 0xa5, 0x63, 0xf0, 0xb3, 0x6c, 0x6d, 0xe7, 0x28, 0xf1, 0xb3, 0xcc, 0x8a, 0xa2,
	0x52, 0x73, 0x3d, 0x25, 0x27, 0x84, 0xe2, 0x77, 0x95, 0xe8, 0x95, 0xd0, 0x0b, 0xcc, 0xcc, 0x8e,
	0x3e, 0xe3, 0x09, 0x9c, 0x34, 0x8b, 0xc9, 0x42, 0x54, 0xbd, 0xc2, 0xf5, 0xf8, 0xc7, 0x0d, 0x1d,
	0xda, 0xdc, 0xda, 0xdb, 0xd0, 0x5f, 0x0f, 0x18, 0x9d, 0xfb, 0x9e, 0xeb, 0xca, 0xce, 0x05, 0xf4,
	0x6e, 0xf3, 0x99, 0x41, 0xa5, 0xb9, 0x17, 0x05, 0xc3, 0xc3, 0xb3, 0xe1, 0xa8, 0x0c, 0x44, 0xbb,
	0x7a, 0xf4, 0xd5, 0x95, 0x7e, 0x11, 0x46, 0xad, 0x92, 0xf2, 0x60, 0x78, 0x0e, 0x4f, 0x37, 0x09,
	0xf6, 0x0c, 0x82, 0x3b, 0x5c, 0x91, 0x59, 0xbb, 0x64, 0x03, 0x38, 0x28, 0xd2, 0xd9, 0x12, 0xe9,
	0x1d, 0x1d, 0x38, 0xf7, 0x3f, 0x7a, 0xb1, 0x84, 0xe7, 0xb5, 0x7b, 0xa8, 0xef, 0xf7, 0xd0, 0x27,
	0x1b, 0xa5, 0xaf, 0x41, 0xd3, 0x97, 0x8d, 0x65, 0x52, 0x55, 0xed, 0x3d, 0x87, 0xb7, 0x55, 0xfe,
	0x12, 0x9c, 0xcb, 0x62, 0x57, 0xfe, 0x36, 0x92, 0x51, 0xd6, 0x3e, 0x3e, 0x19, 0x67, 0xff, 0x7c,
	0x38, 0x26, 0xd5, 0x31, 0xaa, 0x22, 0xcf, 0x90, 0x5d, 0x42, 0xd7, 0x45, 0x8f, 0xbd, 0x6e, 0x76,
	0x59, 0xfb, 0x48, 0xc2, 0xd3, 0x87, 0x68, 0xf2, 0x75, 0x0d, 0x3d, 0x4a, 0x10, 0x3b, 0x6d, 0xcf,
	0x6b, 0x33, 0x87, 0xe1, 0x9b, 0x07, 0x79, 0xd2, 0xfa, 0x04, 0x1d, 0xfb, 0x24, 0xec, 0xe5, 0x8e,
	0x40, 0x84, 0xaf, 0xb6, 0x93, 0x24, 0x71, 0x09, 0x5d, 0x37, 0xb8, 0x76, 0x5f, 0xb5, 0xe1, 0xb7,
	0xfb, 0xaa, 0xcf, 0xfb, 0xc2, 0xff, 0xe6, 0x4f, 0xba, 0xeb, 0x1f, 0xd4, 0x87, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x1c, 0xa7, 0x82, 0xb1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error)
	Inspect(ctx context.Context, in *NetworkInspectRequest, opts ...grpc.CallOption) (*NetworkInspectResponse, error)
	List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error)
	Remove(ctx context.Context, in *NetworkRemoveRequest, opts ...grpc.CallOption) (*NetworkRemoveResponse, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error) {
	out := new(NetworkCreateResponse)
	err := c.cc.Invoke(ctx, "/network.NetworkService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Inspect(ctx context.Context, in *NetworkInspectRequest, opts ...grpc.CallOption) (*NetworkInspectResponse, error) {
	out := new(NetworkInspectResponse)
	err := c.cc.Invoke(ctx, "/network.NetworkService/Inspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error) {
	out := new(NetworkListResponse)
	err := c.cc.Invoke(ctx, "/network.NetworkService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Remove(ctx context.Context, in *NetworkRemoveRequest, opts ...grpc.CallOption) (*NetworkRemoveResponse, error) {
	out := new(NetworkRemoveResponse)
	err := c.cc.Invoke(ctx, "/network.NetworkService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Create(context.Context, *NetworkCreateRequest) (*NetworkCreateResponse, error)
	Inspect(context.Context, *NetworkInspectRequest) (*NetworkInspectResponse, error)
	List(context.Context, *NetworkListRequest) (*NetworkListResponse, error)
	Remove(context.Context, *NetworkRemoveRequest) (*NetworkRemoveResponse, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) Create(ctx context.Context, req *NetworkCreateRequest) (*NetworkCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNetworkServiceServer) Inspect(ctx context.Context, req *NetworkInspectRequest) (*NetworkInspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inspect not implemented")
}
func (*UnimplementedNetworkServiceServer) List(ctx context.Context, req *NetworkListRequest) (*NetworkListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedNetworkServiceServer) Remove(ctx context.Context, req *NetworkRemoveRequest) (*NetworkRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NetworkService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Create(ctx, req.(*NetworkCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Inspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Inspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NetworkService/Inspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Inspect(ctx, req.(*NetworkInspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NetworkService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).List(ctx, req.(*NetworkListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NetworkService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Remove(ctx, req.(*NetworkRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "network.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkService_Create_Handler,
		},
		{
			MethodName: "Inspect",
			Handler:    _NetworkService_Inspect_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkService_List_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _NetworkService_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}
