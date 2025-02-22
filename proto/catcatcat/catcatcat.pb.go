// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: catcatcat/catcatcat.proto

package catcatcat

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CatId         string                 `protobuf:"bytes,1,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Fetcher       string                 `protobuf:"bytes,3,opt,name=fetcher,proto3" json:"fetcher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cat) Reset() {
	*x = Cat{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{0}
}

func (x *Cat) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetFetcher() string {
	if x != nil {
		return x.Fetcher
	}
	return ""
}

type Cdp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ts            int64                  `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"` // unix time in millisecond
	Price         float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	InStock       bool                   `protobuf:"varint,3,opt,name=in_stock,json=inStock,proto3" json:"in_stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cdp) Reset() {
	*x = Cdp{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cdp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cdp) ProtoMessage() {}

func (x *Cdp) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cdp.ProtoReflect.Descriptor instead.
func (*Cdp) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{1}
}

func (x *Cdp) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Cdp) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Cdp) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

type ListCatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCatsRequest) Reset() {
	*x = ListCatsRequest{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsRequest) ProtoMessage() {}

func (x *ListCatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsRequest.ProtoReflect.Descriptor instead.
func (*ListCatsRequest) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{2}
}

type ListCatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cats          []*Cat                 `protobuf:"bytes,1,rep,name=cats,proto3" json:"cats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCatsResponse) Reset() {
	*x = ListCatsResponse{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatsResponse) ProtoMessage() {}

func (x *ListCatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatsResponse.ProtoReflect.Descriptor instead.
func (*ListCatsResponse) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{3}
}

func (x *ListCatsResponse) GetCats() []*Cat {
	if x != nil {
		return x.Cats
	}
	return nil
}

type GetCdpsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	CatId string                 `protobuf:"bytes,1,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	// Below is optional
	LastN         int64 `protobuf:"varint,2,opt,name=last_n,json=lastN,proto3" json:"last_n,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCdpsRequest) Reset() {
	*x = GetCdpsRequest{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCdpsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCdpsRequest) ProtoMessage() {}

func (x *GetCdpsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCdpsRequest.ProtoReflect.Descriptor instead.
func (*GetCdpsRequest) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{4}
}

func (x *GetCdpsRequest) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *GetCdpsRequest) GetLastN() int64 {
	if x != nil {
		return x.LastN
	}
	return 0
}

type GetCdpsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cat           *Cat                   `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
	Cdps          []*Cdp                 `protobuf:"bytes,2,rep,name=cdps,proto3" json:"cdps,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCdpsResponse) Reset() {
	*x = GetCdpsResponse{}
	mi := &file_catcatcat_catcatcat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCdpsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCdpsResponse) ProtoMessage() {}

func (x *GetCdpsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catcatcat_catcatcat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCdpsResponse.ProtoReflect.Descriptor instead.
func (*GetCdpsResponse) Descriptor() ([]byte, []int) {
	return file_catcatcat_catcatcat_proto_rawDescGZIP(), []int{5}
}

func (x *GetCdpsResponse) GetCat() *Cat {
	if x != nil {
		return x.Cat
	}
	return nil
}

func (x *GetCdpsResponse) GetCdps() []*Cdp {
	if x != nil {
		return x.Cdps
	}
	return nil
}

var File_catcatcat_catcatcat_proto protoreflect.FileDescriptor

var file_catcatcat_catcatcat_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x2f, 0x63, 0x61, 0x74, 0x63,
	0x61, 0x74, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x61, 0x74,
	0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x22, 0x46, 0x0a, 0x03, 0x43, 0x64, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x04, 0x63,
	0x61, 0x74, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x64, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x64, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x2e,
	0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x64, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63,
	0x61, 0x74, 0x2e, 0x43, 0x64, 0x70, 0x52, 0x04, 0x63, 0x64, 0x70, 0x73, 0x32, 0xb9, 0x01, 0x0a,
	0x09, 0x43, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x12, 0x52, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63,
	0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x63, 0x61, 0x74, 0x73, 0x12, 0x58,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x64, 0x70, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x74, 0x63,
	0x61, 0x74, 0x63, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x64, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x64, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x63, 0x64, 0x70, 0x73, 0x2f,
	0x7b, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x91, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x42, 0x0e, 0x43, 0x61, 0x74, 0x63,
	0x61, 0x74, 0x63, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6c, 0x69, 0x6e, 0x32, 0x30,
	0x2f, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x43, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0xca, 0x02, 0x09, 0x43,
	0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0xe2, 0x02, 0x15, 0x43, 0x61, 0x74, 0x63, 0x61,
	0x74, 0x63, 0x61, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x09, 0x43, 0x61, 0x74, 0x63, 0x61, 0x74, 0x63, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_catcatcat_catcatcat_proto_rawDescOnce sync.Once
	file_catcatcat_catcatcat_proto_rawDescData []byte
)

func file_catcatcat_catcatcat_proto_rawDescGZIP() []byte {
	file_catcatcat_catcatcat_proto_rawDescOnce.Do(func() {
		file_catcatcat_catcatcat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_catcatcat_catcatcat_proto_rawDesc), len(file_catcatcat_catcatcat_proto_rawDesc)))
	})
	return file_catcatcat_catcatcat_proto_rawDescData
}

var file_catcatcat_catcatcat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_catcatcat_catcatcat_proto_goTypes = []any{
	(*Cat)(nil),              // 0: catcatcat.Cat
	(*Cdp)(nil),              // 1: catcatcat.Cdp
	(*ListCatsRequest)(nil),  // 2: catcatcat.ListCatsRequest
	(*ListCatsResponse)(nil), // 3: catcatcat.ListCatsResponse
	(*GetCdpsRequest)(nil),   // 4: catcatcat.GetCdpsRequest
	(*GetCdpsResponse)(nil),  // 5: catcatcat.GetCdpsResponse
}
var file_catcatcat_catcatcat_proto_depIdxs = []int32{
	0, // 0: catcatcat.ListCatsResponse.cats:type_name -> catcatcat.Cat
	0, // 1: catcatcat.GetCdpsResponse.cat:type_name -> catcatcat.Cat
	1, // 2: catcatcat.GetCdpsResponse.cdps:type_name -> catcatcat.Cdp
	2, // 3: catcatcat.Catcatcat.ListCats:input_type -> catcatcat.ListCatsRequest
	4, // 4: catcatcat.Catcatcat.GetCdps:input_type -> catcatcat.GetCdpsRequest
	3, // 5: catcatcat.Catcatcat.ListCats:output_type -> catcatcat.ListCatsResponse
	5, // 6: catcatcat.Catcatcat.GetCdps:output_type -> catcatcat.GetCdpsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_catcatcat_catcatcat_proto_init() }
func file_catcatcat_catcatcat_proto_init() {
	if File_catcatcat_catcatcat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_catcatcat_catcatcat_proto_rawDesc), len(file_catcatcat_catcatcat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catcatcat_catcatcat_proto_goTypes,
		DependencyIndexes: file_catcatcat_catcatcat_proto_depIdxs,
		MessageInfos:      file_catcatcat_catcatcat_proto_msgTypes,
	}.Build()
	File_catcatcat_catcatcat_proto = out.File
	file_catcatcat_catcatcat_proto_goTypes = nil
	file_catcatcat_catcatcat_proto_depIdxs = nil
}
