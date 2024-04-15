// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: role/v1/role.proto

package rolev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// RoleServiceName is the fully-qualified name of the RoleService service.
	RoleServiceName = "role.v1.RoleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RoleServiceGetRoleListProcedure is the fully-qualified name of the RoleService's GetRoleList RPC.
	RoleServiceGetRoleListProcedure = "/role.v1.RoleService/GetRoleList"
	// RoleServiceGetRoleProcedure is the fully-qualified name of the RoleService's GetRole RPC.
	RoleServiceGetRoleProcedure = "/role.v1.RoleService/GetRole"
	// RoleServiceAddRoleProcedure is the fully-qualified name of the RoleService's AddRole RPC.
	RoleServiceAddRoleProcedure = "/role.v1.RoleService/AddRole"
	// RoleServiceEditRoleProcedure is the fully-qualified name of the RoleService's EditRole RPC.
	RoleServiceEditRoleProcedure = "/role.v1.RoleService/EditRole"
	// RoleServiceRemoveRoleProcedure is the fully-qualified name of the RoleService's RemoveRole RPC.
	RoleServiceRemoveRoleProcedure = "/role.v1.RoleService/RemoveRole"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	roleServiceServiceDescriptor           = v1.File_role_v1_role_proto.Services().ByName("RoleService")
	roleServiceGetRoleListMethodDescriptor = roleServiceServiceDescriptor.Methods().ByName("GetRoleList")
	roleServiceGetRoleMethodDescriptor     = roleServiceServiceDescriptor.Methods().ByName("GetRole")
	roleServiceAddRoleMethodDescriptor     = roleServiceServiceDescriptor.Methods().ByName("AddRole")
	roleServiceEditRoleMethodDescriptor    = roleServiceServiceDescriptor.Methods().ByName("EditRole")
	roleServiceRemoveRoleMethodDescriptor  = roleServiceServiceDescriptor.Methods().ByName("RemoveRole")
)

// RoleServiceClient is a client for the role.v1.RoleService service.
type RoleServiceClient interface {
	GetRoleList(context.Context, *connect.Request[v1.GetRoleListRequest]) (*connect.Response[v1.GetRoleListResponse], error)
	GetRole(context.Context, *connect.Request[v1.GetRoleRequest]) (*connect.Response[v1.GetRoleResponse], error)
	AddRole(context.Context, *connect.Request[v1.AddRoleRequest]) (*connect.Response[v1.AddRoleResponse], error)
	EditRole(context.Context, *connect.Request[v1.EditRoleRequest]) (*connect.Response[v1.EditRoleResponse], error)
	RemoveRole(context.Context, *connect.Request[v1.RemoveRoleRequest]) (*connect.Response[v1.RemoveRoleResponse], error)
}

