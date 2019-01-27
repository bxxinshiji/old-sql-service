// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/goods/good.proto

package good

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Good struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BarCode              string   `protobuf:"bytes,3,opt,name=bar_code,json=barCode,proto3" json:"bar_code,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Standard             string   `protobuf:"bytes,5,opt,name=standard,proto3" json:"standard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Good) Reset()         { *m = Good{} }
func (m *Good) String() string { return proto.CompactTextString(m) }
func (*Good) ProtoMessage()    {}
func (*Good) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0c6740ea846d8, []int{0}
}

func (m *Good) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Good.Unmarshal(m, b)
}
func (m *Good) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Good.Marshal(b, m, deterministic)
}
func (m *Good) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Good.Merge(m, src)
}
func (m *Good) XXX_Size() int {
	return xxx_messageInfo_Good.Size(m)
}
func (m *Good) XXX_DiscardUnknown() {
	xxx_messageInfo_Good.DiscardUnknown(m)
}

var xxx_messageInfo_Good proto.InternalMessageInfo

func (m *Good) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Good) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Good) GetBarCode() string {
	if m != nil {
		return m.BarCode
	}
	return ""
}

func (m *Good) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Good) GetStandard() string {
	if m != nil {
		return m.Standard
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0c6740ea846d8, []int{1}
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

type Response struct {
	Good                 *Good    `protobuf:"bytes,1,opt,name=good,proto3" json:"good,omitempty"`
	Goods                []*Good  `protobuf:"bytes,2,rep,name=goods,proto3" json:"goods,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0c6740ea846d8, []int{2}
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

func (m *Response) GetGood() *Good {
	if m != nil {
		return m.Good
	}
	return nil
}

func (m *Response) GetGoods() []*Good {
	if m != nil {
		return m.Goods
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f0c6740ea846d8, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Good)(nil), "good.Good")
	proto.RegisterType((*Request)(nil), "good.Request")
	proto.RegisterType((*Response)(nil), "good.Response")
	proto.RegisterType((*Error)(nil), "good.Error")
}

func init() { proto.RegisterFile("proto/goods/good.proto", fileDescriptor_a5f0c6740ea846d8) }

var fileDescriptor_a5f0c6740ea846d8 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xdd, 0xb6, 0xd9, 0x3f, 0x53, 0xd8, 0xc3, 0x20, 0x12, 0xf7, 0x20, 0xa5, 0x5e, 0x3c,
	0xc8, 0x0a, 0xeb, 0xd9, 0x93, 0x48, 0xef, 0xf9, 0x02, 0x92, 0x36, 0x83, 0xe4, 0x60, 0xa7, 0x9b,
	0x44, 0xfc, 0xfa, 0x92, 0x69, 0x95, 0xc5, 0x4b, 0x78, 0xf9, 0xcd, 0x83, 0x79, 0x6f, 0xe0, 0x66,
	0x0a, 0x9c, 0xf8, 0xe9, 0x83, 0xd9, 0x45, 0x79, 0x8f, 0x02, 0xb0, 0xca, 0xba, 0xfd, 0x86, 0xaa,
	0x63, 0x76, 0xb8, 0x87, 0xc2, 0x3b, 0xbd, 0x6a, 0x56, 0x0f, 0x3b, 0x53, 0x78, 0x87, 0x08, 0xd5,
	0x68, 0x3f, 0x49, 0x17, 0x42, 0x44, 0xe3, 0x2d, 0x6c, 0x7b, 0x1b, 0xde, 0x07, 0x76, 0xa4, 0x4b,
	0xe1, 0x9b, 0xde, 0x86, 0x57, 0x76, 0x84, 0xd7, 0xa0, 0xa6, 0xe0, 0x07, 0xd2, 0x95, 0xf0, 0xf9,
	0x83, 0x07, 0xd8, 0xc6, 0x64, 0x47, 0x67, 0x83, 0xd3, 0x4a, 0x06, 0x7f, 0xff, 0x76, 0x07, 0x1b,
	0x43, 0xe7, 0x2f, 0x8a, 0xa9, 0x3d, 0xc3, 0xd6, 0x50, 0x9c, 0x78, 0x8c, 0x84, 0x77, 0x20, 0xb9,
	0x24, 0x49, 0x7d, 0x82, 0xa3, 0x04, 0xce, 0x09, 0x8d, 0x70, 0x6c, 0x40, 0x49, 0x13, 0x5d, 0x34,
	0xe5, 0x3f, 0xc3, 0x3c, 0xc0, 0x7b, 0x58, 0x53, 0x08, 0x1c, 0xa2, 0x2e, 0xc5, 0x52, 0xcf, 0x96,
	0xb7, 0xcc, 0xcc, 0x32, 0x6a, 0x5f, 0x40, 0x09, 0xc8, 0x3d, 0xa5, 0x4f, 0xde, 0xa7, 0x8c, 0x68,
	0x6c, 0xa0, 0x76, 0x14, 0x87, 0xe0, 0xa7, 0xe4, 0x79, 0x5c, 0x4e, 0x70, 0x89, 0x4e, 0x8f, 0xa0,
	0xba, 0x65, 0x59, 0xd9, 0x51, 0xc2, 0x8b, 0x18, 0x87, 0xfd, 0xac, 0x7f, 0x1b, 0xb5, 0x57, 0xfd,
	0x5a, 0x0e, 0xfe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x91, 0x6b, 0xea, 0x5b, 0x8a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Goods service

type GoodsClient interface {
	// 获取商品信息
	Get(ctx context.Context, in *Good, opts ...client.CallOption) (*Response, error)
}

type goodsClient struct {
	c           client.Client
	serviceName string
}

func NewGoodsClient(serviceName string, c client.Client) GoodsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "good"
	}
	return &goodsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *goodsClient) Get(ctx context.Context, in *Good, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goods service

type GoodsHandler interface {
	// 获取商品信息
	Get(context.Context, *Good, *Response) error
}

func RegisterGoodsHandler(s server.Server, hdlr GoodsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Goods{hdlr}, opts...))
}

type Goods struct {
	GoodsHandler
}

func (h *Goods) Get(ctx context.Context, in *Good, out *Response) error {
	return h.GoodsHandler.Get(ctx, in, out)
}
