// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: sso/sso.proto

package ssov1

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
	Auth_Register_FullMethodName         = "/auth.Auth/Register"
	Auth_Login_FullMethodName            = "/auth.Auth/Login"
	Auth_IsAdmin_FullMethodName          = "/auth.Auth/IsAdmin"
	Auth_GetOrCreateApp_FullMethodName   = "/auth.Auth/GetOrCreateApp"
	Auth_ActivateUser_FullMethodName     = "/auth.Auth/ActivateUser"
	Auth_RenewAccessToken_FullMethodName = "/auth.Auth/RenewAccessToken"
	Auth_GetUser_FullMethodName          = "/auth.Auth/GetUser"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminResponse, error)
	GetOrCreateApp(ctx context.Context, in *GetOrCreateAppRequest, opts ...grpc.CallOption) (*GetOrCreateAppResponse, error)
	ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error)
	RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Auth_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsAdminResponse)
	err := c.cc.Invoke(ctx, Auth_IsAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetOrCreateApp(ctx context.Context, in *GetOrCreateAppRequest, opts ...grpc.CallOption) (*GetOrCreateAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrCreateAppResponse)
	err := c.cc.Invoke(ctx, Auth_GetOrCreateApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateUserResponse)
	err := c.cc.Invoke(ctx, Auth_ActivateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenewAccessTokenResponse)
	err := c.cc.Invoke(ctx, Auth_RenewAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, Auth_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	IsAdmin(context.Context, *IsAdminRequest) (*IsAdminResponse, error)
	GetOrCreateApp(context.Context, *GetOrCreateAppRequest) (*GetOrCreateAppResponse, error)
	ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error)
	RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) IsAdmin(context.Context, *IsAdminRequest) (*IsAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAdmin not implemented")
}
func (UnimplementedAuthServer) GetOrCreateApp(context.Context, *GetOrCreateAppRequest) (*GetOrCreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateApp not implemented")
}
func (UnimplementedAuthServer) ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (UnimplementedAuthServer) RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccessToken not implemented")
}
func (UnimplementedAuthServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_IsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_IsAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsAdmin(ctx, req.(*IsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetOrCreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetOrCreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetOrCreateApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetOrCreateApp(ctx, req.(*GetOrCreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ActivateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ActivateUser(ctx, req.(*ActivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RenewAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RenewAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RenewAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RenewAccessToken(ctx, req.(*RenewAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "IsAdmin",
			Handler:    _Auth_IsAdmin_Handler,
		},
		{
			MethodName: "GetOrCreateApp",
			Handler:    _Auth_GetOrCreateApp_Handler,
		},
		{
			MethodName: "ActivateUser",
			Handler:    _Auth_ActivateUser_Handler,
		},
		{
			MethodName: "RenewAccessToken",
			Handler:    _Auth_RenewAccessToken_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Auth_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
