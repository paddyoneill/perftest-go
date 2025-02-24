// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: pkg/perftest/perftest.proto

package perftest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Perftest_GetPort_FullMethodName       = "/perftest.Perftest/GetPort"
	Perftest_StartPerftest_FullMethodName = "/perftest.Perftest/StartPerftest"
)

// PerftestClient is the client API for Perftest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PerftestClient interface {
	GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error)
	StartPerftest(ctx context.Context, in *PerftestRequest, opts ...grpc.CallOption) (*PerftestResponse, error)
}

type perftestClient struct {
	cc grpc.ClientConnInterface
}

func NewPerftestClient(cc grpc.ClientConnInterface) PerftestClient {
	return &perftestClient{cc}
}

func (c *perftestClient) GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, Perftest_GetPort_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perftestClient) StartPerftest(ctx context.Context, in *PerftestRequest, opts ...grpc.CallOption) (*PerftestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PerftestResponse)
	err := c.cc.Invoke(ctx, Perftest_StartPerftest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerftestServer is the server API for Perftest service.
// All implementations must embed UnimplementedPerftestServer
// for forward compatibility.
type PerftestServer interface {
	GetPort(context.Context, *PortRequest) (*PortResponse, error)
	StartPerftest(context.Context, *PerftestRequest) (*PerftestResponse, error)
	mustEmbedUnimplementedPerftestServer()
}

// UnimplementedPerftestServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPerftestServer struct{}

func (UnimplementedPerftestServer) GetPort(context.Context, *PortRequest) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPerftestServer) StartPerftest(context.Context, *PerftestRequest) (*PerftestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPerftest not implemented")
}
func (UnimplementedPerftestServer) mustEmbedUnimplementedPerftestServer() {}
func (UnimplementedPerftestServer) testEmbeddedByValue()                  {}

// UnsafePerftestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PerftestServer will
// result in compilation errors.
type UnsafePerftestServer interface {
	mustEmbedUnimplementedPerftestServer()
}

func RegisterPerftestServer(s grpc.ServiceRegistrar, srv PerftestServer) {
	// If the following call pancis, it indicates UnimplementedPerftestServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Perftest_ServiceDesc, srv)
}

func _Perftest_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerftestServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perftest_GetPort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerftestServer).GetPort(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perftest_StartPerftest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerftestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerftestServer).StartPerftest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perftest_StartPerftest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerftestServer).StartPerftest(ctx, req.(*PerftestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Perftest_ServiceDesc is the grpc.ServiceDesc for Perftest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Perftest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "perftest.Perftest",
	HandlerType: (*PerftestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _Perftest_GetPort_Handler,
		},
		{
			MethodName: "StartPerftest",
			Handler:    _Perftest_StartPerftest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/perftest/perftest.proto",
}
