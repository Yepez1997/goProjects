// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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

// Calculate - api that takes two ints
type Calculate struct {
	FirstInt             int32    `protobuf:"varint,1,opt,name=first_int,json=firstInt,proto3" json:"first_int,omitempty"`
	SecondInt            int32    `protobuf:"varint,2,opt,name=second_int,json=secondInt,proto3" json:"second_int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculate) Reset()         { *m = Calculate{} }
func (m *Calculate) String() string { return proto.CompactTextString(m) }
func (*Calculate) ProtoMessage()    {}
func (*Calculate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *Calculate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculate.Unmarshal(m, b)
}
func (m *Calculate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculate.Marshal(b, m, deterministic)
}
func (m *Calculate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculate.Merge(m, src)
}
func (m *Calculate) XXX_Size() int {
	return xxx_messageInfo_Calculate.Size(m)
}
func (m *Calculate) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculate.DiscardUnknown(m)
}

var xxx_messageInfo_Calculate proto.InternalMessageInfo

func (m *Calculate) GetFirstInt() int32 {
	if m != nil {
		return m.FirstInt
	}
	return 0
}

func (m *Calculate) GetSecondInt() int32 {
	if m != nil {
		return m.SecondInt
	}
	return 0
}

// CalculateRequest - request object
type CalculateRequest struct {
	Calculate            *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CalculateRequest) Reset()         { *m = CalculateRequest{} }
func (m *CalculateRequest) String() string { return proto.CompactTextString(m) }
func (*CalculateRequest) ProtoMessage()    {}
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateRequest.Unmarshal(m, b)
}
func (m *CalculateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateRequest.Marshal(b, m, deterministic)
}
func (m *CalculateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateRequest.Merge(m, src)
}
func (m *CalculateRequest) XXX_Size() int {
	return xxx_messageInfo_CalculateRequest.Size(m)
}
func (m *CalculateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateRequest proto.InternalMessageInfo

func (m *CalculateRequest) GetCalculate() *Calculate {
	if m != nil {
		return m.Calculate
	}
	return nil
}

// CalculateResponse - response object
type CalculateResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateResponse) Reset()         { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()    {}
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalculateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateResponse.Unmarshal(m, b)
}
func (m *CalculateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateResponse.Marshal(b, m, deterministic)
}
func (m *CalculateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateResponse.Merge(m, src)
}
func (m *CalculateResponse) XXX_Size() int {
	return xxx_messageInfo_CalculateResponse.Size(m)
}
func (m *CalculateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateResponse proto.InternalMessageInfo

func (m *CalculateResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculate)(nil), "calculator.Calculate")
	proto.RegisterType((*CalculateRequest)(nil), "calculator.CalculateRequest")
	proto.RegisterType((*CalculateResponse)(nil), "calculator.CalculateResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x77, 0x2e, 0x4e, 0x67, 0x28, 0x2f, 0x55, 0x48,
	0x9a, 0x8b, 0x33, 0x2d, 0xb3, 0xa8, 0xb8, 0x24, 0x3e, 0x33, 0xaf, 0x44, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x35, 0x88, 0x03, 0x2c, 0xe0, 0x99, 0x57, 0x22, 0x24, 0xcb, 0xc5, 0x55, 0x9c, 0x9a, 0x9c,
	0x9f, 0x97, 0x02, 0x96, 0x65, 0x02, 0xcb, 0x72, 0x42, 0x44, 0x3c, 0xf3, 0x4a, 0x94, 0xdc, 0xb9,
	0x04, 0xe0, 0x06, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x73, 0x71, 0xc2, 0xac,
	0x4a, 0x05, 0x9b, 0xc7, 0x6d, 0x24, 0xaa, 0x87, 0xe4, 0x1c, 0x84, 0x06, 0x84, 0x3a, 0x25, 0x6d,
	0x2e, 0x41, 0x24, 0x83, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52,
	0x8b, 0x4b, 0x73, 0x60, 0xce, 0x82, 0xf2, 0x8c, 0x12, 0x91, 0x6c, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x15, 0xf2, 0xe5, 0xe2, 0x41, 0x88, 0x95, 0xe6, 0x0a, 0xc9, 0x60, 0xb7, 0x12, 0xe2,
	0x46, 0x29, 0x59, 0x1c, 0xb2, 0x10, 0x8b, 0x95, 0x18, 0x9c, 0xf8, 0xa2, 0x78, 0x90, 0x83, 0x33,
	0x89, 0x0d, 0x1c, 0x88, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x1b, 0x56, 0x27, 0x70,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	// Unary API call
	CalculateSum(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type calculateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculateServiceClient(cc *grpc.ClientConn) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) CalculateSum(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/CalculateSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	// Unary API call
	CalculateSum(context.Context, *CalculateRequest) (*CalculateResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) CalculateSum(ctx context.Context, req *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSum not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_CalculateSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).CalculateSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/CalculateSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).CalculateSum(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateSum",
			Handler:    _CalculateService_CalculateSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
