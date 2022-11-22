// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: communication.proto

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

// BlockchainServiceClient is the client API for BlockchainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainServiceClient interface {
	//用于广播块、交易、交易输入、交易输出和创世区块生成者
	Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	Transactions(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_TransactionsClient, error)
	Inputs(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_InputsClient, error)
	Outputs(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_OutputsClient, error)
	End(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error)
	CreateBy(ctx context.Context, in *CreateByRequest, opts ...grpc.CallOption) (*CreateByResponse, error)
	//用于请求获得所有块
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*CreateByResponse, error)
	//用于广播新交易
	NewTransaction(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TxsResponse, error)
	NewOutPuts(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_NewOutPutsClient, error)
	NewInPuts(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_NewInPutsClient, error)
	NewEnd(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error)
	//用于探询结点是否在线
	Alive(ctx context.Context, in *AliveRequest, opts ...grpc.CallOption) (*AliveResponse, error)
	//广播B币
	Asset(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error)
}

type blockchainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainServiceClient(cc grpc.ClientConnInterface) BlockchainServiceClient {
	return &blockchainServiceClient{cc}
}

func (c *blockchainServiceClient) Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Transactions(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_TransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[0], "/communication.BlockchainService/Transactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceTransactionsClient{stream}
	return x, nil
}

type BlockchainService_TransactionsClient interface {
	Send(*TransactionsRequest) error
	CloseAndRecv() (*TxsResponse, error)
	grpc.ClientStream
}

type blockchainServiceTransactionsClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceTransactionsClient) Send(m *TransactionsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockchainServiceTransactionsClient) CloseAndRecv() (*TxsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TxsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) Inputs(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_InputsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[1], "/communication.BlockchainService/Inputs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceInputsClient{stream}
	return x, nil
}

type BlockchainService_InputsClient interface {
	Send(*InputsRequest) error
	CloseAndRecv() (*InputsResponse, error)
	grpc.ClientStream
}

type blockchainServiceInputsClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceInputsClient) Send(m *InputsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockchainServiceInputsClient) CloseAndRecv() (*InputsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InputsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) Outputs(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_OutputsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[2], "/communication.BlockchainService/Outputs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceOutputsClient{stream}
	return x, nil
}

type BlockchainService_OutputsClient interface {
	Send(*OutputsRequest) error
	CloseAndRecv() (*OutputsResponse, error)
	grpc.ClientStream
}

type blockchainServiceOutputsClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceOutputsClient) Send(m *OutputsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockchainServiceOutputsClient) CloseAndRecv() (*OutputsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OutputsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) End(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error) {
	out := new(EndResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/End", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) CreateBy(ctx context.Context, in *CreateByRequest, opts ...grpc.CallOption) (*CreateByResponse, error) {
	out := new(CreateByResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/CreateBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*CreateByResponse, error) {
	out := new(CreateByResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) NewTransaction(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (*TxsResponse, error) {
	out := new(TxsResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/NewTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) NewOutPuts(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_NewOutPutsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[3], "/communication.BlockchainService/NewOutPuts", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceNewOutPutsClient{stream}
	return x, nil
}

type BlockchainService_NewOutPutsClient interface {
	Send(*OutputsRequest) error
	CloseAndRecv() (*OutputsResponse, error)
	grpc.ClientStream
}

type blockchainServiceNewOutPutsClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceNewOutPutsClient) Send(m *OutputsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockchainServiceNewOutPutsClient) CloseAndRecv() (*OutputsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OutputsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) NewInPuts(ctx context.Context, opts ...grpc.CallOption) (BlockchainService_NewInPutsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockchainService_ServiceDesc.Streams[4], "/communication.BlockchainService/NewInPuts", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockchainServiceNewInPutsClient{stream}
	return x, nil
}

type BlockchainService_NewInPutsClient interface {
	Send(*InputsRequest) error
	CloseAndRecv() (*InputsResponse, error)
	grpc.ClientStream
}

type blockchainServiceNewInPutsClient struct {
	grpc.ClientStream
}

func (x *blockchainServiceNewInPutsClient) Send(m *InputsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockchainServiceNewInPutsClient) CloseAndRecv() (*InputsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InputsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blockchainServiceClient) NewEnd(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error) {
	out := new(EndResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/NewEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Alive(ctx context.Context, in *AliveRequest, opts ...grpc.CallOption) (*AliveResponse, error) {
	out := new(AliveResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/Alive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainServiceClient) Asset(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error) {
	out := new(AssetResponse)
	err := c.cc.Invoke(ctx, "/communication.BlockchainService/Asset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServiceServer is the server API for BlockchainService service.
// All implementations must embed UnimplementedBlockchainServiceServer
// for forward compatibility
type BlockchainServiceServer interface {
	//用于广播块、交易、交易输入、交易输出和创世区块生成者
	Block(context.Context, *BlockRequest) (*BlockResponse, error)
	Transactions(BlockchainService_TransactionsServer) error
	Inputs(BlockchainService_InputsServer) error
	Outputs(BlockchainService_OutputsServer) error
	End(context.Context, *EndRequest) (*EndResponse, error)
	CreateBy(context.Context, *CreateByRequest) (*CreateByResponse, error)
	//用于请求获得所有块
	GetBlock(context.Context, *GetBlockRequest) (*CreateByResponse, error)
	//用于广播新交易
	NewTransaction(context.Context, *TransactionsRequest) (*TxsResponse, error)
	NewOutPuts(BlockchainService_NewOutPutsServer) error
	NewInPuts(BlockchainService_NewInPutsServer) error
	NewEnd(context.Context, *EndRequest) (*EndResponse, error)
	//用于探询结点是否在线
	Alive(context.Context, *AliveRequest) (*AliveResponse, error)
	//广播B币
	Asset(context.Context, *AssetRequest) (*AssetResponse, error)
	mustEmbedUnimplementedBlockchainServiceServer()
}

// UnimplementedBlockchainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServiceServer struct {
}

func (UnimplementedBlockchainServiceServer) Block(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedBlockchainServiceServer) Transactions(BlockchainService_TransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method Transactions not implemented")
}
func (UnimplementedBlockchainServiceServer) Inputs(BlockchainService_InputsServer) error {
	return status.Errorf(codes.Unimplemented, "method Inputs not implemented")
}
func (UnimplementedBlockchainServiceServer) Outputs(BlockchainService_OutputsServer) error {
	return status.Errorf(codes.Unimplemented, "method Outputs not implemented")
}
func (UnimplementedBlockchainServiceServer) End(context.Context, *EndRequest) (*EndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method End not implemented")
}
func (UnimplementedBlockchainServiceServer) CreateBy(context.Context, *CreateByRequest) (*CreateByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBy not implemented")
}
func (UnimplementedBlockchainServiceServer) GetBlock(context.Context, *GetBlockRequest) (*CreateByResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockchainServiceServer) NewTransaction(context.Context, *TransactionsRequest) (*TxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTransaction not implemented")
}
func (UnimplementedBlockchainServiceServer) NewOutPuts(BlockchainService_NewOutPutsServer) error {
	return status.Errorf(codes.Unimplemented, "method NewOutPuts not implemented")
}
func (UnimplementedBlockchainServiceServer) NewInPuts(BlockchainService_NewInPutsServer) error {
	return status.Errorf(codes.Unimplemented, "method NewInPuts not implemented")
}
func (UnimplementedBlockchainServiceServer) NewEnd(context.Context, *EndRequest) (*EndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEnd not implemented")
}
func (UnimplementedBlockchainServiceServer) Alive(context.Context, *AliveRequest) (*AliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedBlockchainServiceServer) Asset(context.Context, *AssetRequest) (*AssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Asset not implemented")
}
func (UnimplementedBlockchainServiceServer) mustEmbedUnimplementedBlockchainServiceServer() {}

// UnsafeBlockchainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServiceServer will
// result in compilation errors.
type UnsafeBlockchainServiceServer interface {
	mustEmbedUnimplementedBlockchainServiceServer()
}

func RegisterBlockchainServiceServer(s grpc.ServiceRegistrar, srv BlockchainServiceServer) {
	s.RegisterService(&BlockchainService_ServiceDesc, srv)
}

func _BlockchainService_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Block(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Transactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockchainServiceServer).Transactions(&blockchainServiceTransactionsServer{stream})
}

type BlockchainService_TransactionsServer interface {
	SendAndClose(*TxsResponse) error
	Recv() (*TransactionsRequest, error)
	grpc.ServerStream
}

type blockchainServiceTransactionsServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceTransactionsServer) SendAndClose(m *TxsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockchainServiceTransactionsServer) Recv() (*TransactionsRequest, error) {
	m := new(TransactionsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockchainService_Inputs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockchainServiceServer).Inputs(&blockchainServiceInputsServer{stream})
}

type BlockchainService_InputsServer interface {
	SendAndClose(*InputsResponse) error
	Recv() (*InputsRequest, error)
	grpc.ServerStream
}

type blockchainServiceInputsServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceInputsServer) SendAndClose(m *InputsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockchainServiceInputsServer) Recv() (*InputsRequest, error) {
	m := new(InputsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockchainService_Outputs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockchainServiceServer).Outputs(&blockchainServiceOutputsServer{stream})
}

type BlockchainService_OutputsServer interface {
	SendAndClose(*OutputsResponse) error
	Recv() (*OutputsRequest, error)
	grpc.ServerStream
}

type blockchainServiceOutputsServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceOutputsServer) SendAndClose(m *OutputsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockchainServiceOutputsServer) Recv() (*OutputsRequest, error) {
	m := new(OutputsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockchainService_End_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).End(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/End",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).End(ctx, req.(*EndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_CreateBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).CreateBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/CreateBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).CreateBy(ctx, req.(*CreateByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_NewTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).NewTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/NewTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).NewTransaction(ctx, req.(*TransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_NewOutPuts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockchainServiceServer).NewOutPuts(&blockchainServiceNewOutPutsServer{stream})
}

type BlockchainService_NewOutPutsServer interface {
	SendAndClose(*OutputsResponse) error
	Recv() (*OutputsRequest, error)
	grpc.ServerStream
}

type blockchainServiceNewOutPutsServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceNewOutPutsServer) SendAndClose(m *OutputsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockchainServiceNewOutPutsServer) Recv() (*OutputsRequest, error) {
	m := new(OutputsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockchainService_NewInPuts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockchainServiceServer).NewInPuts(&blockchainServiceNewInPutsServer{stream})
}

type BlockchainService_NewInPutsServer interface {
	SendAndClose(*InputsResponse) error
	Recv() (*InputsRequest, error)
	grpc.ServerStream
}

type blockchainServiceNewInPutsServer struct {
	grpc.ServerStream
}

func (x *blockchainServiceNewInPutsServer) SendAndClose(m *InputsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockchainServiceNewInPutsServer) Recv() (*InputsRequest, error) {
	m := new(InputsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlockchainService_NewEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).NewEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/NewEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).NewEnd(ctx, req.(*EndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/Alive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Alive(ctx, req.(*AliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainService_Asset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServiceServer).Asset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.BlockchainService/Asset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServiceServer).Asset(ctx, req.(*AssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockchainService_ServiceDesc is the grpc.ServiceDesc for BlockchainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communication.BlockchainService",
	HandlerType: (*BlockchainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Block",
			Handler:    _BlockchainService_Block_Handler,
		},
		{
			MethodName: "End",
			Handler:    _BlockchainService_End_Handler,
		},
		{
			MethodName: "CreateBy",
			Handler:    _BlockchainService_CreateBy_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BlockchainService_GetBlock_Handler,
		},
		{
			MethodName: "NewTransaction",
			Handler:    _BlockchainService_NewTransaction_Handler,
		},
		{
			MethodName: "NewEnd",
			Handler:    _BlockchainService_NewEnd_Handler,
		},
		{
			MethodName: "Alive",
			Handler:    _BlockchainService_Alive_Handler,
		},
		{
			MethodName: "Asset",
			Handler:    _BlockchainService_Asset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transactions",
			Handler:       _BlockchainService_Transactions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Inputs",
			Handler:       _BlockchainService_Inputs_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Outputs",
			Handler:       _BlockchainService_Outputs_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "NewOutPuts",
			Handler:       _BlockchainService_NewOutPuts_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "NewInPuts",
			Handler:       _BlockchainService_NewInPuts_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "communication.proto",
}