// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_suggestion_api

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

// OcpSuggestionApiClient is the client API for OcpSuggestionApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpSuggestionApiClient interface {
	// CreateSuggestionV1 создаёт предложение курса и возвращает id предложения
	CreateSuggestionV1(ctx context.Context, in *CreateSuggestionV1Request, opts ...grpc.CallOption) (*CreateSuggestionV1Response, error)
	// DescribeSuggestionV1 возвращает описание предложения с указанным id
	DescribeSuggestionV1(ctx context.Context, in *DescribeSuggestionV1Request, opts ...grpc.CallOption) (*DescribeSuggestionV1Response, error)
	// ListSuggestionV1 возвращает список предложений
	ListSuggestionV1(ctx context.Context, in *ListSuggestionV1Request, opts ...grpc.CallOption) (*ListSuggestionV1Response, error)
	// UpdateSuggestionV1 обновляет предложение с указанным id
	UpdateSuggestionV1(ctx context.Context, in *UpdateSuggestionV1Request, opts ...grpc.CallOption) (*UpdateSuggestionV1Response, error)
	// RemoveSuggestionV1 удаляет предложение с указанным id
	RemoveSuggestionV1(ctx context.Context, in *RemoveSuggestionV1Request, opts ...grpc.CallOption) (*RemoveSuggestionV1Response, error)
}

type ocpSuggestionApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpSuggestionApiClient(cc grpc.ClientConnInterface) OcpSuggestionApiClient {
	return &ocpSuggestionApiClient{cc}
}

