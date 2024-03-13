// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package audiov1

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

// AudioServiceClient is the client API for AudioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioServiceClient interface {
	SteamAudio(ctx context.Context, in *AudioRequest, opts ...grpc.CallOption) (*AudioResponse, error)
}

type audioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioServiceClient(cc grpc.ClientConnInterface) AudioServiceClient {
	return &audioServiceClient{cc}
}

func (c *audioServiceClient) SteamAudio(ctx context.Context, in *AudioRequest, opts ...grpc.CallOption) (*AudioResponse, error) {
	out := new(AudioResponse)
	err := c.cc.Invoke(ctx, "/music.AudioService/SteamAudio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioServiceServer is the server API for AudioService service.
// All implementations must embed UnimplementedAudioServiceServer
// for forward compatibility
type AudioServiceServer interface {
	SteamAudio(context.Context, *AudioRequest) (*AudioResponse, error)
	mustEmbedUnimplementedAudioServiceServer()
}

// UnimplementedAudioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAudioServiceServer struct {
}

func (UnimplementedAudioServiceServer) SteamAudio(context.Context, *AudioRequest) (*AudioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SteamAudio not implemented")
}
func (UnimplementedAudioServiceServer) mustEmbedUnimplementedAudioServiceServer() {}

// UnsafeAudioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServiceServer will
// result in compilation errors.
type UnsafeAudioServiceServer interface {
	mustEmbedUnimplementedAudioServiceServer()
}

func RegisterAudioServiceServer(s grpc.ServiceRegistrar, srv AudioServiceServer) {
	s.RegisterService(&AudioService_ServiceDesc, srv)
}

func _AudioService_SteamAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AudioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).SteamAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music.AudioService/SteamAudio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).SteamAudio(ctx, req.(*AudioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioService_ServiceDesc is the grpc.ServiceDesc for AudioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music.AudioService",
	HandlerType: (*AudioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SteamAudio",
			Handler:    _AudioService_SteamAudio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audio/audio.proto",
}
