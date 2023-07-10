// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: main.proto

package pb

import (
	_ "go.unistack.org/micro-proto/v3/api"
	_ "go.unistack.org/micro-proto/v3/openapiv3"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

type GetBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName *string `protobuf:"bytes,1,opt,name=book_name,json=bookName,proto3,oneof" json:"book_name,omitempty"`
}

func (x *GetBookReq) Reset() {
	*x = GetBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReq) ProtoMessage() {}

func (x *GetBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReq.ProtoReflect.Descriptor instead.
func (*GetBookReq) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookReq) GetBookName() string {
	if x != nil && x.BookName != nil {
		return *x.BookName
	}
	return ""
}

type PostBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookTitle string `protobuf:"bytes,1,opt,name=book_title,json=bookTitle,proto3" json:"book_title,omitempty"`
	Author    string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Genre     string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Year      string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
	BookBytes []byte `protobuf:"bytes,5,opt,name=book_bytes,json=bookBytes,proto3" json:"book_bytes,omitempty"`
}

func (x *PostBook) Reset() {
	*x = PostBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBook) ProtoMessage() {}

func (x *PostBook) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBook.ProtoReflect.Descriptor instead.
func (*PostBook) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{2}
}

func (x *PostBook) GetBookTitle() string {
	if x != nil {
		return x.BookTitle
	}
	return ""
}

func (x *PostBook) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PostBook) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *PostBook) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *PostBook) GetBookBytes() []byte {
	if x != nil {
		return x.BookBytes
	}
	return nil
}

type GetBookRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Genre  string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Year   string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *GetBookRsp) Reset() {
	*x = GetBookRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRsp) ProtoMessage() {}

func (x *GetBookRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRsp.ProtoReflect.Descriptor instead.
func (*GetBookRsp) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookRsp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetBookRsp) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetBookRsp) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetBookRsp) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

type StatusRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *StatusRsp) Reset() {
	*x = StatusRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRsp) ProtoMessage() {}

func (x *StatusRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRsp.ProtoReflect.Descriptor instead.
func (*StatusRsp) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{4}
}

func (x *StatusRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StatusRsp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type StatusUploadedBookRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *StatusUploadedBookRsp) Reset() {
	*x = StatusUploadedBookRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUploadedBookRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUploadedBookRsp) ProtoMessage() {}

func (x *StatusUploadedBookRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUploadedBookRsp.ProtoReflect.Descriptor instead.
func (*StatusUploadedBookRsp) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{5}
}

func (x *StatusUploadedBookRsp) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *StatusUploadedBookRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x1a, 0x15, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x62,
	0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x41, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44,
	0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x94, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x79, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0x22, 0x3d, 0xaa, 0x84, 0x9e,
	0x03, 0x2f, 0x2a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x42,
	0x20, 0x0a, 0x1e, 0x12, 0x1c, 0x0a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73,
	0x70, 0xb2, 0xea, 0xff, 0xf9, 0x01, 0x03, 0x12, 0x01, 0x2f, 0x12, 0x8e, 0x01, 0x0a, 0x04, 0x50,
	0x55, 0x53, 0x48, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x26,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x43, 0xaa, 0x84, 0x9e, 0x03, 0x2e, 0x2a, 0x0a, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x20, 0x0a, 0x1e, 0x12, 0x1c, 0x0a,
	0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0xb2, 0xea, 0xff, 0xf9, 0x01,
	0x0a, 0x22, 0x05, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x03, 0x47,
	0x45, 0x54, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x3d, 0xaa, 0x84,
	0x9e, 0x03, 0x2b, 0x2a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x20, 0x0a, 0x1e,
	0x12, 0x1c, 0x0a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0xb2, 0xea,
	0xff, 0xf9, 0x01, 0x07, 0x12, 0x05, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x42, 0x20, 0x5a, 0x1e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_main_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: bookservermicro.Empty
	(*GetBookReq)(nil),            // 1: bookservermicro.getBookReq
	(*PostBook)(nil),              // 2: bookservermicro.postBook
	(*GetBookRsp)(nil),            // 3: bookservermicro.getBookRsp
	(*StatusRsp)(nil),             // 4: bookservermicro.StatusRsp
	(*StatusUploadedBookRsp)(nil), // 5: bookservermicro.StatusUploadedBookRsp
}
var file_main_proto_depIdxs = []int32{
	0, // 0: bookservermicro.Server.HOME:input_type -> bookservermicro.Empty
	2, // 1: bookservermicro.Server.PUSH:input_type -> bookservermicro.postBook
	1, // 2: bookservermicro.Server.GET:input_type -> bookservermicro.getBookReq
	4, // 3: bookservermicro.Server.HOME:output_type -> bookservermicro.StatusRsp
	5, // 4: bookservermicro.Server.PUSH:output_type -> bookservermicro.StatusUploadedBookRsp
	3, // 5: bookservermicro.Server.GET:output_type -> bookservermicro.getBookRsp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookReq); i {
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
		file_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBook); i {
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
		file_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRsp); i {
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
		file_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRsp); i {
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
		file_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusUploadedBookRsp); i {
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
	file_main_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}