func (c *ocpSuggestionApiClient) CreateSuggestionV1(ctx context.Context, in *CreateSuggestionV1Request, opts ...grpc.CallOption) (*CreateSuggestionV1Response, error) {
	out := new(CreateSuggestionV1Response)
	err := c.cc.Invoke(ctx, "/ocp.suggestion.api.OcpSuggestionApi/CreateSuggestionV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSuggestionApiClient) DescribeSuggestionV1(ctx context.Context, in *DescribeSuggestionV1Request, opts ...grpc.CallOption) (*DescribeSuggestionV1Response, error) {
	out := new(DescribeSuggestionV1Response)
	err := c.cc.Invoke(ctx, "/ocp.suggestion.api.OcpSuggestionApi/DescribeSuggestionV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSuggestionApiClient) ListSuggestionV1(ctx context.Context, in *ListSuggestionV1Request, opts ...grpc.CallOption) (*ListSuggestionV1Response, error) {
	out := new(ListSuggestionV1Response)
	err := c.cc.Invoke(ctx, "/ocp.suggestion.api.OcpSuggestionApi/ListSuggestionV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSuggestionApiClient) UpdateSuggestionV1(ctx context.Context, in *UpdateSuggestionV1Request, opts ...grpc.CallOption) (*UpdateSuggestionV1Response, error) {
	out := new(UpdateSuggestionV1Response)
	err := c.cc.Invoke(ctx, "/ocp.suggestion.api.OcpSuggestionApi/UpdateSuggestionV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSuggestionApiClient) RemoveSuggestionV1(ctx context.Context, in *RemoveSuggestionV1Request, opts ...grpc.CallOption) (*RemoveSuggestionV1Response, error) {
	out := new(RemoveSuggestionV1Response)
	err := c.cc.Invoke(ctx, "/ocp.suggestion.api.OcpSuggestionApi/RemoveSuggestionV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpSuggestionApiServer is the server API for OcpSuggestionApi service.
// All implementations must embed UnimplementedOcpSuggestionApiServer
// for forward compatibility
type OcpSuggestionApiServer interface {
	// CreateSuggestionV1 создаёт предложение курса и возвращает id предложения
	CreateSuggestionV1(context.Context, *CreateSuggestionV1Request) (*CreateSuggestionV1Response, error)
	// DescribeSuggestionV1 возвращает описание предложения с указанным id
	DescribeSuggestionV1(context.Context, *DescribeSuggestionV1Request) (*DescribeSuggestionV1Response, error)
	// ListSuggestionV1 возвращает список предложений
	ListSuggestionV1(context.Context, *ListSuggestionV1Request) (*ListSuggestionV1Response, error)
	// UpdateSuggestionV1 обновляет предложение с указанным id
	UpdateSuggestionV1(context.Context, *UpdateSuggestionV1Request) (*UpdateSuggestionV1Response, error)
	// RemoveSuggestionV1 удаляет предложение с указанным id
	RemoveSuggestionV1(context.Context, *RemoveSuggestionV1Request) (*RemoveSuggestionV1Response, error)
	mustEmbedUnimplementedOcpSuggestionApiServer()
}

// UnimplementedOcpSuggestionApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpSuggestionApiServer struct {
}

func (UnimplementedOcpSuggestionApiServer) CreateSuggestionV1(context.Context, *CreateSuggestionV1Request) (*CreateSuggestionV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSuggestionV1 not implemented")
}
func (UnimplementedOcpSuggestionApiServer) DescribeSuggestionV1(context.Context, *DescribeSuggestionV1Request) (*DescribeSuggestionV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSuggestionV1 not implemented")
}
func (UnimplementedOcpSuggestionApiServer) ListSuggestionV1(context.Context, *ListSuggestionV1Request) (*ListSuggestionV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSuggestionV1 not implemented")
}
func (UnimplementedOcpSuggestionApiServer) UpdateSuggestionV1(context.Context, *UpdateSuggestionV1Request) (*UpdateSuggestionV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSuggestionV1 not implemented")
}
func (UnimplementedOcpSuggestionApiServer) RemoveSuggestionV1(context.Context, *RemoveSuggestionV1Request) (*RemoveSuggestionV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSuggestionV1 not implemented")
}
func (UnimplementedOcpSuggestionApiServer) mustEmbedUnimplementedOcpSuggestionApiServer() {}

// UnsafeOcpSuggestionApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpSuggestionApiServer will
// result in compilation errors.
type UnsafeOcpSuggestionApiServer interface {
	mustEmbedUnimplementedOcpSuggestionApiServer()
}

func RegisterOcpSuggestionApiServer(s grpc.ServiceRegistrar, srv OcpSuggestionApiServer) {
	s.RegisterService(&OcpSuggestionApi_ServiceDesc, srv)
}

func _OcpSuggestionApi_CreateSuggestionV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSuggestionV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSuggestionApiServer).CreateSuggestionV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.suggestion.api.OcpSuggestionApi/CreateSuggestionV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSuggestionApiServer).CreateSuggestionV1(ctx, req.(*CreateSuggestionV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSuggestionApi_DescribeSuggestionV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSuggestionV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSuggestionApiServer).DescribeSuggestionV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.suggestion.api.OcpSuggestionApi/DescribeSuggestionV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSuggestionApiServer).DescribeSuggestionV1(ctx, req.(*DescribeSuggestionV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSuggestionApi_ListSuggestionV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSuggestionV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSuggestionApiServer).ListSuggestionV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.suggestion.api.OcpSuggestionApi/ListSuggestionV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSuggestionApiServer).ListSuggestionV1(ctx, req.(*ListSuggestionV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSuggestionApi_UpdateSuggestionV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSuggestionV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSuggestionApiServer).UpdateSuggestionV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.suggestion.api.OcpSuggestionApi/UpdateSuggestionV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSuggestionApiServer).UpdateSuggestionV1(ctx, req.(*UpdateSuggestionV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSuggestionApi_RemoveSuggestionV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSuggestionV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSuggestionApiServer).RemoveSuggestionV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.suggestion.api.OcpSuggestionApi/RemoveSuggestionV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSuggestionApiServer).RemoveSuggestionV1(ctx, req.(*RemoveSuggestionV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpSuggestionApi_ServiceDesc is the grpc.ServiceDesc for OcpSuggestionApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpSuggestionApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.suggestion.api.OcpSuggestionApi",
	HandlerType: (*OcpSuggestionApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSuggestionV1",
			Handler:    _OcpSuggestionApi_CreateSuggestionV1_Handler,
		},
		{
			MethodName: "DescribeSuggestionV1",
			Handler:    _OcpSuggestionApi_DescribeSuggestionV1_Handler,
		},
		{
			MethodName: "ListSuggestionV1",
			Handler:    _OcpSuggestionApi_ListSuggestionV1_Handler,
		},
		{
			MethodName: "UpdateSuggestionV1",
			Handler:    _OcpSuggestionApi_UpdateSuggestionV1_Handler,
		},
		{
			MethodName: "RemoveSuggestionV1",
			Handler:    _OcpSuggestionApi_RemoveSuggestionV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-suggestion-api/ocp-suggestion-api.proto",
}