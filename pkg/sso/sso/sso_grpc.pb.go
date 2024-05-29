// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: sso/sso.proto

package sso

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

// SsoClient is the client API for Sso service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SsoClient interface {
	GetInfoByGoogleToken(ctx context.Context, in *GoogleTokenRequest, opts ...grpc.CallOption) (*GoogleTokenResponse, error)
}

type ssoClient struct {
	cc grpc.ClientConnInterface
}

func NewSsoClient(cc grpc.ClientConnInterface) SsoClient {
	return &ssoClient{cc}
}

func (c *ssoClient) GetInfoByGoogleToken(ctx context.Context, in *GoogleTokenRequest, opts ...grpc.CallOption) (*GoogleTokenResponse, error) {
	out := new(GoogleTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Sso/GetInfoByGoogleToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsoServer is the server API for Sso service.
// All implementations must embed UnimplementedSsoServer
// for forward compatibility
type SsoServer interface {
	GetInfoByGoogleToken(context.Context, *GoogleTokenRequest) (*GoogleTokenResponse, error)
	mustEmbedUnimplementedSsoServer()
}

// UnimplementedSsoServer must be embedded to have forward compatible implementations.
type UnimplementedSsoServer struct {
}

func (UnimplementedSsoServer) GetInfoByGoogleToken(context.Context, *GoogleTokenRequest) (*GoogleTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByGoogleToken not implemented")
}
func (UnimplementedSsoServer) mustEmbedUnimplementedSsoServer() {}

// UnsafeSsoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SsoServer will
// result in compilation errors.
type UnsafeSsoServer interface {
	mustEmbedUnimplementedSsoServer()
}

func RegisterSsoServer(s grpc.ServiceRegistrar, srv SsoServer) {
	s.RegisterService(&Sso_ServiceDesc, srv)
}

func _Sso_GetInfoByGoogleToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoServer).GetInfoByGoogleToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Sso/GetInfoByGoogleToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoServer).GetInfoByGoogleToken(ctx, req.(*GoogleTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sso_ServiceDesc is the grpc.ServiceDesc for Sso service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sso_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Sso",
	HandlerType: (*SsoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfoByGoogleToken",
			Handler:    _Sso_GetInfoByGoogleToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
