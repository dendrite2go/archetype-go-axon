// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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

type Greeting struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Acknowledgement struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Acknowledgement) Reset()         { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()    {}
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Acknowledgement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Acknowledgement.Unmarshal(m, b)
}
func (m *Acknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Acknowledgement.Marshal(b, m, deterministic)
}
func (m *Acknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acknowledgement.Merge(m, src)
}
func (m *Acknowledgement) XXX_Size() int {
	return xxx_messageInfo_Acknowledgement.Size(m)
}
func (m *Acknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_Acknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_Acknowledgement proto.InternalMessageInfo

func (m *Acknowledgement) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GreetCommand struct {
	AggregateIdentifier  string    `protobuf:"bytes,1,opt,name=aggregateIdentifier,proto3" json:"aggregateIdentifier,omitempty"`
	Message              *Greeting `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetCommand) Reset()         { *m = GreetCommand{} }
func (m *GreetCommand) String() string { return proto.CompactTextString(m) }
func (*GreetCommand) ProtoMessage()    {}
func (*GreetCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}

func (m *GreetCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetCommand.Unmarshal(m, b)
}
func (m *GreetCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetCommand.Marshal(b, m, deterministic)
}
func (m *GreetCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetCommand.Merge(m, src)
}
func (m *GreetCommand) XXX_Size() int {
	return xxx_messageInfo_GreetCommand.Size(m)
}
func (m *GreetCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetCommand.DiscardUnknown(m)
}

var xxx_messageInfo_GreetCommand proto.InternalMessageInfo

func (m *GreetCommand) GetAggregateIdentifier() string {
	if m != nil {
		return m.AggregateIdentifier
	}
	return ""
}

func (m *GreetCommand) GetMessage() *Greeting {
	if m != nil {
		return m.Message
	}
	return nil
}

type GreetedEvent struct {
	Message              *Greeting `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetedEvent) Reset()         { *m = GreetedEvent{} }
func (m *GreetedEvent) String() string { return proto.CompactTextString(m) }
func (*GreetedEvent) ProtoMessage()    {}
func (*GreetedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{4}
}

func (m *GreetedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetedEvent.Unmarshal(m, b)
}
func (m *GreetedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetedEvent.Marshal(b, m, deterministic)
}
func (m *GreetedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetedEvent.Merge(m, src)
}
func (m *GreetedEvent) XXX_Size() int {
	return xxx_messageInfo_GreetedEvent.Size(m)
}
func (m *GreetedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GreetedEvent proto.InternalMessageInfo

func (m *GreetedEvent) GetMessage() *Greeting {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Greeting)(nil), "org.leialearns.grpc.example.Greeting")
	proto.RegisterType((*Acknowledgement)(nil), "org.leialearns.grpc.example.Acknowledgement")
	proto.RegisterType((*Empty)(nil), "org.leialearns.grpc.example.Empty")
	proto.RegisterType((*GreetCommand)(nil), "org.leialearns.grpc.example.GreetCommand")
	proto.RegisterType((*GreetedEvent)(nil), "org.leialearns.grpc.example.GreetedEvent")
}

