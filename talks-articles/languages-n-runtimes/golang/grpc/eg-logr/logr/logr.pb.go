// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logr.proto

/*
Package logr is a generated protocol buffer package.

It is generated from these files:
	logr.proto

It has these top-level messages:
	LogrRequest
	LogrResponse
	LogrFilter
*/
package logr

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

// to create logr request
type LogrRequest struct {
	App     string                  `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Logline string                  `protobuf:"bytes,2,opt,name=logline" json:"logline,omitempty"`
	Tags    []*LogrRequest_LogrTags `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *LogrRequest) Reset()                    { *m = LogrRequest{} }
func (m *LogrRequest) String() string            { return proto.CompactTextString(m) }
func (*LogrRequest) ProtoMessage()               {}
func (*LogrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogrRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *LogrRequest) GetLogline() string {
	if m != nil {
		return m.Logline
	}
	return ""
}

func (m *LogrRequest) GetTags() []*LogrRequest_LogrTags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type LogrRequest_LogrTags struct {
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
}

func (m *LogrRequest_LogrTags) Reset()                    { *m = LogrRequest_LogrTags{} }
func (m *LogrRequest_LogrTags) String() string            { return proto.CompactTextString(m) }
func (*LogrRequest_LogrTags) ProtoMessage()               {}
func (*LogrRequest_LogrTags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *LogrRequest_LogrTags) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// to create logr response
type LogrResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *LogrResponse) Reset()                    { *m = LogrResponse{} }
func (m *LogrResponse) String() string            { return proto.CompactTextString(m) }
func (*LogrResponse) ProtoMessage()               {}
func (*LogrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogrResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogrResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// to create logr filter
type LogrFilter struct {
	App  string                 `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Tags []*LogrFilter_LogrTags `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *LogrFilter) Reset()                    { *m = LogrFilter{} }
func (m *LogrFilter) String() string            { return proto.CompactTextString(m) }
func (*LogrFilter) ProtoMessage()               {}
func (*LogrFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogrFilter) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *LogrFilter) GetTags() []*LogrFilter_LogrTags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type LogrFilter_LogrTags struct {
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
}

func (m *LogrFilter_LogrTags) Reset()                    { *m = LogrFilter_LogrTags{} }
func (m *LogrFilter_LogrTags) String() string            { return proto.CompactTextString(m) }
func (*LogrFilter_LogrTags) ProtoMessage()               {}
func (*LogrFilter_LogrTags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *LogrFilter_LogrTags) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*LogrRequest)(nil), "logr.LogrRequest")
	proto.RegisterType((*LogrRequest_LogrTags)(nil), "logr.LogrRequest.LogrTags")
	proto.RegisterType((*LogrResponse)(nil), "logr.LogrResponse")
	proto.RegisterType((*LogrFilter)(nil), "logr.LogrFilter")
	proto.RegisterType((*LogrFilter_LogrTags)(nil), "logr.LogrFilter.LogrTags")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Logr service

type LogrClient interface {
	GetLogs(ctx context.Context, in *LogrFilter, opts ...grpc.CallOption) (Logr_GetLogsClient, error)
	CreateLog(ctx context.Context, in *LogrRequest, opts ...grpc.CallOption) (*LogrResponse, error)
}

type logrClient struct {
	cc *grpc.ClientConn
}

func NewLogrClient(cc *grpc.ClientConn) LogrClient {
	return &logrClient{cc}
}

func (c *logrClient) GetLogs(ctx context.Context, in *LogrFilter, opts ...grpc.CallOption) (Logr_GetLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Logr_serviceDesc.Streams[0], c.cc, "/logr.Logr/GetLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logrGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Logr_GetLogsClient interface {
	Recv() (*LogrRequest, error)
	grpc.ClientStream
}

type logrGetLogsClient struct {
	grpc.ClientStream
}

