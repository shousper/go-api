// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/divisionone/go-api/proto/api.proto

/*
Package go_api is a generated protocol buffer package.

It is generated from these files:
	github.com/divisionone/go-api/proto/api.proto

It has these top-level messages:
	Pair
	Request
	Response
	Event
*/
package go_api

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

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// A HTTP request as RPC
type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP response as RPC
type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// A HTTP event as RPC
type Event struct {
	// e.g login
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// uuid
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// event headers
	Header map[string]*Pair `protobuf:"bytes,4,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// the event data
	Data string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Event) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.api.Pair")
	proto.RegisterType((*Request)(nil), "go.api.Request")
	proto.RegisterType((*Response)(nil), "go.api.Response")
	proto.RegisterType((*Event)(nil), "go.api.Event")
}

func init() { proto.RegisterFile("github.com/divisionone/go-api/proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0xae, 0xd3, 0x30,
	0x10, 0x55, 0xe2, 0x24, 0xdc, 0x4c, 0x11, 0x42, 0x5e, 0x20, 0x73, 0xb9, 0x42, 0x55, 0x36, 0x54,
	0x48, 0x37, 0x81, 0x96, 0x05, 0x62, 0x0b, 0x55, 0x59, 0x56, 0xfe, 0x03, 0xb7, 0xb1, 0x12, 0x0b,
	0x12, 0x9b, 0xd8, 0xa9, 0xd4, 0xaf, 0xe2, 0x0b, 0xf8, 0x10, 0xfe, 0x06, 0x65, 0xe2, 0x3e, 0xa8,
	0xca, 0x06, 0xb8, 0xbb, 0x13, 0xfb, 0xcc, 0x99, 0x33, 0x67, 0x1c, 0x78, 0x55, 0x29, 0x57, 0xf7,
	0x9b, 0x7c, 0xab, 0x9b, 0xa2, 0x51, 0xdb, 0x4e, 0x17, 0x95, 0xbe, 0x17, 0x46, 0x15, 0xa6, 0xd3,
	0x4e, 0x17, 0xc2, 0xa8, 0x1c, 0x11, 0x4d, 0x2a, 0x9d, 0x0b, 0xa3, 0xb2, 0x37, 0x10, 0xad, 0x85,
	0xea, 0xe8, 0x53, 0x20, 0x5f, 0xe4, 0x9e, 0x05, 0xd3, 0x60, 0x96, 0xf2, 0x01, 0xd2, 0x67, 0x90,
	0xec, 0xc4, 0xd7, 0x5e, 0x5a, 0x16, 0x4e, 0xc9, 0x2c, 0xe5, 0xfe, 0x2b, 0xfb, 0x4e, 0xe0, 0x11,
	0x97, 0xdf, 0x7a, 0x69, 0xdd, 0xc0, 0x69, 0xa4, 0xab, 0x75, 0xe9, 0x0b, 0xfd, 0x17, 0xa5, 0x10,
	0x19, 0xe1, 0x6a, 0x16, 0xe2, 0x29, 0x62, 0xba, 0x80, 0xa4, 0x96, 0xa2, 0x94, 0x1d, 0x23, 0x53,
	0x32, 0x9b, 0xcc, 0x5f, 0xe4, 0xa3, 0x85, 0xdc, 0x8b, 0xe5, 0x9f, 0xf1, 0x76, 0xd9, 0xba, 0x6e,
	0xcf, 0x3d, 0x95, 0xbe, 0x06, 0x52, 0x49, 0xc7, 0x22, 0xac, 0x60, 0x97, 0x15, 0x2b, 0xe9, 0x46,
	0xfa, 0x40, 0xa2, 0xf7, 0x10, 0x19, 0x6d, 0x1d, 0x8b, 0x91, 0xfc, 0xfc, 0x92, 0xbc, 0xd6, 0xd6,
	0xb3, 0x91, 0x36, 0x78, 0xdc, 0xe8, 0x72, 0xcf, 0x92, 0xd1, 0xe3, 0x80, 0x6f, 0x57, 0x30, 0x39,
	0x73, 0x71, 0x25, 0x94, 0x0c, 0x62, 0x8c, 0x01, 0x27, 0x9b, 0xcc, 0x1f, 0x1f, 0x9a, 0x0c, 0x19,
	0xf2, 0xf1, 0xea, 0x43, 0xf8, 0x3e, 0xb8, 0xfd, 0x04, 0x37, 0x07, 0x73, 0xff, 0xa0, 0xb2, 0x84,
	0xf4, 0xe8, 0xfa, 0xef, 0x65, 0xb2, 0x1f, 0x01, 0xdc, 0x70, 0x69, 0x8d, 0x6e, 0xad, 0xa4, 0x2f,
	0x01, 0xac, 0x13, 0xae, 0xb7, 0x1f, 0x75, 0x29, 0x51, 0x2d, 0xe6, 0x67, 0x27, 0xf4, 0xdd, 0x71,
	0x4d, 0x21, 0xe6, 0x78, 0x77, 0xca, 0x71, 0x54, 0xb8, 0xba, 0xa7, 0x43, 0x98, 0xe4, 0x01, 0xc2,
	0xcc, 0x7e, 0x06, 0x10, 0x2f, 0x77, 0xb2, 0xc5, 0x9d, 0xb5, 0xa2, 0x91, 0x5e, 0x04, 0x31, 0x7d,
	0x02, 0xa1, 0x2a, 0xfd, 0x4b, 0x0b, 0x55, 0x49, 0xef, 0x20, 0x75, 0xaa, 0x91, 0xd6, 0x89, 0xc6,
	0xa0, 0x1f, 0xc2, 0x4f, 0x07, 0xf4, 0xed, 0x71, 0xbc, 0xe8, 0xf7, 0x67, 0x82, 0x0d, 0xfe, 0x34,
	0x5b, 0x29, 0x9c, 0x60, 0xf1, 0xd8, 0x74, 0xc0, 0xff, 0x6d, 0xb6, 0x4d, 0x82, 0xbf, 0xe3, 0xe2,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xdb, 0x92, 0x8c, 0xb9, 0x03, 0x00, 0x00,
}
