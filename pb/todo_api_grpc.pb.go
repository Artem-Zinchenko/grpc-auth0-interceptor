// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: azinchenko/todo/v1/todo_api.proto

package pb

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

// TaskApiClient is the client API for TaskApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskApiClient interface {
	Add(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	MarkDone(ctx context.Context, in *MarkDoneRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type taskApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskApiClient(cc grpc.ClientConnInterface) TaskApiClient {
	return &taskApiClient{cc}
}

func (c *taskApiClient) Add(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/azinchenko.auth.TaskApi/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskApiClient) Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/azinchenko.auth.TaskApi/Archive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskApiClient) MarkDone(ctx context.Context, in *MarkDoneRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/azinchenko.auth.TaskApi/MarkDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskApiServer is the server API for TaskApi service.
// All implementations must embed UnimplementedTaskApiServer
// for forward compatibility
type TaskApiServer interface {
	Add(context.Context, *AddTaskRequest) (*TaskResponse, error)
	Archive(context.Context, *ArchiveRequest) (*TaskResponse, error)
	MarkDone(context.Context, *MarkDoneRequest) (*TaskResponse, error)
	mustEmbedUnimplementedTaskApiServer()
}

// UnimplementedTaskApiServer must be embedded to have forward compatible implementations.
type UnimplementedTaskApiServer struct {
}

func (UnimplementedTaskApiServer) Add(context.Context, *AddTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTaskApiServer) Archive(context.Context, *ArchiveRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Archive not implemented")
}
func (UnimplementedTaskApiServer) MarkDone(context.Context, *MarkDoneRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkDone not implemented")
}
func (UnimplementedTaskApiServer) mustEmbedUnimplementedTaskApiServer() {}

// UnsafeTaskApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskApiServer will
// result in compilation errors.
type UnsafeTaskApiServer interface {
	mustEmbedUnimplementedTaskApiServer()
}

func RegisterTaskApiServer(s grpc.ServiceRegistrar, srv TaskApiServer) {
	s.RegisterService(&TaskApi_ServiceDesc, srv)
}

func _TaskApi_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskApiServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azinchenko.auth.TaskApi/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskApiServer).Add(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskApi_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskApiServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azinchenko.auth.TaskApi/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskApiServer).Archive(ctx, req.(*ArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskApi_MarkDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskApiServer).MarkDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azinchenko.auth.TaskApi/MarkDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskApiServer).MarkDone(ctx, req.(*MarkDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskApi_ServiceDesc is the grpc.ServiceDesc for TaskApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "azinchenko.auth.TaskApi",
	HandlerType: (*TaskApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TaskApi_Add_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _TaskApi_Archive_Handler,
		},
		{
			MethodName: "MarkDone",
			Handler:    _TaskApi_MarkDone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "azinchenko/todo/v1/todo_api.proto",
}
