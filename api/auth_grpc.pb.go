// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: proto/auth.proto

package api

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

// AuthAppServiceClient is the client API for AuthAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthAppServiceClient interface {
	Register(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error)
	Login(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error)
	AddAccount(ctx context.Context, in *AddAppRequest, opts ...grpc.CallOption) (*AddedResponse, error)
}

type authAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAppServiceClient(cc grpc.ClientConnInterface) AuthAppServiceClient {
	return &authAppServiceClient{cc}
}

func (c *authAppServiceClient) Register(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthAppService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAppServiceClient) Login(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthAppService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAppServiceClient) AddAccount(ctx context.Context, in *AddAppRequest, opts ...grpc.CallOption) (*AddedResponse, error) {
	out := new(AddedResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthAppService/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAppServiceServer is the server API for AuthAppService service.
// All implementations must embed UnimplementedAuthAppServiceServer
// for forward compatibility
type AuthAppServiceServer interface {
	Register(context.Context, *AppRequest) (*AppResponse, error)
	Login(context.Context, *AppRequest) (*AppResponse, error)
	AddAccount(context.Context, *AddAppRequest) (*AddedResponse, error)
	mustEmbedUnimplementedAuthAppServiceServer()
}

// UnimplementedAuthAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthAppServiceServer struct {
}

func (UnimplementedAuthAppServiceServer) Register(context.Context, *AppRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthAppServiceServer) Login(context.Context, *AppRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthAppServiceServer) AddAccount(context.Context, *AddAppRequest) (*AddedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedAuthAppServiceServer) mustEmbedUnimplementedAuthAppServiceServer() {}

// UnsafeAuthAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthAppServiceServer will
// result in compilation errors.
type UnsafeAuthAppServiceServer interface {
	mustEmbedUnimplementedAuthAppServiceServer()
}

func RegisterAuthAppServiceServer(s grpc.ServiceRegistrar, srv AuthAppServiceServer) {
	s.RegisterService(&AuthAppService_ServiceDesc, srv)
}

func _AuthAppService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAppServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthAppService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAppServiceServer).Register(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAppService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAppServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthAppService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAppServiceServer).Login(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAppService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAppServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthAppService/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAppServiceServer).AddAccount(ctx, req.(*AddAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthAppService_ServiceDesc is the grpc.ServiceDesc for AuthAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.AuthAppService",
	HandlerType: (*AuthAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthAppService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthAppService_Login_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _AuthAppService_AddAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// AuthGithubServiceClient is the client API for AuthGithubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthGithubServiceClient interface {
	Login(ctx context.Context, in *GitHubRequest, opts ...grpc.CallOption) (*AppResponse, error)
	AddAccount(ctx context.Context, in *AddGitHubRequest, opts ...grpc.CallOption) (*AddedResponse, error)
}

type authGithubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGithubServiceClient(cc grpc.ClientConnInterface) AuthGithubServiceClient {
	return &authGithubServiceClient{cc}
}

func (c *authGithubServiceClient) Login(ctx context.Context, in *GitHubRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthGithubService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authGithubServiceClient) AddAccount(ctx context.Context, in *AddGitHubRequest, opts ...grpc.CallOption) (*AddedResponse, error) {
	out := new(AddedResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthGithubService/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGithubServiceServer is the server API for AuthGithubService service.
// All implementations must embed UnimplementedAuthGithubServiceServer
// for forward compatibility
type AuthGithubServiceServer interface {
	Login(context.Context, *GitHubRequest) (*AppResponse, error)
	AddAccount(context.Context, *AddGitHubRequest) (*AddedResponse, error)
	mustEmbedUnimplementedAuthGithubServiceServer()
}

// UnimplementedAuthGithubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthGithubServiceServer struct {
}

func (UnimplementedAuthGithubServiceServer) Login(context.Context, *GitHubRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthGithubServiceServer) AddAccount(context.Context, *AddGitHubRequest) (*AddedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedAuthGithubServiceServer) mustEmbedUnimplementedAuthGithubServiceServer() {}

// UnsafeAuthGithubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthGithubServiceServer will
// result in compilation errors.
type UnsafeAuthGithubServiceServer interface {
	mustEmbedUnimplementedAuthGithubServiceServer()
}

func RegisterAuthGithubServiceServer(s grpc.ServiceRegistrar, srv AuthGithubServiceServer) {
	s.RegisterService(&AuthGithubService_ServiceDesc, srv)
}

func _AuthGithubService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGithubServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthGithubService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGithubServiceServer).Login(ctx, req.(*GitHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthGithubService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGitHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGithubServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthGithubService/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGithubServiceServer).AddAccount(ctx, req.(*AddGitHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthGithubService_ServiceDesc is the grpc.ServiceDesc for AuthGithubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthGithubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.AuthGithubService",
	HandlerType: (*AuthGithubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthGithubService_Login_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _AuthGithubService_AddAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// AuthGoogleServiceClient is the client API for AuthGoogleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthGoogleServiceClient interface {
	Login(ctx context.Context, in *GoogleRequest, opts ...grpc.CallOption) (*AppResponse, error)
	AddAccount(ctx context.Context, in *AddGoogleRequest, opts ...grpc.CallOption) (*AddedResponse, error)
}

type authGoogleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGoogleServiceClient(cc grpc.ClientConnInterface) AuthGoogleServiceClient {
	return &authGoogleServiceClient{cc}
}

func (c *authGoogleServiceClient) Login(ctx context.Context, in *GoogleRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthGoogleService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authGoogleServiceClient) AddAccount(ctx context.Context, in *AddGoogleRequest, opts ...grpc.CallOption) (*AddedResponse, error) {
	out := new(AddedResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthGoogleService/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGoogleServiceServer is the server API for AuthGoogleService service.
// All implementations must embed UnimplementedAuthGoogleServiceServer
// for forward compatibility
type AuthGoogleServiceServer interface {
	Login(context.Context, *GoogleRequest) (*AppResponse, error)
	AddAccount(context.Context, *AddGoogleRequest) (*AddedResponse, error)
	mustEmbedUnimplementedAuthGoogleServiceServer()
}

// UnimplementedAuthGoogleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthGoogleServiceServer struct {
}

func (UnimplementedAuthGoogleServiceServer) Login(context.Context, *GoogleRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthGoogleServiceServer) AddAccount(context.Context, *AddGoogleRequest) (*AddedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedAuthGoogleServiceServer) mustEmbedUnimplementedAuthGoogleServiceServer() {}

// UnsafeAuthGoogleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthGoogleServiceServer will
// result in compilation errors.
type UnsafeAuthGoogleServiceServer interface {
	mustEmbedUnimplementedAuthGoogleServiceServer()
}

func RegisterAuthGoogleServiceServer(s grpc.ServiceRegistrar, srv AuthGoogleServiceServer) {
	s.RegisterService(&AuthGoogleService_ServiceDesc, srv)
}

func _AuthGoogleService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGoogleServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthGoogleService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGoogleServiceServer).Login(ctx, req.(*GoogleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthGoogleService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoogleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGoogleServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthGoogleService/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGoogleServiceServer).AddAccount(ctx, req.(*AddGoogleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthGoogleService_ServiceDesc is the grpc.ServiceDesc for AuthGoogleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthGoogleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.AuthGoogleService",
	HandlerType: (*AuthGoogleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthGoogleService_Login_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _AuthGoogleService_AddAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// AuthTelegramServiceClient is the client API for AuthTelegramService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthTelegramServiceClient interface {
	Login(ctx context.Context, in *TelegramRequest, opts ...grpc.CallOption) (*AppResponse, error)
}

type authTelegramServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthTelegramServiceClient(cc grpc.ClientConnInterface) AuthTelegramServiceClient {
	return &authTelegramServiceClient{cc}
}

func (c *authTelegramServiceClient) Login(ctx context.Context, in *TelegramRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.AuthTelegramService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthTelegramServiceServer is the server API for AuthTelegramService service.
// All implementations must embed UnimplementedAuthTelegramServiceServer
// for forward compatibility
type AuthTelegramServiceServer interface {
	Login(context.Context, *TelegramRequest) (*AppResponse, error)
	mustEmbedUnimplementedAuthTelegramServiceServer()
}

// UnimplementedAuthTelegramServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthTelegramServiceServer struct {
}

func (UnimplementedAuthTelegramServiceServer) Login(context.Context, *TelegramRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthTelegramServiceServer) mustEmbedUnimplementedAuthTelegramServiceServer() {}

// UnsafeAuthTelegramServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthTelegramServiceServer will
// result in compilation errors.
type UnsafeAuthTelegramServiceServer interface {
	mustEmbedUnimplementedAuthTelegramServiceServer()
}

func RegisterAuthTelegramServiceServer(s grpc.ServiceRegistrar, srv AuthTelegramServiceServer) {
	s.RegisterService(&AuthTelegramService_ServiceDesc, srv)
}

func _AuthTelegramService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelegramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthTelegramServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.AuthTelegramService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthTelegramServiceServer).Login(ctx, req.(*TelegramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthTelegramService_ServiceDesc is the grpc.ServiceDesc for AuthTelegramService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthTelegramService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.AuthTelegramService",
	HandlerType: (*AuthTelegramServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthTelegramService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// CheckAuthServiceClient is the client API for CheckAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckAuthServiceClient interface {
	CheckAuth(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error)
}

type checkAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckAuthServiceClient(cc grpc.ClientConnInterface) CheckAuthServiceClient {
	return &checkAuthServiceClient{cc}
}

func (c *checkAuthServiceClient) CheckAuth(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error) {
	out := new(CheckAuthResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.CheckAuthService/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckAuthServiceServer is the server API for CheckAuthService service.
// All implementations must embed UnimplementedCheckAuthServiceServer
// for forward compatibility
type CheckAuthServiceServer interface {
	CheckAuth(context.Context, *TokenRequest) (*CheckAuthResponse, error)
	mustEmbedUnimplementedCheckAuthServiceServer()
}

// UnimplementedCheckAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckAuthServiceServer struct {
}

func (UnimplementedCheckAuthServiceServer) CheckAuth(context.Context, *TokenRequest) (*CheckAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedCheckAuthServiceServer) mustEmbedUnimplementedCheckAuthServiceServer() {}

// UnsafeCheckAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckAuthServiceServer will
// result in compilation errors.
type UnsafeCheckAuthServiceServer interface {
	mustEmbedUnimplementedCheckAuthServiceServer()
}

func RegisterCheckAuthServiceServer(s grpc.ServiceRegistrar, srv CheckAuthServiceServer) {
	s.RegisterService(&CheckAuthService_ServiceDesc, srv)
}

func _CheckAuthService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckAuthServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.CheckAuthService/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckAuthServiceServer).CheckAuth(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckAuthService_ServiceDesc is the grpc.ServiceDesc for CheckAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.CheckAuthService",
	HandlerType: (*CheckAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _CheckAuthService_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// LogoutServiceClient is the client API for LogoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogoutServiceClient interface {
	Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type logoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogoutServiceClient(cc grpc.ClientConnInterface) LogoutServiceClient {
	return &logoutServiceClient{cc}
}

func (c *logoutServiceClient) Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/logger.v1.LogoutService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogoutServiceServer is the server API for LogoutService service.
// All implementations must embed UnimplementedLogoutServiceServer
// for forward compatibility
type LogoutServiceServer interface {
	Logout(context.Context, *TokenRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedLogoutServiceServer()
}

// UnimplementedLogoutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogoutServiceServer struct {
}

func (UnimplementedLogoutServiceServer) Logout(context.Context, *TokenRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedLogoutServiceServer) mustEmbedUnimplementedLogoutServiceServer() {}

// UnsafeLogoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogoutServiceServer will
// result in compilation errors.
type UnsafeLogoutServiceServer interface {
	mustEmbedUnimplementedLogoutServiceServer()
}

func RegisterLogoutServiceServer(s grpc.ServiceRegistrar, srv LogoutServiceServer) {
	s.RegisterService(&LogoutService_ServiceDesc, srv)
}

func _LogoutService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogoutServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.v1.LogoutService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogoutServiceServer).Logout(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogoutService_ServiceDesc is the grpc.ServiceDesc for LogoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1.LogoutService",
	HandlerType: (*LogoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logout",
			Handler:    _LogoutService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
