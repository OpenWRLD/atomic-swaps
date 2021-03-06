// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AuditRequest struct {
	ContractBytes        string   `protobuf:"bytes,1,opt,name=contractBytes,proto3" json:"contractBytes,omitempty"`
	TransactionBytes     string   `protobuf:"bytes,2,opt,name=transactionBytes,proto3" json:"transactionBytes,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditRequest) Reset()         { *m = AuditRequest{} }
func (m *AuditRequest) String() string { return proto.CompactTextString(m) }
func (*AuditRequest) ProtoMessage()    {}
func (*AuditRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{0}
}

func (m *AuditRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditRequest.Unmarshal(m, b)
}
func (m *AuditRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditRequest.Marshal(b, m, deterministic)
}
func (m *AuditRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditRequest.Merge(m, src)
}
func (m *AuditRequest) XXX_Size() int {
	return xxx_messageInfo_AuditRequest.Size(m)
}
func (m *AuditRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuditRequest proto.InternalMessageInfo

func (m *AuditRequest) GetContractBytes() string {
	if m != nil {
		return m.ContractBytes
	}
	return ""
}

func (m *AuditRequest) GetTransactionBytes() string {
	if m != nil {
		return m.TransactionBytes
	}
	return ""
}

func (m *AuditRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type AuditResponse struct {
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	ContractValue        string   `protobuf:"bytes,2,opt,name=contractValue,proto3" json:"contractValue,omitempty"`
	RecipientAddress     string   `protobuf:"bytes,3,opt,name=recipientAddress,proto3" json:"recipientAddress,omitempty"`
	RefundAddress        string   `protobuf:"bytes,4,opt,name=refundAddress,proto3" json:"refundAddress,omitempty"`
	SecretHash           string   `protobuf:"bytes,5,opt,name=secretHash,proto3" json:"secretHash,omitempty"`
	LockTimeBlock        string   `protobuf:"bytes,6,opt,name=lockTimeBlock,proto3" json:"lockTimeBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditResponse) Reset()         { *m = AuditResponse{} }
func (m *AuditResponse) String() string { return proto.CompactTextString(m) }
func (*AuditResponse) ProtoMessage()    {}
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{1}
}

func (m *AuditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditResponse.Unmarshal(m, b)
}
func (m *AuditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditResponse.Marshal(b, m, deterministic)
}
func (m *AuditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditResponse.Merge(m, src)
}
func (m *AuditResponse) XXX_Size() int {
	return xxx_messageInfo_AuditResponse.Size(m)
}
func (m *AuditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditResponse proto.InternalMessageInfo

func (m *AuditResponse) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *AuditResponse) GetContractValue() string {
	if m != nil {
		return m.ContractValue
	}
	return ""
}

func (m *AuditResponse) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *AuditResponse) GetRefundAddress() string {
	if m != nil {
		return m.RefundAddress
	}
	return ""
}

func (m *AuditResponse) GetSecretHash() string {
	if m != nil {
		return m.SecretHash
	}
	return ""
}

func (m *AuditResponse) GetLockTimeBlock() string {
	if m != nil {
		return m.LockTimeBlock
	}
	return ""
}

func init() {
	proto.RegisterType((*AuditRequest)(nil), "api.AuditRequest")
	proto.RegisterType((*AuditResponse)(nil), "api.AuditResponse")
}

func init() { proto.RegisterFile("audit.proto", fileDescriptor_5594839dd8e38a1b) }

