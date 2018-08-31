// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserRequest) Reset()         { *m = NewUserRequest{} }
func (m *NewUserRequest) String() string { return proto.CompactTextString(m) }
func (*NewUserRequest) ProtoMessage()    {}
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_782e69eea2d045fd, []int{0}
}
func (m *NewUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserRequest.Unmarshal(m, b)
}
func (m *NewUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserRequest.Marshal(b, m, deterministic)
}
func (dst *NewUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserRequest.Merge(dst, src)
}
func (m *NewUserRequest) XXX_Size() int {
	return xxx_messageInfo_NewUserRequest.Size(m)
}
func (m *NewUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserRequest proto.InternalMessageInfo

func (m *NewUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type NewUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserResponse) Reset()         { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()    {}
func (*NewUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_782e69eea2d045fd, []int{1}
}
func (m *NewUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserResponse.Unmarshal(m, b)
}
func (m *NewUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserResponse.Marshal(b, m, deterministic)
}
func (dst *NewUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserResponse.Merge(dst, src)
}
func (m *NewUserResponse) XXX_Size() int {
	return xxx_messageInfo_NewUserResponse.Size(m)
}
func (m *NewUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserResponse proto.InternalMessageInfo

func (m *NewUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserByEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByEmailRequest) Reset()         { *m = GetUserByEmailRequest{} }
func (m *GetUserByEmailRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByEmailRequest) ProtoMessage()    {}
func (*GetUserByEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_782e69eea2d045fd, []int{2}
}
func (m *GetUserByEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByEmailRequest.Unmarshal(m, b)
}
func (m *GetUserByEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByEmailRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserByEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByEmailRequest.Merge(dst, src)
}
func (m *GetUserByEmailRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByEmailRequest.Size(m)
}
func (m *GetUserByEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByEmailRequest proto.InternalMessageInfo

func (m *GetUserByEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_782e69eea2d045fd, []int{3}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*NewUserRequest)(nil), "pb.NewUserRequest")
	proto.RegisterType((*NewUserResponse)(nil), "pb.NewUserResponse")
	proto.RegisterType((*GetUserByEmailRequest)(nil), "pb.GetUserByEmailRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := c.cc.Invoke(ctx, "/pb.User/NewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.User/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*UserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).NewUser(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByEmail(ctx, req.(*GetUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _User_NewUser_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _User_GetUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_782e69eea2d045fd) }

var fileDescriptor_user_782e69eea2d045fd = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x8a, 0xe3, 0xe2, 0xf3, 0x4b,
	0x2d, 0x0f, 0x2d, 0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2,
	0x00, 0xa9, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85,
	0x44, 0xb8, 0x58, 0x53, 0x73, 0x13, 0x33, 0x73, 0x24, 0x98, 0xc0, 0x12, 0x10, 0x0e, 0x48, 0x47,
	0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x33, 0x44, 0x07, 0x8c, 0xaf, 0xa4, 0xc8,
	0xc5, 0x0f, 0x37, 0xbf, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x05,
	0x6a, 0x34, 0x53, 0x66, 0x8a, 0x92, 0x2e, 0x97, 0xa8, 0x7b, 0x6a, 0x09, 0x48, 0x89, 0x53, 0xa5,
	0x2b, 0xc8, 0x40, 0x98, 0x4b, 0xe0, 0xb6, 0x31, 0x22, 0xd9, 0xa6, 0x14, 0xc0, 0xc5, 0x83, 0xcf,
	0x38, 0x14, 0xf7, 0x33, 0xe1, 0x72, 0x3f, 0x33, 0x92, 0x89, 0x46, 0xb5, 0x5c, 0x2c, 0x20, 0x13,
	0x85, 0x4c, 0xb8, 0xd8, 0xa1, 0x6e, 0x15, 0x12, 0xd2, 0x2b, 0x48, 0xd2, 0x43, 0x0d, 0x18, 0x29,
	0x61, 0x14, 0x31, 0x88, 0xed, 0x4a, 0x0c, 0x42, 0xf6, 0x5c, 0x7c, 0xa8, 0xce, 0x17, 0x92, 0x04,
	0x29, 0xc4, 0xea, 0x25, 0x29, 0x01, 0x90, 0x14, 0xaa, 0x01, 0x49, 0x6c, 0xe0, 0xd8, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x69, 0x24, 0xe6, 0x03, 0x9b, 0x01, 0x00, 0x00,
}
