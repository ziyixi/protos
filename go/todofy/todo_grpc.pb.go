// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/todofy/todo.proto

package todofy

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
	TodoService_PopulateTodo_FullMethodName = "/todofy.TodoService/PopulateTodo"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	// Populate the todo using the specified method.
	PopulateTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) PopulateTodo(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_PopulateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	// Populate the todo using the specified method.
	PopulateTodo(context.Context, *TodoRequest) (*TodoResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) PopulateTodo(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopulateTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_PopulateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).PopulateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_PopulateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).PopulateTodo(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todofy.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PopulateTodo",
			Handler:    _TodoService_PopulateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todofy/todo.proto",
}
