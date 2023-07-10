// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: ydb_operation_v1.proto

package Ydb_Operation_V1

import (
	context "context"
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OperationService_GetOperation_FullMethodName    = "/Ydb.Operation.V1.OperationService/GetOperation"
	OperationService_CancelOperation_FullMethodName = "/Ydb.Operation.V1.OperationService/CancelOperation"
	OperationService_ForgetOperation_FullMethodName = "/Ydb.Operation.V1.OperationService/ForgetOperation"
	OperationService_ListOperations_FullMethodName  = "/Ydb.Operation.V1.OperationService/ListOperations"
)

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationServiceClient interface {
	// Check status for a given operation.
	GetOperation(ctx context.Context, in *Ydb_Operations.GetOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.GetOperationResponse, error)
	// Starts cancellation of a long-running operation,
	// Clients can use GetOperation to check whether the cancellation succeeded
	// or whether the operation completed despite cancellation.
	CancelOperation(ctx context.Context, in *Ydb_Operations.CancelOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.CancelOperationResponse, error)
	// Forgets long-running operation. It does not cancel the operation and returns
	// an error if operation was not completed.
	ForgetOperation(ctx context.Context, in *Ydb_Operations.ForgetOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.ForgetOperationResponse, error)
	// Lists operations that match the specified filter in the request.
	ListOperations(ctx context.Context, in *Ydb_Operations.ListOperationsRequest, opts ...grpc.CallOption) (*Ydb_Operations.ListOperationsResponse, error)
}

type operationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationServiceClient(cc grpc.ClientConnInterface) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) GetOperation(ctx context.Context, in *Ydb_Operations.GetOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.GetOperationResponse, error) {
	out := new(Ydb_Operations.GetOperationResponse)
	err := c.cc.Invoke(ctx, OperationService_GetOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) CancelOperation(ctx context.Context, in *Ydb_Operations.CancelOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.CancelOperationResponse, error) {
	out := new(Ydb_Operations.CancelOperationResponse)
	err := c.cc.Invoke(ctx, OperationService_CancelOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) ForgetOperation(ctx context.Context, in *Ydb_Operations.ForgetOperationRequest, opts ...grpc.CallOption) (*Ydb_Operations.ForgetOperationResponse, error) {
	out := new(Ydb_Operations.ForgetOperationResponse)
	err := c.cc.Invoke(ctx, OperationService_ForgetOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) ListOperations(ctx context.Context, in *Ydb_Operations.ListOperationsRequest, opts ...grpc.CallOption) (*Ydb_Operations.ListOperationsResponse, error) {
	out := new(Ydb_Operations.ListOperationsResponse)
	err := c.cc.Invoke(ctx, OperationService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServiceServer is the server API for OperationService service.
// All implementations must embed UnimplementedOperationServiceServer
// for forward compatibility
type OperationServiceServer interface {
	// Check status for a given operation.
	GetOperation(context.Context, *Ydb_Operations.GetOperationRequest) (*Ydb_Operations.GetOperationResponse, error)
	// Starts cancellation of a long-running operation,
	// Clients can use GetOperation to check whether the cancellation succeeded
	// or whether the operation completed despite cancellation.
	CancelOperation(context.Context, *Ydb_Operations.CancelOperationRequest) (*Ydb_Operations.CancelOperationResponse, error)
	// Forgets long-running operation. It does not cancel the operation and returns
	// an error if operation was not completed.
	ForgetOperation(context.Context, *Ydb_Operations.ForgetOperationRequest) (*Ydb_Operations.ForgetOperationResponse, error)
	// Lists operations that match the specified filter in the request.
	ListOperations(context.Context, *Ydb_Operations.ListOperationsRequest) (*Ydb_Operations.ListOperationsResponse, error)
	mustEmbedUnimplementedOperationServiceServer()
}

// UnimplementedOperationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (UnimplementedOperationServiceServer) GetOperation(context.Context, *Ydb_Operations.GetOperationRequest) (*Ydb_Operations.GetOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperation not implemented")
}
func (UnimplementedOperationServiceServer) CancelOperation(context.Context, *Ydb_Operations.CancelOperationRequest) (*Ydb_Operations.CancelOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOperation not implemented")
}
func (UnimplementedOperationServiceServer) ForgetOperation(context.Context, *Ydb_Operations.ForgetOperationRequest) (*Ydb_Operations.ForgetOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetOperation not implemented")
}
func (UnimplementedOperationServiceServer) ListOperations(context.Context, *Ydb_Operations.ListOperationsRequest) (*Ydb_Operations.ListOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedOperationServiceServer) mustEmbedUnimplementedOperationServiceServer() {}

// UnsafeOperationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServiceServer will
// result in compilation errors.
type UnsafeOperationServiceServer interface {
	mustEmbedUnimplementedOperationServiceServer()
}

func RegisterOperationServiceServer(s grpc.ServiceRegistrar, srv OperationServiceServer) {
	s.RegisterService(&OperationService_ServiceDesc, srv)
}

func _OperationService_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Operations.GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationService_GetOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).GetOperation(ctx, req.(*Ydb_Operations.GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Operations.CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationService_CancelOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).CancelOperation(ctx, req.(*Ydb_Operations.CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_ForgetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Operations.ForgetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).ForgetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationService_ForgetOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).ForgetOperation(ctx, req.(*Ydb_Operations.ForgetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Operations.ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).ListOperations(ctx, req.(*Ydb_Operations.ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationService_ServiceDesc is the grpc.ServiceDesc for OperationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Operation.V1.OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperation",
			Handler:    _OperationService_GetOperation_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _OperationService_CancelOperation_Handler,
		},
		{
			MethodName: "ForgetOperation",
			Handler:    _OperationService_ForgetOperation_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _OperationService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_operation_v1.proto",
}
