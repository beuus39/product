// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protocols/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CategoryID int32 `protobuf:"varint,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
}

func (x *ProductQueryRequest) Reset() {
	*x = ProductQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductQueryRequest) ProtoMessage() {}

func (x *ProductQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductQueryRequest.ProtoReflect.Descriptor instead.
func (*ProductQueryRequest) Descriptor() ([]byte, []int) {
	return file_protocols_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductQueryRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductQueryRequest) GetCategoryID() int32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CategoryID  int32   `protobuf:"varint,2,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Image       string  `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	Stock       int32   `protobuf:"varint,6,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Price       float64 `protobuf:"fixed64,7,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_protocols_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductResponse) GetCategoryID() int32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductList []*ProductResponse `protobuf:"bytes,1,rep,name=ProductList,proto3" json:"ProductList,omitempty"`
}

func (x *ProductListResponse) Reset() {
	*x = ProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListResponse) ProtoMessage() {}

func (x *ProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListResponse.ProtoReflect.Descriptor instead.
func (*ProductListResponse) Descriptor() ([]byte, []int) {
	return file_protocols_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductListResponse) GetProductList() []*ProductResponse {
	if x != nil {
		return x.ProductList
	}
	return nil
}

var File_protocols_product_proto protoreflect.FileDescriptor

var file_protocols_product_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xe5, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x07, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_product_proto_rawDescOnce sync.Once
	file_protocols_product_proto_rawDescData = file_protocols_product_proto_rawDesc
)

func file_protocols_product_proto_rawDescGZIP() []byte {
	file_protocols_product_proto_rawDescOnce.Do(func() {
		file_protocols_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_product_proto_rawDescData)
	})
	return file_protocols_product_proto_rawDescData
}

var file_protocols_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocols_product_proto_goTypes = []interface{}{
	(*ProductQueryRequest)(nil), // 0: product.ProductQueryRequest
	(*ProductResponse)(nil),     // 1: product.ProductResponse
	(*ProductListResponse)(nil), // 2: product.ProductListResponse
}
var file_protocols_product_proto_depIdxs = []int32{
	1, // 0: product.ProductListResponse.ProductList:type_name -> product.ProductResponse
	0, // 1: product.ProductService.FindByID:input_type -> product.ProductQueryRequest
	0, // 2: product.ProductService.FindByCategory:input_type -> product.ProductQueryRequest
	0, // 3: product.ProductService.FindAll:input_type -> product.ProductQueryRequest
	1, // 4: product.ProductService.FindByID:output_type -> product.ProductResponse
	1, // 5: product.ProductService.FindByCategory:output_type -> product.ProductResponse
	1, // 6: product.ProductService.FindAll:output_type -> product.ProductResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocols_product_proto_init() }
func file_protocols_product_proto_init() {
	if File_protocols_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocols_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocols_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocols_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocols_product_proto_goTypes,
		DependencyIndexes: file_protocols_product_proto_depIdxs,
		MessageInfos:      file_protocols_product_proto_msgTypes,
	}.Build()
	File_protocols_product_proto = out.File
	file_protocols_product_proto_rawDesc = nil
	file_protocols_product_proto_goTypes = nil
	file_protocols_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	FindByID(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	FindByCategory(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (ProductService_FindByCategoryClient, error)
	FindAll(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (ProductService_FindAllClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) FindByID(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindByCategory(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (ProductService_FindByCategoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[0], "/product.ProductService/FindByCategory", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceFindByCategoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_FindByCategoryClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceFindByCategoryClient struct {
	grpc.ClientStream
}

func (x *productServiceFindByCategoryClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) FindAll(ctx context.Context, in *ProductQueryRequest, opts ...grpc.CallOption) (ProductService_FindAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductService_serviceDesc.Streams[1], "/product.ProductService/FindAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceFindAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_FindAllClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type productServiceFindAllClient struct {
	grpc.ClientStream
}

func (x *productServiceFindAllClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	FindByID(context.Context, *ProductQueryRequest) (*ProductResponse, error)
	FindByCategory(*ProductQueryRequest, ProductService_FindByCategoryServer) error
	FindAll(*ProductQueryRequest, ProductService_FindAllServer) error
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) FindByID(context.Context, *ProductQueryRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (*UnimplementedProductServiceServer) FindByCategory(*ProductQueryRequest, ProductService_FindByCategoryServer) error {
	return status.Errorf(codes.Unimplemented, "method FindByCategory not implemented")
}
func (*UnimplementedProductServiceServer) FindAll(*ProductQueryRequest, ProductService_FindAllServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindByID(ctx, req.(*ProductQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindByCategory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).FindByCategory(m, &productServiceFindByCategoryServer{stream})
}

type ProductService_FindByCategoryServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceFindByCategoryServer struct {
	grpc.ServerStream
}

func (x *productServiceFindByCategoryServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_FindAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).FindAll(m, &productServiceFindAllServer{stream})
}

type ProductService_FindAllServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type productServiceFindAllServer struct {
	grpc.ServerStream
}

func (x *productServiceFindAllServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _ProductService_FindByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindByCategory",
			Handler:       _ProductService_FindByCategory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindAll",
			Handler:       _ProductService_FindAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocols/product.proto",
}
