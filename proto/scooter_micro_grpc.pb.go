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

// ScooterServiceClient is the client API for ScooterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScooterServiceClient interface {
	Register(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ScooterService_RegisterClient, error)
	Receive(ctx context.Context, opts ...grpc.CallOption) (ScooterService_ReceiveClient, error)
	GetAllScooters(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ScooterList, error)
	GetAllScootersByStationID(ctx context.Context, in *StationID, opts ...grpc.CallOption) (*ScooterList, error)
	GetScooterById(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*Scooter, error)
	GetScooterStatus(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*ScooterStatus, error)
	SendCurrentStatus(ctx context.Context, in *SendStatus, opts ...grpc.CallOption) (*Response, error)
	CreateScooterStatusInRent(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*ScooterStatusInRent, error)
	GetStationByID(ctx context.Context, in *StationID, opts ...grpc.CallOption) (*Station, error)
	GetAllStations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*StationList, error)
}

type scooterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScooterServiceClient(cc grpc.ClientConnInterface) ScooterServiceClient {
	return &scooterServiceClient{cc}
}

func (c *scooterServiceClient) Register(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (ScooterService_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScooterService_ServiceDesc.Streams[0], "/proto.ScooterService/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &scooterServiceRegisterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScooterService_RegisterClient interface {
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type scooterServiceRegisterClient struct {
	grpc.ClientStream
}

func (x *scooterServiceRegisterClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scooterServiceClient) Receive(ctx context.Context, opts ...grpc.CallOption) (ScooterService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScooterService_ServiceDesc.Streams[1], "/proto.ScooterService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &scooterServiceReceiveClient{stream}
	return x, nil
}

type ScooterService_ReceiveClient interface {
	Send(*ClientMessage) error
	CloseAndRecv() (*ServerMessage, error)
	grpc.ClientStream
}

type scooterServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *scooterServiceReceiveClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scooterServiceReceiveClient) CloseAndRecv() (*ServerMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scooterServiceClient) GetAllScooters(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ScooterList, error) {
	out := new(ScooterList)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetAllScooters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) GetAllScootersByStationID(ctx context.Context, in *StationID, opts ...grpc.CallOption) (*ScooterList, error) {
	out := new(ScooterList)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetAllScootersByStationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) GetScooterById(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*Scooter, error) {
	out := new(Scooter)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetScooterById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) GetScooterStatus(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*ScooterStatus, error) {
	out := new(ScooterStatus)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetScooterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) SendCurrentStatus(ctx context.Context, in *SendStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/SendCurrentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) CreateScooterStatusInRent(ctx context.Context, in *ScooterID, opts ...grpc.CallOption) (*ScooterStatusInRent, error) {
	out := new(ScooterStatusInRent)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/CreateScooterStatusInRent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) GetStationByID(ctx context.Context, in *StationID, opts ...grpc.CallOption) (*Station, error) {
	out := new(Station)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetStationByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scooterServiceClient) GetAllStations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*StationList, error) {
	out := new(StationList)
	err := c.cc.Invoke(ctx, "/proto.ScooterService/GetAllStations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScooterServiceServer is the server API for ScooterService service.
// All implementations must embed UnimplementedScooterServiceServer
// for forward compatibility
type ScooterServiceServer interface {
	Register(*ClientRequest, ScooterService_RegisterServer) error
	Receive(ScooterService_ReceiveServer) error
	GetAllScooters(context.Context, *Request) (*ScooterList, error)
	GetAllScootersByStationID(context.Context, *StationID) (*ScooterList, error)
	GetScooterById(context.Context, *ScooterID) (*Scooter, error)
	GetScooterStatus(context.Context, *ScooterID) (*ScooterStatus, error)
	SendCurrentStatus(context.Context, *SendStatus) (*Response, error)
	CreateScooterStatusInRent(context.Context, *ScooterID) (*ScooterStatusInRent, error)
	GetStationByID(context.Context, *StationID) (*Station, error)
	GetAllStations(context.Context, *Request) (*StationList, error)
	mustEmbedUnimplementedScooterServiceServer()
}

// UnimplementedScooterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScooterServiceServer struct {
}

func (UnimplementedScooterServiceServer) Register(*ClientRequest, ScooterService_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedScooterServiceServer) Receive(ScooterService_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedScooterServiceServer) GetAllScooters(context.Context, *Request) (*ScooterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllScooters not implemented")
}
func (UnimplementedScooterServiceServer) GetAllScootersByStationID(context.Context, *StationID) (*ScooterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllScootersByStationID not implemented")
}
func (UnimplementedScooterServiceServer) GetScooterById(context.Context, *ScooterID) (*Scooter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScooterById not implemented")
}
func (UnimplementedScooterServiceServer) GetScooterStatus(context.Context, *ScooterID) (*ScooterStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScooterStatus not implemented")
}
func (UnimplementedScooterServiceServer) SendCurrentStatus(context.Context, *SendStatus) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCurrentStatus not implemented")
}
func (UnimplementedScooterServiceServer) CreateScooterStatusInRent(context.Context, *ScooterID) (*ScooterStatusInRent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScooterStatusInRent not implemented")
}
func (UnimplementedScooterServiceServer) GetStationByID(context.Context, *StationID) (*Station, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStationByID not implemented")
}
func (UnimplementedScooterServiceServer) GetAllStations(context.Context, *Request) (*StationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStations not implemented")
}
func (UnimplementedScooterServiceServer) mustEmbedUnimplementedScooterServiceServer() {}

// UnsafeScooterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScooterServiceServer will
// result in compilation errors.
type UnsafeScooterServiceServer interface {
	mustEmbedUnimplementedScooterServiceServer()
}

func RegisterScooterServiceServer(s grpc.ServiceRegistrar, srv ScooterServiceServer) {
	s.RegisterService(&ScooterService_ServiceDesc, srv)
}

func _ScooterService_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScooterServiceServer).Register(m, &scooterServiceRegisterServer{stream})
}

type ScooterService_RegisterServer interface {
	Send(*ServerMessage) error
	grpc.ServerStream
}

type scooterServiceRegisterServer struct {
	grpc.ServerStream
}

func (x *scooterServiceRegisterServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ScooterService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScooterServiceServer).Receive(&scooterServiceReceiveServer{stream})
}

type ScooterService_ReceiveServer interface {
	SendAndClose(*ServerMessage) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type scooterServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *scooterServiceReceiveServer) SendAndClose(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scooterServiceReceiveServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScooterService_GetAllScooters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetAllScooters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetAllScooters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetAllScooters(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_GetAllScootersByStationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetAllScootersByStationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetAllScootersByStationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetAllScootersByStationID(ctx, req.(*StationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_GetScooterById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScooterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetScooterById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetScooterById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetScooterById(ctx, req.(*ScooterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_GetScooterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScooterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetScooterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetScooterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetScooterStatus(ctx, req.(*ScooterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_SendCurrentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).SendCurrentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/SendCurrentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).SendCurrentStatus(ctx, req.(*SendStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_CreateScooterStatusInRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScooterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).CreateScooterStatusInRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/CreateScooterStatusInRent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).CreateScooterStatusInRent(ctx, req.(*ScooterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_GetStationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetStationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetStationByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetStationByID(ctx, req.(*StationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScooterService_GetAllStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScooterServiceServer).GetAllStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ScooterService/GetAllStations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScooterServiceServer).GetAllStations(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ScooterService_ServiceDesc is the grpc.ServiceDesc for ScooterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScooterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ScooterService",
	HandlerType: (*ScooterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllScooters",
			Handler:    _ScooterService_GetAllScooters_Handler,
		},
		{
			MethodName: "GetAllScootersByStationID",
			Handler:    _ScooterService_GetAllScootersByStationID_Handler,
		},
		{
			MethodName: "GetScooterById",
			Handler:    _ScooterService_GetScooterById_Handler,
		},
		{
			MethodName: "GetScooterStatus",
			Handler:    _ScooterService_GetScooterStatus_Handler,
		},
		{
			MethodName: "SendCurrentStatus",
			Handler:    _ScooterService_SendCurrentStatus_Handler,
		},
		{
			MethodName: "CreateScooterStatusInRent",
			Handler:    _ScooterService_CreateScooterStatusInRent_Handler,
		},
		{
			MethodName: "GetStationByID",
			Handler:    _ScooterService_GetStationByID_Handler,
		},
		{
			MethodName: "GetAllStations",
			Handler:    _ScooterService_GetAllStations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ScooterService_Register_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receive",
			Handler:       _ScooterService_Receive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "scooter_service/proto/scooter_micro.proto",
}
