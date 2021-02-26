// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/customerpb/customer.proto

package customerpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Customer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerCode         string   `protobuf:"bytes,2,opt,name=customer_code,json=customerCode,proto3" json:"customer_code,omitempty"`
	CustomerName         string   `protobuf:"bytes,3,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	CustomerAddress      string   `protobuf:"bytes,4,opt,name=customer_address,json=customerAddress,proto3" json:"customer_address,omitempty"`
	CustomerEmail        string   `protobuf:"bytes,5,opt,name=customer_email,json=customerEmail,proto3" json:"customer_email,omitempty"`
	CustomerPhone        string   `protobuf:"bytes,6,opt,name=customer_phone,json=customerPhone,proto3" json:"customer_phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_212f321a2ac7b5a3, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetCustomerCode() string {
	if m != nil {
		return m.CustomerCode
	}
	return ""
}

func (m *Customer) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *Customer) GetCustomerAddress() string {
	if m != nil {
		return m.CustomerAddress
	}
	return ""
}

func (m *Customer) GetCustomerEmail() string {
	if m != nil {
		return m.CustomerEmail
	}
	return ""
}

func (m *Customer) GetCustomerPhone() string {
	if m != nil {
		return m.CustomerPhone
	}
	return ""
}

type CreateCustomerRequest struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerRequest) Reset()         { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()    {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_212f321a2ac7b5a3, []int{1}
}

func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRequest.Unmarshal(m, b)
}
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRequest.Merge(m, src)
}
func (m *CreateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRequest.Size(m)
}
func (m *CreateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRequest proto.InternalMessageInfo

func (m *CreateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerResponse struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerResponse) Reset()         { *m = CreateCustomerResponse{} }
func (m *CreateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerResponse) ProtoMessage()    {}
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_212f321a2ac7b5a3, []int{2}
}

func (m *CreateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerResponse.Unmarshal(m, b)
}
func (m *CreateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerResponse.Merge(m, src)
}
func (m *CreateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerResponse.Size(m)
}
func (m *CreateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerResponse proto.InternalMessageInfo

func (m *CreateCustomerResponse) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "customerpb.Customer")
	proto.RegisterType((*CreateCustomerRequest)(nil), "customerpb.CreateCustomerRequest")
	proto.RegisterType((*CreateCustomerResponse)(nil), "customerpb.CreateCustomerResponse")
}

func init() { proto.RegisterFile("api/v1/customerpb/customer.proto", fileDescriptor_212f321a2ac7b5a3) }

var fileDescriptor_212f321a2ac7b5a3 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x5a, 0xdd, 0x98, 0x4f, 0xed, 0x24, 0xa8, 0x04, 0x11, 0xa9, 0x15, 0x61, 0x5e,
	0x3a, 0x9d, 0x47, 0x4f, 0x5a, 0x3c, 0xe8, 0x41, 0xa4, 0x1e, 0x04, 0x2f, 0x92, 0x35, 0x0f, 0x8c,
	0xd8, 0x26, 0x26, 0xdd, 0xbe, 0xad, 0xdf, 0x45, 0x96, 0x99, 0x94, 0xa8, 0x08, 0xde, 0xd2, 0x7f,
	0x7f, 0xfc, 0xdb, 0x17, 0x1e, 0xa4, 0x4c, 0x89, 0xf1, 0xfc, 0x6c, 0x5c, 0xcd, 0x4c, 0x2b, 0x6b,
	0xd4, 0x6a, 0xea, 0x8f, 0xb9, 0xd2, 0xb2, 0x95, 0x04, 0xba, 0x57, 0xd9, 0x47, 0x04, 0x83, 0xe2,
	0xeb, 0x91, 0x24, 0x10, 0x0b, 0x4e, 0xa3, 0x34, 0x1a, 0xf5, 0xca, 0x58, 0x70, 0x72, 0x04, 0x9b,
	0x8e, 0x3e, 0x57, 0x92, 0x23, 0x8d, 0xd3, 0x68, 0xb4, 0x56, 0x6e, 0xb8, 0xb0, 0x90, 0x1c, 0x03,
	0xd4, 0xb0, 0x1a, 0xe9, 0x4a, 0x88, 0xee, 0x58, 0x8d, 0xe4, 0x04, 0xb6, 0x3c, 0x62, 0x9c, 0x6b,
	0x34, 0x86, 0xae, 0x5a, 0x37, 0x74, 0xf9, 0xe5, 0x32, 0x26, 0xc7, 0x90, 0x78, 0x8a, 0x35, 0x13,
	0x6f, 0xb4, 0x67, 0xa1, 0xff, 0xca, 0xf5, 0x22, 0x0c, 0x98, 0x7a, 0x91, 0x0d, 0xd2, 0x7e, 0xc8,
	0xee, 0x17, 0x61, 0x76, 0x03, 0x3b, 0x85, 0x46, 0xd6, 0xa2, 0x1b, 0xb2, 0xc4, 0xf7, 0x19, 0x9a,
	0x96, 0x9c, 0xc2, 0xc0, 0x49, 0x3b, 0xf1, 0xfa, 0x64, 0x3b, 0xef, 0xee, 0x25, 0xf7, 0xdc, 0xab,
	0xec, 0x16, 0x76, 0xbf, 0x57, 0x19, 0x25, 0x1b, 0x83, 0xff, 0xef, 0x9a, 0xbc, 0xc2, 0xd0, 0xa5,
	0x0f, 0xa8, 0xe7, 0xa2, 0x42, 0xf2, 0x08, 0x49, 0x58, 0x4f, 0x0e, 0x83, 0x92, 0xdf, 0xa6, 0xd8,
	0xcb, 0xfe, 0x22, 0xcb, 0xbf, 0xbb, 0x3a, 0x78, 0xda, 0xff, 0xb1, 0x12, 0x17, 0xdd, 0x71, 0xda,
	0xb7, 0x5b, 0x71, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x6f, 0xbe, 0xef, 0x39, 0x02, 0x00,
	0x00,
}