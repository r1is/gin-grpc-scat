// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pb/scat.proto

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

// AnalyzeToolsClient is the client API for AnalyzeTools service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzeToolsClient interface {
	CheckSmartContractCode(ctx context.Context, in *SourceCodeRequest, opts ...grpc.CallOption) (*CheckResult, error)
}

type analyzeToolsClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzeToolsClient(cc grpc.ClientConnInterface) AnalyzeToolsClient {
	return &analyzeToolsClient{cc}
}

func (c *analyzeToolsClient) CheckSmartContractCode(ctx context.Context, in *SourceCodeRequest, opts ...grpc.CallOption) (*CheckResult, error) {
	out := new(CheckResult)
	err := c.cc.Invoke(ctx, "/pb.AnalyzeTools/checkSmartContractCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzeToolsServer is the server API for AnalyzeTools service.
// All implementations must embed UnimplementedAnalyzeToolsServer
// for forward compatibility
type AnalyzeToolsServer interface {
	CheckSmartContractCode(context.Context, *SourceCodeRequest) (*CheckResult, error)
	mustEmbedUnimplementedAnalyzeToolsServer()
}

// UnimplementedAnalyzeToolsServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyzeToolsServer struct {
}

func (UnimplementedAnalyzeToolsServer) CheckSmartContractCode(context.Context, *SourceCodeRequest) (*CheckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSmartContractCode not implemented")
}
func (UnimplementedAnalyzeToolsServer) mustEmbedUnimplementedAnalyzeToolsServer() {}

// UnsafeAnalyzeToolsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzeToolsServer will
// result in compilation errors.
type UnsafeAnalyzeToolsServer interface {
	mustEmbedUnimplementedAnalyzeToolsServer()
}

func RegisterAnalyzeToolsServer(s grpc.ServiceRegistrar, srv AnalyzeToolsServer) {
	s.RegisterService(&AnalyzeTools_ServiceDesc, srv)
}

func _AnalyzeTools_CheckSmartContractCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzeToolsServer).CheckSmartContractCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AnalyzeTools/checkSmartContractCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzeToolsServer).CheckSmartContractCode(ctx, req.(*SourceCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyzeTools_ServiceDesc is the grpc.ServiceDesc for AnalyzeTools service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyzeTools_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AnalyzeTools",
	HandlerType: (*AnalyzeToolsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkSmartContractCode",
			Handler:    _AnalyzeTools_CheckSmartContractCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/scat.proto",
}
