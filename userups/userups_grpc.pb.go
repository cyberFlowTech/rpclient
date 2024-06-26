// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: define/userups.proto

package userups

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
	UserUps_GetUserPoints_FullMethodName       = "/userups.UserUps/getUserPoints"
	UserUps_GetUserPointDetails_FullMethodName = "/userups.UserUps/getUserPointDetails"
	UserUps_AddUserPoints_FullMethodName       = "/userups.UserUps/addUserPoints"
	UserUps_GetUserTask_FullMethodName         = "/userups.UserUps/getUserTask"
)

// UserUpsClient is the client API for UserUps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserUpsClient interface {
	// 用户积分查询，包括积分，等级
	GetUserPoints(ctx context.Context, in *GetUserPointsReq, opts ...grpc.CallOption) (*GetUserPointsResp, error)
	// 积分明细查询，包括增加，销毁等
	GetUserPointDetails(ctx context.Context, in *GetUserPointDetailsReq, opts ...grpc.CallOption) (*GetUserPointDetailsResp, error)
	// 用户积分信息修改，添加一条明显。并修改记录更新前后的积分情况
	AddUserPoints(ctx context.Context, in *SetUserPointReq, opts ...grpc.CallOption) (*SetUserPointResp, error)
	// 查询当前任务（根据状态分类 新手，日常）
	GetUserTask(ctx context.Context, in *GetUserTaskReq, opts ...grpc.CallOption) (*GetUserTaskResp, error)
}

type userUpsClient struct {
	cc grpc.ClientConnInterface
}

func NewUserUpsClient(cc grpc.ClientConnInterface) UserUpsClient {
	return &userUpsClient{cc}
}

func (c *userUpsClient) GetUserPoints(ctx context.Context, in *GetUserPointsReq, opts ...grpc.CallOption) (*GetUserPointsResp, error) {
	out := new(GetUserPointsResp)
	err := c.cc.Invoke(ctx, UserUps_GetUserPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userUpsClient) GetUserPointDetails(ctx context.Context, in *GetUserPointDetailsReq, opts ...grpc.CallOption) (*GetUserPointDetailsResp, error) {
	out := new(GetUserPointDetailsResp)
	err := c.cc.Invoke(ctx, UserUps_GetUserPointDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userUpsClient) AddUserPoints(ctx context.Context, in *SetUserPointReq, opts ...grpc.CallOption) (*SetUserPointResp, error) {
	out := new(SetUserPointResp)
	err := c.cc.Invoke(ctx, UserUps_AddUserPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userUpsClient) GetUserTask(ctx context.Context, in *GetUserTaskReq, opts ...grpc.CallOption) (*GetUserTaskResp, error) {
	out := new(GetUserTaskResp)
	err := c.cc.Invoke(ctx, UserUps_GetUserTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserUpsServer is the server API for UserUps service.
// All implementations must embed UnimplementedUserUpsServer
// for forward compatibility
type UserUpsServer interface {
	// 用户积分查询，包括积分，等级
	GetUserPoints(context.Context, *GetUserPointsReq) (*GetUserPointsResp, error)
	// 积分明细查询，包括增加，销毁等
	GetUserPointDetails(context.Context, *GetUserPointDetailsReq) (*GetUserPointDetailsResp, error)
	// 用户积分信息修改，添加一条明显。并修改记录更新前后的积分情况
	AddUserPoints(context.Context, *SetUserPointReq) (*SetUserPointResp, error)
	// 查询当前任务（根据状态分类 新手，日常）
	GetUserTask(context.Context, *GetUserTaskReq) (*GetUserTaskResp, error)
	mustEmbedUnimplementedUserUpsServer()
}

// UnimplementedUserUpsServer must be embedded to have forward compatible implementations.
type UnimplementedUserUpsServer struct {
}

func (UnimplementedUserUpsServer) GetUserPoints(context.Context, *GetUserPointsReq) (*GetUserPointsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPoints not implemented")
}
func (UnimplementedUserUpsServer) GetUserPointDetails(context.Context, *GetUserPointDetailsReq) (*GetUserPointDetailsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPointDetails not implemented")
}
func (UnimplementedUserUpsServer) AddUserPoints(context.Context, *SetUserPointReq) (*SetUserPointResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserPoints not implemented")
}
func (UnimplementedUserUpsServer) GetUserTask(context.Context, *GetUserTaskReq) (*GetUserTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTask not implemented")
}
func (UnimplementedUserUpsServer) mustEmbedUnimplementedUserUpsServer() {}

// UnsafeUserUpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserUpsServer will
// result in compilation errors.
type UnsafeUserUpsServer interface {
	mustEmbedUnimplementedUserUpsServer()
}

func RegisterUserUpsServer(s grpc.ServiceRegistrar, srv UserUpsServer) {
	s.RegisterService(&UserUps_ServiceDesc, srv)
}

func _UserUps_GetUserPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPointsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserUpsServer).GetUserPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserUps_GetUserPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserUpsServer).GetUserPoints(ctx, req.(*GetUserPointsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserUps_GetUserPointDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPointDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserUpsServer).GetUserPointDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserUps_GetUserPointDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserUpsServer).GetUserPointDetails(ctx, req.(*GetUserPointDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserUps_AddUserPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPointReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserUpsServer).AddUserPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserUps_AddUserPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserUpsServer).AddUserPoints(ctx, req.(*SetUserPointReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserUps_GetUserTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserUpsServer).GetUserTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserUps_GetUserTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserUpsServer).GetUserTask(ctx, req.(*GetUserTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserUps_ServiceDesc is the grpc.ServiceDesc for UserUps service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserUps_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userups.UserUps",
	HandlerType: (*UserUpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserPoints",
			Handler:    _UserUps_GetUserPoints_Handler,
		},
		{
			MethodName: "getUserPointDetails",
			Handler:    _UserUps_GetUserPointDetails_Handler,
		},
		{
			MethodName: "addUserPoints",
			Handler:    _UserUps_AddUserPoints_Handler,
		},
		{
			MethodName: "getUserTask",
			Handler:    _UserUps_GetUserTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "define/userups.proto",
}
