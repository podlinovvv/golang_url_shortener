// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_shortener_grpc

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

// ShortenerServiceClient is the client API for ShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerServiceClient interface {
	Create(ctx context.Context, in *FullUrl, opts ...grpc.CallOption) (*ShortUrl, error)
	Get(ctx context.Context, in *ShortUrl, opts ...grpc.CallOption) (*FullUrl, error)
}

type shortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerServiceClient(cc grpc.ClientConnInterface) ShortenerServiceClient {
	return &shortenerServiceClient{cc}
}

func (c *shortenerServiceClient) Create(ctx context.Context, in *FullUrl, opts ...grpc.CallOption) (*ShortUrl, error) {
	out := new(ShortUrl)
	err := c.cc.Invoke(ctx, "/shortener.ShortenerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerServiceClient) Get(ctx context.Context, in *ShortUrl, opts ...grpc.CallOption) (*FullUrl, error) {
	out := new(FullUrl)
	err := c.cc.Invoke(ctx, "/shortener.ShortenerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServiceServer is the server API for ShortenerService service.
// All implementations must embed UnimplementedShortenerServiceServer
// for forward compatibility
type ShortenerServiceServer interface {
	Create(context.Context, *FullUrl) (*ShortUrl, error)
	Get(context.Context, *ShortUrl) (*FullUrl, error)
	mustEmbedUnimplementedShortenerServiceServer()
}

// UnimplementedShortenerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServiceServer struct {
}

func (UnimplementedShortenerServiceServer) Create(context.Context, *FullUrl) (*ShortUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShortenerServiceServer) Get(context.Context, *ShortUrl) (*FullUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedShortenerServiceServer) mustEmbedUnimplementedShortenerServiceServer() {}

// UnsafeShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServiceServer will
// result in compilation errors.
type UnsafeShortenerServiceServer interface {
	mustEmbedUnimplementedShortenerServiceServer()
}

func RegisterShortenerServiceServer(s grpc.ServiceRegistrar, srv ShortenerServiceServer) {
	s.RegisterService(&ShortenerService_ServiceDesc, srv)
}

func _ShortenerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.ShortenerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServiceServer).Create(ctx, req.(*FullUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.ShortenerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServiceServer).Get(ctx, req.(*ShortUrl))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenerService_ServiceDesc is the grpc.ServiceDesc for ShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.ShortenerService",
	HandlerType: (*ShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShortenerService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ShortenerService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener.proto",
}
