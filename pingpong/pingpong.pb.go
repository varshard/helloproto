// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pingpong.proto

package pingpong

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Ping struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Ping) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Pong struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "pingpong.Ping")
	proto.RegisterType((*Pong)(nil), "pingpong.Pong")
}

func init() { proto.RegisterFile("pingpong.proto", fileDescriptor_1cfbf639ab46154b) }

var fileDescriptor_1cfbf639ab46154b = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0xcc, 0x4b,
	0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x0c,
	0xb8, 0x58, 0x02, 0x32, 0xf3, 0xd2, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xb0, 0x8e, 0x7c, 0x52, 0x74, 0x18,
	0x59, 0x73, 0x71, 0x80, 0xec, 0x00, 0xeb, 0xd2, 0xe7, 0xe2, 0x0c, 0x2e, 0x49, 0x2c, 0x2a, 0x81,
	0x58, 0xaa, 0x07, 0x77, 0x17, 0x88, 0x2f, 0x85, 0xcc, 0xcf, 0xcf, 0x4b, 0xd7, 0x60, 0x34, 0x60,
	0x4c, 0x62, 0x03, 0xbb, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xe6, 0x40, 0xbc, 0xc3,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	StartPing(ctx context.Context, opts ...grpc.CallOption) (PingPong_StartPingClient, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) StartPing(ctx context.Context, opts ...grpc.CallOption) (PingPong_StartPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PingPong_serviceDesc.Streams[0], "/pingpong.PingPong/StartPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingPongStartPingClient{stream}
	return x, nil
}

type PingPong_StartPingClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type pingPongStartPingClient struct {
	grpc.ClientStream
}

func (x *pingPongStartPingClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingPongStartPingClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	StartPing(PingPong_StartPingServer) error
}

// UnimplementedPingPongServer can be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (*UnimplementedPingPongServer) StartPing(srv PingPong_StartPingServer) error {
	return status.Errorf(codes.Unimplemented, "method StartPing not implemented")
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_StartPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingPongServer).StartPing(&pingPongStartPingServer{stream})
}

type PingPong_StartPingServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type pingPongStartPingServer struct {
	grpc.ServerStream
}

func (x *pingPongStartPingServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingPongStartPingServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartPing",
			Handler:       _PingPong_StartPing_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pingpong.proto",
}
