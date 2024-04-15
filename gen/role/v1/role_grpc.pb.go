// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: role/v1/role.proto

package rolev1

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
	RoleService_GetRoleList_FullMethodName = "/role.v1.RoleService/GetRoleList"
	RoleService_GetRole_FullMethodName     = "/role.v1.RoleService/GetRole"
	RoleService_AddRole_FullMethodName     = "/role.v1.RoleService/AddRole"
	RoleService_EditRole_FullMethodName    = "/role.v1.RoleService/EditRole"
	RoleService_RemoveRole_FullMethodName  = "/role.v1.RoleService/RemoveRole"
)

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetRoleListResponse, error)
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
	AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error)
	EditRole(ctx context.Context, in *EditRoleRequest, opts ...grpc.CallOption) (*EditRoleResponse, error)
	RemoveRole(ctx context.Context, in *RemoveRoleRequest, opts ...grpc.CallOption) (*RemoveRoleResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetRoleListResponse, error) {
	out := new(GetRoleListResponse)
	err := c.cc.Invoke(ctx, RoleService_GetRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	out := new(GetRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error) {
	out := new(AddRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_AddRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) EditRole(ctx context.Context, in *EditRoleRequest, opts ...grpc.CallOption) (*EditRoleResponse, error) {
	out := new(EditRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_EditRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RemoveRole(ctx context.Context, in *RemoveRoleRequest, opts ...grpc.CallOption) (*RemoveRoleResponse, error) {
	out := new(RemoveRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_RemoveRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	GetRoleList(context.Context, *GetRoleListRequest) (*GetRoleListResponse, error)
	GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error)
	AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error)
	EditRole(context.Context, *EditRoleRequest) (*EditRoleResponse, error)
	RemoveRole(context.Context, *RemoveRoleRequest) (*RemoveRoleResponse, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) GetRoleList(context.Context, *GetRoleListRequest) (*GetRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleList not implemented")
}
func (UnimplementedRoleServiceServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRoleServiceServer) AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedRoleServiceServer) EditRole(context.Context, *EditRoleRequest) (*EditRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRole not implemented")
}
func (UnimplementedRoleServiceServer) RemoveRole(context.Context, *RemoveRoleRequest) (*RemoveRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRole not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_GetRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GetRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetRoleList(ctx, req.(*GetRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_AddRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).AddRole(ctx, req.(*AddRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_EditRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).EditRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_EditRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).EditRole(ctx, req.(*EditRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_RemoveRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RemoveRole(ctx, req.(*RemoveRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.v1.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoleList",
			Handler:    _RoleService_GetRoleList_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RoleService_GetRole_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _RoleService_AddRole_Handler,
		},
		{
			MethodName: "EditRole",
			Handler:    _RoleService_EditRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _RoleService_RemoveRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role/v1/role.proto",
}
