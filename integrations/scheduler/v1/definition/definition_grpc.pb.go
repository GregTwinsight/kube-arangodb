// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: integrations/scheduler/v1/definition/definition.proto

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

// SchedulerV1Client is the client API for SchedulerV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerV1Client interface {
	// Creates BatchJob from specification
	CreateBatchJob(ctx context.Context, in *CreateBatchJobRequest, opts ...grpc.CallOption) (*CreateBatchJobResponse, error)
	// Returns BatchJob. If job does not exists, Exists flag is set to false
	GetBatchJob(ctx context.Context, in *GetBatchJobRequest, opts ...grpc.CallOption) (*GetBatchJobResponse, error)
	// Returns list of the BatchJobs
	ListBatchJob(ctx context.Context, in *ListBatchJobRequest, opts ...grpc.CallOption) (*ListBatchJobResponse, error)
	// Deletes BatchJob. If job does not exists, Exists flag is set to false
	DeleteBatchJob(ctx context.Context, in *DeleteBatchJobRequest, opts ...grpc.CallOption) (*DeleteBatchJobResponse, error)
}

type schedulerV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerV1Client(cc grpc.ClientConnInterface) SchedulerV1Client {
	return &schedulerV1Client{cc}
}

func (c *schedulerV1Client) CreateBatchJob(ctx context.Context, in *CreateBatchJobRequest, opts ...grpc.CallOption) (*CreateBatchJobResponse, error) {
	out := new(CreateBatchJobResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV1/CreateBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV1Client) GetBatchJob(ctx context.Context, in *GetBatchJobRequest, opts ...grpc.CallOption) (*GetBatchJobResponse, error) {
	out := new(GetBatchJobResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV1/GetBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV1Client) ListBatchJob(ctx context.Context, in *ListBatchJobRequest, opts ...grpc.CallOption) (*ListBatchJobResponse, error) {
	out := new(ListBatchJobResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV1/ListBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV1Client) DeleteBatchJob(ctx context.Context, in *DeleteBatchJobRequest, opts ...grpc.CallOption) (*DeleteBatchJobResponse, error) {
	out := new(DeleteBatchJobResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV1/DeleteBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerV1Server is the server API for SchedulerV1 service.
// All implementations must embed UnimplementedSchedulerV1Server
// for forward compatibility
type SchedulerV1Server interface {
	// Creates BatchJob from specification
	CreateBatchJob(context.Context, *CreateBatchJobRequest) (*CreateBatchJobResponse, error)
	// Returns BatchJob. If job does not exists, Exists flag is set to false
	GetBatchJob(context.Context, *GetBatchJobRequest) (*GetBatchJobResponse, error)
	// Returns list of the BatchJobs
	ListBatchJob(context.Context, *ListBatchJobRequest) (*ListBatchJobResponse, error)
	// Deletes BatchJob. If job does not exists, Exists flag is set to false
	DeleteBatchJob(context.Context, *DeleteBatchJobRequest) (*DeleteBatchJobResponse, error)
	mustEmbedUnimplementedSchedulerV1Server()
}

// UnimplementedSchedulerV1Server must be embedded to have forward compatible implementations.
type UnimplementedSchedulerV1Server struct {
}

func (UnimplementedSchedulerV1Server) CreateBatchJob(context.Context, *CreateBatchJobRequest) (*CreateBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBatchJob not implemented")
}
func (UnimplementedSchedulerV1Server) GetBatchJob(context.Context, *GetBatchJobRequest) (*GetBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchJob not implemented")
}
func (UnimplementedSchedulerV1Server) ListBatchJob(context.Context, *ListBatchJobRequest) (*ListBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatchJob not implemented")
}
func (UnimplementedSchedulerV1Server) DeleteBatchJob(context.Context, *DeleteBatchJobRequest) (*DeleteBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBatchJob not implemented")
}
func (UnimplementedSchedulerV1Server) mustEmbedUnimplementedSchedulerV1Server() {}

// UnsafeSchedulerV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerV1Server will
// result in compilation errors.
type UnsafeSchedulerV1Server interface {
	mustEmbedUnimplementedSchedulerV1Server()
}

func RegisterSchedulerV1Server(s grpc.ServiceRegistrar, srv SchedulerV1Server) {
	s.RegisterService(&SchedulerV1_ServiceDesc, srv)
}

func _SchedulerV1_CreateBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV1Server).CreateBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV1/CreateBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV1Server).CreateBatchJob(ctx, req.(*CreateBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV1_GetBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV1Server).GetBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV1/GetBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV1Server).GetBatchJob(ctx, req.(*GetBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV1_ListBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV1Server).ListBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV1/ListBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV1Server).ListBatchJob(ctx, req.(*ListBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV1_DeleteBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV1Server).DeleteBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV1/DeleteBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV1Server).DeleteBatchJob(ctx, req.(*DeleteBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulerV1_ServiceDesc is the grpc.ServiceDesc for SchedulerV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulerV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.SchedulerV1",
	HandlerType: (*SchedulerV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBatchJob",
			Handler:    _SchedulerV1_CreateBatchJob_Handler,
		},
		{
			MethodName: "GetBatchJob",
			Handler:    _SchedulerV1_GetBatchJob_Handler,
		},
		{
			MethodName: "ListBatchJob",
			Handler:    _SchedulerV1_ListBatchJob_Handler,
		},
		{
			MethodName: "DeleteBatchJob",
			Handler:    _SchedulerV1_DeleteBatchJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrations/scheduler/v1/definition/definition.proto",
}
