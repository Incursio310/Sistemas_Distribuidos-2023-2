// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: Proto/messages.proto

package Proto

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

// Information_TradesClient is the client API for Information_Trades service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Information_TradesClient interface {
	Notificate(ctx context.Context, in *Continent, opts ...grpc.CallOption) (*NameNodeResponse, error)
	Saves_Name(ctx context.Context, in *NameNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error)
	Get_Name(ctx context.Context, in *NameNodeIDRequest, opts ...grpc.CallOption) (*DataNodeNamesResponse, error)
	ONU(ctx context.Context, in *ONURequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type information_TradesClient struct {
	cc grpc.ClientConnInterface
}

func NewInformation_TradesClient(cc grpc.ClientConnInterface) Information_TradesClient {
	return &information_TradesClient{cc}
}

func (c *information_TradesClient) Notificate(ctx context.Context, in *Continent, opts ...grpc.CallOption) (*NameNodeResponse, error) {
	out := new(NameNodeResponse)
	err := c.cc.Invoke(ctx, "/grpc.Information_Trades/Notificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *information_TradesClient) Saves_Name(ctx context.Context, in *NameNodeRequest, opts ...grpc.CallOption) (*DataNodeResponse, error) {
	out := new(DataNodeResponse)
	err := c.cc.Invoke(ctx, "/grpc.Information_Trades/Saves_Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *information_TradesClient) Get_Name(ctx context.Context, in *NameNodeIDRequest, opts ...grpc.CallOption) (*DataNodeNamesResponse, error) {
	out := new(DataNodeNamesResponse)
	err := c.cc.Invoke(ctx, "/grpc.Information_Trades/Get_Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *information_TradesClient) ONU(ctx context.Context, in *ONURequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/grpc.Information_Trades/ONU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Information_TradesServer is the server API for Information_Trades service.
// All implementations must embed UnimplementedInformation_TradesServer
// for forward compatibility
type Information_TradesServer interface {
	Notificate(context.Context, *Continent) (*NameNodeResponse, error)
	Saves_Name(context.Context, *NameNodeRequest) (*DataNodeResponse, error)
	Get_Name(context.Context, *NameNodeIDRequest) (*DataNodeNamesResponse, error)
	ONU(context.Context, *ONURequest) (*StatusResponse, error)
	mustEmbedUnimplementedInformation_TradesServer()
}

// UnimplementedInformation_TradesServer must be embedded to have forward compatible implementations.
type UnimplementedInformation_TradesServer struct {
}

func (UnimplementedInformation_TradesServer) Notificate(context.Context, *Continent) (*NameNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notificate not implemented")
}
func (UnimplementedInformation_TradesServer) Saves_Name(context.Context, *NameNodeRequest) (*DataNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Saves_Name not implemented")
}
func (UnimplementedInformation_TradesServer) Get_Name(context.Context, *NameNodeIDRequest) (*DataNodeNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get_Name not implemented")
}
func (UnimplementedInformation_TradesServer) ONU(context.Context, *ONURequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ONU not implemented")
}
func (UnimplementedInformation_TradesServer) mustEmbedUnimplementedInformation_TradesServer() {}

// UnsafeInformation_TradesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Information_TradesServer will
// result in compilation errors.
type UnsafeInformation_TradesServer interface {
	mustEmbedUnimplementedInformation_TradesServer()
}

func RegisterInformation_TradesServer(s grpc.ServiceRegistrar, srv Information_TradesServer) {
	s.RegisterService(&Information_Trades_ServiceDesc, srv)
}

func _Information_Trades_Notificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Continent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Information_TradesServer).Notificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Information_Trades/Notificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Information_TradesServer).Notificate(ctx, req.(*Continent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_Trades_Saves_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Information_TradesServer).Saves_Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Information_Trades/Saves_Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Information_TradesServer).Saves_Name(ctx, req.(*NameNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_Trades_Get_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameNodeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Information_TradesServer).Get_Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Information_Trades/Get_Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Information_TradesServer).Get_Name(ctx, req.(*NameNodeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_Trades_ONU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ONURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Information_TradesServer).ONU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Information_Trades/ONU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Information_TradesServer).ONU(ctx, req.(*ONURequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Information_Trades_ServiceDesc is the grpc.ServiceDesc for Information_Trades service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Information_Trades_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Information_Trades",
	HandlerType: (*Information_TradesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notificate",
			Handler:    _Information_Trades_Notificate_Handler,
		},
		{
			MethodName: "Saves_Name",
			Handler:    _Information_Trades_Saves_Name_Handler,
		},
		{
			MethodName: "Get_Name",
			Handler:    _Information_Trades_Get_Name_Handler,
		},
		{
			MethodName: "ONU",
			Handler:    _Information_Trades_ONU_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/messages.proto",
}
