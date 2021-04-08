// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wallet

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

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	GetUserWallet(ctx context.Context, in *GetUserWalletRequest, opts ...grpc.CallOption) (*GetUserWalletResponse, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetUserWallet(ctx context.Context, in *GetUserWalletRequest, opts ...grpc.CallOption) (*GetUserWalletResponse, error) {
	out := new(GetUserWalletResponse)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetUserWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	GetUserWallet(context.Context, *GetUserWalletRequest) (*GetUserWalletResponse, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) GetUserWallet(context.Context, *GetUserWalletRequest) (*GetUserWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWallet not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_GetUserWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetUserWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetUserWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetUserWallet(ctx, req.(*GetUserWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserWallet",
			Handler:    _Wallet_GetUserWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/wallet.proto",
}
