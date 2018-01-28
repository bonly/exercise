// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Proto.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	Proto.proto

It has these top-level messages:
	ReqLogin
	RepLogin
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ReqLogin struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *ReqLogin) Reset()                    { *m = ReqLogin{} }
func (m *ReqLogin) String() string            { return proto1.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()               {}
func (*ReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqLogin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqLogin) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type RepLogin struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *RepLogin) Reset()                    { *m = RepLogin{} }
func (m *RepLogin) String() string            { return proto1.CompactTextString(m) }
func (*RepLogin) ProtoMessage()               {}
func (*RepLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepLogin) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *RepLogin) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto1.RegisterType((*ReqLogin)(nil), "proto.ReqLogin")
	proto1.RegisterType((*RepLogin)(nil), "proto.RepLogin")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*RepLogin, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*RepLogin, error) {
	out := new(RepLogin)
	err := grpc.Invoke(ctx, "/proto.User/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	Login(context.Context, *ReqLogin) (*RepLogin, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*ReqLogin))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto.proto",
}

// Client API for Game service

type GameClient interface {
	Logout(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*RepLogin, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Logout(ctx context.Context, in *ReqLogin, opts ...grpc.CallOption) (*RepLogin, error) {
	out := new(RepLogin)
	err := grpc.Invoke(ctx, "/proto.Game/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Game service

type GameServer interface {
	Logout(context.Context, *ReqLogin) (*RepLogin, error)
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Game/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Logout(ctx, req.(*ReqLogin))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logout",
			Handler:    _Game_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto.proto",
}

func init() { proto1.RegisterFile("Proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2b, 0x00, 0x91, 0x42, 0xac, 0x60, 0x4a, 0xc9, 0x8c, 0x8b, 0x23, 0x28, 0xb5, 0xd0,
	0x27, 0x3f, 0x3d, 0x33, 0x4f, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x48, 0x2c, 0x2e, 0x2e, 0x4f, 0x91,
	0x60, 0x02, 0x8b, 0x42, 0x79, 0x4a, 0x7a, 0x20, 0x7d, 0x05, 0x10, 0x7d, 0x02, 0x5c, 0xcc, 0x45,
	0xa9, 0x25, 0x60, 0x6d, 0xac, 0x41, 0x20, 0x26, 0x48, 0x24, 0xb7, 0x38, 0x1d, 0xaa, 0x05, 0xc4,
	0x34, 0x32, 0xe6, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0xd2, 0xe6, 0x62, 0x85, 0x68, 0xe2, 0x87,
	0xb8, 0x43, 0x0f, 0x66, 0xbb, 0x14, 0x42, 0x00, 0x62, 0xac, 0x12, 0x83, 0x91, 0x09, 0x17, 0x8b,
	0x3b, 0xc8, 0x11, 0x3a, 0x5c, 0x6c, 0x3e, 0xf9, 0xe9, 0xf9, 0xa5, 0x25, 0xc4, 0xe8, 0x4a, 0x62,
	0x03, 0x8b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x00, 0x70, 0x08, 0xef, 0x00, 0x00,
	0x00,
}
