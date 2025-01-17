// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: public/api/public/v1/public.proto

package v1

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

const (
	Public_VerifyUser_FullMethodName = "/v2.Public/VerifyUser"
	Public_GetUser_FullMethodName    = "/v2.Public/GetUser"
)

// PublicClient is the client API for Public service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicClient interface {
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type publicClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicClient(cc grpc.ClientConnInterface) PublicClient {
	return &publicClient{cc}
}

func (c *publicClient) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, Public_VerifyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Public_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServer is the server API for Public service.
// All implementations must embed UnimplementedPublicServer
// for forward compatibility
type PublicServer interface {
	VerifyUser(context.Context, *VerifyUserRequest) (*VerifyResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	mustEmbedUnimplementedPublicServer()
}

// UnimplementedPublicServer must be embedded to have forward compatible implementations.
type UnimplementedPublicServer struct {
}

func (UnimplementedPublicServer) VerifyUser(context.Context, *VerifyUserRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedPublicServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedPublicServer) mustEmbedUnimplementedPublicServer() {}

// UnsafePublicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicServer will
// result in compilation errors.
type UnsafePublicServer interface {
	mustEmbedUnimplementedPublicServer()
}

func RegisterPublicServer(s grpc.ServiceRegistrar, srv PublicServer) {
	s.RegisterService(&Public_ServiceDesc, srv)
}

func _Public_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Public_VerifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).VerifyUser(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Public_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Public_ServiceDesc is the grpc.ServiceDesc for Public service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Public_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.Public",
	HandlerType: (*PublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyUser",
			Handler:    _Public_VerifyUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Public_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/api/public/v1/public.proto",
}
