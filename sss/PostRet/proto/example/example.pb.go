// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostRet

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//web---srv
type Request struct {
	//	手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	//	密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	//	短信验证码
	SmsCode              string   `protobuf:"bytes,3,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

//srv---web
type Response struct {
	//	错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//	错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//	sessionid
	Sessionid            string   `protobuf:"bytes,3,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostRet.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostRet.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xad, 0x62, 0xb3, 0x9d, 0xe3, 0x50, 0x64, 0xad, 0x1f, 0xc8, 0x9e, 0x3c, 0x45, 0xd0,
	0x9f, 0x20, 0x5e, 0x3c, 0x49, 0x40, 0xf1, 0x26, 0x6d, 0x33, 0x2c, 0x81, 0xcd, 0x4e, 0x9c, 0x89,
	0x1f, 0x3f, 0x5f, 0xcc, 0x46, 0xbd, 0xd8, 0x53, 0x78, 0xde, 0x17, 0xde, 0xe4, 0x09, 0x9c, 0x24,
	0xe1, 0xcc, 0x57, 0xf4, 0xb9, 0x8e, 0x69, 0xa0, 0x9f, 0xd3, 0x96, 0x14, 0x97, 0x3d, 0xdb, 0x18,
	0xb6, 0xc2, 0x56, 0xe5, 0xdd, 0x3e, 0xb0, 0x66, 0x47, 0xb9, 0x7b, 0x06, 0xe3, 0xe8, 0xf5, 0x8d,
	0x34, 0xe3, 0x11, 0xcc, 0x23, 0x6f, 0xc2, 0x40, 0xed, 0xec, 0x62, 0x76, 0xb9, 0x70, 0x95, 0x70,
	0x05, 0x4d, 0x5a, 0xab, 0x7e, 0xb0, 0xf8, 0x76, 0xbf, 0x34, 0xbf, 0x8c, 0xc7, 0xd0, 0x68, 0xd4,
	0x97, 0x2d, 0x7b, 0x6a, 0x0f, 0x4a, 0x67, 0x34, 0xea, 0x2d, 0x7b, 0xea, 0x9e, 0xa0, 0x71, 0xa4,
	0x89, 0x47, 0x25, 0x5c, 0xc2, 0x21, 0x89, 0x8c, 0x5c, 0x97, 0x27, 0xf8, 0xbe, 0x90, 0x44, 0xa2,
	0xf6, 0x75, 0xb6, 0x12, 0x9e, 0xc2, 0x42, 0x49, 0x35, 0xf0, 0x18, 0x7c, 0x5d, 0xfd, 0x0b, 0xae,
	0x1f, 0xc1, 0xdc, 0x4d, 0x62, 0x78, 0x0f, 0xa6, 0x7a, 0xe0, 0x99, 0xfd, 0x4f, 0xcf, 0x56, 0xb7,
	0xd5, 0xf9, 0xae, 0x7a, 0x7a, 0x60, 0xb7, 0xb7, 0x99, 0x97, 0x5f, 0xba, 0xf9, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x2e, 0x1e, 0x10, 0x9c, 0x44, 0x01, 0x00, 0x00,
}