var fileDescriptor_5594839dd8e38a1b = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0x96, 0x54, 0xd4, 0x50, 0x7e, 0x2c, 0x86, 0x28, 0x42, 0x08, 0x45, 0x0c, 0x55,
	0x87, 0x44, 0xc0, 0xc6, 0xd6, 0x8a, 0x81, 0x85, 0xa5, 0x42, 0xec, 0x17, 0xe7, 0x52, 0x2c, 0x82,
	0x6d, 0x6c, 0x07, 0xa9, 0x12, 0x13, 0xaf, 0xc0, 0xa3, 0xf1, 0x0a, 0xbc, 0x03, 0x2b, 0xb2, 0xdd,
	0x54, 0x49, 0xbb, 0xd9, 0x9f, 0xce, 0x3d, 0xf7, 0xd8, 0x87, 0xec, 0x41, 0x5d, 0x72, 0x9b, 0x2b,
	0x2d, 0xad, 0xa4, 0x7d, 0x50, 0x3c, 0x3d, 0x5d, 0x48, 0xb9, 0xa8, 0xb0, 0x00, 0xc5, 0x0b, 0x10,
	0x42, 0x5a, 0xb0, 0x5c, 0x0a, 0x13, 0x24, 0xd9, 0x27, 0xd9, 0x9f, 0xba, 0x89, 0x39, 0xbe, 0xd7,
	0x68, 0x2c, 0xbd, 0x20, 0x23, 0x26, 0x85, 0xd5, 0xc0, 0xec, 0x6c, 0x69, 0xd1, 0x24, 0xd1, 0x79,
	0x34, 0x1e, 0xce, 0xbb, 0x90, 0x4e, 0xc8, 0x91, 0xd5, 0x20, 0x0c, 0x30, 0xe7, 0x15, 0x84, 0x3d,
	0x2f, 0xdc, 0xe2, 0x34, 0x25, 0xbb, 0xac, 0xd6, 0x1a, 0x05, 0x5b, 0x26, 0x7d, 0xaf, 0x59, 0xdf,
	0xb3, 0xbf, 0x88, 0x8c, 0x56, 0xeb, 0x8d, 0x92, 0xc2, 0x20, 0x1d, 0x93, 0xc3, 0x66, 0xd5, 0xb4,
	0x2c, 0x35, 0x9a, 0x26, 0xc1, 0x26, 0x6e, 0x27, 0x7d, 0x84, 0xaa, 0xc6, 0x55, 0x80, 0x2e, 0x74,
	0x49, 0x35, 0x32, 0xae, 0x38, 0x8a, 0xb5, 0x61, 0x48, 0xb1, 0xc5, 0x9d, 0xa3, 0xc6, 0xe7, 0x5a,
	0x94, 0x8d, 0x70, 0x27, 0x38, 0x76, 0x20, 0x3d, 0x23, 0xc4, 0x20, 0xd3, 0x68, 0xef, 0xc0, 0xbc,
	0x24, 0xb1, 0x97, 0xb4, 0x88, 0x73, 0xa9, 0x24, 0x7b, 0x7d, 0xe0, 0x6f, 0x38, 0x73, 0x87, 0x64,
	0x10, 0x5c, 0x3a, 0xf0, 0xea, 0x9e, 0xc4, 0xfe, 0xe1, 0xf4, 0x96, 0xc4, 0xbe, 0x32, 0x7a, 0x9c,
	0x83, 0xe2, 0x79, 0xbb, 0x8c, 0x94, 0xb6, 0x51, 0xf8, 0xa0, 0xec, 0xe4, 0xeb, 0xe7, 0xf7, 0xbb,
	0x77, 0x90, 0x0d, 0x8b, 0x8f, 0xcb, 0xc2, 0x4f, 0xde, 0x44, 0x93, 0xa7, 0x81, 0x6f, 0xf3, 0xfa,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x0d, 0x90, 0xf0, 0xff, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuditClient is the client API for Audit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditClient interface {
	Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error)
}

type auditClient struct {
	cc *grpc.ClientConn
}

func NewAuditClient(cc *grpc.ClientConn) AuditClient {
	return &auditClient{cc}
}

func (c *auditClient) Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error) {
	out := new(AuditResponse)
	err := c.cc.Invoke(ctx, "/api.Audit/audit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServer is the server API for Audit service.
type AuditServer interface {
	Audit(context.Context, *AuditRequest) (*AuditResponse, error)
}

// UnimplementedAuditServer can be embedded to have forward compatible implementations.
type UnimplementedAuditServer struct {
}

func (*UnimplementedAuditServer) Audit(ctx context.Context, req *AuditRequest) (*AuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Audit not implemented")
}

func RegisterAuditServer(s *grpc.Server, srv AuditServer) {
	s.RegisterService(&_Audit_serviceDesc, srv)
}

func _Audit_Audit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).Audit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Audit/Audit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).Audit(ctx, req.(*AuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Audit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Audit",
	HandlerType: (*AuditServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "audit",
			Handler:    _Audit_Audit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit.proto",
}
