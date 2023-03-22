// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.21.9
// source: remote.proto

package remote

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

// QueryTimeSeriesServiceClient is the client API for QueryTimeSeriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryTimeSeriesServiceClient interface {
	Raw(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (QueryTimeSeriesService_RawClient, error)
}

type queryTimeSeriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryTimeSeriesServiceClient(cc grpc.ClientConnInterface) QueryTimeSeriesServiceClient {
	return &queryTimeSeriesServiceClient{cc}
}

func (c *queryTimeSeriesServiceClient) Raw(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (QueryTimeSeriesService_RawClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryTimeSeriesService_ServiceDesc.Streams[0], "/remote.QueryTimeSeriesService/Raw", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryTimeSeriesServiceRawClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryTimeSeriesService_RawClient interface {
	Recv() (*TimeSeries, error)
	grpc.ClientStream
}

type queryTimeSeriesServiceRawClient struct {
	grpc.ClientStream
}

func (x *queryTimeSeriesServiceRawClient) Recv() (*TimeSeries, error) {
	m := new(TimeSeries)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryTimeSeriesServiceServer is the server API for QueryTimeSeriesService service.
// All implementations must embed UnimplementedQueryTimeSeriesServiceServer
// for forward compatibility
type QueryTimeSeriesServiceServer interface {
	Raw(*FilterRequest, QueryTimeSeriesService_RawServer) error
	mustEmbedUnimplementedQueryTimeSeriesServiceServer()
}

// UnimplementedQueryTimeSeriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryTimeSeriesServiceServer struct {
}

func (UnimplementedQueryTimeSeriesServiceServer) Raw(*FilterRequest, QueryTimeSeriesService_RawServer) error {
	return status.Errorf(codes.Unimplemented, "method Raw not implemented")
}
func (UnimplementedQueryTimeSeriesServiceServer) mustEmbedUnimplementedQueryTimeSeriesServiceServer() {
}

// UnsafeQueryTimeSeriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryTimeSeriesServiceServer will
// result in compilation errors.
type UnsafeQueryTimeSeriesServiceServer interface {
	mustEmbedUnimplementedQueryTimeSeriesServiceServer()
}

func RegisterQueryTimeSeriesServiceServer(s grpc.ServiceRegistrar, srv QueryTimeSeriesServiceServer) {
	s.RegisterService(&QueryTimeSeriesService_ServiceDesc, srv)
}

func _QueryTimeSeriesService_Raw_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilterRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryTimeSeriesServiceServer).Raw(m, &queryTimeSeriesServiceRawServer{stream})
}

type QueryTimeSeriesService_RawServer interface {
	Send(*TimeSeries) error
	grpc.ServerStream
}

type queryTimeSeriesServiceRawServer struct {
	grpc.ServerStream
}

func (x *queryTimeSeriesServiceRawServer) Send(m *TimeSeries) error {
	return x.ServerStream.SendMsg(m)
}

// QueryTimeSeriesService_ServiceDesc is the grpc.ServiceDesc for QueryTimeSeriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryTimeSeriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.QueryTimeSeriesService",
	HandlerType: (*QueryTimeSeriesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Raw",
			Handler:       _QueryTimeSeriesService_Raw_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote.proto",
}