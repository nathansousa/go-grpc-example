// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

package grpc

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

type AddBookResponse struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBookResponse) Reset()         { *m = AddBookResponse{} }
func (m *AddBookResponse) String() string { return proto.CompactTextString(m) }
func (*AddBookResponse) ProtoMessage()    {}
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{0}
}

func (m *AddBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBookResponse.Unmarshal(m, b)
}
func (m *AddBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBookResponse.Marshal(b, m, deterministic)
}
func (m *AddBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBookResponse.Merge(m, src)
}
func (m *AddBookResponse) XXX_Size() int {
	return xxx_messageInfo_AddBookResponse.Size(m)
}
func (m *AddBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBookResponse proto.InternalMessageInfo

func (m *AddBookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *AddBookResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetBookRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{1}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetBookResponse struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookResponse) Reset()         { *m = GetBookResponse{} }
func (m *GetBookResponse) String() string { return proto.CompactTextString(m) }
func (*GetBookResponse) ProtoMessage()    {}
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{2}
}

func (m *GetBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookResponse.Unmarshal(m, b)
}
func (m *GetBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookResponse.Marshal(b, m, deterministic)
}
func (m *GetBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookResponse.Merge(m, src)
}
func (m *GetBookResponse) XXX_Size() int {
	return xxx_messageInfo_GetBookResponse.Size(m)
}
func (m *GetBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookResponse proto.InternalMessageInfo

func (m *GetBookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *GetBookResponse) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *GetBookResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Book struct {
	BookID               string   `protobuf:"bytes,1,opt,name=bookID,proto3" json:"bookID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	IsDigital            bool     `protobuf:"varint,5,opt,name=isDigital,proto3" json:"isDigital,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{4}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetBookID() string {
	if m != nil {
		return m.BookID
	}
	return ""
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetIsDigital() bool {
	if m != nil {
		return m.IsDigital
	}
	return false
}

func init() {
	proto.RegisterType((*AddBookResponse)(nil), "grpc.AddBookResponse")
	proto.RegisterType((*GetBookRequest)(nil), "grpc.GetBookRequest")
	proto.RegisterType((*GetBookResponse)(nil), "grpc.GetBookResponse")
	proto.RegisterType((*Error)(nil), "grpc.Error")
	proto.RegisterType((*Book)(nil), "grpc.Book")
}

func init() { proto.RegisterFile("book.proto", fileDescriptor_1e89d0eaa98dc5d8) }

var fileDescriptor_1e89d0eaa98dc5d8 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x49, 0x9b, 0xb4, 0xcd, 0x04, 0xbe, 0xc2, 0xf0, 0x29, 0xa1, 0x88, 0xd4, 0xe0, 0xa1,
	0x07, 0xc9, 0xa1, 0xe2, 0x03, 0x44, 0x2a, 0xe2, 0xd1, 0xd5, 0x17, 0x48, 0x93, 0x21, 0x0d, 0xb5,
	0xd9, 0x74, 0x77, 0x23, 0x78, 0xf1, 0xd9, 0x65, 0x37, 0x9b, 0x44, 0x0b, 0x82, 0xb7, 0xf9, 0x4f,
	0x7e, 0x99, 0xf9, 0xcf, 0xcc, 0x02, 0x6c, 0x39, 0xdf, 0xc7, 0xb5, 0xe0, 0x8a, 0xa3, 0x5b, 0x88,
	0x3a, 0x8b, 0x5e, 0x61, 0x9e, 0xe4, 0xf9, 0x3d, 0xe7, 0x7b, 0x46, 0xb2, 0xe6, 0x95, 0x24, 0xbc,
	0x04, 0x57, 0x63, 0xa1, 0xb3, 0x74, 0x56, 0xc1, 0x1a, 0x62, 0xcd, 0xc5, 0x86, 0x30, 0x79, 0xbc,
	0x02, 0x8f, 0x84, 0xe0, 0x22, 0x1c, 0x19, 0x20, 0x68, 0x81, 0x07, 0x9d, 0x62, 0xed, 0x97, 0xe8,
	0x1a, 0xfe, 0x3d, 0x92, 0x6a, 0xab, 0x1e, 0x1b, 0x92, 0x0a, 0x11, 0xdc, 0x2a, 0x3d, 0x90, 0x29,
	0xea, 0x33, 0x13, 0x47, 0x35, 0xcc, 0x7b, 0xea, 0x8f, 0xbd, 0x17, 0x30, 0x3b, 0x36, 0x69, 0xa5,
	0x4a, 0xf5, 0x61, 0xda, 0x8f, 0x59, 0xaf, 0x07, 0x5f, 0xe3, 0x5f, 0x7d, 0xdd, 0x81, 0x67, 0xb4,
	0xb6, 0x93, 0xf1, 0xbc, 0xb7, 0xa3, 0x63, 0x0c, 0x61, 0x7a, 0x20, 0x29, 0xd3, 0x82, 0x4c, 0x69,
	0x9f, 0x75, 0x32, 0xda, 0x81, 0xab, 0x3d, 0xe0, 0x39, 0x4c, 0xb4, 0x8b, 0xa7, 0x8d, 0xfd, 0xcf,
	0xaa, 0x7e, 0xb8, 0xd1, 0x30, 0x9c, 0x66, 0xd3, 0x46, 0xed, 0xac, 0x1d, 0x9f, 0x59, 0x85, 0x17,
	0xe0, 0x97, 0x72, 0x53, 0x16, 0xa5, 0x4a, 0xdf, 0x42, 0x6f, 0xe9, 0xac, 0x66, 0x6c, 0x48, 0xac,
	0x3f, 0x21, 0xd0, 0x9d, 0x5e, 0x48, 0xbc, 0x97, 0x19, 0xe1, 0x0d, 0x4c, 0xed, 0x75, 0xf0, 0xdb,
	0x2e, 0x16, 0x67, 0x6d, 0x7c, 0x7a, 0xb8, 0x04, 0xd0, 0xee, 0x33, 0xa9, 0xf2, 0xe7, 0x6e, 0x2d,
	0xff, 0x5b, 0xf8, 0xe7, 0x3d, 0xba, 0x12, 0x27, 0xfb, 0xdf, 0x4e, 0xcc, 0xdb, 0xb8, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x6c, 0xa2, 0xdd, 0x29, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*AddBookResponse, error)
	GetBookAndQuantity(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*AddBookResponse, error) {
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, "/grpc.BookService/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookAndQuantity(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/grpc.BookService/GetBookAndQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	AddBook(context.Context, *Book) (*AddBookResponse, error)
	GetBookAndQuantity(context.Context, *GetBookRequest) (*GetBookResponse, error)
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) AddBook(ctx context.Context, req *Book) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (*UnimplementedBookServiceServer) GetBookAndQuantity(ctx context.Context, req *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookAndQuantity not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BookService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookAndQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookAndQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BookService/GetBookAndQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookAndQuantity(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BookService_AddBook_Handler,
		},
		{
			MethodName: "GetBookAndQuantity",
			Handler:    _BookService_GetBookAndQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}