// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authz.proto

/*
Package istio_v1_authz is a generated protocol buffer package.

It is generated from these files:
	authz.proto

It has these top-level messages:
	Request
	Response
*/
package istio_v1_authz

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Response_Status_Code int32

const (
	// https://cloud.google.com/appengine/docs/admin-api/reference/rpc/google.rpc#google.rpc.Code
	Response_Status_OK                  Response_Status_Code = 0
	Response_Status_CANCELLED           Response_Status_Code = 1
	Response_Status_UNKNOWN             Response_Status_Code = 2
	Response_Status_INVALID_ARGUMENT    Response_Status_Code = 3
	Response_Status_DEADLINE_EXCEEDED   Response_Status_Code = 4
	Response_Status_NOT_FOUND           Response_Status_Code = 5
	Response_Status_ALREADY_EXISTS      Response_Status_Code = 6
	Response_Status_PERMISSION_DENIED   Response_Status_Code = 7
	Response_Status_UNAUTHENTICATED     Response_Status_Code = 8
	Response_Status_RESOURCE_EXHAUSTED  Response_Status_Code = 9
	Response_Status_FAILED_PRECONDITION Response_Status_Code = 10
	Response_Status_ABORTED             Response_Status_Code = 11
	Response_Status_OUT_OF_RANGE        Response_Status_Code = 12
	Response_Status_UNIMPLEMENTED       Response_Status_Code = 13
	Response_Status_INTERNAL            Response_Status_Code = 14
	Response_Status_UNAVAILABLE         Response_Status_Code = 15
	Response_Status_DATA_LOSS           Response_Status_Code = 16
)

var Response_Status_Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	8:  "UNAUTHENTICATED",
	9:  "RESOURCE_EXHAUSTED",
	10: "FAILED_PRECONDITION",
	11: "ABORTED",
	12: "OUT_OF_RANGE",
	13: "UNIMPLEMENTED",
	14: "INTERNAL",
	15: "UNAVAILABLE",
	16: "DATA_LOSS",
}
var Response_Status_Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     8,
	"RESOURCE_EXHAUSTED":  9,
	"FAILED_PRECONDITION": 10,
	"ABORTED":             11,
	"OUT_OF_RANGE":        12,
	"UNIMPLEMENTED":       13,
	"INTERNAL":            14,
	"UNAVAILABLE":         15,
	"DATA_LOSS":           16,
}

func (x Response_Status_Code) String() string {
	return proto.EnumName(Response_Status_Code_name, int32(x))
}
func (Response_Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0, 0} }

type Request struct {
	Subject *Request_Subject `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	Action  *Request_Action  `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetSubject() *Request_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Request) GetAction() *Request_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

type Request_Subject struct {
	ServiceAccount       string            `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	Namespace            string            `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Service              string            `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Pod                  string            `protobuf:"bytes,4,opt,name=pod" json:"pod,omitempty"`
	IpAddress            string            `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	Port                 string            `protobuf:"bytes,6,opt,name=port" json:"port,omitempty"`
	ServiceAccountLabels map[string]string `protobuf:"bytes,7,rep,name=service_account_labels,json=serviceAccountLabels" json:"service_account_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request_Subject) Reset()                    { *m = Request_Subject{} }
func (m *Request_Subject) String() string            { return proto.CompactTextString(m) }
func (*Request_Subject) ProtoMessage()               {}
func (*Request_Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Request_Subject) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *Request_Subject) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Request_Subject) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request_Subject) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *Request_Subject) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Request_Subject) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Request_Subject) GetServiceAccountLabels() map[string]string {
	if m != nil {
		return m.ServiceAccountLabels
	}
	return nil
}

type Request_Action struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Service   string `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Port      string `protobuf:"bytes,3,opt,name=port" json:"port,omitempty"`
	IpAddress string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
}