// NewRoleServiceClient constructs a client for the role.v1.RoleService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRoleServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RoleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &roleServiceClient{
		getRoleList: connect.NewClient[v1.GetRoleListRequest, v1.GetRoleListResponse](
			httpClient,
			baseURL+RoleServiceGetRoleListProcedure,
			connect.WithSchema(roleServiceGetRoleListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getRole: connect.NewClient[v1.GetRoleRequest, v1.GetRoleResponse](
			httpClient,
			baseURL+RoleServiceGetRoleProcedure,
			connect.WithSchema(roleServiceGetRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addRole: connect.NewClient[v1.AddRoleRequest, v1.AddRoleResponse](
			httpClient,
			baseURL+RoleServiceAddRoleProcedure,
			connect.WithSchema(roleServiceAddRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		editRole: connect.NewClient[v1.EditRoleRequest, v1.EditRoleResponse](
			httpClient,
			baseURL+RoleServiceEditRoleProcedure,
			connect.WithSchema(roleServiceEditRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeRole: connect.NewClient[v1.RemoveRoleRequest, v1.RemoveRoleResponse](
			httpClient,
			baseURL+RoleServiceRemoveRoleProcedure,
			connect.WithSchema(roleServiceRemoveRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// roleServiceClient implements RoleServiceClient.
type roleServiceClient struct {
	getRoleList *connect.Client[v1.GetRoleListRequest, v1.GetRoleListResponse]
	getRole     *connect.Client[v1.GetRoleRequest, v1.GetRoleResponse]
	addRole     *connect.Client[v1.AddRoleRequest, v1.AddRoleResponse]
	editRole    *connect.Client[v1.EditRoleRequest, v1.EditRoleResponse]
	removeRole  *connect.Client[v1.RemoveRoleRequest, v1.RemoveRoleResponse]
}

// GetRoleList calls role.v1.RoleService.GetRoleList.
func (c *roleServiceClient) GetRoleList(ctx context.Context, req *connect.Request[v1.GetRoleListRequest]) (*connect.Response[v1.GetRoleListResponse], error) {
	return c.getRoleList.CallUnary(ctx, req)
}

// GetRole calls role.v1.RoleService.GetRole.
func (c *roleServiceClient) GetRole(ctx context.Context, req *connect.Request[v1.GetRoleRequest]) (*connect.Response[v1.GetRoleResponse], error) {
	return c.getRole.CallUnary(ctx, req)
}

// AddRole calls role.v1.RoleService.AddRole.
func (c *roleServiceClient) AddRole(ctx context.Context, req *connect.Request[v1.AddRoleRequest]) (*connect.Response[v1.AddRoleResponse], error) {
	return c.addRole.CallUnary(ctx, req)
}

// EditRole calls role.v1.RoleService.EditRole.
func (c *roleServiceClient) EditRole(ctx context.Context, req *connect.Request[v1.EditRoleRequest]) (*connect.Response[v1.EditRoleResponse], error) {
	return c.editRole.CallUnary(ctx, req)
}

// RemoveRole calls role.v1.RoleService.RemoveRole.
func (c *roleServiceClient) RemoveRole(ctx context.Context, req *connect.Request[v1.RemoveRoleRequest]) (*connect.Response[v1.RemoveRoleResponse], error) {
	return c.removeRole.CallUnary(ctx, req)
}

// RoleServiceHandler is an implementation of the role.v1.RoleService service.
type RoleServiceHandler interface {
	GetRoleList(context.Context, *connect.Request[v1.GetRoleListRequest]) (*connect.Response[v1.GetRoleListResponse], error)
	GetRole(context.Context, *connect.Request[v1.GetRoleRequest]) (*connect.Response[v1.GetRoleResponse], error)
	AddRole(context.Context, *connect.Request[v1.AddRoleRequest]) (*connect.Response[v1.AddRoleResponse], error)
	EditRole(context.Context, *connect.Request[v1.EditRoleRequest]) (*connect.Response[v1.EditRoleResponse], error)
	RemoveRole(context.Context, *connect.Request[v1.RemoveRoleRequest]) (*connect.Response[v1.RemoveRoleResponse], error)
}

// NewRoleServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRoleServiceHandler(svc RoleServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	roleServiceGetRoleListHandler := connect.NewUnaryHandler(
		RoleServiceGetRoleListProcedure,
		svc.GetRoleList,
		connect.WithSchema(roleServiceGetRoleListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	roleServiceGetRoleHandler := connect.NewUnaryHandler(
		RoleServiceGetRoleProcedure,
		svc.GetRole,
		connect.WithSchema(roleServiceGetRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	roleServiceAddRoleHandler := connect.NewUnaryHandler(
		RoleServiceAddRoleProcedure,
		svc.AddRole,
		connect.WithSchema(roleServiceAddRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	roleServiceEditRoleHandler := connect.NewUnaryHandler(
		RoleServiceEditRoleProcedure,
		svc.EditRole,
		connect.WithSchema(roleServiceEditRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	roleServiceRemoveRoleHandler := connect.NewUnaryHandler(
		RoleServiceRemoveRoleProcedure,
		svc.RemoveRole,
		connect.WithSchema(roleServiceRemoveRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/role.v1.RoleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RoleServiceGetRoleListProcedure:
			roleServiceGetRoleListHandler.ServeHTTP(w, r)
		case RoleServiceGetRoleProcedure:
			roleServiceGetRoleHandler.ServeHTTP(w, r)
		case RoleServiceAddRoleProcedure:
			roleServiceAddRoleHandler.ServeHTTP(w, r)
		case RoleServiceEditRoleProcedure:
			roleServiceEditRoleHandler.ServeHTTP(w, r)
		case RoleServiceRemoveRoleProcedure:
			roleServiceRemoveRoleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRoleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRoleServiceHandler struct{}

func (UnimplementedRoleServiceHandler) GetRoleList(context.Context, *connect.Request[v1.GetRoleListRequest]) (*connect.Response[v1.GetRoleListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("role.v1.RoleService.GetRoleList is not implemented"))
}

func (UnimplementedRoleServiceHandler) GetRole(context.Context, *connect.Request[v1.GetRoleRequest]) (*connect.Response[v1.GetRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("role.v1.RoleService.GetRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) AddRole(context.Context, *connect.Request[v1.AddRoleRequest]) (*connect.Response[v1.AddRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("role.v1.RoleService.AddRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) EditRole(context.Context, *connect.Request[v1.EditRoleRequest]) (*connect.Response[v1.EditRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("role.v1.RoleService.EditRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) RemoveRole(context.Context, *connect.Request[v1.RemoveRoleRequest]) (*connect.Response[v1.RemoveRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("role.v1.RoleService.RemoveRole is not implemented"))
}
