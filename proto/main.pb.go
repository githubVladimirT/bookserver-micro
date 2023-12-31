// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: main.proto

package pb

import (
	_ "go.unistack.org/micro-proto/v3/api"
	_ "go.unistack.org/micro-proto/v3/codec"
	_ "go.unistack.org/micro-proto/v3/openapiv3"
	_ "go.unistack.org/micro-proto/v3/tag"
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

type GetBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookTitle string `protobuf:"bytes,1,opt,name=book_title,json=bookTitle,proto3" json:"book_title,omitempty"`
}

func (x *GetBook) Reset() {
	*x = GetBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBook) ProtoMessage() {}

func (x *GetBook) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetBook.ProtoReflect.Descriptor instead.
func (*GetBook) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *GetBook) GetBookTitle() string {
	if x != nil {
		return x.BookTitle
	}
	return ""
}

type SortType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortType string `protobuf:"bytes,1,opt,name=sort_type,json=sortType,proto3" json:"sort_type,omitempty"`
}

func (x *SortType) Reset() {
	*x = SortType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortType) ProtoMessage() {}

func (x *SortType) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SortType.ProtoReflect.Descriptor instead.
func (*SortType) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{2}
}

func (x *SortType) GetSortType() string {
	if x != nil {
		return x.SortType
	}
	return ""
}

type GetAllBooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []string `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetAllBooks) Reset() {
	*x = GetAllBooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooks) ProtoMessage() {}

func (x *GetAllBooks) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllBooks.ProtoReflect.Descriptor instead.
func (*GetAllBooks) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllBooks) GetBooks() []string {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetAllBooksAndSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*GetBookRsp `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetAllBooksAndSort) Reset() {
	*x = GetAllBooksAndSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBooksAndSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksAndSort) ProtoMessage() {}

func (x *GetAllBooksAndSort) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllBooksAndSort.ProtoReflect.Descriptor instead.
func (*GetAllBooksAndSort) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllBooksAndSort) GetBooks() []*GetBookRsp {
	if x != nil {
		return x.Books
	}
	return nil
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
		mi := &file_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBook) ProtoMessage() {}

func (x *PostBook) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PostBook.ProtoReflect.Descriptor instead.
func (*PostBook) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{5}
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
		mi := &file_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRsp) ProtoMessage() {}

func (x *GetBookRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[6]
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
	return file_main_proto_rawDescGZIP(), []int{6}
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

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *StatusRsp) Reset() {
	*x = StatusRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRsp) ProtoMessage() {}

func (x *StatusRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[7]
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
	return file_main_proto_rawDescGZIP(), []int{7}
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
}

func (x *StatusUploadedBookRsp) Reset() {
	*x = StatusUploadedBookRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUploadedBookRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUploadedBookRsp) ProtoMessage() {}

func (x *StatusUploadedBookRsp) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[8]
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
	return file_main_proto_rawDescGZIP(), []int{8}
}

func (x *StatusUploadedBookRsp) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x1a, 0x15, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x07,
	0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x27, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x47, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x8a, 0x01,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x22, 0x2d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x30, 0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49,
	0x64, 0x32, 0xbc, 0x05, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x7c, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x73, 0x70, 0x22, 0x40, 0xaa, 0x84,
	0x9e, 0x03, 0x32, 0x2a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x3a, 0x2a, 0x0a, 0x28, 0x12, 0x26, 0x0a,
	0x24, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0e, 0x12, 0x0c, 0x0a, 0x0a, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x73, 0x70, 0xb2, 0xea, 0xff, 0xf9, 0x01, 0x03, 0x12, 0x01, 0x2f, 0x12, 0x5b,
	0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x10, 0xb2, 0xea, 0xff, 0xf9, 0x01,
	0x0a, 0x22, 0x05, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x1b,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x45, 0xaa, 0x84, 0x9e,
	0x03, 0x33, 0x2a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x2b, 0x0a, 0x29, 0x12, 0x27, 0x0a, 0x25,
	0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0f, 0x12, 0x0d, 0x0a, 0x0b, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x73, 0x70, 0xb2, 0xea, 0xff, 0xf9, 0x01, 0x07, 0x12, 0x05, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x12, 0x96, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x51, 0xaa, 0x84, 0x9e, 0x03, 0x3b, 0x2a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x3a, 0x2c, 0x0a, 0x2a,
	0x12, 0x28, 0x0a, 0x26, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x10, 0x12, 0x0e, 0x0a, 0x0c, 0x2e, 0x67,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0xb2, 0xea, 0xff, 0xf9, 0x01, 0x0b,
	0x12, 0x09, 0x2f, 0x61, 0x6c, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0xb2, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x23, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x6f,
	0x72, 0x74, 0x22, 0x5c, 0xaa, 0x84, 0x9e, 0x03, 0x49, 0x2a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x3a, 0x33, 0x0a,
	0x31, 0x12, 0x2f, 0x0a, 0x2d, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x17, 0x12, 0x15, 0x0a, 0x13, 0x2e,
	0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x6f,
	0x72, 0x74, 0xb2, 0xea, 0xff, 0xf9, 0x01, 0x08, 0x12, 0x06, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x56, 0x6c, 0x61, 0x64, 0x69, 0x6d, 0x69, 0x72, 0x54, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_main_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: bookservermicro.Empty
	(*GetBook)(nil),               // 1: bookservermicro.getBook
	(*SortType)(nil),              // 2: bookservermicro.sortType
	(*GetAllBooks)(nil),           // 3: bookservermicro.getAllBooks
	(*GetAllBooksAndSort)(nil),    // 4: bookservermicro.getAllBooksAndSort
	(*PostBook)(nil),              // 5: bookservermicro.postBook
	(*GetBookRsp)(nil),            // 6: bookservermicro.getBookRsp
	(*StatusRsp)(nil),             // 7: bookservermicro.StatusRsp
	(*StatusUploadedBookRsp)(nil), // 8: bookservermicro.statusUploadedBookRsp
}
var file_main_proto_depIdxs = []int32{
	6, // 0: bookservermicro.getAllBooksAndSort.books:type_name -> bookservermicro.getBookRsp
	0, // 1: bookservermicro.BookServer.Home:input_type -> bookservermicro.Empty
	5, // 2: bookservermicro.BookServer.Push:input_type -> bookservermicro.postBook
	1, // 3: bookservermicro.BookServer.Book:input_type -> bookservermicro.getBook
	0, // 4: bookservermicro.BookServer.GetAllBooks:input_type -> bookservermicro.Empty
	2, // 5: bookservermicro.BookServer.GetAllBooksAndSort:input_type -> bookservermicro.sortType
	7, // 6: bookservermicro.BookServer.Home:output_type -> bookservermicro.StatusRsp
	8, // 7: bookservermicro.BookServer.Push:output_type -> bookservermicro.statusUploadedBookRsp
	6, // 8: bookservermicro.BookServer.Book:output_type -> bookservermicro.getBookRsp
	3, // 9: bookservermicro.BookServer.GetAllBooks:output_type -> bookservermicro.getAllBooks
	4, // 10: bookservermicro.BookServer.GetAllBooksAndSort:output_type -> bookservermicro.getAllBooksAndSort
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			switch v := v.(*GetBook); i {
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
			switch v := v.(*SortType); i {
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
			switch v := v.(*GetAllBooks); i {
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
			switch v := v.(*GetAllBooksAndSort); i {
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
		file_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_main_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_main_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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
