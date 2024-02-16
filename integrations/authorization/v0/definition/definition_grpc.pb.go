// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: integrations/authorization/v0/definition/definition.proto

package definition

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

// AuthorizationV0Client is the client API for AuthorizationV0 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationV0Client interface {
	// Can validate if specified request is allowed
	Can(ctx context.Context, in *CanRequest, opts ...grpc.CallOption) (*CanResponse, error)
}

type authorizationV0Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationV0Client(cc grpc.ClientConnInterface) AuthorizationV0Client {
	return &authorizationV0Client{cc}
}

func (c *authorizationV0Client) Can(ctx context.Context, in *CanRequest, opts ...grpc.CallOption) (*CanResponse, error) {
	out := new(CanResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationV0/Can", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationV0Server is the server API for AuthorizationV0 service.
// All implementations must embed UnimplementedAuthorizationV0Server
// for forward compatibility
type AuthorizationV0Server interface {
	// Can validate if specified request is allowed
	Can(context.Context, *CanRequest) (*CanResponse, error)
	mustEmbedUnimplementedAuthorizationV0Server()
}

// UnimplementedAuthorizationV0Server must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationV0Server struct {
}

func (UnimplementedAuthorizationV0Server) Can(context.Context, *CanRequest) (*CanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Can not implemented")
}
func (UnimplementedAuthorizationV0Server) mustEmbedUnimplementedAuthorizationV0Server() {}

// UnsafeAuthorizationV0Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationV0Server will
// result in compilation errors.
type UnsafeAuthorizationV0Server interface {
	mustEmbedUnimplementedAuthorizationV0Server()
}

func RegisterAuthorizationV0Server(s grpc.ServiceRegistrar, srv AuthorizationV0Server) {
	s.RegisterService(&AuthorizationV0_ServiceDesc, srv)
}

func _AuthorizationV0_Can_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationV0Server).Can(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationV0/Can",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationV0Server).Can(ctx, req.(*CanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationV0_ServiceDesc is the grpc.ServiceDesc for AuthorizationV0 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationV0_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.AuthorizationV0",
	HandlerType: (*AuthorizationV0Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Can",
			Handler:    _AuthorizationV0_Can_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrations/authorization/v0/definition/definition.proto",
}
