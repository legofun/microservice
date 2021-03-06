// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BaseResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Sex                  bool     `protobuf:"varint,4,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	IdCard               string   `protobuf:"bytes,6,opt,name=IdCard,proto3" json:"IdCard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

type SearchUserParam struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	IdCard               string   `protobuf:"bytes,3,opt,name=IdCard,proto3" json:"IdCard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchUserParam) Reset()         { *m = SearchUserParam{} }
func (m *SearchUserParam) String() string { return proto.CompactTextString(m) }
func (*SearchUserParam) ProtoMessage()    {}
func (*SearchUserParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *SearchUserParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserParam.Unmarshal(m, b)
}
func (m *SearchUserParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserParam.Marshal(b, m, deterministic)
}
func (m *SearchUserParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserParam.Merge(m, src)
}
func (m *SearchUserParam) XXX_Size() int {
	return xxx_messageInfo_SearchUserParam.Size(m)
}
func (m *SearchUserParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserParam.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserParam proto.InternalMessageInfo

func (m *SearchUserParam) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SearchUserParam) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SearchUserParam) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

type UserAddInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Sex                  bool     `protobuf:"varint,4,opt,name=Sex,proto3" json:"Sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddInfo) Reset()         { *m = UserAddInfo{} }
func (m *UserAddInfo) String() string { return proto.CompactTextString(m) }
func (*UserAddInfo) ProtoMessage()    {}
func (*UserAddInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserAddInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddInfo.Unmarshal(m, b)
}
func (m *UserAddInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddInfo.Marshal(b, m, deterministic)
}
func (m *UserAddInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddInfo.Merge(m, src)
}
func (m *UserAddInfo) XXX_Size() int {
	return xxx_messageInfo_UserAddInfo.Size(m)
}
func (m *UserAddInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddInfo proto.InternalMessageInfo

func (m *UserAddInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserAddInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserAddInfo) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

type UserList struct {
	Index                int32    `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	PageSize             int32    `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Count                int32    `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
	Users                []*User  `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *UserList) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *UserList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseResponse)(nil), "proto.BaseResponse")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*SearchUserParam)(nil), "proto.SearchUserParam")
	proto.RegisterType((*UserAddInfo)(nil), "proto.UserAddInfo")
	proto.RegisterType((*UserList)(nil), "proto.UserList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0xf1, 0x9f, 0x05, 0x3a, 0x54, 0xa5, 0x9a, 0x56, 0x68, 0xc5, 0xc9, 0xf5, 0xc9, 0x87,
	0x0a, 0xb5, 0xf4, 0xda, 0x0b, 0xa0, 0x28, 0xb2, 0x14, 0x22, 0x62, 0x2b, 0x97, 0xdc, 0x16, 0x76,
	0x42, 0x2c, 0x05, 0x1b, 0x79, 0x4d, 0x84, 0x72, 0xc8, 0x67, 0xc9, 0x47, 0x8d, 0x76, 0xd7, 0x10,
	0x27, 0x11, 0x87, 0x9c, 0x3c, 0x6f, 0xf6, 0xcd, 0xcf, 0x6f, 0xff, 0x00, 0xec, 0x14, 0x95, 0xa3,
	0x6d, 0x59, 0x54, 0x05, 0x32, 0xf3, 0x09, 0xff, 0xc3, 0xd7, 0xa9, 0x50, 0x94, 0x90, 0xda, 0x16,
	0xb9, 0x22, 0x44, 0xf0, 0x67, 0x85, 0x24, 0xee, 0x04, 0x4e, 0xc4, 0x12, 0x53, 0x23, 0x87, 0xce,
	0x9c, 0x94, 0x12, 0x6b, 0xe2, 0x6e, 0xe0, 0x44, 0x5f, 0x92, 0x83, 0x0c, 0x9f, 0xc0, 0xbf, 0x56,
	0x54, 0xe2, 0x37, 0x70, 0x63, 0x59, 0xcf, 0xb8, 0xb1, 0xd4, 0x94, 0x4b, 0xb1, 0x39, 0xd8, 0x4d,
	0x8d, 0xdf, 0xc1, 0x9b, 0xac, 0x89, 0x7b, 0xc6, 0xa4, 0x4b, 0xdd, 0x49, 0x69, 0xcf, 0xfd, 0xc0,
	0x89, 0xba, 0x89, 0x2e, 0x71, 0x00, 0xed, 0x79, 0xb1, 0xcc, 0xee, 0x89, 0x33, 0x33, 0x59, 0x2b,
	0xdd, 0x8f, 0xe5, 0x4c, 0x94, 0x92, 0xb7, 0x6d, 0xdf, 0xaa, 0xf0, 0x0a, 0xfa, 0x29, 0x89, 0x72,
	0x75, 0xa7, 0x53, 0x2c, 0x44, 0x29, 0x36, 0x1f, 0xa2, 0xbc, 0x22, 0xdd, 0x13, 0x48, 0xef, 0x0d,
	0xf2, 0x0c, 0x7a, 0x1a, 0x36, 0x91, 0x32, 0xce, 0x6f, 0x8b, 0xe3, 0x4e, 0x9c, 0xcf, 0xed, 0x24,
	0xdc, 0x41, 0x57, 0x63, 0x2e, 0x32, 0x55, 0xe1, 0x4f, 0x60, 0x71, 0x2e, 0x69, 0x6f, 0x12, 0xb0,
	0xc4, 0x0a, 0x1c, 0x42, 0x77, 0x21, 0xd6, 0x94, 0x66, 0x8f, 0x07, 0xd4, 0x51, 0xeb, 0x89, 0x59,
	0xb1, 0xcb, 0x2b, 0x43, 0x64, 0x89, 0x15, 0xf8, 0x0b, 0x98, 0xbe, 0x40, 0xc5, 0x59, 0xe0, 0x45,
	0xbd, 0x71, 0xcf, 0xde, 0xe4, 0x48, 0xff, 0x27, 0xb1, 0x2b, 0xe3, 0x67, 0xc7, 0xc6, 0x4f, 0xa9,
	0x7c, 0xc8, 0x56, 0x84, 0xbf, 0xc1, 0x3b, 0xa7, 0x0a, 0x07, 0xb5, 0xf5, 0xdd, 0x61, 0x0d, 0x9b,
	0x88, 0xb0, 0x85, 0x7f, 0xc1, 0x37, 0x81, 0x4f, 0xd9, 0xfb, 0x0d, 0xbb, 0x36, 0x86, 0x2d, 0xfc,
	0x03, 0xde, 0x44, 0x4a, 0xc4, 0xc6, 0x4a, 0x7d, 0x74, 0xc3, 0x1f, 0x75, 0xaf, 0xf9, 0xbe, 0xc2,
	0xd6, 0xb4, 0x73, 0x63, 0x9f, 0xde, 0xb2, 0x6d, 0x3e, 0xff, 0x5e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x95, 0x32, 0x4b, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//获取指定用户接口
	Get(ctx context.Context, in *SearchUserParam, opts ...grpc.CallOption) (*User, error)
	//获取用户列表，分页
	List(ctx context.Context, in *SearchUserParam, opts ...grpc.CallOption) (*UserList, error)
	//添加用户
	Add(ctx context.Context, in *UserAddInfo, opts ...grpc.CallOption) (*BaseResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *SearchUserParam, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) List(ctx context.Context, in *SearchUserParam, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/proto.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Add(ctx context.Context, in *UserAddInfo, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//获取指定用户接口
	Get(context.Context, *SearchUserParam) (*User, error)
	//获取用户列表，分页
	List(context.Context, *SearchUserParam) (*UserList, error)
	//添加用户
	Add(context.Context, *UserAddInfo) (*BaseResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *SearchUserParam) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) List(ctx context.Context, req *SearchUserParam) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedUserServiceServer) Add(ctx context.Context, req *UserAddInfo) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*SearchUserParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*SearchUserParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Add(ctx, req.(*UserAddInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _UserService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
