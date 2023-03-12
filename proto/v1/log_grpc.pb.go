// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: v1/log.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LogService_CreateLogLine_FullMethodName      = "/v1.LogService/CreateLogLine"
	LogService_CreateLogLines_FullMethodName     = "/v1.LogService/CreateLogLines"
	LogService_GetLogLineTotal_FullMethodName    = "/v1.LogService/GetLogLineTotal"
	LogService_GetLogLinesHistory_FullMethodName = "/v1.LogService/GetLogLinesHistory"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	CreateLogLine(ctx context.Context, in *CreateLogLineRequest, opts ...grpc.CallOption) (*CreateLogLineResponse, error)
	CreateLogLines(ctx context.Context, in *CreateLogLinesRequest, opts ...grpc.CallOption) (*CreateLogLinesResponse, error)
	GetLogLineTotal(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLogLineTotalResponse, error)
	GetLogLinesHistory(ctx context.Context, in *GetLogLinesHistoryRequest, opts ...grpc.CallOption) (*GetLogLinesHistoryResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) CreateLogLine(ctx context.Context, in *CreateLogLineRequest, opts ...grpc.CallOption) (*CreateLogLineResponse, error) {
	out := new(CreateLogLineResponse)
	err := c.cc.Invoke(ctx, LogService_CreateLogLine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) CreateLogLines(ctx context.Context, in *CreateLogLinesRequest, opts ...grpc.CallOption) (*CreateLogLinesResponse, error) {
	out := new(CreateLogLinesResponse)
	err := c.cc.Invoke(ctx, LogService_CreateLogLines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetLogLineTotal(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLogLineTotalResponse, error) {
	out := new(GetLogLineTotalResponse)
	err := c.cc.Invoke(ctx, LogService_GetLogLineTotal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetLogLinesHistory(ctx context.Context, in *GetLogLinesHistoryRequest, opts ...grpc.CallOption) (*GetLogLinesHistoryResponse, error) {
	out := new(GetLogLinesHistoryResponse)
	err := c.cc.Invoke(ctx, LogService_GetLogLinesHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	CreateLogLine(context.Context, *CreateLogLineRequest) (*CreateLogLineResponse, error)
	CreateLogLines(context.Context, *CreateLogLinesRequest) (*CreateLogLinesResponse, error)
	GetLogLineTotal(context.Context, *emptypb.Empty) (*GetLogLineTotalResponse, error)
	GetLogLinesHistory(context.Context, *GetLogLinesHistoryRequest) (*GetLogLinesHistoryResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) CreateLogLine(context.Context, *CreateLogLineRequest) (*CreateLogLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogLine not implemented")
}
func (UnimplementedLogServiceServer) CreateLogLines(context.Context, *CreateLogLinesRequest) (*CreateLogLinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogLines not implemented")
}
func (UnimplementedLogServiceServer) GetLogLineTotal(context.Context, *emptypb.Empty) (*GetLogLineTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogLineTotal not implemented")
}
func (UnimplementedLogServiceServer) GetLogLinesHistory(context.Context, *GetLogLinesHistoryRequest) (*GetLogLinesHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogLinesHistory not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_CreateLogLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CreateLogLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CreateLogLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CreateLogLine(ctx, req.(*CreateLogLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_CreateLogLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogLinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CreateLogLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CreateLogLines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CreateLogLines(ctx, req.(*CreateLogLinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetLogLineTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetLogLineTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_GetLogLineTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetLogLineTotal(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetLogLinesHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogLinesHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetLogLinesHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_GetLogLinesHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetLogLinesHistory(ctx, req.(*GetLogLinesHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLogLine",
			Handler:    _LogService_CreateLogLine_Handler,
		},
		{
			MethodName: "CreateLogLines",
			Handler:    _LogService_CreateLogLines_Handler,
		},
		{
			MethodName: "GetLogLineTotal",
			Handler:    _LogService_GetLogLineTotal_Handler,
		},
		{
			MethodName: "GetLogLinesHistory",
			Handler:    _LogService_GetLogLinesHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/log.proto",
}