// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: v1/rpc/contacts.proto

package contacts

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
	Contacts_GetFriendRequestList_FullMethodName = "/contacts.Contacts/GetFriendRequestList"
	Contacts_SetFriendRequest_FullMethodName     = "/contacts.Contacts/SetFriendRequest"
	Contacts_GetContacts_FullMethodName          = "/contacts.Contacts/GetContacts"
	Contacts_SetContact_FullMethodName           = "/contacts.Contacts/SetContact"
	Contacts_GetGroupContacts_FullMethodName     = "/contacts.Contacts/GetGroupContacts"
	Contacts_SetGroupContact_FullMethodName      = "/contacts.Contacts/SetGroupContact"
	Contacts_GetBlackList_FullMethodName         = "/contacts.Contacts/GetBlackList"
	Contacts_SetBlackList_FullMethodName         = "/contacts.Contacts/SetBlackList"
)

// ContactsClient is the client API for Contacts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactsClient interface {
	// 获取申请list
	GetFriendRequestList(ctx context.Context, in *GetFriendRequestListReq, opts ...grpc.CallOption) (*GetFriendRequestListResp, error)
	SetFriendRequest(ctx context.Context, in *SetFriendRequestReq, opts ...grpc.CallOption) (*SetFriendRequestResp, error)
	// 获取通讯录,管理通讯录
	GetContacts(ctx context.Context, in *GetContactsReq, opts ...grpc.CallOption) (*GetContactsResp, error)
	SetContact(ctx context.Context, in *SetContactReq, opts ...grpc.CallOption) (*SetContactResp, error)
	// 查询我等群聊
	GetGroupContacts(ctx context.Context, in *GetGroupContactsReq, opts ...grpc.CallOption) (*GetGroupContactsResp, error)
	SetGroupContact(ctx context.Context, in *SetGroupContactReq, opts ...grpc.CallOption) (*SetGroupContactResp, error)
	// 设置黑名单
	GetBlackList(ctx context.Context, in *GetBlackListReq, opts ...grpc.CallOption) (*GetBlackListResp, error)
	SetBlackList(ctx context.Context, in *SetBlackListReq, opts ...grpc.CallOption) (*SetBlackListResp, error)
}

type contactsClient struct {
	cc grpc.ClientConnInterface
}

func NewContactsClient(cc grpc.ClientConnInterface) ContactsClient {
	return &contactsClient{cc}
}

