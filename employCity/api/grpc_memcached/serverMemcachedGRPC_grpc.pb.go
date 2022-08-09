// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/grpc_memcached/serverMemcachedGRPC.proto

package grpc_memcached

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

// MyGRPCClient is the client API for MyGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyGRPCClient interface {
	Get(ctx context.Context, in *DeleteGetRequest, opts ...grpc.CallOption) (*ReplyGet, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *DeleteGetRequest, opts ...grpc.CallOption) (*Reply, error)
}

type myGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMyGRPCClient(cc grpc.ClientConnInterface) MyGRPCClient {
	return &myGRPCClient{cc}
}

func (c *myGRPCClient) Get(ctx context.Context, in *DeleteGetRequest, opts ...grpc.CallOption) (*ReplyGet, error) {
	out := new(ReplyGet)
	err := c.cc.Invoke(ctx, "/MyGRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myGRPCClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/MyGRPC/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myGRPCClient) Delete(ctx context.Context, in *DeleteGetRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/MyGRPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyGRPCServer is the server API for MyGRPC service.
// All implementations must embed UnimplementedMyGRPCServer
// for forward compatibility
type MyGRPCServer interface {
	Get(context.Context, *DeleteGetRequest) (*ReplyGet, error)
	Set(context.Context, *SetRequest) (*Reply, error)
	Delete(context.Context, *DeleteGetRequest) (*Reply, error)
	mustEmbedUnimplementedMyGRPCServer()
}

// UnimplementedMyGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMyGRPCServer struct {
}

func (UnimplementedMyGRPCServer) Get(context.Context, *DeleteGetRequest) (*ReplyGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMyGRPCServer) Set(context.Context, *SetRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedMyGRPCServer) Delete(context.Context, *DeleteGetRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMyGRPCServer) mustEmbedUnimplementedMyGRPCServer() {}

// UnsafeMyGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyGRPCServer will
// result in compilation errors.
type UnsafeMyGRPCServer interface {
	mustEmbedUnimplementedMyGRPCServer()
}

func RegisterMyGRPCServer(s grpc.ServiceRegistrar, srv MyGRPCServer) {
	s.RegisterService(&MyGRPC_ServiceDesc, srv)
}

func _MyGRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MyGRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGRPCServer).Get(ctx, req.(*DeleteGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyGRPC_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGRPCServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MyGRPC/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGRPCServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyGRPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGRPCServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MyGRPC/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGRPCServer).Delete(ctx, req.(*DeleteGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyGRPC_ServiceDesc is the grpc.ServiceDesc for MyGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MyGRPC",
	HandlerType: (*MyGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MyGRPC_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _MyGRPC_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MyGRPC_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc_memcached/serverMemcachedGRPC.proto",
}
