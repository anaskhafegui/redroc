// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: grpc/protos/download.proto

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

// DownloadPhotoClient is the client API for DownloadPhoto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadPhotoClient interface {
	// RPC for download a photo
	Download(ctx context.Context, in *DownloadPhotoRequest, opts ...grpc.CallOption) (*DownloadPhotoResponse, error)
}

type downloadPhotoClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadPhotoClient(cc grpc.ClientConnInterface) DownloadPhotoClient {
	return &downloadPhotoClient{cc}
}

func (c *downloadPhotoClient) Download(ctx context.Context, in *DownloadPhotoRequest, opts ...grpc.CallOption) (*DownloadPhotoResponse, error) {
	out := new(DownloadPhotoResponse)
	err := c.cc.Invoke(ctx, "/grpc.DownloadPhoto/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadPhotoServer is the server API for DownloadPhoto service.
// All implementations must embed UnimplementedDownloadPhotoServer
// for forward compatibility
type DownloadPhotoServer interface {
	// RPC for download a photo
	Download(context.Context, *DownloadPhotoRequest) (*DownloadPhotoResponse, error)
	mustEmbedUnimplementedDownloadPhotoServer()
}

// UnimplementedDownloadPhotoServer must be embedded to have forward compatible implementations.
type UnimplementedDownloadPhotoServer struct {
}

func (UnimplementedDownloadPhotoServer) Download(context.Context, *DownloadPhotoRequest) (*DownloadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDownloadPhotoServer) mustEmbedUnimplementedDownloadPhotoServer() {}

// UnsafeDownloadPhotoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadPhotoServer will
// result in compilation errors.
type UnsafeDownloadPhotoServer interface {
	mustEmbedUnimplementedDownloadPhotoServer()
}

func RegisterDownloadPhotoServer(s grpc.ServiceRegistrar, srv DownloadPhotoServer) {
	s.RegisterService(&DownloadPhoto_ServiceDesc, srv)
}

func _DownloadPhoto_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadPhotoServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DownloadPhoto/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadPhotoServer).Download(ctx, req.(*DownloadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DownloadPhoto_ServiceDesc is the grpc.ServiceDesc for DownloadPhoto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DownloadPhoto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DownloadPhoto",
	HandlerType: (*DownloadPhotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _DownloadPhoto_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/protos/download.proto",
}
