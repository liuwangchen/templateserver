// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

package health

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// The definition of health info.
type Health struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{0}
}

func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (m *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(m, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Health)(nil), "health.Health")
}

func init() { proto.RegisterFile("health.proto", fileDescriptor_fdbebe66dda7cb29) }

var fileDescriptor_fdbebe66dda7cb29 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xaa, 0xa4, 0xa4, 0xa1, 0xb2, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e,
	0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0x52, 0x49, 0x89, 0x8b, 0xcd, 0x03, 0x6c, 0x88, 0x90, 0x04,
	0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x8c, 0x6b, 0x14, 0xce, 0xc5, 0x0b, 0x51, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4,
	0xc6, 0xc5, 0x9d, 0x9c, 0x91, 0x9a, 0x9c, 0x0d, 0xd5, 0x29, 0xa6, 0x07, 0xb1, 0x41, 0x0f, 0x66,
	0x83, 0x9e, 0x2b, 0xc8, 0x06, 0x29, 0x3e, 0x3d, 0xa8, 0x6b, 0x21, 0xea, 0x94, 0xf8, 0x9b, 0x2e,
	0x3f, 0x99, 0xcc, 0xc4, 0x29, 0xc4, 0xae, 0x0f, 0x11, 0x77, 0x52, 0x8d, 0x52, 0x2e, 0x49, 0xcd,
	0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0xc8, 0x4e, 0xd7,
	0x2f, 0x2a, 0x48, 0x86, 0xb8, 0x15, 0xaa, 0x2c, 0x89, 0x0d, 0xcc, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xff, 0x01, 0x7a, 0x77, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	CheckHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error)
}

type healthServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthServiceClient(cc *grpc.ClientConn) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) CheckHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/health.HealthService/checkHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	CheckHealth(context.Context, *empty.Empty) (*Health, error)
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.HealthService/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).CheckHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkHealth",
			Handler:    _HealthService_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}
