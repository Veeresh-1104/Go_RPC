// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: dummy.proto

package proto

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

// AddTwoNumbrsClient is the client API for AddTwoNumbrs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddTwoNumbrsClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type addTwoNumbrsClient struct {
	cc grpc.ClientConnInterface
}

func NewAddTwoNumbrsClient(cc grpc.ClientConnInterface) AddTwoNumbrsClient {
	return &addTwoNumbrsClient{cc}
}

func (c *addTwoNumbrsClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/greet.AddTwoNumbrs/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddTwoNumbrsServer is the server API for AddTwoNumbrs service.
// All implementations must embed UnimplementedAddTwoNumbrsServer
// for forward compatibility
type AddTwoNumbrsServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedAddTwoNumbrsServer()
}

// UnimplementedAddTwoNumbrsServer must be embedded to have forward compatible implementations.
type UnimplementedAddTwoNumbrsServer struct {
}

func (UnimplementedAddTwoNumbrsServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAddTwoNumbrsServer) mustEmbedUnimplementedAddTwoNumbrsServer() {}

// UnsafeAddTwoNumbrsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddTwoNumbrsServer will
// result in compilation errors.
type UnsafeAddTwoNumbrsServer interface {
	mustEmbedUnimplementedAddTwoNumbrsServer()
}

func RegisterAddTwoNumbrsServer(s grpc.ServiceRegistrar, srv AddTwoNumbrsServer) {
	s.RegisterService(&AddTwoNumbrs_ServiceDesc, srv)
}

func _AddTwoNumbrs_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddTwoNumbrsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.AddTwoNumbrs/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddTwoNumbrsServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AddTwoNumbrs_ServiceDesc is the grpc.ServiceDesc for AddTwoNumbrs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddTwoNumbrs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.AddTwoNumbrs",
	HandlerType: (*AddTwoNumbrsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddTwoNumbrs_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dummy.proto",
}