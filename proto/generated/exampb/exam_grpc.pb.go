// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: exam.proto

package exampb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExamService_GetExamResult_FullMethodName = "/exam.ExamService/GetExamResult"
)

// ExamServiceClient is the client API for ExamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamServiceClient interface {
	GetExamResult(ctx context.Context, in *GetExamResultRequest, opts ...grpc.CallOption) (*GetExamResultResponse, error)
}

type examServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExamServiceClient(cc grpc.ClientConnInterface) ExamServiceClient {
	return &examServiceClient{cc}
}

func (c *examServiceClient) GetExamResult(ctx context.Context, in *GetExamResultRequest, opts ...grpc.CallOption) (*GetExamResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExamResultResponse)
	err := c.cc.Invoke(ctx, ExamService_GetExamResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExamServiceServer is the server API for ExamService service.
// All implementations must embed UnimplementedExamServiceServer
// for forward compatibility.
type ExamServiceServer interface {
	GetExamResult(context.Context, *GetExamResultRequest) (*GetExamResultResponse, error)
	mustEmbedUnimplementedExamServiceServer()
}

// UnimplementedExamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExamServiceServer struct{}

func (UnimplementedExamServiceServer) GetExamResult(context.Context, *GetExamResultRequest) (*GetExamResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExamResult not implemented")
}
func (UnimplementedExamServiceServer) mustEmbedUnimplementedExamServiceServer() {}
func (UnimplementedExamServiceServer) testEmbeddedByValue()                     {}

// UnsafeExamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExamServiceServer will
// result in compilation errors.
type UnsafeExamServiceServer interface {
	mustEmbedUnimplementedExamServiceServer()
}

func RegisterExamServiceServer(s grpc.ServiceRegistrar, srv ExamServiceServer) {
	// If the following call pancis, it indicates UnimplementedExamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExamService_ServiceDesc, srv)
}

func _ExamService_GetExamResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExamResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).GetExamResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExamService_GetExamResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).GetExamResult(ctx, req.(*GetExamResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExamService_ServiceDesc is the grpc.ServiceDesc for ExamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exam.ExamService",
	HandlerType: (*ExamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExamResult",
			Handler:    _ExamService_GetExamResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exam.proto",
}
