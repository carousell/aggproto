// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package complex_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComplexUsecaseServiceClient is the client API for ComplexUsecaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplexUsecaseServiceClient interface {
	InvokeComplexUsecase(ctx context.Context, in *ComplexUsecaseRequest, opts ...grpc.CallOption) (*ComplexUsecaseResponse, error)
}

type complexUsecaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplexUsecaseServiceClient(cc grpc.ClientConnInterface) ComplexUsecaseServiceClient {
	return &complexUsecaseServiceClient{cc}
}

func (c *complexUsecaseServiceClient) InvokeComplexUsecase(ctx context.Context, in *ComplexUsecaseRequest, opts ...grpc.CallOption) (*ComplexUsecaseResponse, error) {
	out := new(ComplexUsecaseResponse)
	err := c.cc.Invoke(ctx, "/complex_v1.ComplexUsecaseService/InvokeComplexUsecase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplexUsecaseServiceServer is the server API for ComplexUsecaseService service.
// All implementations must embed UnimplementedComplexUsecaseServiceServer
// for forward compatibility
type ComplexUsecaseServiceServer interface {
	InvokeComplexUsecase(context.Context, *ComplexUsecaseRequest) (*ComplexUsecaseResponse, error)
	mustEmbedUnimplementedComplexUsecaseServiceServer()
}

// UnimplementedComplexUsecaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComplexUsecaseServiceServer struct {
}

func (UnimplementedComplexUsecaseServiceServer) InvokeComplexUsecase(context.Context, *ComplexUsecaseRequest) (*ComplexUsecaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeComplexUsecase not implemented")
}
func (UnimplementedComplexUsecaseServiceServer) mustEmbedUnimplementedComplexUsecaseServiceServer() {}

// UnsafeComplexUsecaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplexUsecaseServiceServer will
// result in compilation errors.
type UnsafeComplexUsecaseServiceServer interface {
	mustEmbedUnimplementedComplexUsecaseServiceServer()
}

func RegisterComplexUsecaseServiceServer(s grpc.ServiceRegistrar, srv ComplexUsecaseServiceServer) {
	s.RegisterService(&ComplexUsecaseService_ServiceDesc, srv)
}

func _ComplexUsecaseService_InvokeComplexUsecase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplexUsecaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplexUsecaseServiceServer).InvokeComplexUsecase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/complex_v1.ComplexUsecaseService/InvokeComplexUsecase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplexUsecaseServiceServer).InvokeComplexUsecase(ctx, req.(*ComplexUsecaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplexUsecaseService_ServiceDesc is the grpc.ServiceDesc for ComplexUsecaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplexUsecaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "complex_v1.ComplexUsecaseService",
	HandlerType: (*ComplexUsecaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvokeComplexUsecase",
			Handler:    _ComplexUsecaseService_InvokeComplexUsecase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "complex_v1/complex_usecase.proto",
}