// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/transactions/myserver.proto

package grpc_transactions

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

// TransactionsGRPCClient is the client API for TransactionsGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionsGRPCClient interface {
	CreateTransaction(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	VerifyTransaction(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyReplyGet, error)
}

type transactionsGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsGRPCClient(cc grpc.ClientConnInterface) TransactionsGRPCClient {
	return &transactionsGRPCClient{cc}
}

func (c *transactionsGRPCClient) CreateTransaction(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/TransactionsGRPC/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsGRPCClient) VerifyTransaction(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyReplyGet, error) {
	out := new(VerifyReplyGet)
	err := c.cc.Invoke(ctx, "/TransactionsGRPC/VerifyTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsGRPCServer is the server API for TransactionsGRPC service.
// All implementations must embed UnimplementedTransactionsGRPCServer
// for forward compatibility
type TransactionsGRPCServer interface {
	CreateTransaction(context.Context, *CreateRequest) (*CreateReply, error)
	VerifyTransaction(context.Context, *VerifyRequest) (*VerifyReplyGet, error)
	mustEmbedUnimplementedTransactionsGRPCServer()
}

// UnimplementedTransactionsGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionsGRPCServer struct {
}

func (UnimplementedTransactionsGRPCServer) CreateTransaction(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionsGRPCServer) VerifyTransaction(context.Context, *VerifyRequest) (*VerifyReplyGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTransaction not implemented")
}
func (UnimplementedTransactionsGRPCServer) mustEmbedUnimplementedTransactionsGRPCServer() {}

// UnsafeTransactionsGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsGRPCServer will
// result in compilation errors.
type UnsafeTransactionsGRPCServer interface {
	mustEmbedUnimplementedTransactionsGRPCServer()
}

func RegisterTransactionsGRPCServer(s grpc.ServiceRegistrar, srv TransactionsGRPCServer) {
	s.RegisterService(&TransactionsGRPC_ServiceDesc, srv)
}

func _TransactionsGRPC_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsGRPCServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionsGRPC/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsGRPCServer).CreateTransaction(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsGRPC_VerifyTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsGRPCServer).VerifyTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionsGRPC/VerifyTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsGRPCServer).VerifyTransaction(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionsGRPC_ServiceDesc is the grpc.ServiceDesc for TransactionsGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionsGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionsGRPC",
	HandlerType: (*TransactionsGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionsGRPC_CreateTransaction_Handler,
		},
		{
			MethodName: "VerifyTransaction",
			Handler:    _TransactionsGRPC_VerifyTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/transactions/myserver.proto",
}