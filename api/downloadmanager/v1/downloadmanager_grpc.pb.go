// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/downloadmanager/v1/downloadmanager.proto

package v1

import (
	context "context"
	common "github.com/staroffish/am/api/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DownloadmanagerClient is the client API for Downloadmanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadmanagerClient interface {
	ScanTaskAndDownload(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ScanTaskAndDownloadReply, error)
	ScanTask(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ScanTaskReply, error)
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
	ListTask(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ListTaskReply, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskReply, error)
}

type downloadmanagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadmanagerClient(cc grpc.ClientConnInterface) DownloadmanagerClient {
	return &downloadmanagerClient{cc}
}

func (c *downloadmanagerClient) ScanTaskAndDownload(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ScanTaskAndDownloadReply, error) {
	out := new(ScanTaskAndDownloadReply)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/ScanTaskAndDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) ScanTask(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ScanTaskReply, error) {
	out := new(ScanTaskReply)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/ScanTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) ListTask(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*ListTaskReply, error) {
	out := new(ListTaskReply)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/ListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloadmanagerClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskReply, error) {
	out := new(GetTaskReply)
	err := c.cc.Invoke(ctx, "/api.downloadmanager.v1.Downloadmanager/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadmanagerServer is the server API for Downloadmanager service.
// All implementations must embed UnimplementedDownloadmanagerServer
// for forward compatibility
type DownloadmanagerServer interface {
	ScanTaskAndDownload(context.Context, *common.Empty) (*ScanTaskAndDownloadReply, error)
	ScanTask(context.Context, *common.Empty) (*ScanTaskReply, error)
	AddTask(context.Context, *AddTaskRequest) (*common.Empty, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*common.Empty, error)
	ListTask(context.Context, *common.Empty) (*ListTaskReply, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error)
	mustEmbedUnimplementedDownloadmanagerServer()
}

// UnimplementedDownloadmanagerServer must be embedded to have forward compatible implementations.
type UnimplementedDownloadmanagerServer struct {
}

func (UnimplementedDownloadmanagerServer) ScanTaskAndDownload(context.Context, *common.Empty) (*ScanTaskAndDownloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanTaskAndDownload not implemented")
}
func (UnimplementedDownloadmanagerServer) ScanTask(context.Context, *common.Empty) (*ScanTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanTask not implemented")
}
func (UnimplementedDownloadmanagerServer) AddTask(context.Context, *AddTaskRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedDownloadmanagerServer) UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedDownloadmanagerServer) DeleteTask(context.Context, *DeleteTaskRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedDownloadmanagerServer) ListTask(context.Context, *common.Empty) (*ListTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTask not implemented")
}
func (UnimplementedDownloadmanagerServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedDownloadmanagerServer) mustEmbedUnimplementedDownloadmanagerServer() {}

// UnsafeDownloadmanagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadmanagerServer will
// result in compilation errors.
type UnsafeDownloadmanagerServer interface {
	mustEmbedUnimplementedDownloadmanagerServer()
}

func RegisterDownloadmanagerServer(s grpc.ServiceRegistrar, srv DownloadmanagerServer) {
	s.RegisterService(&Downloadmanager_ServiceDesc, srv)
}

func _Downloadmanager_ScanTaskAndDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).ScanTaskAndDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/ScanTaskAndDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).ScanTaskAndDownload(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_ScanTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).ScanTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/ScanTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).ScanTask(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).AddTask(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_ListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).ListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/ListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).ListTask(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloadmanager_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadmanagerServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.downloadmanager.v1.Downloadmanager/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadmanagerServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Downloadmanager_ServiceDesc is the grpc.ServiceDesc for Downloadmanager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Downloadmanager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.downloadmanager.v1.Downloadmanager",
	HandlerType: (*DownloadmanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanTaskAndDownload",
			Handler:    _Downloadmanager_ScanTaskAndDownload_Handler,
		},
		{
			MethodName: "ScanTask",
			Handler:    _Downloadmanager_ScanTask_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _Downloadmanager_AddTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Downloadmanager_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Downloadmanager_DeleteTask_Handler,
		},
		{
			MethodName: "ListTask",
			Handler:    _Downloadmanager_ListTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Downloadmanager_GetTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/downloadmanager/v1/downloadmanager.proto",
}
