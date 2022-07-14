// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: ydb_scheme_v1.proto

package Ydb_Scheme_V1

import (
	context "context"
	Ydb_Scheme "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scheme"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SchemeServiceClient is the client API for SchemeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemeServiceClient interface {
	// Make Directory.
	MakeDirectory(ctx context.Context, in *Ydb_Scheme.MakeDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.MakeDirectoryResponse, error)
	// Remove Directory.
	RemoveDirectory(ctx context.Context, in *Ydb_Scheme.RemoveDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.RemoveDirectoryResponse, error)
	// Returns information about given directory and objects inside it.
	ListDirectory(ctx context.Context, in *Ydb_Scheme.ListDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.ListDirectoryResponse, error)
	// Returns information about object with given path.
	DescribePath(ctx context.Context, in *Ydb_Scheme.DescribePathRequest, opts ...grpc.CallOption) (*Ydb_Scheme.DescribePathResponse, error)
	// Modify permissions.
	ModifyPermissions(ctx context.Context, in *Ydb_Scheme.ModifyPermissionsRequest, opts ...grpc.CallOption) (*Ydb_Scheme.ModifyPermissionsResponse, error)
}

type schemeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemeServiceClient(cc grpc.ClientConnInterface) SchemeServiceClient {
	return &schemeServiceClient{cc}
}

func (c *schemeServiceClient) MakeDirectory(ctx context.Context, in *Ydb_Scheme.MakeDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.MakeDirectoryResponse, error) {
	out := new(Ydb_Scheme.MakeDirectoryResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Scheme.V1.SchemeService/MakeDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemeServiceClient) RemoveDirectory(ctx context.Context, in *Ydb_Scheme.RemoveDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.RemoveDirectoryResponse, error) {
	out := new(Ydb_Scheme.RemoveDirectoryResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Scheme.V1.SchemeService/RemoveDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemeServiceClient) ListDirectory(ctx context.Context, in *Ydb_Scheme.ListDirectoryRequest, opts ...grpc.CallOption) (*Ydb_Scheme.ListDirectoryResponse, error) {
	out := new(Ydb_Scheme.ListDirectoryResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Scheme.V1.SchemeService/ListDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemeServiceClient) DescribePath(ctx context.Context, in *Ydb_Scheme.DescribePathRequest, opts ...grpc.CallOption) (*Ydb_Scheme.DescribePathResponse, error) {
	out := new(Ydb_Scheme.DescribePathResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Scheme.V1.SchemeService/DescribePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemeServiceClient) ModifyPermissions(ctx context.Context, in *Ydb_Scheme.ModifyPermissionsRequest, opts ...grpc.CallOption) (*Ydb_Scheme.ModifyPermissionsResponse, error) {
	out := new(Ydb_Scheme.ModifyPermissionsResponse)
	err := c.cc.Invoke(ctx, "/Ydb.Scheme.V1.SchemeService/ModifyPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemeServiceServer is the server API for SchemeService service.
// All implementations must embed UnimplementedSchemeServiceServer
// for forward compatibility
type SchemeServiceServer interface {
	// Make Directory.
	MakeDirectory(context.Context, *Ydb_Scheme.MakeDirectoryRequest) (*Ydb_Scheme.MakeDirectoryResponse, error)
	// Remove Directory.
	RemoveDirectory(context.Context, *Ydb_Scheme.RemoveDirectoryRequest) (*Ydb_Scheme.RemoveDirectoryResponse, error)
	// Returns information about given directory and objects inside it.
	ListDirectory(context.Context, *Ydb_Scheme.ListDirectoryRequest) (*Ydb_Scheme.ListDirectoryResponse, error)
	// Returns information about object with given path.
	DescribePath(context.Context, *Ydb_Scheme.DescribePathRequest) (*Ydb_Scheme.DescribePathResponse, error)
	// Modify permissions.
	ModifyPermissions(context.Context, *Ydb_Scheme.ModifyPermissionsRequest) (*Ydb_Scheme.ModifyPermissionsResponse, error)
	mustEmbedUnimplementedSchemeServiceServer()
}

// UnimplementedSchemeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemeServiceServer struct {
}

func (UnimplementedSchemeServiceServer) MakeDirectory(context.Context, *Ydb_Scheme.MakeDirectoryRequest) (*Ydb_Scheme.MakeDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeDirectory not implemented")
}
func (UnimplementedSchemeServiceServer) RemoveDirectory(context.Context, *Ydb_Scheme.RemoveDirectoryRequest) (*Ydb_Scheme.RemoveDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDirectory not implemented")
}
func (UnimplementedSchemeServiceServer) ListDirectory(context.Context, *Ydb_Scheme.ListDirectoryRequest) (*Ydb_Scheme.ListDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDirectory not implemented")
}
func (UnimplementedSchemeServiceServer) DescribePath(context.Context, *Ydb_Scheme.DescribePathRequest) (*Ydb_Scheme.DescribePathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePath not implemented")
}
func (UnimplementedSchemeServiceServer) ModifyPermissions(context.Context, *Ydb_Scheme.ModifyPermissionsRequest) (*Ydb_Scheme.ModifyPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPermissions not implemented")
}
func (UnimplementedSchemeServiceServer) mustEmbedUnimplementedSchemeServiceServer() {}

// UnsafeSchemeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemeServiceServer will
// result in compilation errors.
type UnsafeSchemeServiceServer interface {
	mustEmbedUnimplementedSchemeServiceServer()
}

func RegisterSchemeServiceServer(s grpc.ServiceRegistrar, srv SchemeServiceServer) {
	s.RegisterService(&SchemeService_ServiceDesc, srv)
}

func _SchemeService_MakeDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scheme.MakeDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeServiceServer).MakeDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Scheme.V1.SchemeService/MakeDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeServiceServer).MakeDirectory(ctx, req.(*Ydb_Scheme.MakeDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemeService_RemoveDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scheme.RemoveDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeServiceServer).RemoveDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Scheme.V1.SchemeService/RemoveDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeServiceServer).RemoveDirectory(ctx, req.(*Ydb_Scheme.RemoveDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemeService_ListDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scheme.ListDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeServiceServer).ListDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Scheme.V1.SchemeService/ListDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeServiceServer).ListDirectory(ctx, req.(*Ydb_Scheme.ListDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemeService_DescribePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scheme.DescribePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeServiceServer).DescribePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Scheme.V1.SchemeService/DescribePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeServiceServer).DescribePath(ctx, req.(*Ydb_Scheme.DescribePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemeService_ModifyPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Scheme.ModifyPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeServiceServer).ModifyPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ydb.Scheme.V1.SchemeService/ModifyPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeServiceServer).ModifyPermissions(ctx, req.(*Ydb_Scheme.ModifyPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemeService_ServiceDesc is the grpc.ServiceDesc for SchemeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Scheme.V1.SchemeService",
	HandlerType: (*SchemeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeDirectory",
			Handler:    _SchemeService_MakeDirectory_Handler,
		},
		{
			MethodName: "RemoveDirectory",
			Handler:    _SchemeService_RemoveDirectory_Handler,
		},
		{
			MethodName: "ListDirectory",
			Handler:    _SchemeService_ListDirectory_Handler,
		},
		{
			MethodName: "DescribePath",
			Handler:    _SchemeService_DescribePath_Handler,
		},
		{
			MethodName: "ModifyPermissions",
			Handler:    _SchemeService_ModifyPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_scheme_v1.proto",
}