func init() {
	proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6)
}

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x86, 0x17, 0x61, 0xce, 0x1d, 0xbf, 0x20, 0x8a, 0x8c, 0x79, 0x33, 0x82, 0xc2, 0x40, 0x49,
	0xc7, 0xfc, 0x01, 0x43, 0x65, 0x88, 0x57, 0x82, 0xde, 0xed, 0x42, 0x88, 0xed, 0x31, 0x94, 0x35,
	0x1f, 0x9c, 0x84, 0xa9, 0x97, 0xfe, 0x3f, 0x7f, 0x94, 0xac, 0xb6, 0x0c, 0x45, 0xca, 0x76, 0xd9,
	0xe6, 0x79, 0xdf, 0x3c, 0xe7, 0x10, 0xd8, 0xc7, 0x77, 0x65, 0x7c, 0x81, 0xd2, 0x93, 0x8b, 0x8e,
	0x9f, 0x3a, 0xd2, 0xb2, 0xc0, 0x5c, 0x15, 0xa8, 0xc8, 0x06, 0xa9, 0xc9, 0xa7, 0xb2, 0x42, 0xc4,
	0x19, 0xec, 0xdc, 0x11, 0x62, 0xcc, 0xad, 0xe6, 0x3d, 0xe8, 0x18, 0x0c, 0x41, 0x69, 0xec, 0xb1,
	0x01, 0x1b, 0x76, 0x1f, 0xeb, 0x4f, 0x71, 0x01, 0x87, 0xd7, 0xe9, 0xdc, 0xba, 0xb7, 0x02, 0x33,
	0x8d, 0x06, 0x6d, 0x6c, 0x80, 0x3b, 0xd0, 0x9e, 0x1a, 0x1f, 0x3f, 0xc4, 0x27, 0x83, 0xbd, 0xb2,
	0xfc, 0xd6, 0x19, 0xa3, 0x6c, 0xc6, 0x47, 0x70, 0xa4, 0xb4, 0x26, 0xd4, 0x2a, 0xe2, 0x7d, 0x86,
	0x36, 0xe6, 0xaf, 0x39, 0x52, 0x95, 0xff, 0xef, 0x88, 0x4f, 0x56, 0xb7, 0x6c, 0x0d, 0xd8, 0x70,
	0x77, 0x7c, 0x2e, 0x1b, 0xa6, 0x91, 0xf5, 0x28, 0x2b, 0x99, 0x87, 0x4a, 0x01, 0xb3, 0xe9, 0x62,
	0xa9, 0x3d, 0xf9, 0xad, 0xbd, 0x71, 0xe1, 0xf8, 0x8b, 0xc1, 0xc1, 0x4f, 0x23, 0x3d, 0x21, 0x2d,
	0xf2, 0x14, 0xf9, 0x33, 0xb4, 0xcb, 0x3f, 0x7c, 0xbd, 0xae, 0xfe, 0x65, 0x23, 0xf6, 0x67, 0xd1,
	0xa2, 0xc5, 0x67, 0xd0, 0xad, 0xb3, 0x81, 0x8b, 0xc6, 0x70, 0xb9, 0xf8, 0xfe, 0x7a, 0x1e, 0xa2,
	0x35, 0x62, 0x37, 0x27, 0xb3, 0xe3, 0x40, 0x69, 0xe2, 0xe7, 0x3a, 0x59, 0x42, 0x49, 0x05, 0xbd,
	0x6c, 0x97, 0x6f, 0xe7, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xab, 0xa2, 0x6e, 0x9b, 0x4c, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
	// Greets AxonServer.
	Greet(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (*Acknowledgement, error)
	Greetings(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GreeterService_GreetingsClient, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Greet(ctx context.Context, in *Greeting, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, "/org.leialearns.grpc.example.GreeterService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) Greetings(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GreeterService_GreetingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreeterService_serviceDesc.Streams[0], "/org.leialearns.grpc.example.GreeterService/Greetings", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceGreetingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterService_GreetingsClient interface {
	Recv() (*Greeting, error)
	grpc.ClientStream
}

type greeterServiceGreetingsClient struct {
	grpc.ClientStream
}

func (x *greeterServiceGreetingsClient) Recv() (*Greeting, error) {
	m := new(Greeting)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
	// Greets AxonServer.
	Greet(context.Context, *Greeting) (*Acknowledgement, error)
	Greetings(*Empty, GreeterService_GreetingsServer) error
}

// UnimplementedGreeterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (*UnimplementedGreeterServiceServer) Greet(ctx context.Context, req *Greeting) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreeterServiceServer) Greetings(req *Empty, srv GreeterService_GreetingsServer) error {
	return status.Errorf(codes.Unimplemented, "method Greetings not implemented")
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

func _GreeterService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Greeting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.leialearns.grpc.example.GreeterService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Greet(ctx, req.(*Greeting))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_Greetings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServiceServer).Greetings(m, &greeterServiceGreetingsServer{stream})
}

type GreeterService_GreetingsServer interface {
	Send(*Greeting) error
	grpc.ServerStream
}

type greeterServiceGreetingsServer struct {
	grpc.ServerStream
}

func (x *greeterServiceGreetingsServer) Send(m *Greeting) error {
	return x.ServerStream.SendMsg(m)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.leialearns.grpc.example.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreeterService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greetings",
			Handler:       _GreeterService_Greetings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
