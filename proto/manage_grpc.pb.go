// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ManageClient is the client API for Manage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageClient interface {
	CheckServer(ctx context.Context, in *CheckServerRequest, opts ...grpc.CallOption) (*CheckServerResponse, error)
	InitialData(ctx context.Context, in *InitialDataRequest, opts ...grpc.CallOption) (*InitialDataResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error)
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	Tunnel(ctx context.Context, opts ...grpc.CallOption) (Manage_TunnelClient, error)
}

type manageClient struct {
	cc grpc.ClientConnInterface
}

func NewManageClient(cc grpc.ClientConnInterface) ManageClient {
	return &manageClient{cc}
}

func (c *manageClient) CheckServer(ctx context.Context, in *CheckServerRequest, opts ...grpc.CallOption) (*CheckServerResponse, error) {
	out := new(CheckServerResponse)
	err := c.cc.Invoke(ctx, "/Manage/CheckServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) InitialData(ctx context.Context, in *InitialDataRequest, opts ...grpc.CallOption) (*InitialDataResponse, error) {
	out := new(InitialDataResponse)
	err := c.cc.Invoke(ctx, "/Manage/InitialData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Manage/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error) {
	out := new(SetPasswordResponse)
	err := c.cc.Invoke(ctx, "/Manage/SetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/Manage/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Tunnel(ctx context.Context, opts ...grpc.CallOption) (Manage_TunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manage_ServiceDesc.Streams[0], "/Manage/Tunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageTunnelClient{stream}
	return x, nil
}

type Manage_TunnelClient interface {
	Send(*TunnelData) error
	Recv() (*TunnelData, error)
	grpc.ClientStream
}

type manageTunnelClient struct {
	grpc.ClientStream
}

func (x *manageTunnelClient) Send(m *TunnelData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *manageTunnelClient) Recv() (*TunnelData, error) {
	m := new(TunnelData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManageServer is the server API for Manage service.
// All implementations should embed UnimplementedManageServer
// for forward compatibility
type ManageServer interface {
	CheckServer(context.Context, *CheckServerRequest) (*CheckServerResponse, error)
	InitialData(context.Context, *InitialDataRequest) (*InitialDataResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error)
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
	Tunnel(Manage_TunnelServer) error
}

// UnimplementedManageServer should be embedded to have forward compatible implementations.
type UnimplementedManageServer struct {
}

func (UnimplementedManageServer) CheckServer(context.Context, *CheckServerRequest) (*CheckServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckServer not implemented")
}
func (UnimplementedManageServer) InitialData(context.Context, *InitialDataRequest) (*InitialDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialData not implemented")
}
func (UnimplementedManageServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedManageServer) SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedManageServer) Config(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedManageServer) Tunnel(Manage_TunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method Tunnel not implemented")
}

// UnsafeManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageServer will
// result in compilation errors.
type UnsafeManageServer interface {
	mustEmbedUnimplementedManageServer()
}

func RegisterManageServer(s grpc.ServiceRegistrar, srv ManageServer) {
	s.RegisterService(&Manage_ServiceDesc, srv)
}

func _Manage_CheckServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).CheckServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/CheckServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).CheckServer(ctx, req.(*CheckServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_InitialData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitialDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).InitialData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/InitialData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).InitialData(ctx, req.(*InitialDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Tunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManageServer).Tunnel(&manageTunnelServer{stream})
}

type Manage_TunnelServer interface {
	Send(*TunnelData) error
	Recv() (*TunnelData, error)
	grpc.ServerStream
}

type manageTunnelServer struct {
	grpc.ServerStream
}

func (x *manageTunnelServer) Send(m *TunnelData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *manageTunnelServer) Recv() (*TunnelData, error) {
	m := new(TunnelData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Manage_ServiceDesc is the grpc.ServiceDesc for Manage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Manage",
	HandlerType: (*ManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckServer",
			Handler:    _Manage_CheckServer_Handler,
		},
		{
			MethodName: "InitialData",
			Handler:    _Manage_InitialData_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Manage_CreateUser_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _Manage_SetPassword_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _Manage_Config_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tunnel",
			Handler:       _Manage_Tunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/manage.proto",
}