func (x *logrGetLogsClient) Recv() (*LogrRequest, error) {
	m := new(LogrRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logrClient) CreateLog(ctx context.Context, in *LogrRequest, opts ...grpc.CallOption) (*LogrResponse, error) {
	out := new(LogrResponse)
	err := grpc.Invoke(ctx, "/logr.Logr/CreateLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logr service

type LogrServer interface {
	GetLogs(*LogrFilter, Logr_GetLogsServer) error
	CreateLog(context.Context, *LogrRequest) (*LogrResponse, error)
}

func RegisterLogrServer(s *grpc.Server, srv LogrServer) {
	s.RegisterService(&_Logr_serviceDesc, srv)
}

func _Logr_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogrFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogrServer).GetLogs(m, &logrGetLogsServer{stream})
}

type Logr_GetLogsServer interface {
	Send(*LogrRequest) error
	grpc.ServerStream
}

type logrGetLogsServer struct {
	grpc.ServerStream
}

func (x *logrGetLogsServer) Send(m *LogrRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Logr_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogrServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logr.Logr/CreateLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogrServer).CreateLog(ctx, req.(*LogrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logr.Logr",
	HandlerType: (*LogrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLog",
			Handler:    _Logr_CreateLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _Logr_GetLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logr.proto",
}

func init() { proto.RegisterFile("logr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0xc3, 0x40,
	0x0c, 0x6e, 0x9a, 0x88, 0xb6, 0x2e, 0x03, 0xdc, 0x14, 0x22, 0x86, 0xea, 0xa6, 0x2e, 0x44, 0x28,
	0x30, 0xb1, 0x22, 0xc1, 0x92, 0x29, 0xe2, 0x05, 0x8e, 0xca, 0xb2, 0x2a, 0x42, 0xef, 0x38, 0xbb,
	0xcf, 0xc0, 0x6b, 0x23, 0x5f, 0x1a, 0x51, 0x51, 0x96, 0x6e, 0x9f, 0x9d, 0xef, 0x2f, 0x3e, 0x80,
	0xde, 0x53, 0xac, 0x43, 0xf4, 0xe2, 0x4d, 0xa1, 0xd8, 0x7e, 0x67, 0xb0, 0x6c, 0x3d, 0xc5, 0x0e,
	0xbf, 0xf6, 0xc8, 0x62, 0xae, 0x20, 0x77, 0x21, 0x94, 0xd9, 0x2a, 0x5b, 0x2f, 0x3a, 0x85, 0xa6,
	0x84, 0x59, 0xef, 0xa9, 0xdf, 0xee, 0xb0, 0x9c, 0xa6, 0xed, 0x38, 0x9a, 0x1a, 0x0a, 0x71, 0xc4,
	0x65, 0xbe, 0xca, 0xd7, 0xcb, 0xa6, 0xaa, 0x93, 0xf9, 0x91, 0x59, 0xc2, 0x6f, 0x8e, 0xb8, 0x4b,
	0xbc, 0xea, 0x16, 0xe6, 0xe3, 0x46, 0x73, 0xc4, 0xd1, 0x98, 0x23, 0x8e, 0xec, 0x13, 0x5c, 0x0e,
	0x5a, 0x0e, 0x7e, 0xc7, 0xa8, 0xb9, 0xbc, 0xdf, 0x6c, 0x90, 0x39, 0xb1, 0xe6, 0xdd, 0x38, 0xaa,
	0xf6, 0x93, 0xe9, 0xd0, 0x46, 0xa1, 0xfd, 0x00, 0x50, 0xed, 0xcb, 0xb6, 0x17, 0x8c, 0xff, 0xfc,
	0xc3, 0xdd, 0xa1, 0xe9, 0x34, 0x35, 0xbd, 0xf9, 0x6d, 0x3a, 0x28, 0xce, 0x2a, 0xda, 0x04, 0x28,
	0xf4, 0xab, 0x69, 0x60, 0xf6, 0x8a, 0xd2, 0x7a, 0x25, 0xfd, 0x75, 0xac, 0xae, 0x4f, 0xae, 0x61,
	0x27, 0xf7, 0x99, 0x79, 0x84, 0xc5, 0x73, 0x44, 0x27, 0xd8, 0x7a, 0x32, 0xa7, 0x9c, 0xca, 0x1c,
	0xaf, 0x86, 0x43, 0xd8, 0xc9, 0xfb, 0x45, 0x7a, 0xb1, 0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7d, 0x3f, 0x12, 0x8e, 0xbf, 0x01, 0x00, 0x00,
}
