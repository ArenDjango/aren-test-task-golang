// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: rates_service.proto

package rates

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

// RatesServiceClient is the client API for RatesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatesServiceClient interface {
	GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error)
}

type ratesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatesServiceClient(cc grpc.ClientConnInterface) RatesServiceClient {
	return &ratesServiceClient{cc}
}

func (c *ratesServiceClient) GetRates(ctx context.Context, in *GetRatesRequest, opts ...grpc.CallOption) (*GetRatesResponse, error) {
	out := new(GetRatesResponse)
	err := c.cc.Invoke(ctx, "/ratesService/GetRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatesServiceServer is the server API for RatesService service.
// All implementations must embed UnimplementedRatesServiceServer
// for forward compatibility
type RatesServiceServer interface {
	GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error)
	mustEmbedUnimplementedRatesServiceServer()
}

// UnimplementedRatesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatesServiceServer struct {
}

func (UnimplementedRatesServiceServer) GetRates(context.Context, *GetRatesRequest) (*GetRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}
func (UnimplementedRatesServiceServer) mustEmbedUnimplementedRatesServiceServer() {}

// UnsafeRatesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatesServiceServer will
// result in compilation errors.
type UnsafeRatesServiceServer interface {
	mustEmbedUnimplementedRatesServiceServer()
}

func RegisterRatesServiceServer(s grpc.ServiceRegistrar, srv RatesServiceServer) {
	s.RegisterService(&RatesService_ServiceDesc, srv)
}

func _RatesService_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatesServiceServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratesService/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatesServiceServer).GetRates(ctx, req.(*GetRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatesService_ServiceDesc is the grpc.ServiceDesc for RatesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ratesService",
	HandlerType: (*RatesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _RatesService_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rates_service.proto",
}
