// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: pingpongservice.proto

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

const (
	PingPongService_Play_FullMethodName = "/pb.PingPongService/Play"
)

// PingPongServiceClient is the client API for PingPongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongServiceClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (PingPongService_PlayClient, error)
}

type pingPongServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongServiceClient(cc grpc.ClientConnInterface) PingPongServiceClient {
	return &pingPongServiceClient{cc}
}

func (c *pingPongServiceClient) Play(ctx context.Context, opts ...grpc.CallOption) (PingPongService_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingPongService_ServiceDesc.Streams[0], PingPongService_Play_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pingPongServicePlayClient{stream}
	return x, nil
}

type PingPongService_PlayClient interface {
	Send(*Ball) error
	Recv() (*Ball, error)
	grpc.ClientStream
}

type pingPongServicePlayClient struct {
	grpc.ClientStream
}

func (x *pingPongServicePlayClient) Send(m *Ball) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingPongServicePlayClient) Recv() (*Ball, error) {
	m := new(Ball)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongServiceServer is the server API for PingPongService service.
// All implementations should embed UnimplementedPingPongServiceServer
// for forward compatibility
type PingPongServiceServer interface {
	Play(PingPongService_PlayServer) error
}

// UnimplementedPingPongServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPingPongServiceServer struct {
}

func (UnimplementedPingPongServiceServer) Play(PingPongService_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}

// UnsafePingPongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServiceServer will
// result in compilation errors.
type UnsafePingPongServiceServer interface {
	mustEmbedUnimplementedPingPongServiceServer()
}

func RegisterPingPongServiceServer(s grpc.ServiceRegistrar, srv PingPongServiceServer) {
	s.RegisterService(&PingPongService_ServiceDesc, srv)
}

func _PingPongService_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingPongServiceServer).Play(&pingPongServicePlayServer{stream})
}

type PingPongService_PlayServer interface {
	Send(*Ball) error
	Recv() (*Ball, error)
	grpc.ServerStream
}

type pingPongServicePlayServer struct {
	grpc.ServerStream
}

func (x *pingPongServicePlayServer) Send(m *Ball) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingPongServicePlayServer) Recv() (*Ball, error) {
	m := new(Ball)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongService_ServiceDesc is the grpc.ServiceDesc for PingPongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PingPongService",
	HandlerType: (*PingPongServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _PingPongService_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pingpongservice.proto",
}