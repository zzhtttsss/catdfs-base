// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: Get.proto

package pb

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

// MasterGetServiceClient is the client API for MasterGetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterGetServiceClient interface {
	// CheckAndGet Called by client.
	// Check whether the path and file name entered by the user in the Add operation are legal.
	CheckAndGet(ctx context.Context, in *CheckAndGetArgs, opts ...grpc.CallOption) (*CheckAndGetReply, error)
	// GetDataNodes4Get Called by client.
	// Allocate some DataNode to store a Chunk and select the primary DataNode
	GetDataNodes4Get(ctx context.Context, in *GetDataNodes4GetArgs, opts ...grpc.CallOption) (*GetDataNodes4GetReply, error)
	// SetupStream2DataNode Called by client.
	// Set up the stream between client and chunkserver, then chunkserver call rpc send data
	SetupStream2DataNode(ctx context.Context, in *SetupStream2DataNodeArgs, opts ...grpc.CallOption) (*SetupStream2DataNodeReply, error)
	// ReleaseLease4Add Called by client.
	// Release the lease of a chunk.
	ReleaseLease4Get(ctx context.Context, in *ReleaseLease4GetArgs, opts ...grpc.CallOption) (*ReleaseLease4GetReply, error)
}

type masterGetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterGetServiceClient(cc grpc.ClientConnInterface) MasterGetServiceClient {
	return &masterGetServiceClient{cc}
}

func (c *masterGetServiceClient) CheckAndGet(ctx context.Context, in *CheckAndGetArgs, opts ...grpc.CallOption) (*CheckAndGetReply, error) {
	out := new(CheckAndGetReply)
	err := c.cc.Invoke(ctx, "/pb.MasterGetService/CheckAndGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGetServiceClient) GetDataNodes4Get(ctx context.Context, in *GetDataNodes4GetArgs, opts ...grpc.CallOption) (*GetDataNodes4GetReply, error) {
	out := new(GetDataNodes4GetReply)
	err := c.cc.Invoke(ctx, "/pb.MasterGetService/GetDataNodes4Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGetServiceClient) SetupStream2DataNode(ctx context.Context, in *SetupStream2DataNodeArgs, opts ...grpc.CallOption) (*SetupStream2DataNodeReply, error) {
	out := new(SetupStream2DataNodeReply)
	err := c.cc.Invoke(ctx, "/pb.MasterGetService/SetupStream2DataNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGetServiceClient) ReleaseLease4Get(ctx context.Context, in *ReleaseLease4GetArgs, opts ...grpc.CallOption) (*ReleaseLease4GetReply, error) {
	out := new(ReleaseLease4GetReply)
	err := c.cc.Invoke(ctx, "/pb.MasterGetService/ReleaseLease4Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterGetServiceServer is the server API for MasterGetService service.
// All implementations must embed UnimplementedMasterGetServiceServer
// for forward compatibility
type MasterGetServiceServer interface {
	// CheckAndGet Called by client.
	// Check whether the path and file name entered by the user in the Add operation are legal.
	CheckAndGet(context.Context, *CheckAndGetArgs) (*CheckAndGetReply, error)
	// GetDataNodes4Get Called by client.
	// Allocate some DataNode to store a Chunk and select the primary DataNode
	GetDataNodes4Get(context.Context, *GetDataNodes4GetArgs) (*GetDataNodes4GetReply, error)
	// SetupStream2DataNode Called by client.
	// Set up the stream between client and chunkserver, then chunkserver call rpc send data
	SetupStream2DataNode(context.Context, *SetupStream2DataNodeArgs) (*SetupStream2DataNodeReply, error)
	// ReleaseLease4Add Called by client.
	// Release the lease of a chunk.
	ReleaseLease4Get(context.Context, *ReleaseLease4GetArgs) (*ReleaseLease4GetReply, error)
	mustEmbedUnimplementedMasterGetServiceServer()
}

// UnimplementedMasterGetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterGetServiceServer struct {
}

func (UnimplementedMasterGetServiceServer) CheckAndGet(context.Context, *CheckAndGetArgs) (*CheckAndGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAndGet not implemented")
}
func (UnimplementedMasterGetServiceServer) GetDataNodes4Get(context.Context, *GetDataNodes4GetArgs) (*GetDataNodes4GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataNodes4Get not implemented")
}
func (UnimplementedMasterGetServiceServer) SetupStream2DataNode(context.Context, *SetupStream2DataNodeArgs) (*SetupStream2DataNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupStream2DataNode not implemented")
}
func (UnimplementedMasterGetServiceServer) ReleaseLease4Get(context.Context, *ReleaseLease4GetArgs) (*ReleaseLease4GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseLease4Get not implemented")
}
func (UnimplementedMasterGetServiceServer) mustEmbedUnimplementedMasterGetServiceServer() {}

// UnsafeMasterGetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterGetServiceServer will
// result in compilation errors.
type UnsafeMasterGetServiceServer interface {
	mustEmbedUnimplementedMasterGetServiceServer()
}

func RegisterMasterGetServiceServer(s grpc.ServiceRegistrar, srv MasterGetServiceServer) {
	s.RegisterService(&MasterGetService_ServiceDesc, srv)
}

func _MasterGetService_CheckAndGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAndGetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGetServiceServer).CheckAndGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterGetService/CheckAndGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGetServiceServer).CheckAndGet(ctx, req.(*CheckAndGetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGetService_GetDataNodes4Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataNodes4GetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGetServiceServer).GetDataNodes4Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterGetService/GetDataNodes4Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGetServiceServer).GetDataNodes4Get(ctx, req.(*GetDataNodes4GetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGetService_SetupStream2DataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupStream2DataNodeArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGetServiceServer).SetupStream2DataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterGetService/SetupStream2DataNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGetServiceServer).SetupStream2DataNode(ctx, req.(*SetupStream2DataNodeArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGetService_ReleaseLease4Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseLease4GetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGetServiceServer).ReleaseLease4Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterGetService/ReleaseLease4Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGetServiceServer).ReleaseLease4Get(ctx, req.(*ReleaseLease4GetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterGetService_ServiceDesc is the grpc.ServiceDesc for MasterGetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterGetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MasterGetService",
	HandlerType: (*MasterGetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAndGet",
			Handler:    _MasterGetService_CheckAndGet_Handler,
		},
		{
			MethodName: "GetDataNodes4Get",
			Handler:    _MasterGetService_GetDataNodes4Get_Handler,
		},
		{
			MethodName: "SetupStream2DataNode",
			Handler:    _MasterGetService_SetupStream2DataNode_Handler,
		},
		{
			MethodName: "ReleaseLease4Get",
			Handler:    _MasterGetService_ReleaseLease4Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Get.proto",
}