func (c *contactsClient) GetFriendRequestList(ctx context.Context, in *GetFriendRequestListReq, opts ...grpc.CallOption) (*GetFriendRequestListResp, error) {
	out := new(GetFriendRequestListResp)
	err := c.cc.Invoke(ctx, Contacts_GetFriendRequestList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SetFriendRequest(ctx context.Context, in *SetFriendRequestReq, opts ...grpc.CallOption) (*SetFriendRequestResp, error) {
	out := new(SetFriendRequestResp)
	err := c.cc.Invoke(ctx, Contacts_SetFriendRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) GetContacts(ctx context.Context, in *GetContactsReq, opts ...grpc.CallOption) (*GetContactsResp, error) {
	out := new(GetContactsResp)
	err := c.cc.Invoke(ctx, Contacts_GetContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SetContact(ctx context.Context, in *SetContactReq, opts ...grpc.CallOption) (*SetContactResp, error) {
	out := new(SetContactResp)
	err := c.cc.Invoke(ctx, Contacts_SetContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) GetGroupContacts(ctx context.Context, in *GetGroupContactsReq, opts ...grpc.CallOption) (*GetGroupContactsResp, error) {
	out := new(GetGroupContactsResp)
	err := c.cc.Invoke(ctx, Contacts_GetGroupContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SetGroupContact(ctx context.Context, in *SetGroupContactReq, opts ...grpc.CallOption) (*SetGroupContactResp, error) {
	out := new(SetGroupContactResp)
	err := c.cc.Invoke(ctx, Contacts_SetGroupContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) GetBlackList(ctx context.Context, in *GetBlackListReq, opts ...grpc.CallOption) (*GetBlackListResp, error) {
	out := new(GetBlackListResp)
	err := c.cc.Invoke(ctx, Contacts_GetBlackList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SetBlackList(ctx context.Context, in *SetBlackListReq, opts ...grpc.CallOption) (*SetBlackListResp, error) {
	out := new(SetBlackListResp)
	err := c.cc.Invoke(ctx, Contacts_SetBlackList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactsServer is the server API for Contacts service.
// All implementations must embed UnimplementedContactsServer
// for forward compatibility
type ContactsServer interface {
	// 获取申请list
	GetFriendRequestList(context.Context, *GetFriendRequestListReq) (*GetFriendRequestListResp, error)
	SetFriendRequest(context.Context, *SetFriendRequestReq) (*SetFriendRequestResp, error)
	// 获取通讯录,管理通讯录
	GetContacts(context.Context, *GetContactsReq) (*GetContactsResp, error)
	SetContact(context.Context, *SetContactReq) (*SetContactResp, error)
	// 查询我等群聊
	GetGroupContacts(context.Context, *GetGroupContactsReq) (*GetGroupContactsResp, error)
	SetGroupContact(context.Context, *SetGroupContactReq) (*SetGroupContactResp, error)
	// 设置黑名单
	GetBlackList(context.Context, *GetBlackListReq) (*GetBlackListResp, error)
	SetBlackList(context.Context, *SetBlackListReq) (*SetBlackListResp, error)
	mustEmbedUnimplementedContactsServer()
}

// UnimplementedContactsServer must be embedded to have forward compatible implementations.
type UnimplementedContactsServer struct {
}

func (UnimplementedContactsServer) GetFriendRequestList(context.Context, *GetFriendRequestListReq) (*GetFriendRequestListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendRequestList not implemented")
}
func (UnimplementedContactsServer) SetFriendRequest(context.Context, *SetFriendRequestReq) (*SetFriendRequestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriendRequest not implemented")
}
func (UnimplementedContactsServer) GetContacts(context.Context, *GetContactsReq) (*GetContactsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedContactsServer) SetContact(context.Context, *SetContactReq) (*SetContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetContact not implemented")
}
func (UnimplementedContactsServer) GetGroupContacts(context.Context, *GetGroupContactsReq) (*GetGroupContactsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupContacts not implemented")
}
func (UnimplementedContactsServer) SetGroupContact(context.Context, *SetGroupContactReq) (*SetGroupContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroupContact not implemented")
}
func (UnimplementedContactsServer) GetBlackList(context.Context, *GetBlackListReq) (*GetBlackListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlackList not implemented")
}
func (UnimplementedContactsServer) SetBlackList(context.Context, *SetBlackListReq) (*SetBlackListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBlackList not implemented")
}
func (UnimplementedContactsServer) mustEmbedUnimplementedContactsServer() {}

// UnsafeContactsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactsServer will
// result in compilation errors.
type UnsafeContactsServer interface {
	mustEmbedUnimplementedContactsServer()
}

func RegisterContactsServer(s grpc.ServiceRegistrar, srv ContactsServer) {
	s.RegisterService(&Contacts_ServiceDesc, srv)
}

func _Contacts_GetFriendRequestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendRequestListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).GetFriendRequestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_GetFriendRequestList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).GetFriendRequestList(ctx, req.(*GetFriendRequestListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SetFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendRequestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_SetFriendRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetFriendRequest(ctx, req.(*SetFriendRequestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_GetContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).GetContacts(ctx, req.(*GetContactsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_SetContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetContact(ctx, req.(*SetContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_GetGroupContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupContactsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).GetGroupContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_GetGroupContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).GetGroupContacts(ctx, req.(*GetGroupContactsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SetGroupContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetGroupContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_SetGroupContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetGroupContact(ctx, req.(*SetGroupContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_GetBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).GetBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_GetBlackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).GetBlackList(ctx, req.(*GetBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SetBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contacts_SetBlackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetBlackList(ctx, req.(*SetBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Contacts_ServiceDesc is the grpc.ServiceDesc for Contacts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contacts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contacts.Contacts",
	HandlerType: (*ContactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriendRequestList",
			Handler:    _Contacts_GetFriendRequestList_Handler,
		},
		{
			MethodName: "SetFriendRequest",
			Handler:    _Contacts_SetFriendRequest_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _Contacts_GetContacts_Handler,
		},
		{
			MethodName: "SetContact",
			Handler:    _Contacts_SetContact_Handler,
		},
		{
			MethodName: "GetGroupContacts",
			Handler:    _Contacts_GetGroupContacts_Handler,
		},
		{
			MethodName: "SetGroupContact",
			Handler:    _Contacts_SetGroupContact_Handler,
		},
		{
			MethodName: "GetBlackList",
			Handler:    _Contacts_GetBlackList_Handler,
		},
		{
			MethodName: "SetBlackList",
			Handler:    _Contacts_SetBlackList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/rpc/contacts.proto",
}