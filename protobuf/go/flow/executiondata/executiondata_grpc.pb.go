// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: flow/executiondata/executiondata.proto

package access

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

// ExecutionDataAPIClient is the client API for ExecutionDataAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionDataAPIClient interface {
	GetExecutionDataByBlockID(ctx context.Context, in *GetExecutionDataByBlockIDRequest, opts ...grpc.CallOption) (*GetExecutionDataByBlockIDResponse, error)
}

type executionDataAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionDataAPIClient(cc grpc.ClientConnInterface) ExecutionDataAPIClient {
	return &executionDataAPIClient{cc}
}

func (c *executionDataAPIClient) GetExecutionDataByBlockID(ctx context.Context, in *GetExecutionDataByBlockIDRequest, opts ...grpc.CallOption) (*GetExecutionDataByBlockIDResponse, error) {
	out := new(GetExecutionDataByBlockIDResponse)
	err := c.cc.Invoke(ctx, "/flow.access.ExecutionDataAPI/GetExecutionDataByBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionDataAPIServer is the server API for ExecutionDataAPI service.
// All implementations should embed UnimplementedExecutionDataAPIServer
// for forward compatibility
type ExecutionDataAPIServer interface {
	GetExecutionDataByBlockID(context.Context, *GetExecutionDataByBlockIDRequest) (*GetExecutionDataByBlockIDResponse, error)
}

// UnimplementedExecutionDataAPIServer should be embedded to have forward compatible implementations.
type UnimplementedExecutionDataAPIServer struct {
}

func (UnimplementedExecutionDataAPIServer) GetExecutionDataByBlockID(context.Context, *GetExecutionDataByBlockIDRequest) (*GetExecutionDataByBlockIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecutionDataByBlockID not implemented")
}

// UnsafeExecutionDataAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionDataAPIServer will
// result in compilation errors.
type UnsafeExecutionDataAPIServer interface {
	mustEmbedUnimplementedExecutionDataAPIServer()
}

func RegisterExecutionDataAPIServer(s grpc.ServiceRegistrar, srv ExecutionDataAPIServer) {
	s.RegisterService(&ExecutionDataAPI_ServiceDesc, srv)
}

func _ExecutionDataAPI_GetExecutionDataByBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExecutionDataByBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionDataAPIServer).GetExecutionDataByBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.ExecutionDataAPI/GetExecutionDataByBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionDataAPIServer).GetExecutionDataByBlockID(ctx, req.(*GetExecutionDataByBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecutionDataAPI_ServiceDesc is the grpc.ServiceDesc for ExecutionDataAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutionDataAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.access.ExecutionDataAPI",
	HandlerType: (*ExecutionDataAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExecutionDataByBlockID",
			Handler:    _ExecutionDataAPI_GetExecutionDataByBlockID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flow/executiondata/executiondata.proto",
}