func (m *Request_Action) Reset()                    { *m = Request_Action{} }
func (m *Request_Action) String() string            { return proto.CompactTextString(m) }
func (*Request_Action) ProtoMessage()               {}
func (*Request_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Request_Action) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Request_Action) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request_Action) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Request_Action) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type Response struct {
	Status *Response_Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetStatus() *Response_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// TODO Maybe replace with googleapis
type Response_Status struct {
	Code    Response_Status_Code `protobuf:"varint,1,opt,name=code,enum=istio.v1.authz.Response_Status_Code" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response_Status) Reset()                    { *m = Response_Status{} }
func (m *Response_Status) String() string            { return proto.CompactTextString(m) }
func (*Response_Status) ProtoMessage()               {}
func (*Response_Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Status) GetCode() Response_Status_Code {
	if m != nil {
		return m.Code
	}
	return Response_Status_OK
}

func (m *Response_Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "istio.v1.authz.Request")
	proto.RegisterType((*Request_Subject)(nil), "istio.v1.authz.Request.Subject")
	proto.RegisterType((*Request_Action)(nil), "istio.v1.authz.Request.Action")
	proto.RegisterType((*Response)(nil), "istio.v1.authz.Response")
	proto.RegisterType((*Response_Status)(nil), "istio.v1.authz.Response.Status")
	proto.RegisterEnum("istio.v1.authz.Response_Status_Code", Response_Status_Code_name, Response_Status_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authorization service

type AuthorizationClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/istio.v1.authz.Authorization/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationServer interface {
	Check(context.Context, *Request) (*Response, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.v1.authz.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.v1.authz.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz.proto",
}

func init() { proto.RegisterFile("authz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0x37, 0x5f, 0x0e, 0x39, 0x21, 0xc9, 0x70, 0x60, 0xc1, 0x1b, 0xed, 0x07, 0x42, 0x2b,
	0x2d, 0x57, 0x96, 0x36, 0x95, 0x5a, 0xe0, 0x6e, 0xb0, 0x07, 0xb0, 0x30, 0x63, 0x34, 0xb6, 0x29,
	0xbd, 0xb2, 0x8c, 0x33, 0x2a, 0x29, 0x10, 0xa7, 0x19, 0x07, 0x09, 0x1e, 0xa4, 0xaf, 0xd2, 0x97,
	0x69, 0x1f, 0xa0, 0x6f, 0x51, 0xf9, 0x23, 0x6d, 0x83, 0x9a, 0x72, 0x37, 0xf3, 0x9f, 0xdf, 0x39,
	0x73, 0xfe, 0xff, 0x4c, 0x0c, 0xed, 0x68, 0x96, 0x5e, 0x3f, 0x1a, 0x93, 0x69, 0x92, 0x26, 0xd8,
	0x1d, 0xa9, 0x74, 0x94, 0x18, 0xf7, 0xff, 0x1b, 0xb9, 0xba, 0xf3, 0xb9, 0x0e, 0x4d, 0x21, 0xdf,
	0xcf, 0xa4, 0x4a, 0x71, 0x1f, 0x9a, 0x6a, 0x76, 0xf5, 0x4e, 0xc6, 0xa9, 0x5e, 0xd9, 0xae, 0xec,
	0xb6, 0x07, 0xff, 0x18, 0x8b, 0xb4, 0x51, 0x92, 0x86, 0x57, 0x60, 0x62, 0xce, 0xe3, 0x4b, 0xd0,
	0xa2, 0x38, 0x1d, 0x25, 0x63, 0xbd, 0x9a, 0x57, 0xfe, 0xbd, 0xac, 0x92, 0xe6, 0x94, 0x28, 0xe9,
	0xfe, 0x97, 0x2a, 0x34, 0xcb, 0x66, 0xf8, 0x1f, 0xf4, 0x94, 0x9c, 0xde, 0x8f, 0x62, 0x19, 0x46,
	0x71, 0x9c, 0xcc, 0xc6, 0xc5, 0x18, 0x2d, 0xd1, 0x2d, 0x65, 0x5a, 0xa8, 0xf8, 0x27, 0xb4, 0xc6,
	0xd1, 0x9d, 0x54, 0x93, 0x28, 0x96, 0xf9, 0x7d, 0x2d, 0xf1, 0x5d, 0x40, 0x1d, 0x9a, 0x25, 0xaf,
	0xd7, 0xf2, 0xb3, 0xf9, 0x16, 0x09, 0xd4, 0x26, 0xc9, 0x50, 0xaf, 0xe7, 0x6a, 0xb6, 0xc4, 0xbf,
	0x00, 0x46, 0x93, 0x30, 0x1a, 0x0e, 0xa7, 0x52, 0x29, 0xbd, 0x51, 0xb4, 0x1a, 0x4d, 0x68, 0x21,
	0x20, 0x42, 0x7d, 0x92, 0x4c, 0x53, 0x5d, 0xcb, 0x0f, 0xf2, 0x35, 0x26, 0xb0, 0xf9, 0x64, 0xca,
	0xf0, 0x36, 0xba, 0x92, 0xb7, 0x4a, 0x6f, 0x6e, 0xd7, 0x76, 0xdb, 0x83, 0xfd, 0x67, 0x32, 0x33,
	0xbc, 0x05, 0x33, 0x4e, 0x5e, 0xcb, 0xc6, 0xe9, 0xf4, 0x41, 0x6c, 0xa8, 0x9f, 0x1c, 0xf5, 0x8f,
	0xe1, 0x8f, 0xa5, 0x25, 0x99, 0xa5, 0x1b, 0xf9, 0x50, 0xe6, 0x94, 0x2d, 0x71, 0x03, 0x1a, 0xf7,
	0xd1, 0xed, 0x6c, 0x1e, 0x4c, 0xb1, 0x39, 0xa8, 0xee, 0x55, 0xfa, 0x0a, 0xb4, 0x22, 0xfd, 0xc5,
	0x00, 0x2b, 0xbf, 0x08, 0xb0, 0xba, 0x18, 0xe0, 0x3c, 0x8f, 0xda, 0x0f, 0x79, 0x2c, 0x46, 0x58,
	0x7f, 0x12, 0xe1, 0xce, 0xa7, 0x1a, 0xac, 0x08, 0xa9, 0x26, 0xc9, 0x58, 0x49, 0x7c, 0x05, 0x9a,
	0x4a, 0xa3, 0x74, 0xa6, 0x96, 0xbf, 0xaf, 0x82, 0x34, 0xbc, 0x1c, 0x13, 0x25, 0xde, 0xff, 0x50,
	0x03, 0xad, 0x90, 0x70, 0x0f, 0xea, 0x71, 0x32, 0x2c, 0xc6, 0xee, 0x0e, 0xfe, 0x7d, 0xa6, 0x83,
	0x61, 0x26, 0x43, 0x29, 0xf2, 0x8a, 0xcc, 0xd7, 0x9d, 0x54, 0x2a, 0x7a, 0xfb, 0xcd, 0x57, 0xb9,
	0xdd, 0xf9, 0x58, 0x85, 0x7a, 0x06, 0xa2, 0x06, 0x55, 0xf7, 0x94, 0xfc, 0x86, 0x1d, 0x68, 0x99,
	0x94, 0x9b, 0xcc, 0x71, 0x98, 0x45, 0x2a, 0xd8, 0x86, 0x66, 0xc0, 0x4f, 0xb9, 0xfb, 0x9a, 0x93,
	0x2a, 0x6e, 0x00, 0xb1, 0xf9, 0x05, 0x75, 0x6c, 0x2b, 0xa4, 0xe2, 0x38, 0x38, 0x63, 0xdc, 0x27,
	0x35, 0xfc, 0x1d, 0xd6, 0x2c, 0x46, 0x2d, 0xc7, 0xe6, 0x2c, 0x64, 0x97, 0x26, 0x63, 0x16, 0xb3,
	0x48, 0x3d, 0x6b, 0xc4, 0x5d, 0x3f, 0x3c, 0x72, 0x03, 0x6e, 0x91, 0x06, 0x22, 0x74, 0xa9, 0x23,
	0x18, 0xb5, 0xde, 0x84, 0xec, 0xd2, 0xf6, 0x7c, 0x8f, 0x68, 0x59, 0xe5, 0x39, 0x13, 0x67, 0xb6,
	0xe7, 0xd9, 0x2e, 0x0f, 0x2d, 0xc6, 0x6d, 0x66, 0x91, 0x26, 0xae, 0x43, 0x2f, 0xe0, 0x34, 0xf0,
	0x4f, 0x18, 0xf7, 0x6d, 0x93, 0xfa, 0xcc, 0x22, 0x2b, 0xb8, 0x09, 0x28, 0x98, 0xe7, 0x06, 0xc2,
	0xcc, 0x6e, 0x39, 0xa1, 0x81, 0x97, 0xe9, 0x2d, 0xdc, 0x82, 0xf5, 0x23, 0x6a, 0x3b, 0xcc, 0x0a,
	0xcf, 0x05, 0x33, 0x5d, 0x6e, 0xd9, 0xbe, 0xed, 0x72, 0x02, 0xd9, 0xe4, 0xf4, 0xd0, 0x15, 0x19,
	0xd5, 0x46, 0x02, 0xab, 0x6e, 0xe0, 0x87, 0xee, 0x51, 0x28, 0x28, 0x3f, 0x66, 0x64, 0x15, 0xd7,
	0xa0, 0x13, 0x70, 0xfb, 0xec, 0xdc, 0x61, 0x99, 0x0d, 0x66, 0x91, 0x0e, 0xae, 0xc2, 0x8a, 0xcd,
	0x7d, 0x26, 0x38, 0x75, 0x48, 0x17, 0x7b, 0xd0, 0x0e, 0x38, 0xbd, 0xa0, 0xb6, 0x43, 0x0f, 0x1d,
	0x46, 0x7a, 0x99, 0x21, 0x8b, 0xfa, 0x34, 0x74, 0x5c, 0xcf, 0x23, 0x64, 0x70, 0x0a, 0x1d, 0x3a,
	0x4b, 0xaf, 0x93, 0xe9, 0xe8, 0x31, 0xca, 0x9f, 0xd6, 0x01, 0x34, 0xcc, 0x6b, 0x19, 0xdf, 0xe0,
	0xd6, 0x92, 0xff, 0x41, 0x5f, 0x5f, 0xf6, 0x93, 0x5d, 0x69, 0xf9, 0x27, 0xea, 0xc5, 0xd7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x3c, 0x5a, 0x42, 0xb1, 0x04, 0x00, 0x00,
}
